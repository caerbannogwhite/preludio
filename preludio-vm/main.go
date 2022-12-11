package main

import (
	"github.com/alexflint/go-arg"
)

func main() {

	var args struct {
		InputPath  string `arg:"positional"`
		DebugLevel int    `arg:"-d, --debug-level" help:"debug level" default:"0"`
		Verbose    bool   `arg:"-v, --verbose" help:"verbosity level" default:"false"`
		Warnings   bool   `arg:"-w, --warnings" help:"print warnings" defaut:"true"`
	}

	arg.MustParse(&args)
	// fmt.Println(args.Foo, args.Bar)

	vm := NewPreludioVM(&PreludioVMParams{
		PrintWarnings: args.Warnings,
		DebugLevel:    args.DebugLevel,
		InputPath:     args.InputPath,
	})

	vm.ReadPreludioBytecode()

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
