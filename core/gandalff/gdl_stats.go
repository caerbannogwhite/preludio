package gandalff

import "math"

func __gdl_sum__(s GDLSeries) float64 {
	sum := 0.0
	switch series := s.(type) {
	case GDLSeriesBool:
		if series.isNullable {
			for i := 0; i < series.Len(); i++ {
				if !series.IsNull(i) && series.Get(i).(bool) {
					sum += 1.0
				}
			}
			return sum
		} else {
			for i := 0; i < series.Len(); i++ {
				if series.Get(i).(bool) {
					sum += 1.0
				}
			}
			return sum
		}

	case GDLSeriesInt32:
		if series.isNullable {
			for i := 0; i < series.Len(); i++ {
				if !series.IsNull(i) {
					sum += float64(series.Get(i).(int))
				}
			}
			return sum
		} else {
			for i := 0; i < series.Len(); i++ {
				sum += float64(series.Get(i).(int))
			}
			return sum
		}

	case GDLSeriesFloat64:
		if series.isNullable {
			for i := 0; i < series.Len(); i++ {
				if !series.IsNull(i) {
					sum += series.Get(i).(float64)
				}
			}
			return sum
		} else {
			for i := 0; i < series.Len(); i++ {
				sum += series.Get(i).(float64)
			}
			return sum
		}

	default:
		return 0.0
	}
}

func __gdl_sum_grouped__(s GDLSeries, groups [][]int) []float64 {
	sum := make([]float64, len(groups))
	switch series := s.(type) {
	case GDLSeriesBool:
		if series.isNullable {
			for gi, group := range groups {
				for _, i := range group {
					if !series.IsNull(i) && series.Get(i).(bool) {
						sum[gi] += 1.0
					}
				}
			}
			return sum
		} else {
			for gi, group := range groups {
				for _, i := range group {
					if series.Get(i).(bool) {
						sum[gi] += 1.0
					}
				}
			}
			return sum
		}

	case GDLSeriesInt32:
		if series.isNullable {
			for gi, group := range groups {
				for _, i := range group {
					if !series.IsNull(i) {
						sum[gi] += float64(series.Get(i).(int))
					}
				}
			}
			return sum
		} else {
			for gi, group := range groups {
				for _, i := range group {
					sum[gi] += float64(series.Get(i).(int))
				}
			}
			return sum
		}

	case GDLSeriesFloat64:
		if series.isNullable {
			for gi, group := range groups {
				for _, i := range group {
					if !series.IsNull(i) {
						sum[gi] += series.Get(i).(float64)
					}
				}
			}
			return sum
		} else {
			for gi, group := range groups {
				for _, i := range group {
					sum[gi] += series.Get(i).(float64)
				}
			}
			return sum
		}

	default:
		return sum
	}
}

func __gdl_min__(s GDLSeries) float64 {
	min := math.MaxFloat64
	switch series := s.(type) {
	case GDLSeriesBool:
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

	case GDLSeriesInt32:
		if series.isNullable {
			for i := 0; i < series.Len(); i++ {
				if !series.IsNull(i) {
					min = math.Min(min, float64(series.Get(i).(int32)))
				}
			}
			return min
		} else {
			for i := 0; i < series.Len(); i++ {
				min = math.Min(min, float64(series.Get(i).(int32)))
			}
			return min
		}

	case GDLSeriesFloat64:
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

func __gdl_min_grouped__(s GDLSeries, groups [][]int) []float64 {
	min := make([]float64, len(groups))
	for i := range min {
		min[i] = math.MaxFloat64
	}
	switch series := s.(type) {
	case GDLSeriesBool:
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

	case GDLSeriesInt32:
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

	case GDLSeriesFloat64:
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

func __gdl_max__(s GDLSeries) float64 {
	max := -math.MaxFloat64
	switch series := s.(type) {
	case GDLSeriesBool:
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

	case GDLSeriesInt32:
		if series.isNullable {
			for i := 0; i < series.Len(); i++ {
				if !series.IsNull(i) {
					max = math.Max(max, float64(series.Get(i).(int32)))
				}
			}
			return max
		} else {
			for i := 0; i < series.Len(); i++ {
				max = math.Max(max, float64(series.Get(i).(int32)))
			}
			return max
		}

	case GDLSeriesFloat64:
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

func __gdl_max_grouped__(s GDLSeries, groups [][]int) []float64 {
	max := make([]float64, len(groups))
	for i := range max {
		max[i] = -math.MaxFloat64
	}
	switch series := s.(type) {
	case GDLSeriesBool:
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

	case GDLSeriesInt32:
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

	case GDLSeriesFloat64:
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

func __gdl_mean__(s GDLSeries) float64 {
	sum := 0.0
	switch series := s.(type) {
	case GDLSeriesBool:
		if series.isNullable {
			for i := 0; i < series.Len(); i++ {
				if !series.IsNull(i) && series.Get(i).(bool) {
					sum += 1.0
				}
			}
			return sum / float64(series.NonNullCount())
		} else {
			for i := 0; i < series.Len(); i++ {
				if series.Get(i).(bool) {
					sum += 1.0
				}
			}
			return sum / float64(series.Len())
		}

	case GDLSeriesInt32:
		if series.isNullable {
			for i := 0; i < series.Len(); i++ {
				if !series.IsNull(i) {
					sum += float64(series.Get(i).(int))
				}
			}
			return sum / float64(series.NonNullCount())
		} else {
			for i := 0; i < series.Len(); i++ {
				sum += float64(series.Get(i).(int))
			}
			return sum / float64(series.Len())
		}

	case GDLSeriesFloat64:
		if series.isNullable {
			for i := 0; i < series.Len(); i++ {
				if !series.IsNull(i) {
					sum += series.Get(i).(float64)
				}
			}
			return sum / float64(series.NonNullCount())
		} else {
			for i := 0; i < series.Len(); i++ {
				sum += series.Get(i).(float64)
			}
			return sum / float64(series.Len())
		}

	default:
		return 0.0
	}
}

func __gdl_mean_grouped__(s GDLSeries, groups [][]int) []float64 {
	sum := make([]float64, len(groups))
	switch series := s.(type) {
	case GDLSeriesBool:
		if series.isNullable {
			for gi, group := range groups {
				for _, i := range group {
					if !series.IsNull(i) && series.Get(i).(bool) {
						sum[gi] += 1.0
					}
				}
				sum[gi] /= float64(len(group))
			}
			return sum
		} else {
			for gi, group := range groups {
				for _, i := range group {
					if series.Get(i).(bool) {
						sum[gi] += 1.0
					}
				}
				sum[gi] /= float64(len(group))
			}
			return sum
		}

	case GDLSeriesInt32:
		if series.isNullable {
			for gi, group := range groups {
				for _, i := range group {
					if !series.IsNull(i) {
						sum[gi] += float64(series.Get(i).(int))
					}
				}
				sum[gi] /= float64(len(group))
			}
			return sum
		} else {
			for gi, group := range groups {
				for _, i := range group {
					sum[gi] += float64(series.Get(i).(int))
				}
				sum[gi] /= float64(len(group))
			}
			return sum
		}

	case GDLSeriesFloat64:
		if series.isNullable {
			for gi, group := range groups {
				for _, i := range group {
					if !series.IsNull(i) {
						sum[gi] += series.Get(i).(float64)
					}
				}
				sum[gi] /= float64(len(group))
			}
			return sum
		} else {
			for gi, group := range groups {
				for _, i := range group {
					sum[gi] += series.Get(i).(float64)
				}
				sum[gi] /= float64(len(group))
			}
			return sum
		}

	default:
		return sum
	}
}

func __gdl_std__(s GDLSeries) float64 {
	mean := __gdl_mean__(s)
	sum := 0.0
	switch series := s.(type) {
	case GDLSeriesBool:
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

	case GDLSeriesInt32:
		if series.isNullable {
			for i := 0; i < series.Len(); i++ {
				if !series.IsNull(i) {
					sum += math.Pow(float64(series.Get(i).(int32))-mean, 2)
				}
			}
			return math.Sqrt(sum / float64(series.Len()))
		} else {
			for i := 0; i < series.Len(); i++ {
				sum += math.Pow(float64(series.Get(i).(int32))-mean, 2)
			}
			return math.Sqrt(sum / float64(series.Len()))
		}

	case GDLSeriesFloat64:
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

func __gdl_std_grouped__(s GDLSeries, groups [][]int) []float64 {
	stddev := make([]float64, len(groups))
	for i := range stddev {
		stddev[i] = 0.0
	}
	mean := __gdl_mean_grouped__(s, groups)
	switch series := s.(type) {
	case GDLSeriesBool:
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

	case GDLSeriesInt32:
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

	case GDLSeriesFloat64:
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
