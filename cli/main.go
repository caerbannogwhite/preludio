package main

import (
	"compiler"
	"fmt"
	"os"
	"preludio"

	"github.com/alexflint/go-arg"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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

const BANNER_TEXT = `
PRELUDIO
`

const KEY_MAP = `
	KEY MAP
	
	ctrl+c		quit
	ctrl+s		save
	ctrl+r		show/hide key map
`

// DEFAULT SETTINGS
const (
	INIT_ROWS_NUM    = 100
	VISIBLE_ROWS_NUM = 20
	ROW_WIDTH        = 60
)

// COLOR SCHEME

var (
	BANNER_STYLE = lipgloss.NewStyle().
			Background(lipgloss.Color("##61edd0")).
			Bold(true).
			PaddingLeft(4).
			Width(ROW_WIDTH)

	EDITOR_ROW_PROMPT_STYLE = lipgloss.NewStyle().
				Foreground(lipgloss.Color("240"))

	focusedStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("#7D56F4"))

	blurredStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	cursorStyle  = focusedStyle.Copy()
	noStyle      = lipgloss.NewStyle()
	helpStyle    = blurredStyle.Copy()
)

type CodeEditor struct {
	saved           bool
	showKeyMap      bool
	currentRow      int
	visibleRows     int
	firstVisibleRow int
	lastRowInUse    int
	rows            []textinput.Model
	footerMessage   string
}

func NewCodeEditor() CodeEditor {

	rows := make([]textinput.Model, INIT_ROWS_NUM)
	for idx := range rows {
		rows[idx] = NewCodeEditorRow(idx)
	}

	editor := CodeEditor{
		saved:           false,
		showKeyMap:      true,
		currentRow:      0,
		visibleRows:     VISIBLE_ROWS_NUM,
		firstVisibleRow: 0,
		rows:            rows,
		footerMessage:   KEY_MAP,
	}

	return editor
}

func NewCodeEditorRow(idx int) textinput.Model {
	row := textinput.New()

	row.Placeholder = ""
	row.CharLimit = 0
	row.Width = ROW_WIDTH

	row.CursorStyle = cursorStyle

	row.Prompt = fmt.Sprintf(" %4d \t", idx+1)
	row.PromptStyle = EDITOR_ROW_PROMPT_STYLE

	return row
}

func (editor CodeEditor) Init() tea.Cmd {

	cmds := make([]tea.Cmd, 0)

	// Focus on first row
	cmds = append(cmds, editor.rows[0].Focus())
	editor.rows[0].PromptStyle = focusedStyle
	editor.rows[0].TextStyle = focusedStyle

	// Blicking cursor
	cmds = append(cmds, textinput.Blink)

	return tea.Batch(cmds...)
}

func (editor CodeEditor) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		// case "esc":

		case "up":
			// move visible region up
			if editor.currentRow == (editor.firstVisibleRow+1) && editor.firstVisibleRow > 0 {
				editor.firstVisibleRow--
			}

			editor.currentRow--
			if editor.currentRow < 0 {
				editor.currentRow = 0
			}

			return editor.updateFocus()

		case "down":

			// move visible region down
			if tmp := (editor.firstVisibleRow + editor.visibleRows - 2); editor.currentRow == tmp && tmp < len(editor.rows)-1 {
				editor.firstVisibleRow++
			}

			editor.currentRow++
			if editor.currentRow == len(editor.rows) {
				editor.currentRow = len(editor.rows) - 1
			}

			return editor.updateFocus()

		// case "left":

		// case "right":

		case "backspace":
			// if editor.currentRow.Cursor.

		case "enter":

			// add a new empty row
			values := make([]string, len(editor.rows)-editor.currentRow-1)
			for idx, row := range editor.rows[editor.currentRow+1 : len(editor.rows)] {
				values[idx] = row.Value()
			}

			editor.rows = append(editor.rows, NewCodeEditorRow(len(editor.rows)))

			editor.rows[editor.currentRow+1].SetValue("")
			for idx := range editor.rows[editor.currentRow+3 : len(editor.rows)] {
				editor.rows[editor.currentRow+2+idx].SetValue(values[idx])
			}

			// move visible region down
			if editor.currentRow == (editor.firstVisibleRow + editor.visibleRows - 2) {
				editor.firstVisibleRow++
			}

			editor.currentRow++

			return editor.updateFocus()

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

	cmd = editor.updateInputs(msg)

	return editor, cmd
}

func (editor CodeEditor) updateFocus() (tea.Model, tea.Cmd) {
	cmds := make([]tea.Cmd, len(editor.rows))
	for i := 0; i <= len(editor.rows)-1; i++ {
		if i == editor.currentRow {
			// Set focused state
			cmds[i] = editor.rows[i].Focus()
			editor.rows[i].PromptStyle = focusedStyle
			editor.rows[i].TextStyle = focusedStyle
			continue
		}

		// Remove focused state
		editor.rows[i].Blur()
		editor.rows[i].PromptStyle = noStyle
		editor.rows[i].TextStyle = noStyle
	}

	return editor, tea.Batch(cmds...)
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
	out := BANNER_STYLE.Render(BANNER_TEXT)
	out += "\n"

	lastVisibleRow := editor.firstVisibleRow + editor.visibleRows
	for idx := editor.firstVisibleRow; idx < lastVisibleRow && idx < len(editor.rows); idx++ {
		out += editor.rows[idx].View() + "\n"
	}

	out += fmt.Sprintf("\n%s", helpStyle.Render(editor.footerMessage))

	return out
}
