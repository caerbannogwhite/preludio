package gandalff

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
					sum += float64(series.Get(i).(int32))
				}
			}
			return sum / float64(series.NonNullCount())
		} else {
			for i := 0; i < series.Len(); i++ {
				sum += float64(series.Get(i).(int32))
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
