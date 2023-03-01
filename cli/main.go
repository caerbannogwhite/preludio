package main

import (
	"compiler"
	"fmt"
	"os"
	"preludio"

	"github.com/alexflint/go-arg"
	tea "github.com/charmbracelet/bubbletea"
)

type CliArgs struct {
	InputPath  string `arg:"-i, --input" help:"source file input path" default:""`
	DebugLevel int    `arg:"-d, --debug-level" help:"debug level" default:"0"`
	Verbose    bool   `arg:"-v, --verbose" help:"verbosity level" default:"false"`
	Warnings   bool   `arg:"-w, --warnings" help:"print warnings" defaut:"true"`
}

func main() {

	var args CliArgs

	arg.MustParse(&args)

	be := new(preludio.ByteEater).
		InitVM().
		SetPrintWarning(args.Warnings).
		SetDebugLevel(args.DebugLevel)

	if args.InputPath != "" {
		var err error
		var file *os.File

		file, err = os.Open(args.InputPath)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		buff := make([]byte, 0)
		_, err = file.Read(buff)
		if err != nil {
			fmt.Println(err)
			return
		}

		bytecode := compiler.Compile(string(buff))
		if args.Verbose {
			fmt.Println("Bytecode generated")
		}
		be.ReadBytecode(bytecode)

	} else {
		REPL(args)
	}
}

func REPL(args CliArgs) {

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

	// columns := []table.Column{
	// 	{Title: "Function", Width: 15},
	// 	{Title: "Description", Width: 30},
	// }

	// rows := []table.Row{
	// 	// {"derive", ""},
	// 	{"describe", "Describe the current table."},
	// 	{"from", "Set a table to transform."},
	// 	{"exportCSV", "Export the current table to a CSV file."},
	// 	{"importCSV", "Import the current table from a CSV file."},
	// 	{"new", "Create a new table."},
	// 	{"select", "Select a subset of columns."},
	// 	{"sort", "Sort the rows according to a sub set of columns."},
	// 	{"take", "Take a range of rows."},
	// 	{"asBool", "Coerce to boolean"},
	// 	{"asInteger", "Coerce to integer."},
	// 	{"asFloat", "Coerce to float."},
	// 	{"asString", "Convert to string."},
	// }

	// t := table.New(
	// 	table.WithColumns(columns),
	// 	table.WithRows(rows),
	// 	table.WithFocused(true),
	// 	table.WithHeight(7),
	// )

	// s := table.DefaultStyles()
	// s.Header = s.Header.
	// 	BorderStyle(lipgloss.NormalBorder()).
	// 	BorderForeground(lipgloss.Color("240")).
	// 	BorderBottom(true).
	// 	Bold(false)
	// s.Selected = s.Selected.
	// 	Foreground(lipgloss.Color("229")).
	// 	Background(lipgloss.Color("57")).
	// 	Bold(false)
	// t.SetStyles(s)

	// m := model{t}
	// if _, err := tea.NewProgram(m).Run(); err != nil {
	// 	fmt.Println("Error running program:", err)
	// 	os.Exit(1)
	// }
}
