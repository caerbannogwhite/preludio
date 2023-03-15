package main

import (
	"bufio"
	"fmt"
	"os"
	"preludiocli"
	"preludiocompiler"
	"preludiocore"
	"strings"

	"github.com/alexflint/go-arg"
	tea "github.com/charmbracelet/bubbletea"
)

const VERSION = "0.1.0-alpha"

type CliArgs struct {
	InputPath  string `arg:"-i, --input" help:"source file input path" default:""`
	DebugLevel int    `arg:"-d, --debug-level" help:"debug level" default:"0"`
	Editor     bool   `arg:"-e, --editor" help:"launch the text editor" default:"false"`
	Verbose    bool   `arg:"-v, --verbose" help:"verbosity level" default:"false"`
	Warnings   bool   `arg:"-w, --warnings" help:"print warnings" defaut:"true"`
}

func main() {

	var args CliArgs

	arg.MustParse(&args)

	if args.InputPath != "" {
		be := new(preludiocore.ByteEater).
			InitVM().
			SetPrintWarning(args.Warnings).
			SetDebugLevel(args.DebugLevel)

		bytecode := preludiocompiler.CompileFile(args.InputPath)
		if args.Verbose {
			fmt.Println("Bytecode generated")
		}
		be.RunBytecode(bytecode)

	} else if args.Editor {
		LaunchCodeEditor(args)
	} else {
		LaunchRepl(args)
	}
}

func LaunchCodeEditor(args CliArgs) {

	be := new(preludiocore.ByteEater).
		InitVM().
		SetPrintWarning(args.Warnings).
		SetDebugLevel(args.DebugLevel)

	codeEditor := preludiocli.NewCodeEditor().
		SetPreludioByteEater(*be)

	if _, err := tea.NewProgram(codeEditor).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

func LaunchRepl(args CliArgs) {

	fmt.Println("Welcome to the Preludio REPL!")
	fmt.Println("Version:", VERSION)

	be := new(preludiocore.ByteEater).
		InitVM().
		SetPrintWarning(args.Warnings).
		SetFullOutput(false).
		SetDebugLevel(args.DebugLevel).
		SetVerbose(args.Verbose)

	if args.Verbose {
		fmt.Printf("\nPreludio VM initialized\n")
		fmt.Printf("%15s %t\n", "Print warnings:", args.Warnings)
		fmt.Printf("%15s %d\n", "Debug level:", args.DebugLevel)
		fmt.Printf("%15s %t\n", "Verbose:", args.Verbose)
	}

	in := bufio.NewReader(os.Stdin)

	readerStart := true
	code := ""
	for {
		if readerStart {
			fmt.Print(">>> ")
		} else {
			fmt.Print("... ")
		}

		line, err := in.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				fmt.Println("Bye!")
				os.Exit(0)
			}
			fmt.Println("Error reading input:", err)
		}

		line = strings.TrimSpace(line)

		if line == "" {
			bytecode := preludiocompiler.CompileSource(code)
			be.RunBytecode(bytecode)

			res := be.GetOutput()
			for _, log := range res.Log {
				if log.LogType == preludiocore.LOG_DEBUG {
					if int(log.Level) < args.DebugLevel {
						fmt.Println(log.Message)
					}
				} else {
					fmt.Println(log.Message)
				}
			}

			for _, c := range res.Data {
				prettyPrint(10, c)
			}

			code = ""
			readerStart = true
		} else {
			readerStart = false
		}

		code += line + "\n"
	}
}

func truncate(s string, n int) string {
	if len(s) > n {
		return s[:n-3] + "..."
	}
	return s
}

func prettyPrint(colSize int, columnar []preludiocore.Columnar) {

	actualColSize := colSize + 3
	fmtString := fmt.Sprintf("| %%-%ds ", colSize)

	// header
	fmt.Printf("    ")
	for i := 0; i < len(columnar)*actualColSize; i++ {
		if i%actualColSize == 0 {
			fmt.Print("+")
		} else {
			fmt.Print("-")
		}
	}
	fmt.Println("+")

	// column names
	fmt.Printf("    ")
	for _, c := range columnar {
		fmt.Printf(fmtString, truncate(c.Name, colSize))
	}
	fmt.Println("|")

	// separator
	fmt.Printf("    ")
	for i := 0; i < len(columnar)*actualColSize; i++ {
		if i%actualColSize == 0 {
			fmt.Print("+")
		} else {
			fmt.Print("-")
		}
	}
	fmt.Println("+")

	// column types
	fmt.Printf("    ")
	for _, c := range columnar {
		fmt.Printf(fmtString, truncate(c.Type, colSize))
	}
	fmt.Println("|")

	// separator
	fmt.Printf("    ")
	for i := 0; i < len(columnar)*actualColSize; i++ {
		if i%actualColSize == 0 {
			fmt.Print("+")
		} else {
			fmt.Print("-")
		}
	}
	fmt.Println("+")

	// data
	for i := 0; i < len(columnar[0].Data); i++ {
		fmt.Printf("    ")
		for _, c := range columnar {
			fmt.Printf(fmtString, truncate(c.Data[i], colSize))
		}
		fmt.Println("|")
	}

	// separator
	fmt.Printf("    ")
	for i := 0; i < len(columnar)*actualColSize; i++ {
		if i%actualColSize == 0 {
			fmt.Print("+")
		} else {
			fmt.Print("-")
		}
	}
	fmt.Println("+")
}
