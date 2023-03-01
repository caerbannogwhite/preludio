package main

import (
	"compiler"
	"fmt"
	"os"
	"preludio"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const BANNER_TEXT = `
PRELUDIO
`

const KEY_MAP = `
	KEY MAP
	
	ctrl+c		quit				•	ctrl+r		run the code
	ctrl+s		save
	ctrl+h		show/hide key map
`

// DEFAULT SETTINGS
const (
	INIT_ROWS_NUM    = 100
	VISIBLE_ROWS_NUM = 20
	ROW_WIDTH        = 60

	EDITOR_LINE_ENDING = "\n"
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

	EDITOR_ROW_FOCUSED_STYLE = lipgloss.NewStyle().
					Background(lipgloss.Color("#7D56F4"))

	BLURRED_STYLE = lipgloss.NewStyle().
			Foreground(lipgloss.Color("240"))

	EDITOR_ERROR_MESSAGE_STYLE = lipgloss.NewStyle().
					Foreground(lipgloss.Color("211"))

	NO_STYLE = lipgloss.NewStyle()
)

type CodeEditor struct {
	saved           bool
	showKeyMap      bool
	currentRowIdx   int
	visibleRows     int
	firstVisibleRow int
	lastRowInUse    int
	rows            []textinput.Model
	errorMessage    string
	footerMessage   string
	byteEater       preludio.ByteEater
}

func NewCodeEditor() CodeEditor {

	rows := make([]textinput.Model, INIT_ROWS_NUM)
	for idx := range rows {
		rows[idx] = NewCodeEditorRow(idx)
	}

	editor := CodeEditor{
		saved:           false,
		showKeyMap:      true,
		currentRowIdx:   0,
		visibleRows:     VISIBLE_ROWS_NUM,
		firstVisibleRow: 0,
		rows:            rows,
		errorMessage:    "",
		footerMessage:   KEY_MAP,
	}

	return editor
}

func (editor CodeEditor) SetPreludioByteEater(be preludio.ByteEater) CodeEditor {
	editor.byteEater = be
	return editor
}

func (editor *CodeEditor) MoveCurrentRowUp() {
	// move visible region up
	if editor.currentRowIdx == (editor.firstVisibleRow+1) && editor.firstVisibleRow > 0 {
		editor.firstVisibleRow--
	}

	editor.currentRowIdx--
	if editor.currentRowIdx < 0 {
		editor.currentRowIdx = 0
	}
}

func (editor *CodeEditor) MoveCurrentRowDown() {
	// move visible region down
	if editor.currentRowIdx == (editor.firstVisibleRow + editor.visibleRows - 2) {
		editor.firstVisibleRow++
	}

	editor.currentRowIdx++
	if editor.currentRowIdx == len(editor.rows) {
		editor.currentRowIdx = len(editor.rows) - 1
	}
}

func (editor *CodeEditor) GetFirstInUseRoxIndex() int {
	idx := 0
	for editor.rows[idx].Value() == "" {
		idx++
	}
	return idx
}

func (editor *CodeEditor) GetLastInUseRowIndex() int {
	idx := len(editor.rows) - 1
	for idx > 0 && editor.rows[idx].Value() == "" {
		idx--
	}
	return idx
}

func (editor *CodeEditor) GetCurrentEditorCode() string {
	code := ""
	for _, row := range editor.rows[editor.GetFirstInUseRoxIndex() : editor.GetLastInUseRowIndex()+1] {
		code += row.Value() + EDITOR_LINE_ENDING
	}
	return code
}

func NewCodeEditorRow(idx int) textinput.Model {
	row := textinput.New()

	row.Placeholder = ""
	row.CharLimit = 0
	row.Width = ROW_WIDTH

	row.CursorStyle = EDITOR_ROW_FOCUSED_STYLE

	row.Prompt = fmt.Sprintf(" %4d ", idx+1)
	row.PromptStyle = EDITOR_ROW_PROMPT_STYLE

	return row
}

func (editor CodeEditor) Init() tea.Cmd {

	cmds := make([]tea.Cmd, 0)

	// Focus on first row
	cmds = append(cmds, editor.rows[0].Focus())
	editor.rows[0].PromptStyle = EDITOR_ROW_FOCUSED_STYLE
	editor.rows[0].TextStyle = EDITOR_ROW_FOCUSED_STYLE
	editor.rows[0].Cursor.TextStyle = EDITOR_ROW_FOCUSED_STYLE

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
			editor.MoveCurrentRowUp()
			return editor.updateFocus()

		case "down":
			editor.MoveCurrentRowDown()
			return editor.updateFocus()

		case "left":
			if editor.rows[editor.currentRowIdx].Position() == 0 {
				editor.MoveCurrentRowUp()
				return editor.updateFocus()
			}

		case "right":
			row := editor.rows[editor.currentRowIdx]
			if (row.Value() == "" && row.Position() == 0) || row.Position() == len(row.Value()) {
				editor.MoveCurrentRowDown()
				return editor.updateFocus()
			}

		case "backspace":
			editor.saved = false

			if editor.rows[editor.currentRowIdx].Position() == 0 && editor.currentRowIdx > 0 {
				lastRow := editor.GetLastInUseRowIndex()
				if lastRow >= editor.currentRowIdx {

					// Copy all the values from the current row (where the focus is) to the last row
					// in use (it can be the same row, so +1)
					// the +2 at the end: empty the last row with the for loop
					values := make([]string, lastRow-editor.currentRowIdx+2)
					for idx, row := range editor.rows[editor.currentRowIdx : lastRow+2] {
						values[idx] = row.Value()
					}

					if editor.rows[editor.currentRowIdx-1].Value() != "" {
						val := editor.rows[editor.currentRowIdx-1].Value()
						editor.rows[editor.currentRowIdx-1].SetValue(val + values[0])

						for idx, value := range values[1:] {
							editor.rows[editor.currentRowIdx+idx].SetValue(value)
						}
					} else {
						for idx, value := range values {
							editor.rows[editor.currentRowIdx+idx-1].SetValue(value)
						}

						editor.rows[editor.currentRowIdx-1].SetCursor(0)
					}
				}

				editor.MoveCurrentRowUp()

				return editor.updateFocus()
			}

		case "enter":
			editor.saved = false

			// add a new empty row
			lastRow := editor.GetLastInUseRowIndex() + 1
			if lastRow > editor.currentRowIdx {
				values := make([]string, lastRow-editor.currentRowIdx-1)
				for idx, row := range editor.rows[editor.currentRowIdx+1 : lastRow] {
					values[idx] = row.Value()
				}

				// deal with the current row
				currentRow := editor.rows[editor.currentRowIdx]

				currentValue := currentRow.Value()
				editor.rows[editor.currentRowIdx].SetValue(currentValue[0:currentRow.Position()])
				editor.rows[editor.currentRowIdx+1].SetValue(currentValue[currentRow.Position():])

				editor.rows = append(editor.rows, NewCodeEditorRow(len(editor.rows)))

				for idx, value := range values {
					editor.rows[editor.currentRowIdx+2+idx].SetValue(value)
				}

				// Put the cursor at the beginning
				editor.rows[editor.currentRowIdx+1].SetCursor(0)
			}

			editor.MoveCurrentRowDown()

			return editor.updateFocus()

		// RUN CODE
		case "ctrl+r":
			code := editor.GetCurrentEditorCode()

			bytecode := compiler.Compile(code)
			editor.errorMessage = "Bytecode generated"
			editor.byteEater.ReadBytecode(bytecode)

		// SHOW/HIDE KEY MAP
		case "ctrl+h":
			if editor.showKeyMap {
				editor.footerMessage = ""
			} else {
				editor.footerMessage = KEY_MAP
			}
			editor.showKeyMap = !editor.showKeyMap

		// SAVE
		case "ctrl+s":
			code := editor.GetCurrentEditorCode()

			err := os.WriteFile(".tmp.prql", []byte(code), 0644)
			if err != nil {
				editor.errorMessage = err.Error()
			}

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
		if i == editor.currentRowIdx {
			// Set focused state
			cmds[i] = editor.rows[i].Focus()
			editor.rows[i].PromptStyle = EDITOR_ROW_FOCUSED_STYLE
			editor.rows[i].TextStyle = EDITOR_ROW_FOCUSED_STYLE
			editor.rows[i].Cursor.TextStyle = EDITOR_ROW_FOCUSED_STYLE
			continue
		}

		// Remove focused state
		editor.rows[i].Blur()
		editor.rows[i].PromptStyle = EDITOR_ROW_PROMPT_STYLE
		editor.rows[i].TextStyle = NO_STYLE
		editor.rows[i].Cursor.TextStyle = NO_STYLE
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

	out += fmt.Sprintf("\n%s", EDITOR_ERROR_MESSAGE_STYLE.Render(editor.errorMessage))

	out += fmt.Sprintf("\n%s", BLURRED_STYLE.Render(editor.footerMessage))

	return out
}
