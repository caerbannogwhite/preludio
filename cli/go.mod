module cli

go 1.19

replace preludio => ../preludio/

require preludio v0.0.0-00010101000000-000000000000

require github.com/go-gota/gota v0.12.0 // indirect

require (
	github.com/alexflint/go-arg v1.4.3
	github.com/alexflint/go-scalar v1.1.0 // indirect
	golang.org/x/net v0.0.0-20210423184538-5f58ad60dda6 // indirect
	gonum.org/v1/gonum v0.9.1 // indirect
)
