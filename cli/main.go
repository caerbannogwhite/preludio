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

	be := new(preludio.ByteEater).
		InitVM().
		SetPrintWarning(args.Warnings).
		SetDebugLevel(args.DebugLevel)

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

			be.PrintLog()

			code = ""
			readerStart = true
		} else {
			readerStart = false
		}

		code += line + "\n"
	}
}
