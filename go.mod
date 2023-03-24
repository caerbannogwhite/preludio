module preludio

go 1.20

replace preludiocore => ./core

replace bytefeeder => ./core/bytefeeder

replace types => ./core/types

replace preludiocli => ./cli

require (
	bytefeeder v0.0.0-00010101000000-000000000000
	github.com/alexflint/go-arg v1.4.3
	github.com/charmbracelet/bubbletea v0.23.2
	preludiocli v0.0.0-00010101000000-000000000000
	preludiocore v0.0.0-00010101000000-000000000000
)

require (
	github.com/alexflint/go-scalar v1.1.0 // indirect
	github.com/antlr/antlr4/runtime/Go/antlr/v4 v4.0.0-20230305170008-8188dc5388df // indirect
	github.com/atotto/clipboard v0.1.4 // indirect
	github.com/aymanbagabas/go-osc52/v2 v2.0.1 // indirect
	github.com/charmbracelet/bubbles v0.15.0 // indirect
	github.com/charmbracelet/lipgloss v0.7.1 // indirect
	github.com/containerd/console v1.0.3 // indirect
	github.com/go-gota/gota v0.12.0 // indirect
	github.com/lucasb-eyer/go-colorful v1.2.0 // indirect
	github.com/mattn/go-isatty v0.0.17 // indirect
	github.com/mattn/go-localereader v0.0.1 // indirect
	github.com/mattn/go-runewidth v0.0.14 // indirect
	github.com/muesli/ansi v0.0.0-20211018074035-2e021307bc4b // indirect
	github.com/muesli/cancelreader v0.2.2 // indirect
	github.com/muesli/reflow v0.3.0 // indirect
	github.com/muesli/termenv v0.15.1 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	golang.org/x/exp v0.0.0-20220722155223-a9213eeb770e // indirect
	golang.org/x/net v0.0.0-20210423184538-5f58ad60dda6 // indirect
	golang.org/x/sync v0.1.0 // indirect
	golang.org/x/sys v0.6.0 // indirect
	golang.org/x/term v0.0.0-20210927222741-03fcf44c2211 // indirect
	golang.org/x/text v0.3.7 // indirect
	gonum.org/v1/gonum v0.9.1 // indirect
	types v0.0.0-00010101000000-000000000000 // indirect
)
