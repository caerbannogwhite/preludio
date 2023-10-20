module preludiocore

go 1.21.3

replace bytefeeder => ./bytefeeder

replace gandalff => ../../gandalff

replace preludiometa => ../../preludiometa

require (
	bytefeeder v0.0.0-00010101000000-000000000000
	gandalff v0.0.0-00010101000000-000000000000
	preludiometa v0.0.0-00010101000000-000000000000

)

require (
	github.com/antlr/antlr4/runtime/Go/antlr/v4 v4.0.0-20230305170008-8188dc5388df // indirect
	github.com/caerbannogwhite/preludiometa v0.0.0-20231018224909-5ea6906bfec2 // indirect
	golang.org/x/exp v0.0.0-20220722155223-a9213eeb770e // indirect
)
