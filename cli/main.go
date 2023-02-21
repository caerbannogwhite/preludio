package main

import (
	"compiler"
	"fmt"
	"os"
	"preludio"

	"github.com/alexflint/go-arg"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {

	var args struct {
		InputPath  string `arg:"-i, --input" help:"source file input path" default:""`
		DebugLevel int    `arg:"-d, --debug-level" help:"debug level" default:"0"`
		Verbose    bool   `arg:"-v, --verbose" help:"verbosity level" default:"false"`
		Warnings   bool   `arg:"-w, --warnings" help:"print warnings" defaut:"true"`
	}

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
		REPL()
	}
}

func REPL() {

	codeEditor := NewCodeEditor()

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

const KEY_MAP = `
	KEY MAP
	
	ctrl+c		quit
	ctrl+s		save
	ctrl+r		show/hide key map
`

type CodeEditor struct {
	saved         bool
	showKeyMap    bool
	currentRow    int
	rows          []textinput.Model
	footerMessage string
}

func NewCodeEditor() CodeEditor {

	rows := []textinput.Model{NewCodeEditorRow()}
	// rows[0].CursorStart()

	editor := CodeEditor{
		saved:         false,
		showKeyMap:    true,
		currentRow:    0,
		rows:          rows,
		footerMessage: KEY_MAP,
	}

	return editor
}

func NewCodeEditorRow() textinput.Model {
	row := textinput.New()
	row.Placeholder = ""
	// row.Focus()
	row.CharLimit = 156
	row.Width = 20

	return row
}

func (editor CodeEditor) Init() tea.Cmd {
	return nil
}

func (editor CodeEditor) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		// case "esc":
		// 	if m.table.Focused() {
		// 		m.table.Blur()
		// 	} else {
		// 		m.table.Focus()
		// 	}

		case "down":
			editor.rows[editor.currentRow].Reset()
			editor.currentRow--
			if editor.currentRow < 0 {
				editor.currentRow = len(editor.rows) - 1
			}
			editor.rows[editor.currentRow].Focus()

		case "up":
			editor.rows[editor.currentRow].Reset()
			editor.currentRow++
			if editor.currentRow == len(editor.rows) {
				editor.currentRow = 0
			}
			editor.rows[editor.currentRow].Focus()

		case "left":

		case "right":

		case "enter":
			editor.rows[editor.currentRow].Reset()
			editor.rows = append(editor.rows, NewCodeEditorRow())
			editor.currentRow++
			editor.rows[editor.currentRow].Focus()

		// SHOW/HIDE KEY MAP
		case "ctrl+r":
			if editor.showKeyMap {
				editor.footerMessage = ""
			} else {
				editor.footerMessage = KEY_MAP
			}
			editor.showKeyMap = !editor.showKeyMap

		// SAVE
		case "ctrl+s":
			editor.saved = true

		// QUIT
		case "ctrl+c":
			// if editor.saved {
			return editor, tea.Quit
			// } else {
			// }
		default:
		}
	}

	editor.rows[editor.currentRow].Focus()

	cmd = editor.updateInputs(msg)

	// m.table, cmd = m.table.Update(msg)
	return editor, cmd
}

func (editor CodeEditor) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(editor.rows))

	// Only text inputs with Focus() set will respond, so it's safe to simply
	// update all of them here without any further logic.
	for i := range editor.rows {
		editor.rows[i], cmds[i] = editor.rows[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func (editor CodeEditor) View() string {
	out := "	Welcome To Preludio!\n\n"

	for idx, row := range editor.rows {
		out += fmt.Sprintf("%3d | %s\n", idx, row.Value())
	}

	out += fmt.Sprintf("\n%s", editor.footerMessage)

	return out
}

// type FunctionSelector struct {
// 	currentPattern    string
// 	selectedFunctions []string
// }

// var baseStyle = lipgloss.NewStyle().
// 	BorderStyle(lipgloss.NormalBorder()).
// 	BorderForeground(lipgloss.Color("240"))

// type model struct {
// 	table table.Model
// }

// func (m model) Init() tea.Cmd {
// 	return nil
// }

// func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
// 	var cmd tea.Cmd
// 	switch msg := msg.(type) {
// 	case tea.KeyMsg:
// 		switch msg.String() {
// 		case "esc":
// 			if m.table.Focused() {
// 				m.table.Blur()
// 			} else {
// 				m.table.Focus()
// 			}
// 		case "ctrl+c":
// 			return m, tea.Quit
// 		case "enter":
// 			return m, tea.Batch(
// 				tea.Printf("%s", m.table.SelectedRow()[0]),
// 			)
// 		default:
// 		}
// 	}
// 	m.table, cmd = m.table.Update(msg)
// 	return m, cmd
// }

// func (m model) View() string {
// 	return baseStyle.Render(m.table.View()) + "\n"
// }
