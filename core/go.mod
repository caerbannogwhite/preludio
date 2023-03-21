module preludiocore

go 1.20

replace preludiocompiler => ./compiler

require (
	github.com/go-gota/gota v0.12.0
	preludiocompiler v0.0.0-00010101000000-000000000000
)

require (
	github.com/antlr/antlr4/runtime/Go/antlr/v4 v4.0.0-20230305170008-8188dc5388df // indirect
	golang.org/x/exp v0.0.0-20220722155223-a9213eeb770e // indirect
	golang.org/x/net v0.0.0-20210423184538-5f58ad60dda6 // indirect
	gonum.org/v1/gonum v0.9.1 // indirect
)
