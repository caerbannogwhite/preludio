package gandalff

import "fmt"

func seriesTakePreprocess(type_ string, size int, params ...int) ([]int, error) {
	switch len(params) {
	case 1:
		// only one parameter, so it must be the number of elements to take
		if params[0] < 0 {
			params[0] = size + params[0]
		}
		if params[0] < 0 {
			return nil, fmt.Errorf("%s.Take: number of elements to take (%d) is less than zero", type_, params[0])
		}
		if params[0] > size {
			return nil, fmt.Errorf("%s.Take: number of elements to take (%d) is greater than the size of the series (%d)", type_, params[0], size)
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
			return nil, fmt.Errorf("%s.Take: start index (%d) is greater than end index (%d)", type_, params[0], params[1])
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
			return nil, fmt.Errorf("%s.Take: start index (%d) is greater than end index (%d)", type_, params[0], params[1])
		}
		if params[2] == 0 {
			return nil, fmt.Errorf("%s.Take: step cannot be zero", type_)
		}
		indeces := make([]int, (params[1]-params[0])/params[2])
		for i := 0; i < (params[1]-params[0])/params[2]; i++ {
			indeces[i] = i*params[2] + params[0]
		}
		return indeces, nil

	default:
		return nil, fmt.Errorf("%s.Take: invalid number of parameters: %d", type_, len(params))
	}
}

func debugPrintPartition(p SeriesPartition, series ...Series) {
	map_ := p.getMap()

	header := ""
	separators := ""
	for i := range series {
		header += fmt.Sprintf("| %-10d ", i)
		separators += "|------------"
	}

	fmt.Println()
	fmt.Printf("    | %-20s %s | %-20s |\n", "Key", header, "Indeces")
	fmt.Printf("    |%s%s-|%s|\n", "----------------------", separators, "----------------------")
	for k, v := range map_ {
		vals := ""
		for _, s := range series {
			vals += fmt.Sprintf("| %-10s ", s.GetAsString(v[0]))
		}

		indeces := ""
		for _, i := range v {
			indeces += fmt.Sprintf("%d ", i)
		}
		fmt.Printf("    | %-20d %s | %-20s |\n", k, vals, indeces)
	}
	fmt.Println()
}
