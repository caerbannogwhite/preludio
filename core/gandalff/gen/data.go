package main

import "typesys"

type OperationApplyTo struct {
	SeriesType string
	InnerType  typesys.BaseType
}

type Operation struct {
	OpCode  typesys.OPCODE
	ApplyTo []OperationApplyTo
}

type SeriesFile struct {
	SeriesType string
	InnerType  typesys.BaseType
	Operations map[string]Operation
}

var DATA = map[string]SeriesFile{
	"gdl_series_int32.go": {
		SeriesType: "GDLSeriesInt32",
		InnerType:  typesys.Int32Type,
		Operations: map[string]Operation{
			"Mul": {
				OpCode: typesys.OP_BINARY_MUL,
				ApplyTo: []OperationApplyTo{
					{
						SeriesType: "GDLSeriesInt32",
						InnerType:  typesys.Int32Type,
					},
					{
						SeriesType: "GDLSeriesFloat64",
						InnerType:  typesys.Float64Type,
					},
				},
			}},
	},

	"gdl_series_float64.go": {
		SeriesType: "GDLSeriesFloat64",
		InnerType:  typesys.Float64Type,
		Operations: map[string]Operation{
			"Mul": {
				OpCode: typesys.OP_BINARY_MUL,
				ApplyTo: []OperationApplyTo{
					{
						SeriesType: "GDLSeriesInt32",
						InnerType:  typesys.Int32Type,
					},
					{
						SeriesType: "GDLSeriesFloat64",
						InnerType:  typesys.Float64Type,
					},
				},
			}},
	},
}
