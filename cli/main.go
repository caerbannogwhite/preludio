package main

import (
	"preludio"

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

	vm := new(preludio.ByteEater).
		InitVM().
		SetPrintWarning(args.Warnings).
		SetDebugLevel(args.DebugLevel).
		SetInputPath(args.InputPath)

	vm.ReadFileBytecode()
}
