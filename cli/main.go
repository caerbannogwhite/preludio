package main

import (
	"bufio"
	"compiler"
	"fmt"
	"os"
	"preludio"
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
		be := new(preludio.ByteEater).
			InitVM().
			SetPrintWarning(args.Warnings).
			SetDebugLevel(args.DebugLevel)

		bytecode := compiler.CompileFile(args.InputPath)
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

	be := new(preludio.ByteEater).
		InitVM().
		SetPrintWarning(args.Warnings).
		SetDebugLevel(args.DebugLevel)

	codeEditor := NewCodeEditor().
		SetPreludioByteEater(*be)

	if _, err := tea.NewProgram(codeEditor).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

func LaunchRepl(args CliArgs) {

	fmt.Println("Welcome to the Preludio REPL!")
	fmt.Println("Version:", VERSION)

	be := new(preludio.ByteEater).
		InitVM().
		SetPrintWarning(args.Warnings).
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
			bytecode := compiler.CompileSource(code)
			be.RunBytecode(bytecode)

			res := be.GetResult()
			for _, log := range res.Log {
				if log.LogType == preludio.LOG_DEBUG {
					if int(log.Level) < args.DebugLevel {
						fmt.Println(log.Message)
					}
				} else {
					fmt.Println(log.Message)
				}
			}

			fmt.Println()
			fmt.Println(res.Data)

			code = ""
			readerStart = true
		} else {
			readerStart = false
		}

		code += line + "\n"
	}
}
