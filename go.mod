module preludio

go 1.24

replace preludiocore => ./src

replace bytefeeder => ./src/bytefeeder

require (
	bytefeeder v0.0.0-00010101000000-000000000000
	github.com/alexflint/go-arg v1.4.3
	github.com/caerbannogwhite/aargh v0.1.2
	github.com/charmbracelet/lipgloss v0.9.1
	preludiocore v0.0.0-00010101000000-000000000000
)

require (
	github.com/alexflint/go-scalar v1.1.0 // indirect
	github.com/antlr/antlr4/runtime/Go/antlr/v4 v4.0.0-20230305170008-8188dc5388df // indirect
	github.com/aymanbagabas/go-osc52/v2 v2.0.1 // indirect
	github.com/lucasb-eyer/go-colorful v1.2.0 // indirect
	github.com/mattn/go-isatty v0.0.18 // indirect
	github.com/mattn/go-runewidth v0.0.15 // indirect
	github.com/muesli/reflow v0.3.0 // indirect
	github.com/muesli/termenv v0.15.2 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/tealeg/xlsx v1.0.5 // indirect
	golang.org/x/exp v0.0.0-20220722155223-a9213eeb770e // indirect
	golang.org/x/sys v0.12.0 // indirect
)
