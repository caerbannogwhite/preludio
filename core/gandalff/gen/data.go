package main

import (
	"fmt"
	"go/ast"
	"typesys"
)

type MakeOperationType func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr

type OperationApplyTo struct {
	SeriesType    string
	InnerType     typesys.BaseType
	MakeOperation MakeOperationType
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
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] * %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesType: "GDLSeriesFloat64",
						InnerType:  typesys.Float64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = float64(%s.data[%s]) * %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
				},
			},
		},
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
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] * float64(%s.data[%s])", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesType: "GDLSeriesFloat64",
						InnerType:  typesys.Float64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] * %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
				},
			},
		},
	},
}
