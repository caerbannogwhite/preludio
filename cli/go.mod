module preludio-vm

go 1.19

replace preludiovm => ../vm/

require github.com/go-gota/gota v0.12.0 // indirect

require (
	github.com/alexflint/go-arg v1.4.3
	github.com/alexflint/go-scalar v1.1.0 // indirect
	golang.org/x/net v0.0.0-20210423184538-5f58ad60dda6 // indirect
	gonum.org/v1/gonum v0.9.1 // indirect
	preludiovm v1.0.0
)
