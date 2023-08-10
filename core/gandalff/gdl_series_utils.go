package gandalff

import "fmt"

func seriesTakePreprocess(size int, params ...int) ([]int, error) {
	switch len(params) {
	case 1:
		// only one parameter, so it must be the number of elements to take
		if params[0] < 0 {
			params[0] = size + params[0]
		}
		if params[0] < 0 {
			return nil, fmt.Errorf("series.Take: number of elements to take (%d) is less than zero", params[0])
		}
		if params[0] > size {
			return nil, fmt.Errorf("series.Take: number of elements to take (%d) is greater than the size of the series (%d)", params[0], size)
		}
		indeces := make([]int, params[0])
		for i := 0; i < params[0]; i++ {
			indeces[i] = i
		}
		return indeces, nil

	case 2:
		// two parameters, so it must be the start and end indeces
		if params[0] < 0 {
			params[0] = size + params[0]
		}
		if params[1] < 0 {
			params[1] = size + params[1]
		}
		if params[0] > params[1] {
			return nil, fmt.Errorf("series.Take: start index (%d) is greater than end index (%d)", params[0], params[1])
		}
		indeces := make([]int, params[1]-params[0])
		for i := 0; i < params[1]-params[0]; i++ {
			indeces[i] = i + params[0]
		}
		return indeces, nil

	case 3:
		// three parameters, so it must be the start index, end index, and step
		if params[0] < 0 {
			params[0] = size + params[0]
		}
		if params[1] < 0 {
			params[1] = size + params[1]
		}
		if params[2] < 0 {
			params[2] = -params[2]
		}
		if params[0] > params[1] {
			return nil, fmt.Errorf("series.Take: start index (%d) is greater than end index (%d)", params[0], params[1])
		}
		if params[2] == 0 {
			return nil, fmt.Errorf("series.Take: step cannot be zero")
		}
		indeces := make([]int, (params[1]-params[0])/params[2])
		for i := 0; i < (params[1]-params[0])/params[2]; i++ {
			indeces[i] = i*params[2] + params[0]
		}
		return indeces, nil

	default:
		return nil, fmt.Errorf("series.Take: invalid number of parameters: %d", len(params))
	}
}

func debugPrintPartition(p SeriesPartition, series ...Series) {
	map_ := p.GetMap()

	header := ""
	separators := ""
	for _, s := range series {
		header += fmt.Sprintf("| %-6s", s.Name())
		separators += "|-------"
	}

	fmt.Println()
	fmt.Printf("    | %-20s %s | %-10s |\n", "Key", header, "Indeces")
	fmt.Printf("    |%s%s-|%s|\n", "----------------------", separators, "------------")
	for k, v := range map_ {
		vals := ""
		for _, s := range series {
			vals += fmt.Sprintf("| %-6s", s.GetString(v[0]))
		}

		indeces := ""
		for _, i := range v {
			indeces += fmt.Sprintf("%d ", i)
		}
		fmt.Printf("    | %-20d %s | %-10s |\n", k, vals, indeces)
	}
	fmt.Println()
}
