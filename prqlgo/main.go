package main

import "os"

func main() {
	inputPath := os.Args[1]

	vm := NewPrqlVirtualMachine(&PrqlVirtualMachineParams{
		PrintWarnings: true,
		DebugLevel:    0,
		InputPath:     inputPath,
	})

	vm.ReadPrqlBytecode()

	// f, _ := os.Open("C:\\Users\\massi\\Downloads\\Cars.csv")

	// cars := dataframe.ReadCSV(f, dataframe.WithDelimiter(';'))

	// switch cars.Col("Car").Type() {
	// case series.Bool:
	// 	fmt.Println("it's bool")
	// case series.Int:
	// 	fmt.Println("it's int")
	// case series.String:
	// 	fmt.Println("it's string")
	// }

	// mean := func(s series.Series) series.Series {
	// 	floats := s.Float()

	// 	sum := 0.0
	// 	for _, f := range floats {
	// 		sum += f
	// 	}
	// 	return series.Floats(sum / float64(len(floats)))
	// }

	// cars2 := cars.Select(2).Capply(mean)
	// fmt.Println(cars2)

	// fmt.Printf("%t %T\n", false, true)
}
