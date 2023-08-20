package gandalff

import (
	"math"
	"sync"
)

func __series_groupby(
	threadNum, minParallelSize, dataLen int, hasNulls bool,
	worker func(threadNum, start, end int, map_ map[int64][]int),
	workerNulls func(threadNum, start, end int, map_ map[int64][]int, nulls *[]int),
) map[int64][]int {

	// If the data is too small, just run the worker function
	if dataLen < minParallelSize {
		map_ := make(map[int64][]int)
		if hasNulls {
			nulls := make([]int, 0)
			workerNulls(0, 0, dataLen, map_, &nulls)

			// Add the nulls to the map
			if len(nulls) > 0 {
				nullKey := __series_get_nullkey(map_, HASH_NULL_KEY)
				map_[nullKey] = nulls
			}
		} else {
			worker(0, 0, dataLen, map_)
		}

		return map_
	}

	// Initialize the maps
	maps := make([]map[int64][]int, THREADS_NUMBER)
	for i := 0; i < THREADS_NUMBER; i++ {
		maps[i] = make(map[int64][]int, DEFAULT_HASH_MAP_INITIAL_CAPACITY)
	}

	var nulls [][]int
	if hasNulls {
		nulls = make([][]int, THREADS_NUMBER)
		for i := 0; i < THREADS_NUMBER; i++ {
			nulls[i] = make([]int, 0)
		}
	}

	// Initialize the wait groups array, one for each level where level is the
	// log2 of the number of threads.
	// The first lever of wait groups has THREADS_NUMBER/2 wait groups, and each
	// wait group waits for two threads.
	//
	// Example: if THREADS_NUMBER = 16, then
	//	- the FIRST level has 8 wait groups (each wait group waits for 2 threads)
	//	 	- the 1st element waits for threads 0 and 1
	//	 	- the 2nd element waits for threads 2 and 3
	//	 	- ...
	//    When workers 0 and 1 are done, the 1st element is notified so the fist
	//    merger of the second level can start
	//
	// 	- the second level has 4 wait groups
	// 	- ...
	levels := int(math.Log2(float64(threadNum)))
	wg := make([][]sync.WaitGroup, levels)
	for i := 0; i < levels; i++ {
		wg[i] = make([]sync.WaitGroup, threadNum/(1<<uint(i+1)))
		for j := 0; j < threadNum/(1<<uint(i+1)); j++ {
			wg[i][j].Add(2)
		}
	}

	// Edge case for the last level
	wg = append(wg, make([]sync.WaitGroup, 1))
	wg[len(wg)-1][0].Add(1)

	// Define the actual worker function
	// and the merger function
	var actualWorker func(int)
	var merger func(int, int, int)
	if hasNulls {
		actualWorker = func(idx int) {
			start := idx * dataLen / threadNum
			end := (idx + 1) * dataLen / threadNum
			if idx == threadNum-1 {
				end = dataLen
			}

			workerNulls(idx, start, end, maps[idx], &nulls[idx])

			// Notify the wait groups at the first level
			wg[0][idx/2].Done()
		}

		merger = func(level, idx1, idx2 int) {
			// Example: if THREADS_NUMBER = 16 and level = 0, then
			// 	- idx1 =  0, idx2 =  1 -> wait for wg[0][0]
			// 	- idx1 =  2, idx2 =  3 -> wait for wg[0][1]
			// 	- ...
			// 	- idx1 = 14, idx2 = 15 -> wait for wg[0][7]
			wg[level][idx1>>uint(level+1)].Wait()

			for k, v := range maps[idx2] {
				maps[idx1][k] = append(maps[idx1][k], v...)
			}

			if nulls != nil {
				nulls[idx1] = append(nulls[idx1], nulls[idx2]...)
			}

			// Notify the wait groups at the next level
			//
			// Example: if THREADS_NUMBER = 16 and level = 0, then
			// 	- idx1 =  0, idx2 =  1 -> notify wg[1][0]
			// 	- idx1 =  2, idx2 =  3 -> notify wg[1][0]
			// 	- ...
			// 	- idx1 = 14, idx2 = 15 -> notify wg[1][3]
			wg[level+1][idx1>>uint(level+2)].Done()
		}
	} else {
		actualWorker = func(idx int) {
			start := idx * dataLen / threadNum
			end := (idx + 1) * dataLen / threadNum
			if idx == threadNum-1 {
				end = dataLen
			}

			worker(idx, start, end, maps[idx])

			// Notify the wait groups at the first level
			wg[0][idx/2].Done()
		}

		merger = func(level, idx1, idx2 int) {
			// Example: if THREADS_NUMBER = 16 and level = 0, then
			// 	- idx1 =  0, idx2 =  1 -> wait for wg[0][0]
			// 	- idx1 =  2, idx2 =  3 -> wait for wg[0][1]
			// 	- ...
			// 	- idx1 = 14, idx2 = 15 -> wait for wg[0][7]
			wg[level][idx1>>uint(level+1)].Wait()

			for k, v := range maps[idx2] {
				maps[idx1][k] = append(maps[idx1][k], v...)
			}

			// Notify the wait groups at the next level
			//
			// Example: if THREADS_NUMBER = 16 and level = 0, then
			// 	- idx1 =  0, idx2 =  1 -> notify wg[1][0]
			// 	- idx1 =  2, idx2 =  3 -> notify wg[1][0]
			// 	- ...
			// 	- idx1 = 14, idx2 = 15 -> notify wg[1][3]
			wg[level+1][idx1>>uint(level+2)].Done()
		}
	}

	// Compute the submaps
	for i := 0; i < threadNum; i++ {
		go actualWorker(i)
	}

	// Merge the submaps
	for level := 0; level < levels; level++ {
		for i := 0; i < threadNum; i += (1 << uint(level+1)) {
			go merger(level, i, i+(1<<level))
		}
	}

	// Wait for the last level (there is only one wait group)
	wg[len(wg)-1][0].Wait()

	// Add the nulls to the map
	if hasNulls && len(nulls[0]) > 0 {
		nullKey := __series_get_nullkey(maps[0], HASH_NULL_KEY)
		maps[0][nullKey] = nulls[0]
	}

	return maps[0]
}

func __series_groupby_multithreaded(
	threadNum, dataLen int, maps []map[int64][]int, nulls [][]int,
	worker func(threadNum, start, end int),
) {
	// Initialize the wait groups array, one for each level where level is the
	// log2 of the number of threads.
	// The first lever of wait groups has THREADS_NUMBER/2 wait groups, and each
	// wait group waits for two threads.
	//
	// Example: if THREADS_NUMBER = 16, then
	//	- the FIRST level has 8 wait groups (each wait group waits for 2 threads)
	//	 	- the 1st element waits for threads 0 and 1
	//	 	- the 2nd element waits for threads 2 and 3
	//	 	- ...
	//    When workers 0 and 1 are done, the 1st element is notified so the fist
	//    merger of the second level can start
	//
	// 	- the second level has 4 wait groups
	// 	- ...
	levels := int(math.Log2(float64(threadNum)))
	wg := make([][]sync.WaitGroup, levels)
	for i := 0; i < levels; i++ {
		wg[i] = make([]sync.WaitGroup, threadNum/(1<<uint(i+1)))
		for j := 0; j < threadNum/(1<<uint(i+1)); j++ {
			wg[i][j].Add(2)
		}
	}

	// Edge case for the last level
	wg = append(wg, make([]sync.WaitGroup, 1))
	wg[len(wg)-1][0].Add(1)

	// Define the worker and merger functions
	actualWorker := func(idx int) {
		start := idx * dataLen / threadNum
		end := (idx + 1) * dataLen / threadNum
		if idx == threadNum-1 {
			end = dataLen
		}

		worker(idx, start, end)

		// Notify the wait groups at the first level
		wg[0][idx/2].Done()
	}

	merger := func(level, idx1, idx2 int) {
		// Example: if THREADS_NUMBER = 16 and level = 0, then
		// 	- idx1 =  0, idx2 =  1 -> wait for wg[0][0]
		// 	- idx1 =  2, idx2 =  3 -> wait for wg[0][1]
		// 	- ...
		// 	- idx1 = 14, idx2 = 15 -> wait for wg[0][7]
		wg[level][idx1>>uint(level+1)].Wait()

		for k, v := range maps[idx2] {
			maps[idx1][k] = append(maps[idx1][k], v...)
		}

		if nulls != nil {
			nulls[idx1] = append(nulls[idx1], nulls[idx2]...)
		}

		// Notify the wait groups at the next level
		//
		// Example: if THREADS_NUMBER = 16 and level = 0, then
		// 	- idx1 =  0, idx2 =  1 -> notify wg[1][0]
		// 	- idx1 =  2, idx2 =  3 -> notify wg[1][0]
		// 	- ...
		// 	- idx1 = 14, idx2 = 15 -> notify wg[1][3]
		wg[level+1][idx1>>uint(level+2)].Done()
	}

	// Compute the submaps
	for i := 0; i < threadNum; i++ {
		go actualWorker(i)
	}

	// Merge the submaps
	for level := 0; level < levels; level++ {
		for i := 0; i < threadNum; i += (1 << uint(level+1)) {
			go merger(level, i, i+(1<<level))
		}
	}

	// Wait for the last level (there is only one wait group)
	wg[len(wg)-1][0].Wait()
}

func __series_get_nullkey(map_ map[int64][]int, seed int64) int64 {
	for {
		if _, ok := map_[seed]; !ok {
			break
		}
		seed = HASH_MAGIC_NUMBER + (seed << 13) + (seed >> 4)
	}
	return seed
}
