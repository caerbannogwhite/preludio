package gandalff

import (
	"math"
	"sync"
)

type __stats_thread_data struct {
	op      AggregateType
	gi      int
	indeces []int
	series  Series
	res     []float64
}

func __stats_worker(wg *sync.WaitGroup, buffer <-chan __stats_thread_data) {
	for td := range buffer {
		switch td.op {
		case AGGREGATE_SUM:
			switch series := td.series.(type) {
			case SeriesBool:

			case SeriesInt:
				sum_ := int(0)
				data := series.getDataPtr()
				for _, i := range td.indeces {
					sum_ += (*data)[i]
				}
				td.res[td.gi] = float64(sum_)

			case SeriesInt64:
				sum_ := int64(0)
				data := series.getDataPtr()
				for _, i := range td.indeces {
					sum_ += (*data)[i]
				}
				td.res[td.gi] = float64(sum_)

			case SeriesFloat64:
				sum_ := float64(0)
				data := series.getDataPtr()
				for _, i := range td.indeces {
					sum_ += (*data)[i]
				}
				td.res[td.gi] = sum_
			}
		case AGGREGATE_MIN:

		case AGGREGATE_MAX:

		case AGGREGATE_MEAN:
			switch series := td.series.(type) {
			case SeriesBool:

			case SeriesInt:
				sum_ := int(0)
				data := series.getDataPtr()
				for _, i := range td.indeces {
					sum_ += (*data)[i]
				}
				td.res[td.gi] = float64(sum_) / float64(len(td.indeces))

			case SeriesInt64:
				sum_ := int64(0)
				data := series.getDataPtr()
				for _, i := range td.indeces {
					sum_ += (*data)[i]
				}
				td.res[td.gi] = float64(sum_) / float64(len(td.indeces))

			case SeriesFloat64:
				sum_ := float64(0)
				data := series.getDataPtr()
				for _, i := range td.indeces {
					sum_ += (*data)[i]
				}
				td.res[td.gi] = sum_ / float64(len(td.indeces))
			}
		}
	}
	wg.Done()
}

func __gdl_sum__(s Series) float64 {
	sum := 0.0
	switch series := s.(type) {
	case SeriesBool:
		data := *series.getDataPtr()
		for i := 0; i < series.Len(); i++ {
			if data[i] {
				sum += 1.0
			}
		}
		return sum

	case SeriesInt:
		data := *series.getDataPtr()
		sum_ := int(0)
		for i := 0; i < series.Len(); i++ {
			sum_ += data[i]
		}
		return float64(sum_)

	case SeriesInt64:
		data := *series.getDataPtr()
		sum_ := int64(0)
		for i := 0; i < series.Len(); i++ {
			sum_ += data[i]
		}
		return float64(sum_)

	case SeriesFloat64:
		data := *series.getDataPtr()
		for i := 0; i < series.Len(); i++ {
			sum += data[i]
		}
		return sum

	default:
		return 0.0
	}
}

func __gdl_sum_grouped__(s Series, groups [][]int) []float64 {
	sum := make([]float64, len(groups))
	switch series := s.(type) {
	case SeriesBool:
		data := *series.getDataPtr()
		for gi, group := range groups {
			for _, i := range group {
				if data[i] {
					sum[gi] += 1.0
				}
			}
		}
		return sum

	case SeriesInt:
		data := *series.getDataPtr()

		// SINGLE THREAD
		if len(data) < MINIMUM_PARALLEL_SIZE_2 || len(groups) < THREADS_NUMBER {
			for gi, group := range groups {
				sum_ := int(0)
				for _, i := range group {
					sum_ += data[i]
				}
				sum[gi] = float64(sum_)
			}
			return sum
		}

		type threadData struct {
			gi      int
			indeces []int
		}

		var wg sync.WaitGroup
		wg.Add(THREADS_NUMBER)

		buffer := make(chan threadData)

		worker := func() {
			for td := range buffer {
				sum_ := int(0)
				for _, i := range td.indeces {
					sum_ += data[i]
				}
				sum[td.gi] = float64(sum_)
			}
			wg.Done()
		}

		for i := 0; i < THREADS_NUMBER; i++ {
			go worker()
		}

		for gi, group := range groups {
			buffer <- threadData{gi, group}
		}

		close(buffer)
		wg.Wait()

		return sum

	case SeriesInt64:
		data := *series.getDataPtr()

		// SINGLE THREAD
		if len(data) < MINIMUM_PARALLEL_SIZE_2 || len(groups) < THREADS_NUMBER {
			for gi, group := range groups {
				sum_ := int64(0)
				for _, i := range group {
					sum_ += data[i]
				}
				sum[gi] = float64(sum_)
			}
			return sum
		}

		type threadData struct {
			gi      int
			indeces []int
		}

		var wg sync.WaitGroup
		wg.Add(THREADS_NUMBER)

		buffer := make(chan threadData)

		worker := func() {
			for td := range buffer {
				sum_ := int64(0)
				for _, i := range td.indeces {
					sum_ += data[i]
				}
				sum[td.gi] = float64(sum_)
			}
			wg.Done()
		}

		for i := 0; i < THREADS_NUMBER; i++ {
			go worker()
		}

		for gi, group := range groups {
			buffer <- threadData{gi, group}
		}

		close(buffer)
		wg.Wait()

		return sum

	case SeriesFloat64:
		data := *series.getDataPtr()

		// SINGLE THREAD
		if len(data) < MINIMUM_PARALLEL_SIZE_2 || len(groups) < THREADS_NUMBER {
			for gi, group := range groups {
				sum_ := float64(0)
				for _, i := range group {
					sum_ += data[i]
				}
				sum[gi] = sum_
			}
			return sum
		}

		// MULTI THREAD
		type threadData struct {
			gi      int
			indeces []int
		}

		var wg sync.WaitGroup
		wg.Add(THREADS_NUMBER)

		buffer := make(chan threadData)

		worker := func() {
			for td := range buffer {
				sum_ := float64(0)
				for _, i := range td.indeces {
					sum_ += data[i]
				}
				sum[td.gi] = float64(sum_)
			}
			wg.Done()
		}

		for i := 0; i < THREADS_NUMBER; i++ {
			go worker()
		}

		for gi, group := range groups {
			buffer <- threadData{gi, group}
		}

		close(buffer)
		wg.Wait()

		return sum

	default:
		return sum
	}
}

func __gdl_min__(s Series) float64 {
	min := math.MaxFloat64
	switch series := s.(type) {
	case SeriesBool:
		if series.isNullable {
			for i := 0; i < series.Len(); i++ {
				if !series.IsNull(i) && !series.Get(i).(bool) {
					min = 0.0
					break
				}
			}
			return min
		} else {
			for i := 0; i < series.Len(); i++ {
				if !series.Get(i).(bool) {
					min = 0.0
					break
				}
			}
			return min
		}

	case SeriesInt:
		if series.isNullable {
			for i := 0; i < series.Len(); i++ {
				if !series.IsNull(i) {
					min = math.Min(min, float64(series.Get(i).(int)))
				}
			}
			return min
		} else {
			for i := 0; i < series.Len(); i++ {
				min = math.Min(min, float64(series.Get(i).(int)))
			}
			return min
		}

	case SeriesInt64:
		if series.isNullable {
			for i := 0; i < series.Len(); i++ {
				if !series.IsNull(i) {
					min = math.Min(min, float64(series.Get(i).(int64)))
				}
			}
			return min
		} else {
			for i := 0; i < series.Len(); i++ {
				min = math.Min(min, float64(series.Get(i).(int64)))
			}
			return min
		}

	case SeriesFloat64:
		if series.isNullable {
			for i := 0; i < series.Len(); i++ {
				if !series.IsNull(i) {
					min = math.Min(min, series.Get(i).(float64))
				}
			}
			return min
		} else {
			for i := 0; i < series.Len(); i++ {
				min = math.Min(min, series.Get(i).(float64))
			}
			return min
		}

	default:
		return 0.0
	}
}

func __gdl_min_grouped__(s Series, groups [][]int) []float64 {
	min := make([]float64, len(groups))
	for i := range min {
		min[i] = math.MaxFloat64
	}
	switch series := s.(type) {
	case SeriesBool:
		if series.isNullable {
			for gi, group := range groups {
				for _, i := range group {
					if !series.IsNull(i) && !series.Get(i).(bool) {
						min[gi] = 0.0
						break
					}
				}
			}
			return min
		} else {
			for gi, group := range groups {
				for _, i := range group {
					if !series.Get(i).(bool) {
						min[gi] = 0.0
						break
					}
				}
			}
			return min
		}

	case SeriesInt:
		if series.isNullable {
			for gi, group := range groups {
				for _, i := range group {
					if !series.IsNull(i) {
						min[gi] = math.Min(min[gi], float64(series.Get(i).(int)))
					}
				}
			}
			return min
		} else {
			for gi, group := range groups {
				for _, i := range group {
					min[gi] = math.Min(min[gi], float64(series.Get(i).(int)))
				}
			}
			return min
		}

	case SeriesInt64:
		if series.isNullable {
			for gi, group := range groups {
				for _, i := range group {
					if !series.IsNull(i) {
						min[gi] = math.Min(min[gi], float64(series.Get(i).(int64)))
					}
				}
			}
			return min
		} else {
			for gi, group := range groups {
				for _, i := range group {
					min[gi] = math.Min(min[gi], float64(series.Get(i).(int64)))
				}
			}
			return min
		}

	case SeriesFloat64:
		if series.isNullable {
			for gi, group := range groups {
				for _, i := range group {
					if !series.IsNull(i) {
						min[gi] = math.Min(min[gi], series.Get(i).(float64))
					}
				}
			}
			return min
		} else {
			for gi, group := range groups {
				for _, i := range group {
					min[gi] = math.Min(min[gi], series.Get(i).(float64))
				}
			}
			return min
		}

	default:
		for i := range min {
			min[i] = 0.0
		}
		return min
	}
}

func __gdl_max__(s Series) float64 {
	max := -math.MaxFloat64
	switch series := s.(type) {
	case SeriesBool:
		if series.isNullable {
			for i := 0; i < series.Len(); i++ {
				if !series.IsNull(i) && series.Get(i).(bool) {
					max = 1.0
					break
				}
			}
			return max
		} else {
			for i := 0; i < series.Len(); i++ {
				if series.Get(i).(bool) {
					max = 1.0
					break
				}
			}
			return max
		}

	case SeriesInt:
		if series.isNullable {
			for i := 0; i < series.Len(); i++ {
				if !series.IsNull(i) {
					max = math.Max(max, float64(series.Get(i).(int)))
				}
			}
			return max
		} else {
			for i := 0; i < series.Len(); i++ {
				max = math.Max(max, float64(series.Get(i).(int)))
			}
			return max
		}

	case SeriesInt64:
		if series.isNullable {
			for i := 0; i < series.Len(); i++ {
				if !series.IsNull(i) {
					max = math.Max(max, float64(series.Get(i).(int64)))
				}
			}
			return max
		} else {
			for i := 0; i < series.Len(); i++ {
				max = math.Max(max, float64(series.Get(i).(int64)))
			}
			return max
		}

	case SeriesFloat64:
		if series.isNullable {
			for i := 0; i < series.Len(); i++ {
				if !series.IsNull(i) {
					max = math.Max(max, series.Get(i).(float64))
				}
			}
			return max
		} else {
			for i := 0; i < series.Len(); i++ {
				max = math.Max(max, series.Get(i).(float64))
			}
			return max
		}

	default:
		return 0.0
	}
}

func __gdl_max_grouped__(s Series, groups [][]int) []float64 {
	max := make([]float64, len(groups))
	for i := range max {
		max[i] = -math.MaxFloat64
	}
	switch series := s.(type) {
	case SeriesBool:
		if series.isNullable {
			for gi, group := range groups {
				for _, i := range group {
					if !series.IsNull(i) && series.Get(i).(bool) {
						max[gi] = 1.0
						break
					}
				}
			}
			return max
		} else {
			for gi, group := range groups {
				for _, i := range group {
					if series.Get(i).(bool) {
						max[gi] = 1.0
						break
					}
				}
			}
			return max
		}

	case SeriesInt:
		if series.isNullable {
			for gi, group := range groups {
				for _, i := range group {
					if !series.IsNull(i) {
						max[gi] = math.Max(max[gi], float64(series.Get(i).(int)))
					}
				}
			}
			return max
		} else {
			for gi, group := range groups {
				for _, i := range group {
					max[gi] = math.Max(max[gi], float64(series.Get(i).(int)))
				}
			}
			return max
		}

	case SeriesInt64:
		if series.isNullable {
			for gi, group := range groups {
				for _, i := range group {
					if !series.IsNull(i) {
						max[gi] = math.Max(max[gi], float64(series.Get(i).(int64)))
					}
				}
			}
			return max
		} else {
			for gi, group := range groups {
				for _, i := range group {
					max[gi] = math.Max(max[gi], float64(series.Get(i).(int64)))
				}
			}
			return max
		}

	case SeriesFloat64:
		if series.isNullable {
			for gi, group := range groups {
				for _, i := range group {
					if !series.IsNull(i) {
						max[gi] = math.Max(max[gi], series.Get(i).(float64))
					}
				}
			}
			return max
		} else {
			for gi, group := range groups {
				for _, i := range group {
					max[gi] = math.Max(max[gi], series.Get(i).(float64))
				}
			}
			return max
		}

	default:
		for i := range max {
			max[i] = 0.0
		}
		return max
	}
}

func __gdl_mean__(s Series) float64 {
	switch series := s.(type) {
	case SeriesBool:
		sum := 0.0
		data := *series.getDataPtr()
		for i := 0; i < series.Len(); i++ {
			if data[i] {
				sum += 1.0
			}
		}
		return sum / float64(series.Len())

	case SeriesInt:
		sum_ := int(0)
		for i := 0; i < series.Len(); i++ {
			sum_ += series.Get(i).(int)
		}
		return float64(sum_) / float64(series.Len())

	case SeriesInt64:
		sum_ := int64(0)
		for i := 0; i < series.Len(); i++ {
			sum_ += series.Get(i).(int64)
		}
		return float64(sum_) / float64(series.Len())

	case SeriesFloat64:
		sum := 0.0
		for i := 0; i < series.Len(); i++ {
			sum += series.Get(i).(float64)
		}
		return sum / float64(series.Len())

	default:
		return 0.0
	}
}

func __gdl_mean_grouped__(s Series, groups [][]int) []float64 {
	sum := make([]float64, len(groups))
	switch series := s.(type) {
	case SeriesBool:
		data := *series.getDataPtr()
		for gi, group := range groups {
			for _, i := range group {
				if data[i] {
					sum[gi] += 1.0
				}
			}
			sum[gi] /= float64(len(group))
		}
		return sum

	case SeriesInt:
		data := *series.getDataPtr()

		// SINGLE THREAD
		if len(data) < MINIMUM_PARALLEL_SIZE_2 || len(groups) < THREADS_NUMBER {
			for gi, group := range groups {
				sum_ := int(0)
				for _, i := range group {
					sum_ += data[i]
				}
				sum[gi] = float64(sum_) / float64(len(group))
			}
			return sum
		}

		// MULTI THREAD
		type threadData struct {
			gi      int
			indeces []int
		}

		var wg sync.WaitGroup
		wg.Add(THREADS_NUMBER)

		buffer := make(chan threadData)

		worker := func() {
			for td := range buffer {
				sum_ := int(0)
				for _, i := range td.indeces {
					sum_ += data[i]
				}
				sum[td.gi] = float64(sum_) / float64(len(td.indeces))
			}
			wg.Done()
		}

		for i := 0; i < THREADS_NUMBER; i++ {
			go worker()
		}

		for gi, group := range groups {
			buffer <- threadData{gi, group}
		}

		close(buffer)
		wg.Wait()

		return sum

	case SeriesInt64:
		data := *series.getDataPtr()

		// SINGLE THREAD
		if len(data) < MINIMUM_PARALLEL_SIZE_2 || len(groups) < THREADS_NUMBER {
			for gi, group := range groups {
				sum_ := int64(0)
				for _, i := range group {
					sum_ += data[i]
				}
				sum[gi] = float64(sum_) / float64(len(group))
			}
			return sum
		}

		// MULTI THREAD
		type threadData struct {
			gi      int
			indeces []int
		}

		var wg sync.WaitGroup
		wg.Add(THREADS_NUMBER)

		buffer := make(chan threadData)

		worker := func() {
			for td := range buffer {
				sum_ := int64(0)
				for _, i := range td.indeces {
					sum_ += data[i]
				}
				sum[td.gi] = float64(sum_) / float64(len(td.indeces))
			}
			wg.Done()
		}

		for i := 0; i < THREADS_NUMBER; i++ {
			go worker()
		}

		for gi, group := range groups {
			buffer <- threadData{gi, group}
		}

		close(buffer)
		wg.Wait()

		return sum

	case SeriesFloat64:
		data := *series.getDataPtr()

		// SINGLE THREAD
		if len(data) < MINIMUM_PARALLEL_SIZE_2 || len(groups) < THREADS_NUMBER {
			for gi, group := range groups {
				sum_ := float64(0)
				for _, i := range group {
					sum_ += data[i]
				}
				sum[gi] = float64(sum_) / float64(len(group))
			}
			return sum
		}

		// MULTI THREAD
		type threadData struct {
			gi      int
			indeces []int
		}

		var wg sync.WaitGroup
		wg.Add(THREADS_NUMBER)

		buffer := make(chan threadData)

		worker := func() {
			for td := range buffer {
				sum_ := float64(0)
				for _, i := range td.indeces {
					sum_ += data[i]
				}
				sum[td.gi] = float64(sum_) / float64(len(td.indeces))
			}
			wg.Done()
		}

		for i := 0; i < THREADS_NUMBER; i++ {
			go worker()
		}

		for gi, group := range groups {
			buffer <- threadData{gi, group}
		}

		close(buffer)
		wg.Wait()

		return sum

	default:
		return sum
	}
}

func __gdl_std__(s Series) float64 {
	mean := __gdl_mean__(s)
	sum := 0.0
	switch series := s.(type) {
	case SeriesBool:
		if series.isNullable {
			for i := 0; i < series.Len(); i++ {
				if !series.IsNull(i) {
					if series.Get(i).(bool) {
						sum += math.Pow(1.0-mean, 2)
					} else {
						sum += math.Pow(-mean, 2)
					}
				}
			}
			return math.Sqrt(sum / float64(series.Len()))
		} else {
			for i := 0; i < series.Len(); i++ {
				if series.Get(i).(bool) {
					sum += math.Pow(1.0-mean, 2)
				} else {
					sum += math.Pow(-mean, 2)
				}
			}
			return math.Sqrt(sum / float64(series.Len()))
		}

	case SeriesInt:
		if series.isNullable {
			for i := 0; i < series.Len(); i++ {
				if !series.IsNull(i) {
					sum += math.Pow(float64(series.Get(i).(int))-mean, 2)
				}
			}
			return math.Sqrt(sum / float64(series.Len()))
		} else {
			for i := 0; i < series.Len(); i++ {
				sum += math.Pow(float64(series.Get(i).(int))-mean, 2)
			}
			return math.Sqrt(sum / float64(series.Len()))
		}

	case SeriesFloat64:
		if series.isNullable {
			for i := 0; i < series.Len(); i++ {
				if !series.IsNull(i) {
					sum += math.Pow(series.Get(i).(float64)-mean, 2)
				}
			}
			return math.Sqrt(sum / float64(series.Len()))
		} else {
			for i := 0; i < series.Len(); i++ {
				sum += math.Pow(series.Get(i).(float64)-mean, 2)
			}
			return math.Sqrt(sum / float64(series.Len()))
		}

	default:
		return 0.0
	}
}

func __gdl_std_grouped__(s Series, groups [][]int) []float64 {
	stddev := make([]float64, len(groups))
	for i := range stddev {
		stddev[i] = 0.0
	}
	mean := __gdl_mean_grouped__(s, groups)
	switch series := s.(type) {
	case SeriesBool:
		if series.isNullable {
			for gi, group := range groups {
				sum := 0.0
				for _, i := range group {
					if !series.IsNull(i) {
						if series.Get(i).(bool) {
							sum += math.Pow(1.0-mean[gi], 2)
						} else {
							sum += math.Pow(-mean[gi], 2)
						}
					}
				}
				stddev[gi] = math.Sqrt(sum / float64(len(group)))
			}
			return stddev
		} else {
			for gi, group := range groups {
				sum := 0.0
				for _, i := range group {
					if series.Get(i).(bool) {
						sum += math.Pow(1.0-mean[gi], 2)
					} else {
						sum += math.Pow(-mean[gi], 2)
					}
				}
				stddev[gi] = math.Sqrt(sum / float64(len(group)))
			}
			return stddev
		}

	case SeriesInt:
		if series.isNullable {
			for gi, group := range groups {
				sum := 0.0
				for _, i := range group {
					if !series.IsNull(i) {
						sum += math.Pow(float64(series.Get(i).(int))-mean[gi], 2)
					}
				}
				stddev[gi] = math.Sqrt(sum / float64(len(group)))
			}
			return stddev
		} else {
			for gi, group := range groups {
				sum := 0.0
				for _, i := range group {
					sum += math.Pow(float64(series.Get(i).(int))-mean[gi], 2)
				}
				stddev[gi] = math.Sqrt(sum / float64(len(group)))
			}
			return stddev
		}

	case SeriesFloat64:
		if series.isNullable {
			for gi, group := range groups {
				sum := 0.0
				for _, i := range group {
					if !series.IsNull(i) {
						sum += math.Pow(series.Get(i).(float64)-mean[gi], 2)
					}
				}
				stddev[gi] = math.Sqrt(sum / float64(len(group)))
			}
			return stddev
		} else {
			for gi, group := range groups {
				sum := 0.0
				for _, i := range group {
					sum += math.Pow(series.Get(i).(float64)-mean[gi], 2)
				}
				stddev[gi] = math.Sqrt(sum / float64(len(group)))
			}
			return stddev
		}

	default:
		return stddev
	}
}
