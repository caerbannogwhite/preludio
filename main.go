package main

import (
	"bufio"
	"fmt"
	"os"
	"preludiocli"
	"preludiocompiler"
	"preludiocore"
	"strconv"
	"strings"

	"github.com/alexflint/go-arg"
	tea "github.com/charmbracelet/bubbletea"
)

const VERSION = "0.1.0-alpha"

type CliArgs struct {
	SourceCode string `arg:"-s, --source" help:"source code to execute" default:""`
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
			SetParamPrintWarning(args.Warnings).
			SetParamDebugLevel(args.DebugLevel)

		bytecode := preludiocompiler.CompileFile(args.InputPath)
		if args.Verbose {
			fmt.Println("Bytecode generated")
		}

		be.RunBytecode(bytecode)

	} else if args.SourceCode != "" {
		be := new(preludiocore.ByteEater).
			InitVM().
			SetParamPrintWarning(args.Warnings).
			SetParamDebugLevel(args.DebugLevel)

		bytecode := preludiocompiler.CompileSource(args.SourceCode)
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
		SetParamPrintWarning(args.Warnings).
		SetParamDebugLevel(args.DebugLevel)

	codeEditor := preludiocli.NewCodeEditor().
		SetPreludioByteEater(*be)

	if _, err := tea.NewProgram(codeEditor).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

func LaunchRepl(args CliArgs) {

	var outputColumnSize = 10

	fmt.Println("Welcome to the Preludio REPL!")
	fmt.Println("Version:", VERSION)

	be := new(preludiocore.ByteEater).
		InitVM().
		SetParamPrintWarning(args.Warnings).
		SetParamFullOutput(false).
		SetParamDebugLevel(args.DebugLevel).
		SetParamVerbose(args.Verbose)

	if args.Verbose {
		fmt.Printf("\nPreludio VM initialized\n")
		fmt.Printf("%15s %t\n", "Print warnings:", be.GetParamPrintWarning())
		fmt.Printf("%15s %t\n", "Full output:", be.GetParamFullOutput())
		fmt.Printf("%15s %d\n", "Debug level:", be.GetParamDebugLevel())
		fmt.Printf("%15s %t\n", "Verbose:", be.GetParamVerbose())
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

		// look for magic commands
		// ignore the rest
		if strings.HasPrefix(line, "%") {
			spt := strings.Split(strings.Trim(line, "\t\n\r "), " ")
			switch spt[0] {
			case "%setenv":
				if len(spt) != 3 {
					fmt.Println("Usage: %setenv <key> <value>")
					continue
				}

				switch spt[1] {
				case "ENV_WARNINGS":
					if spt[2] == "true" {
						be.SetParamPrintWarning(true)
					} else if spt[2] == "false" {
						be.SetParamPrintWarning(false)
					}
					fmt.Printf("Print warnings set to \"%t\"", be.GetParamPrintWarning())
				case "ENV_DEBUG_LEVEL":
					l, err := strconv.Atoi(spt[2])
					if err != nil {
						fmt.Println("Error parsing debug level:", err)
						continue
					}
					be.SetParamDebugLevel(l)
					fmt.Println("Debug level set to", be.GetParamDebugLevel())
				case "ENV_VERBOSE":
					if spt[2] == "true" {
						be.SetParamVerbose(true)
					} else if spt[2] == "false" {
						be.SetParamVerbose(false)
					}
					fmt.Printf("Verbose set to \"%t\"", be.GetParamVerbose())
				case "ENV_FULL_OUTPUT":
					if spt[2] == "true" {
						be.SetParamFullOutput(true)
					} else if spt[2] == "false" {
						be.SetParamFullOutput(false)
					}
					fmt.Printf("Full output set to \"%t\"", be.GetParamFullOutput())
				case "ENV_OUTPUT_COLUMN_SIZE":
					l, err := strconv.Atoi(spt[2])
					if err != nil {
						fmt.Println("Error parsing output column size:", err)
						continue
					}
					outputColumnSize = l
					fmt.Println("Output column size set to", outputColumnSize)
				default:
					fmt.Println("Unknown environment variable:", spt[1])
				}

			// case "%getenv":
			// 	if len(spt) != 2 {
			// 		fmt.Println("Usage: %getenv <key>")
			// 		continue
			// 	}
			// 	fmt.Println(be.GetEnv(spt[1]))

			default:
				fmt.Println("Unknown magic command:", spt[0])
			}
			continue
		}

		line = strings.TrimSpace(line)

		if line == "" {
			bytecode := preludiocompiler.CompileSource(code)
			be.RunBytecode(bytecode)

			res := be.GetOutput()
			for _, log := range res.Log {
				if log.LogType == preludiocore.LOG_DEBUG {
					if int(log.Level) < be.GetParamDebugLevel() {
						fmt.Println(log.Message)
					}
				} else {
					fmt.Println(log.Message)
				}
			}

			for _, c := range res.Data {
				prettyPrint(outputColumnSize, c)
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

	if len(columnar) == 0 {
		return
	}

	actualColSize := colSize + 3
	fmtString := fmt.Sprintf("| %%%ds ", colSize)

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
	// check if there are any column names
	colNames := false
	for _, c := range columnar {
		if c.Name != "" {
			colNames = true
			break
		}
	}

	// only print column names if there are any
	if colNames {
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
	}

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
