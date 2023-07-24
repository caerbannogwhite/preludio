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
		SeriesType: "SeriesInt32",
		InnerType:  typesys.Int32Type,
		Operations: map[string]Operation{
			"Mul": {
				OpCode: typesys.OP_BINARY_MUL,
				ApplyTo: []OperationApplyTo{
					{
						SeriesType: "SeriesInt32",
						InnerType:  typesys.Int32Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] * %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesType: "SeriesInt64",
						InnerType:  typesys.Int64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = int64(%s.data[%s]) * %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesType: "SeriesFloat64",
						InnerType:  typesys.Float64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = float64(%s.data[%s]) * %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
				},
			},

			"Div": {
				OpCode: typesys.OP_BINARY_DIV,
				ApplyTo: []OperationApplyTo{
					{
						SeriesType: "SeriesInt32",
						InnerType:  typesys.Int32Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] / %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesType: "SeriesInt64",
						InnerType:  typesys.Int64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = int64(%s.data[%s]) / %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesType: "SeriesFloat64",
						InnerType:  typesys.Float64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = float64(%s.data[%s]) / %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
				},
			},

			"Mod": {
				OpCode: typesys.OP_BINARY_MOD,
				ApplyTo: []OperationApplyTo{
					{
						SeriesType: "SeriesInt32",
						InnerType:  typesys.Int32Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] %% %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesType: "SeriesInt64",
						InnerType:  typesys.Int64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = int64(%s.data[%s]) %% %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesType: "SeriesFloat64",
						InnerType:  typesys.Float64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = int64(%s.data[%s]) %% int64(%s.data[%s])", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
				},
			},

			"Add": {
				OpCode: typesys.OP_BINARY_ADD,
				ApplyTo: []OperationApplyTo{
					{
						SeriesType: "SeriesInt32",
						InnerType:  typesys.Int32Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] + %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesType: "SeriesInt64",
						InnerType:  typesys.Int64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = int64(%s.data[%s]) + %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesType: "SeriesFloat64",
						InnerType:  typesys.Float64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = float64(%s.data[%s]) + %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesType: "SeriesString",
						InnerType:  typesys.StringType,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.pool.Put(intToString(int64(%s.data[%s])) + *%s.data[%s])", res, resIndex, op2, op1, op1Index, op2, op2Index)}
						},
					},
				},
			},

			"Sub": {
				OpCode: typesys.OP_BINARY_SUB,
				ApplyTo: []OperationApplyTo{
					{
						SeriesType: "SeriesInt32",
						InnerType:  typesys.Int32Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] - %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesType: "SeriesInt64",
						InnerType:  typesys.Int64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = int64(%s.data[%s]) - %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesType: "SeriesFloat64",
						InnerType:  typesys.Float64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = float64(%s.data[%s]) - %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
				},
			},

			"Eq": {
				OpCode: typesys.OP_BINARY_EQ,
				ApplyTo: []OperationApplyTo{
					{
						SeriesType: "SeriesInt32",
						InnerType:  typesys.Int32Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] == %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesType: "SeriesInt64",
						InnerType:  typesys.Int64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = int64(%s.data[%s]) == %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesType: "SeriesFloat64",
						InnerType:  typesys.Float64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = float64(%s.data[%s]) == %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
				},
			},

			"Ne": {
				OpCode: typesys.OP_BINARY_NE,
				ApplyTo: []OperationApplyTo{
					{
						SeriesType: "SeriesInt32",
						InnerType:  typesys.Int32Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] != %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesType: "SeriesInt64",
						InnerType:  typesys.Int64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = int64(%s.data[%s]) != %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesType: "SeriesFloat64",
						InnerType:  typesys.Float64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = float64(%s.data[%s]) != %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
				},
			},

			"Lt": {
				OpCode: typesys.OP_BINARY_LT,
				ApplyTo: []OperationApplyTo{
					{
						SeriesType: "SeriesInt32",
						InnerType:  typesys.Int32Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] < %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesType: "SeriesInt64",
						InnerType:  typesys.Int64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = int64(%s.data[%s]) < %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesType: "SeriesFloat64",
						InnerType:  typesys.Float64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = float64(%s.data[%s]) < %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
				},
			},

			"Le": {
				OpCode: typesys.OP_BINARY_LE,
				ApplyTo: []OperationApplyTo{
					{
						SeriesType: "SeriesInt32",
						InnerType:  typesys.Int32Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] <= %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesType: "SeriesInt64",
						InnerType:  typesys.Int64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = int64(%s.data[%s]) <= %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesType: "SeriesFloat64",
						InnerType:  typesys.Float64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = float64(%s.data[%s]) <= %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
				},
			},

			"Gt": {
				OpCode: typesys.OP_BINARY_GT,
				ApplyTo: []OperationApplyTo{
					{
						SeriesType: "SeriesInt32",
						InnerType:  typesys.Int32Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] > %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesType: "SeriesInt64",
						InnerType:  typesys.Int64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = int64(%s.data[%s]) > %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesType: "SeriesFloat64",
						InnerType:  typesys.Float64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = float64(%s.data[%s]) > %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
				},
			},

			"Ge": {
				OpCode: typesys.OP_BINARY_GE,
				ApplyTo: []OperationApplyTo{
					{
						SeriesType: "SeriesInt32",
						InnerType:  typesys.Int32Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] >= %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesType: "SeriesInt64",
						InnerType:  typesys.Int64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = int64(%s.data[%s]) >= %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesType: "SeriesFloat64",
						InnerType:  typesys.Float64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = float64(%s.data[%s]) >= %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
				},
			},
		},
	},

	"gdl_series_int64.go": {
		SeriesType: "SeriesInt64",
		InnerType:  typesys.Int64Type,
		Operations: map[string]Operation{
			"Mul": {
				OpCode: typesys.OP_BINARY_MUL,
				ApplyTo: []OperationApplyTo{
					{
						SeriesType: "SeriesInt32",
						InnerType:  typesys.Int32Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] * int64(%s.data[%s])", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesType: "SeriesInt64",
						InnerType:  typesys.Int64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] * %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesType: "SeriesFloat64",
						InnerType:  typesys.Float64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = float64(%s.data[%s]) * %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
				},
			},

			"Div": {
				OpCode: typesys.OP_BINARY_DIV,
				ApplyTo: []OperationApplyTo{
					{
						SeriesType: "SeriesInt32",
						InnerType:  typesys.Int32Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] / int64(%s.data[%s])", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesType: "SeriesInt64",
						InnerType:  typesys.Int64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] / %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesType: "SeriesFloat64",
						InnerType:  typesys.Float64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = float64(%s.data[%s]) / %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
				},
			},

			"Mod": {
				OpCode: typesys.OP_BINARY_MOD,
				ApplyTo: []OperationApplyTo{
					{
						SeriesType: "SeriesInt32",
						InnerType:  typesys.Int32Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] %% int64(%s.data[%s])", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesType: "SeriesInt64",
						InnerType:  typesys.Int64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] %% %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesType: "SeriesFloat64",
						InnerType:  typesys.Float64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] %% int64(%s.data[%s])", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
				},
			},

			"Add": {
				OpCode: typesys.OP_BINARY_ADD,
				ApplyTo: []OperationApplyTo{
					{
						SeriesType: "SeriesInt32",
						InnerType:  typesys.Int32Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] + int64(%s.data[%s])", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesType: "SeriesInt64",
						InnerType:  typesys.Int64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] + %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesType: "SeriesFloat64",
						InnerType:  typesys.Float64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = float64(%s.data[%s]) + %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesType: "SeriesString",
						InnerType:  typesys.StringType,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.pool.Put(intToString(%s.data[%s]) + *%s.data[%s])", res, resIndex, op2, op1, op1Index, op2, op2Index)}
						},
					},
				},
			},

			"Sub": {
				OpCode: typesys.OP_BINARY_SUB,
				ApplyTo: []OperationApplyTo{
					{
						SeriesType: "SeriesInt32",
						InnerType:  typesys.Int32Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] - int64(%s.data[%s])", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesType: "SeriesInt64",
						InnerType:  typesys.Int64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] - %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesType: "SeriesFloat64",
						InnerType:  typesys.Float64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = float64(%s.data[%s]) - %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
				},
			},

			"Eq": {
				OpCode: typesys.OP_BINARY_EQ,
				ApplyTo: []OperationApplyTo{
					{
						SeriesType: "SeriesInt32",
						InnerType:  typesys.Int32Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] == int64(%s.data[%s])", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesType: "SeriesInt64",
						InnerType:  typesys.Int64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] == %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesType: "SeriesFloat64",
						InnerType:  typesys.Float64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = float64(%s.data[%s]) == %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
				},
			},

			"Ne": {
				OpCode: typesys.OP_BINARY_NE,
				ApplyTo: []OperationApplyTo{
					{
						SeriesType: "SeriesInt32",
						InnerType:  typesys.Int32Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] != int64(%s.data[%s])", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesType: "SeriesInt64",
						InnerType:  typesys.Int64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] != %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesType: "SeriesFloat64",
						InnerType:  typesys.Float64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = float64(%s.data[%s]) != %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
				},
			},

			"Lt": {
				OpCode: typesys.OP_BINARY_LT,
				ApplyTo: []OperationApplyTo{
					{
						SeriesType: "SeriesInt32",
						InnerType:  typesys.Int32Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = %s.data[%s] < int64(%s.data[%s])`, res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesType: "SeriesInt64",
						InnerType:  typesys.Int64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = %s.data[%s] < %s.data[%s]`, res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesType: "SeriesFloat64",
						InnerType:  typesys.Float64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = float64(%s.data[%s]) < %s.data[%s]`, res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
				},
			},

			"Le": {
				OpCode: typesys.OP_BINARY_LE,
				ApplyTo: []OperationApplyTo{
					{
						SeriesType: "SeriesInt32",
						InnerType:  typesys.Int32Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = %s.data[%s] <= int64(%s.data[%s])`, res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesType: "SeriesInt64",
						InnerType:  typesys.Int64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = %s.data[%s] <= %s.data[%s]`, res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesType: "SeriesFloat64",
						InnerType:  typesys.Float64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = float64(%s.data[%s]) <= %s.data[%s]`, res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
				},
			},

			"Gt": {
				OpCode: typesys.OP_BINARY_GT,
				ApplyTo: []OperationApplyTo{
					{
						SeriesType: "SeriesInt32",
						InnerType:  typesys.Int32Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = %s.data[%s] > int64(%s.data[%s])`, res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesType: "SeriesInt64",
						InnerType:  typesys.Int64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = %s.data[%s] > %s.data[%s]`, res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesType: "SeriesFloat64",
						InnerType:  typesys.Float64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = float64(%s.data[%s]) > %s.data[%s]`, res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
				},
			},

			"Ge": {
				OpCode: typesys.OP_BINARY_GE,
				ApplyTo: []OperationApplyTo{
					{
						SeriesType: "SeriesInt32",
						InnerType:  typesys.Int32Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = %s.data[%s] >= int64(%s.data[%s])`, res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesType: "SeriesInt64",
						InnerType:  typesys.Int64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = %s.data[%s] >= %s.data[%s]`, res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesType: "SeriesFloat64",
						InnerType:  typesys.Float64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = float64(%s.data[%s]) >= %s.data[%s]`, res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
				},
			},
		},
	},

	"gdl_series_float64.go": {
		SeriesType: "SeriesFloat64",
		InnerType:  typesys.Float64Type,
		Operations: map[string]Operation{
			"Mul": {
				OpCode: typesys.OP_BINARY_MUL,
				ApplyTo: []OperationApplyTo{
					{
						SeriesType: "SeriesInt32",
						InnerType:  typesys.Int32Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] * float64(%s.data[%s])", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesType: "SeriesInt64",
						InnerType:  typesys.Int64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] * float64(%s.data[%s])", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesType: "SeriesFloat64",
						InnerType:  typesys.Float64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] * %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
				},
			},

			"Div": {
				OpCode: typesys.OP_BINARY_DIV,
				ApplyTo: []OperationApplyTo{
					{
						SeriesType: "SeriesInt32",
						InnerType:  typesys.Int32Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] / float64(%s.data[%s])", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesType: "SeriesInt64",
						InnerType:  typesys.Int64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] / float64(%s.data[%s])", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesType: "SeriesFloat64",
						InnerType:  typesys.Float64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] / %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
				},
			},

			"Mod": {
				OpCode: typesys.OP_BINARY_MOD,
				ApplyTo: []OperationApplyTo{
					{
						SeriesType: "SeriesInt32",
						InnerType:  typesys.Int32Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = int64(%s.data[%s]) %% int64(%s.data[%s])", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesType: "SeriesInt64",
						InnerType:  typesys.Int64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = int64(%s.data[%s]) %% %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesType: "SeriesFloat64",
						InnerType:  typesys.Float64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = int64(%s.data[%s]) %% int64(%s.data[%s])", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
				},
			},

			"Add": {
				OpCode: typesys.OP_BINARY_ADD,
				ApplyTo: []OperationApplyTo{
					{
						SeriesType: "SeriesInt32",
						InnerType:  typesys.Int32Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] + float64(%s.data[%s])", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesType: "SeriesInt64",
						InnerType:  typesys.Int64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] + float64(%s.data[%s])", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesType: "SeriesFloat64",
						InnerType:  typesys.Float64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] + %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesType: "SeriesString",
						InnerType:  typesys.StringType,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.pool.Put(floatToString(%s.data[%s]) + *%s.data[%s])", res, resIndex, op2, op1, op1Index, op2, op2Index)}
						},
					},
				},
			},

			"Sub": {
				OpCode: typesys.OP_BINARY_SUB,
				ApplyTo: []OperationApplyTo{
					{
						SeriesType: "SeriesInt32",
						InnerType:  typesys.Int32Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] - float64(%s.data[%s])", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesType: "SeriesInt64",
						InnerType:  typesys.Int64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] - float64(%s.data[%s])", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesType: "SeriesFloat64",
						InnerType:  typesys.Float64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] - %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
				},
			},

			"Eq": {
				OpCode: typesys.OP_BINARY_EQ,
				ApplyTo: []OperationApplyTo{
					{
						SeriesType: "SeriesInt32",
						InnerType:  typesys.Int32Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = %s.data[%s] == float64(%s.data[%s])`, res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesType: "SeriesInt64",
						InnerType:  typesys.Int64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = %s.data[%s] == float64(%s.data[%s])`, res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesType: "SeriesFloat64",
						InnerType:  typesys.Float64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = %s.data[%s] == %s.data[%s]`, res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
				},
			},

			"Ne": {
				OpCode: typesys.OP_BINARY_NE,
				ApplyTo: []OperationApplyTo{
					{
						SeriesType: "SeriesInt32",
						InnerType:  typesys.Int32Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = %s.data[%s] != float64(%s.data[%s])`, res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesType: "SeriesInt64",
						InnerType:  typesys.Int64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = %s.data[%s] != float64(%s.data[%s])`, res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesType: "SeriesFloat64",
						InnerType:  typesys.Float64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = %s.data[%s] != %s.data[%s]`, res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
				},
			},

			"Lt": {
				OpCode: typesys.OP_BINARY_LT,
				ApplyTo: []OperationApplyTo{
					{
						SeriesType: "SeriesInt32",
						InnerType:  typesys.Int32Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = %s.data[%s] < float64(%s.data[%s])`, res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesType: "SeriesInt64",
						InnerType:  typesys.Int64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = %s.data[%s] < float64(%s.data[%s])`, res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesType: "SeriesFloat64",
						InnerType:  typesys.Float64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = %s.data[%s] < %s.data[%s]`, res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
				},
			},

			"Le": {
				OpCode: typesys.OP_BINARY_LE,
				ApplyTo: []OperationApplyTo{
					{
						SeriesType: "SeriesInt32",
						InnerType:  typesys.Int32Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = %s.data[%s] <= float64(%s.data[%s])`, res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesType: "SeriesInt64",
						InnerType:  typesys.Int64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = %s.data[%s] <= float64(%s.data[%s])`, res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesType: "SeriesFloat64",
						InnerType:  typesys.Float64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = %s.data[%s] <= %s.data[%s]`, res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
				},
			},

			"Gt": {
				OpCode: typesys.OP_BINARY_GT,
				ApplyTo: []OperationApplyTo{
					{
						SeriesType: "SeriesInt32",
						InnerType:  typesys.Int32Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = %s.data[%s] > float64(%s.data[%s])`, res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesType: "SeriesInt64",
						InnerType:  typesys.Int64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = %s.data[%s] > float64(%s.data[%s])`, res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesType: "SeriesFloat64",
						InnerType:  typesys.Float64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = %s.data[%s] > %s.data[%s]`, res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
				},
			},

			"Ge": {
				OpCode: typesys.OP_BINARY_GE,
				ApplyTo: []OperationApplyTo{
					{
						SeriesType: "SeriesInt32",
						InnerType:  typesys.Int32Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = %s.data[%s] >= float64(%s.data[%s])`, res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesType: "SeriesInt64",
						InnerType:  typesys.Int64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = %s.data[%s] >= float64(%s.data[%s])`, res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesType: "SeriesFloat64",
						InnerType:  typesys.Float64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = %s.data[%s] >= %s.data[%s]`, res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
				},
			},
		},
	},

	"gdl_series_string.go": {
		SeriesType: "SeriesString",
		InnerType:  typesys.StringType,
		Operations: map[string]Operation{
			"Add": {
				OpCode: typesys.OP_BINARY_ADD,
				ApplyTo: []OperationApplyTo{
					{
						SeriesType: "SeriesInt32",
						InnerType:  typesys.Int32Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.pool.Put(*%s.data[%s] + intToString(int64(%s.data[%s])))", res, resIndex, op1, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesType: "SeriesInt64",
						InnerType:  typesys.Int64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.pool.Put(*%s.data[%s] + intToString(%s.data[%s]))", res, resIndex, op1, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesType: "SeriesFloat64",
						InnerType:  typesys.Float64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.pool.Put(*%s.data[%s] + floatToString(%s.data[%s]))", res, resIndex, op1, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesType: "SeriesString",
						InnerType:  typesys.StringType,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.pool.Put(*%s.data[%s] + *%s.data[%s])", res, resIndex, op1, op1, op1Index, op2, op2Index)}
						},
					},
				},
			},

			"Eq": {
				OpCode: typesys.OP_BINARY_EQ,
				ApplyTo: []OperationApplyTo{
					{
						SeriesType: "SeriesString",
						InnerType:  typesys.StringType,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = *%s.data[%s] == *%s.data[%s]`, res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
				},
			},

			"Ne": {
				OpCode: typesys.OP_BINARY_NE,
				ApplyTo: []OperationApplyTo{
					{
						SeriesType: "SeriesString",
						InnerType:  typesys.StringType,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = *%s.data[%s] != *%s.data[%s]`, res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
				},
			},

			"Lt": {
				OpCode: typesys.OP_BINARY_LT,
				ApplyTo: []OperationApplyTo{
					{
						SeriesType: "SeriesString",
						InnerType:  typesys.StringType,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = *%s.data[%s] < *%s.data[%s]`, res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
				},
			},

			"Le": {
				OpCode: typesys.OP_BINARY_LE,
				ApplyTo: []OperationApplyTo{
					{
						SeriesType: "SeriesString",
						InnerType:  typesys.StringType,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = *%s.data[%s] <= *%s.data[%s]`, res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
				},
			},

			"Gt": {
				OpCode: typesys.OP_BINARY_GT,
				ApplyTo: []OperationApplyTo{
					{
						SeriesType: "SeriesString",
						InnerType:  typesys.StringType,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = *%s.data[%s] > *%s.data[%s]`, res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
				},
			},

			"Ge": {
				OpCode: typesys.OP_BINARY_GE,
				ApplyTo: []OperationApplyTo{
					{
						SeriesType: "SeriesString",
						InnerType:  typesys.StringType,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = *%s.data[%s] >= *%s.data[%s]`, res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
				},
			},
		},
	},
}
