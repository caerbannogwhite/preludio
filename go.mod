module preludio

go 1.20

replace preludiocore => ./core

replace bytefeeder => ./core/bytefeeder

replace typesys => ./core/typesys

replace gandalff => ./core/gandalff

require (
	bytefeeder v0.0.0-00010101000000-000000000000
	github.com/alexflint/go-arg v1.4.3
	github.com/charmbracelet/lipgloss v0.7.1
	preludiocore v0.0.0-00010101000000-000000000000
	typesys v0.0.0-00010101000000-000000000000
)

require (
	gandalff v0.0.0-00010101000000-000000000000 // indirect
	github.com/alexflint/go-scalar v1.1.0 // indirect
	github.com/antlr/antlr4/runtime/Go/antlr/v4 v4.0.0-20230305170008-8188dc5388df // indirect
	github.com/aymanbagabas/go-osc52/v2 v2.0.1 // indirect
	github.com/lucasb-eyer/go-colorful v1.2.0 // indirect
	github.com/mattn/go-isatty v0.0.17 // indirect
	github.com/mattn/go-runewidth v0.0.14 // indirect
	github.com/muesli/reflow v0.3.0 // indirect
	github.com/muesli/termenv v0.15.1 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	golang.org/x/exp v0.0.0-20220722155223-a9213eeb770e // indirect
	golang.org/x/sys v0.6.0 // indirect
)
