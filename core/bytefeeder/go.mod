module bytefeeder

go 1.20

replace types => ../types

require (
	github.com/antlr/antlr4/runtime/Go/antlr/v4 v4.0.0-20230305170008-8188dc5388df
	types v0.0.0-00010101000000-000000000000
)

require golang.org/x/exp v0.0.0-20220722155223-a9213eeb770e // indirect
