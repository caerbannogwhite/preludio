package gandalff

import (
	"math"
	"sync"
)

func __series_groupby_multithreaded(
	threadNum, dataLen int,
	maps *[]map[uint64][]int,
	worker func(start, end int, map_ *map[uint64][]int),

) {
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

		map_ := (*maps)[idx]
		if idx == threadNum-1 {
			end = dataLen
		}

		worker(start, end, &map_)

		// Notify the wait groups at the first level
		wg[0][idx/2].Done()
	}

	merger := func(level, idx1, idx2 int) {
		wg[level][idx1>>uint(level+1)].Wait()
		wg[level][idx2>>uint(level+1)].Wait()

		for k, v := range (*maps)[idx2] {
			(*maps)[idx1][k] = append((*maps)[idx1][k], v...)
		}

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
