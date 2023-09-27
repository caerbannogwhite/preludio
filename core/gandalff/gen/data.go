package main

import (
	"fmt"
	"go/ast"
	"typesys"
)

type MakeOperationType func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr

type OperationApplyTo struct {
	SeriesName    string
	SeriesType    typesys.BaseType
	MakeOperation MakeOperationType
}

type Operation struct {
	OpCode  typesys.OPCODE
	ApplyTo []OperationApplyTo
}

type SeriesFile struct {
	SeriesName      string
	SeriesType      typesys.BaseType
	SeriesTypeStr   string
	SeriesGoTypeStr string
	IsGoTypePtr     bool
	IsTimeType      bool
	Operations      map[string]Operation
}

var DATA_BASE_METHODS = map[string]SeriesFile{
	"gdl_series_bool_base.go": {
		SeriesName:      "SeriesBool",
		SeriesTypeStr:   "BoolType",
		SeriesGoTypeStr: "bool",
	},

	"gdl_series_int32_base.go": {
		SeriesName:      "SeriesInt32",
		SeriesTypeStr:   "Int32Type",
		SeriesGoTypeStr: "int32",
	},

	"gdl_series_int64_base.go": {
		SeriesName:      "SeriesInt64",
		SeriesTypeStr:   "Int64Type",
		SeriesGoTypeStr: "int64",
	},

	"gdl_series_float64_base.go": {
		SeriesName:      "SeriesFloat64",
		SeriesTypeStr:   "Float64Type",
		SeriesGoTypeStr: "float64",
	},

	"gdl_series_string_base.go": {
		SeriesName:      "SeriesString",
		SeriesTypeStr:   "StringType",
		SeriesGoTypeStr: "*string",
		IsGoTypePtr:     true,
	},

	"gdl_series_time_base.go": {
		SeriesName:      "SeriesTime",
		SeriesTypeStr:   "TimeType",
		SeriesGoTypeStr: "time.Time",
		IsTimeType:      true,
	},

	"gdl_series_duration_base.go": {
		SeriesName:      "SeriesDuration",
		SeriesTypeStr:   "DurationType",
		SeriesGoTypeStr: "time.Duration",
	},
}

var DATA_OPERATIONS = map[string]SeriesFile{
	"gdl_series_bool_ops.go": {
		SeriesName: "SeriesBool",
		SeriesType: typesys.BoolType,
		Operations: map[string]Operation{
			"Mul": {
				OpCode: typesys.OP_BINARY_MUL,
				ApplyTo: []OperationApplyTo{
					{
						SeriesName: "SeriesBool",
						SeriesType: typesys.BoolType,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("if %s.data[%s] && %s.data[%s] { %s[%s] = 1 }", op1, op1Index, op2, op2Index, res, resIndex)}
						},
					},
					{
						SeriesName: "SeriesInt32",
						SeriesType: typesys.Int32Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("if %s.data[%s] { %s[%s] = %s.data[%s] }", op1, op1Index, res, resIndex, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesInt64",
						SeriesType: typesys.Int64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("if %s.data[%s] { %s[%s] = %s.data[%s] }", op1, op1Index, res, resIndex, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesFloat64",
						SeriesType: typesys.Float64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("if %s.data[%s] { %s[%s] = %s.data[%s] }", op1, op1Index, res, resIndex, op2, op2Index)}
						},
					},
				},
			},

			"Div": {
				OpCode: typesys.OP_BINARY_DIV,
				ApplyTo: []OperationApplyTo{
					{
						SeriesName: "SeriesBool",
						SeriesType: typesys.BoolType,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("b1 := float64(0)\nb2 := float64(0)\nif %s.data[%s] { b1 = 1 }\nif %s.data[%s] { b2 = 1 }\n%s[%s] = b1 / b2", op1, op1Index, op2, op2Index, res, resIndex)}
						},
					},
					{
						SeriesName: "SeriesInt32",
						SeriesType: typesys.Int32Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("b1 := float64(0)\nif %s.data[%s] { b1 = 1 }\n%s[%s] = b1 / float64(%s.data[%s])", op1, op1Index, res, resIndex, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesInt64",
						SeriesType: typesys.Int64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("b1 := float64(0)\nif %s.data[%s] { b1 = 1 }\n%s[%s] = b1 / float64(%s.data[%s])", op1, op1Index, res, resIndex, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesFloat64",
						SeriesType: typesys.Float64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("b1 := float64(0)\nif %s.data[%s] { b1 = 1 }\n%s[%s] = b1 / %s.data[%s]", op1, op1Index, res, resIndex, op2, op2Index)}
						},
					},
				},
			},

			"Mod": {
				OpCode: typesys.OP_BINARY_MOD,
				ApplyTo: []OperationApplyTo{
					{
						SeriesName: "SeriesBool",
						SeriesType: typesys.BoolType,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("b1 := float64(0)\nb2 := float64(0)\nif %s.data[%s] { b1 = 1 }\nif %s.data[%s] { b2 = 1 }\n%s[%s] = math.Mod(b1, b2)", op1, op1Index, op2, op2Index, res, resIndex)}
						},
					},
					{
						SeriesName: "SeriesInt32",
						SeriesType: typesys.Int32Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("b1 := float64(0)\nif %s.data[%s] { b1 = 1 }\n%s[%s] = math.Mod(b1, float64(%s.data[%s]))", op1, op1Index, res, resIndex, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesInt64",
						SeriesType: typesys.Int64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("b1 := float64(0)\nif %s.data[%s] { b1 = 1 }\n%s[%s] = math.Mod(b1, float64(%s.data[%s]))", op1, op1Index, res, resIndex, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesFloat64",
						SeriesType: typesys.Float64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("b1 := float64(0)\nif %s.data[%s] { b1 = 1 }\n%s[%s] = math.Mod(b1, float64(%s.data[%s]))", op1, op1Index, res, resIndex, op2, op2Index)}
						},
					},
				},
			},

			"Pow": {
				OpCode: typesys.OP_BINARY_POW,
				ApplyTo: []OperationApplyTo{
					{
						SeriesName: "SeriesBool",
						SeriesType: typesys.BoolType,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("b1 := float64(0)\nb2 := float64(0)\nif %s.data[%s] { b1 = 1 }\nif %s.data[%s] { b2 = 1 }\n%s[%s] = int64(math.Pow(b1, b2))", op1, op1Index, op2, op2Index, res, resIndex)}
						},
					},
					{
						SeriesName: "SeriesInt32",
						SeriesType: typesys.Int32Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("b1 := float64(0)\nif %s.data[%s] { b1 = 1 }\n%s[%s] = int64(math.Pow(b1, float64(%s.data[%s])))", op1, op1Index, res, resIndex, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesInt64",
						SeriesType: typesys.Int64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("b1 := float64(0)\nif %s.data[%s] { b1 = 1 }\n%s[%s] = int64(math.Pow(b1, float64(%s.data[%s])))", op1, op1Index, res, resIndex, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesFloat64",
						SeriesType: typesys.Float64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("b1 := float64(0)\nif %s.data[%s] { b1 = 1 }\n%s[%s] = float64(math.Pow(b1, float64(%s.data[%s])))", op1, op1Index, res, resIndex, op2, op2Index)}
						},
					},
				},
			},

			"Add": {
				OpCode: typesys.OP_BINARY_ADD,
				ApplyTo: []OperationApplyTo{
					{
						SeriesName: "SeriesBool",
						SeriesType: typesys.BoolType,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("b1 := int64(0)\nb2 := int64(0)\nif %s.data[%s] { b1 = 1 }\nif %s.data[%s] { b2 = 1 }\n%s[%s] = b1 + b2", op1, op1Index, op2, op2Index, res, resIndex)}
						},
					},
					{
						SeriesName: "SeriesInt32",
						SeriesType: typesys.Int32Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("b1 := int32(0)\nif %s.data[%s] { b1 = 1 }\n%s[%s] = b1 + %s.data[%s]", op1, op1Index, res, resIndex, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesInt64",
						SeriesType: typesys.Int64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("b1 := int64(0)\nif %s.data[%s] { b1 = 1 }\n%s[%s] = b1 + %s.data[%s]", op1, op1Index, res, resIndex, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesFloat64",
						SeriesType: typesys.Float64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("b1 := float64(0)\nif %s.data[%s] { b1 = 1 }\n%s[%s] = b1 + %s.data[%s]", op1, op1Index, res, resIndex, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesString",
						SeriesType: typesys.StringType,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.pool.Put(boolToString(%s.data[%s]) + *%s.data[%s])", res, resIndex, op2, op1, op1Index, op2, op2Index)}
						},
					},
				},
			},

			"Sub": {
				OpCode: typesys.OP_BINARY_SUB,
				ApplyTo: []OperationApplyTo{
					{
						SeriesName: "SeriesBool",
						SeriesType: typesys.BoolType,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("b1 := int64(0)\nb2 := int64(0)\nif %s.data[%s] { b1 = 1 }\nif %s.data[%s] { b2 = 1 }\n%s[%s] = b1 - b2", op1, op1Index, op2, op2Index, res, resIndex)}
						},
					},
					{
						SeriesName: "SeriesInt32",
						SeriesType: typesys.Int32Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("b1 := int32(0)\nif %s.data[%s] { b1 = 1 }\n%s[%s] = b1 - %s.data[%s]", op1, op1Index, res, resIndex, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesInt64",
						SeriesType: typesys.Int64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("b1 := int64(0)\nif %s.data[%s] { b1 = 1 }\n%s[%s] = b1 - %s.data[%s]", op1, op1Index, res, resIndex, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesFloat64",
						SeriesType: typesys.Float64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("b1 := float64(0)\nif %s.data[%s] { b1 = 1 }\n%s[%s] = b1 - %s.data[%s]", op1, op1Index, res, resIndex, op2, op2Index)}
						},
					},
				},
			},

			"Eq": {
				OpCode: typesys.OP_BINARY_EQ,
				ApplyTo: []OperationApplyTo{
					{
						SeriesName: "SeriesBool",
						SeriesType: typesys.BoolType,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] == %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
				},
			},

			"Ne": {
				OpCode: typesys.OP_BINARY_NE,
				ApplyTo: []OperationApplyTo{
					{
						SeriesName: "SeriesBool",
						SeriesType: typesys.BoolType,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] != %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
				},
			},

			"Lt": {
				OpCode: typesys.OP_BINARY_LT,
				ApplyTo: []OperationApplyTo{
					{
						SeriesName: "SeriesBool",
						SeriesType: typesys.BoolType,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("b1 := int64(0)\nb2 := int64(0)\nif %s.data[%s] { b1 = 1 }\nif %s.data[%s] { b2 = 1 }\n%s[%s] = b1 < b2", op1, op1Index, op2, op2Index, res, resIndex)}
						},
					},
					{
						SeriesName: "SeriesInt32",
						SeriesType: typesys.Int32Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("b1 := int32(0)\nif %s.data[%s] { b1 = 1 }\n%s[%s] = b1 < %s.data[%s]", op1, op1Index, res, resIndex, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesInt64",
						SeriesType: typesys.Int64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("b1 := int64(0)\nif %s.data[%s] { b1 = 1 }\n%s[%s] = b1 < %s.data[%s]", op1, op1Index, res, resIndex, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesFloat64",
						SeriesType: typesys.Float64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("b1 := float64(0)\nif %s.data[%s] { b1 = 1 }\n%s[%s] = b1 < %s.data[%s]", op1, op1Index, res, resIndex, op2, op2Index)}
						},
					},
				},
			},

			"Le": {
				OpCode: typesys.OP_BINARY_LE,
				ApplyTo: []OperationApplyTo{
					{
						SeriesName: "SeriesBool",
						SeriesType: typesys.BoolType,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("b1 := int64(0)\nb2 := int64(0)\nif %s.data[%s] { b1 = 1 }\nif %s.data[%s] { b2 = 1 }\n%s[%s] = b1 <= b2", op1, op1Index, op2, op2Index, res, resIndex)}
						},
					},
					{
						SeriesName: "SeriesInt32",
						SeriesType: typesys.Int32Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("b1 := int32(0)\nif %s.data[%s] { b1 = 1 }\n%s[%s] = b1 <= %s.data[%s]", op1, op1Index, res, resIndex, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesInt64",
						SeriesType: typesys.Int64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("b1 := int64(0)\nif %s.data[%s] { b1 = 1 }\n%s[%s] = b1 <= %s.data[%s]", op1, op1Index, res, resIndex, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesFloat64",
						SeriesType: typesys.Float64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("b1 := float64(0)\nif %s.data[%s] { b1 = 1 }\n%s[%s] = b1 <= %s.data[%s]", op1, op1Index, res, resIndex, op2, op2Index)}
						},
					},
				},
			},

			"Gt": {
				OpCode: typesys.OP_BINARY_GT,
				ApplyTo: []OperationApplyTo{
					{
						SeriesName: "SeriesBool",
						SeriesType: typesys.BoolType,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("b1 := int64(0)\nb2 := int64(0)\nif %s.data[%s] { b1 = 1 }\nif %s.data[%s] { b2 = 1 }\n%s[%s] = b1 > b2", op1, op1Index, op2, op2Index, res, resIndex)}
						},
					},
					{
						SeriesName: "SeriesInt32",
						SeriesType: typesys.Int32Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("b1 := int32(0)\nif %s.data[%s] { b1 = 1 }\n%s[%s] = b1 > %s.data[%s]", op1, op1Index, res, resIndex, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesInt64",
						SeriesType: typesys.Int64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("b1 := int64(0)\nif %s.data[%s] { b1 = 1 }\n%s[%s] = b1 > %s.data[%s]", op1, op1Index, res, resIndex, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesFloat64",
						SeriesType: typesys.Float64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("b1 := float64(0)\nif %s.data[%s] { b1 = 1 }\n%s[%s] = b1 > %s.data[%s]", op1, op1Index, res, resIndex, op2, op2Index)}
						},
					},
				},
			},

			"Ge": {
				OpCode: typesys.OP_BINARY_GE,
				ApplyTo: []OperationApplyTo{
					{
						SeriesName: "SeriesBool",
						SeriesType: typesys.BoolType,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("b1 := int64(0)\nb2 := int64(0)\nif %s.data[%s] { b1 = 1 }\nif %s.data[%s] { b2 = 1 }\n%s[%s] = b1 >= b2", op1, op1Index, op2, op2Index, res, resIndex)}
						},
					},
					{
						SeriesName: "SeriesInt32",
						SeriesType: typesys.Int32Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("b1 := int32(0)\nif %s.data[%s] { b1 = 1 }\n%s[%s] = b1 >= %s.data[%s]", op1, op1Index, res, resIndex, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesInt64",
						SeriesType: typesys.Int64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("b1 := int64(0)\nif %s.data[%s] { b1 = 1 }\n%s[%s] = b1 >= %s.data[%s]", op1, op1Index, res, resIndex, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesFloat64",
						SeriesType: typesys.Float64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("b1 := float64(0)\nif %s.data[%s] { b1 = 1 }\n%s[%s] = b1 >= %s.data[%s]", op1, op1Index, res, resIndex, op2, op2Index)}
						},
					},
				},
			},

			"And": {
				OpCode: typesys.OP_BINARY_AND,
				ApplyTo: []OperationApplyTo{
					{
						SeriesName: "SeriesBool",
						SeriesType: typesys.BoolType,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] && %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
				},
			},

			"Or": {
				OpCode: typesys.OP_BINARY_OR,
				ApplyTo: []OperationApplyTo{
					{
						SeriesName: "SeriesBool",
						SeriesType: typesys.BoolType,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] || %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
				},
			},
		},
	},

	"gdl_series_int32_ops.go": {
		SeriesName: "SeriesInt32",
		SeriesType: typesys.Int32Type,
		Operations: map[string]Operation{
			"Mul": {
				OpCode: typesys.OP_BINARY_MUL,
				ApplyTo: []OperationApplyTo{
					{
						SeriesName: "SeriesBool",
						SeriesType: typesys.BoolType,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("if %s.data[%s] { %s[%s] = %s.data[%s] }", op2, op2Index, res, resIndex, op1, op1Index)}
						},
					},
					{
						SeriesName: "SeriesInt32",
						SeriesType: typesys.Int32Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] * %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesInt64",
						SeriesType: typesys.Int64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = int64(%s.data[%s]) * %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesFloat64",
						SeriesType: typesys.Float64Type,
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
						SeriesName: "SeriesBool",
						SeriesType: typesys.BoolType,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("b2 := float64(0)\nif %s.data[%s] { b2 = 1 }\n%s[%s] = float64(%s.data[%s]) / b2", op2, op2Index, res, resIndex, op1, op1Index)}
						},
					},
					{
						SeriesName: "SeriesInt32",
						SeriesType: typesys.Int32Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = float64(%s.data[%s]) / float64(%s.data[%s])", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesInt64",
						SeriesType: typesys.Int64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = float64(%s.data[%s]) / float64(%s.data[%s])", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesFloat64",
						SeriesType: typesys.Float64Type,
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
						SeriesName: "SeriesBool",
						SeriesType: typesys.BoolType,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("b2 := float64(0)\nif %s.data[%s] { b2 = 1 }\n%s[%s] = math.Mod(float64(%s.data[%s]), b2)", op2, op2Index, res, resIndex, op1, op1Index)}
						},
					},
					{
						SeriesName: "SeriesInt32",
						SeriesType: typesys.Int32Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = math.Mod(float64(%s.data[%s]), float64(%s.data[%s]))", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesInt64",
						SeriesType: typesys.Int64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = math.Mod(float64(%s.data[%s]), float64(%s.data[%s]))", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesFloat64",
						SeriesType: typesys.Float64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = math.Mod(float64(%s.data[%s]), float64(%s.data[%s]))", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
				},
			},

			"Pow": {
				OpCode: typesys.OP_BINARY_POW,
				ApplyTo: []OperationApplyTo{
					{
						SeriesName: "SeriesBool",
						SeriesType: typesys.BoolType,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("b2 := float64(0)\nif %s.data[%s] { b2 = 1 }\n%s[%s] = int64(math.Pow(float64(%s.data[%s]), b2))", op2, op2Index, res, resIndex, op1, op1Index)}
						},
					},
					{
						SeriesName: "SeriesInt32",
						SeriesType: typesys.Int32Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = int64(math.Pow(float64(%s.data[%s]), float64(%s.data[%s])))", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesInt64",
						SeriesType: typesys.Int64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = int64(math.Pow(float64(%s.data[%s]), float64(%s.data[%s])))", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesFloat64",
						SeriesType: typesys.Float64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = math.Pow(float64(%s.data[%s]), float64(%s.data[%s]))", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
				},
			},

			"Add": {
				OpCode: typesys.OP_BINARY_ADD,
				ApplyTo: []OperationApplyTo{
					{
						SeriesName: "SeriesBool",
						SeriesType: typesys.BoolType,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("b2 := int32(0)\nif %s.data[%s] { b2 = 1 }\n%s[%s] = %s.data[%s] + b2", op2, op2Index, res, resIndex, op1, op1Index)}
						},
					},
					{
						SeriesName: "SeriesInt32",
						SeriesType: typesys.Int32Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] + %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesInt64",
						SeriesType: typesys.Int64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = int64(%s.data[%s]) + %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesFloat64",
						SeriesType: typesys.Float64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = float64(%s.data[%s]) + %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesString",
						SeriesType: typesys.StringType,
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
						SeriesName: "SeriesBool",
						SeriesType: typesys.BoolType,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("b2 := int32(0)\nif %s.data[%s] { b2 = 1 }\n%s[%s] = %s.data[%s] - b2", op2, op2Index, res, resIndex, op1, op1Index)}
						},
					},
					{
						SeriesName: "SeriesInt32",
						SeriesType: typesys.Int32Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] - %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesInt64",
						SeriesType: typesys.Int64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = int64(%s.data[%s]) - %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesFloat64",
						SeriesType: typesys.Float64Type,
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
						SeriesName: "SeriesInt32",
						SeriesType: typesys.Int32Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] == %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesInt64",
						SeriesType: typesys.Int64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = int64(%s.data[%s]) == %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesFloat64",
						SeriesType: typesys.Float64Type,
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
						SeriesName: "SeriesInt32",
						SeriesType: typesys.Int32Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] != %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesInt64",
						SeriesType: typesys.Int64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = int64(%s.data[%s]) != %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesFloat64",
						SeriesType: typesys.Float64Type,
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
						SeriesName: "SeriesBool",
						SeriesType: typesys.BoolType,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("b2 := int32(0)\nif %s.data[%s] { b2 = 1 }\n%s[%s] = %s.data[%s] < b2", op2, op2Index, res, resIndex, op1, op1Index)}
						},
					},
					{
						SeriesName: "SeriesInt32",
						SeriesType: typesys.Int32Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] < %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesInt64",
						SeriesType: typesys.Int64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = int64(%s.data[%s]) < %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesFloat64",
						SeriesType: typesys.Float64Type,
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
						SeriesName: "SeriesBool",
						SeriesType: typesys.BoolType,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("b2 := int32(0)\nif %s.data[%s] { b2 = 1 }\n%s[%s] = %s.data[%s] <= b2", op2, op2Index, res, resIndex, op1, op1Index)}
						},
					},
					{
						SeriesName: "SeriesInt32",
						SeriesType: typesys.Int32Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] <= %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesInt64",
						SeriesType: typesys.Int64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = int64(%s.data[%s]) <= %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesFloat64",
						SeriesType: typesys.Float64Type,
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
						SeriesName: "SeriesBool",
						SeriesType: typesys.BoolType,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("b2 := int32(0)\nif %s.data[%s] { b2 = 1 }\n%s[%s] = %s.data[%s] > b2", op2, op2Index, res, resIndex, op1, op1Index)}
						},
					},
					{
						SeriesName: "SeriesInt32",
						SeriesType: typesys.Int32Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] > %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesInt64",
						SeriesType: typesys.Int64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = int64(%s.data[%s]) > %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesFloat64",
						SeriesType: typesys.Float64Type,
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
						SeriesName: "SeriesBool",
						SeriesType: typesys.BoolType,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("b2 := int32(0)\nif %s.data[%s] { b2 = 1 }\n%s[%s] = %s.data[%s] >= b2", op2, op2Index, res, resIndex, op1, op1Index)}
						},
					},
					{
						SeriesName: "SeriesInt32",
						SeriesType: typesys.Int32Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] >= %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesInt64",
						SeriesType: typesys.Int64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = int64(%s.data[%s]) >= %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesFloat64",
						SeriesType: typesys.Float64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = float64(%s.data[%s]) >= %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
				},
			},
		},
	},

	"gdl_series_int64_ops.go": {
		SeriesName: "SeriesInt64",
		SeriesType: typesys.Int64Type,
		Operations: map[string]Operation{
			"Mul": {
				OpCode: typesys.OP_BINARY_MUL,
				ApplyTo: []OperationApplyTo{
					{
						SeriesName: "SeriesBool",
						SeriesType: typesys.BoolType,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("if %s.data[%s] { %s[%s] = %s.data[%s] }", op2, op2Index, res, resIndex, op1, op1Index)}
						},
					},
					{
						SeriesName: "SeriesInt32",
						SeriesType: typesys.Int32Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] * int64(%s.data[%s])", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesInt64",
						SeriesType: typesys.Int64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] * %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesFloat64",
						SeriesType: typesys.Float64Type,
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
						SeriesName: "SeriesBool",
						SeriesType: typesys.BoolType,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("b2 := float64(0)\nif %s.data[%s] { b2 = 1 }\n%s[%s] = float64(%s.data[%s]) / b2", op2, op2Index, res, resIndex, op1, op1Index)}
						},
					},
					{
						SeriesName: "SeriesInt32",
						SeriesType: typesys.Int32Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = float64(%s.data[%s]) / float64(%s.data[%s])", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesInt64",
						SeriesType: typesys.Int64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = float64(%s.data[%s]) / float64(%s.data[%s])", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesFloat64",
						SeriesType: typesys.Float64Type,
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
						SeriesName: "SeriesBool",
						SeriesType: typesys.BoolType,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("b2 := float64(0)\nif %s.data[%s] { b2 = 1 }\n%s[%s] = math.Mod(float64(%s.data[%s]), b2)", op2, op2Index, res, resIndex, op1, op1Index)}
						},
					},
					{
						SeriesName: "SeriesInt32",
						SeriesType: typesys.Int32Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = math.Mod(float64(%s.data[%s]), float64(%s.data[%s]))", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesInt64",
						SeriesType: typesys.Int64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = math.Mod(float64(%s.data[%s]), float64(%s.data[%s]))", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesFloat64",
						SeriesType: typesys.Float64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = math.Mod(float64(%s.data[%s]), float64(%s.data[%s]))", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
				},
			},

			"Pow": {
				OpCode: typesys.OP_BINARY_POW,
				ApplyTo: []OperationApplyTo{
					{
						SeriesName: "SeriesBool",
						SeriesType: typesys.BoolType,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("b2 := float64(0)\nif %s.data[%s] { b2 = 1 }\n%s[%s] = int64(math.Pow(float64(%s.data[%s]), b2))", op2, op2Index, res, resIndex, op1, op1Index)}
						},
					},
					{
						SeriesName: "SeriesInt32",
						SeriesType: typesys.Int32Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = int64(math.Pow(float64(%s.data[%s]), float64(%s.data[%s])))", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesInt64",
						SeriesType: typesys.Int64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = int64(math.Pow(float64(%s.data[%s]), float64(%s.data[%s])))", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesFloat64",
						SeriesType: typesys.Float64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = math.Pow(float64(%s.data[%s]), float64(%s.data[%s]))", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
				},
			},

			"Add": {
				OpCode: typesys.OP_BINARY_ADD,
				ApplyTo: []OperationApplyTo{
					{
						SeriesName: "SeriesBool",
						SeriesType: typesys.BoolType,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("b2 := int64(0)\nif %s.data[%s] { b2 = 1 }\n%s[%s] = %s.data[%s] + b2", op2, op2Index, res, resIndex, op1, op1Index)}
						},
					},
					{
						SeriesName: "SeriesInt32",
						SeriesType: typesys.Int32Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] + int64(%s.data[%s])", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesInt64",
						SeriesType: typesys.Int64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] + %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesFloat64",
						SeriesType: typesys.Float64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = float64(%s.data[%s]) + %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesString",
						SeriesType: typesys.StringType,
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
						SeriesName: "SeriesBool",
						SeriesType: typesys.BoolType,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("b2 := int64(0)\nif %s.data[%s] { b2 = 1 }\n%s[%s] = %s.data[%s] - b2", op2, op2Index, res, resIndex, op1, op1Index)}
						},
					},
					{
						SeriesName: "SeriesInt32",
						SeriesType: typesys.Int32Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] - int64(%s.data[%s])", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesInt64",
						SeriesType: typesys.Int64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] - %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesFloat64",
						SeriesType: typesys.Float64Type,
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
						SeriesName: "SeriesInt32",
						SeriesType: typesys.Int32Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] == int64(%s.data[%s])", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesInt64",
						SeriesType: typesys.Int64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] == %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesFloat64",
						SeriesType: typesys.Float64Type,
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
						SeriesName: "SeriesInt32",
						SeriesType: typesys.Int32Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] != int64(%s.data[%s])", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesInt64",
						SeriesType: typesys.Int64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] != %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesFloat64",
						SeriesType: typesys.Float64Type,
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
						SeriesName: "SeriesBool",
						SeriesType: typesys.BoolType,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("b2 := int64(0)\nif %s.data[%s] { b2 = 1 }\n%s[%s] = %s.data[%s] < b2", op2, op2Index, res, resIndex, op1, op1Index)}
						},
					},
					{
						SeriesName: "SeriesInt32",
						SeriesType: typesys.Int32Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = %s.data[%s] < int64(%s.data[%s])`, res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesInt64",
						SeriesType: typesys.Int64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = %s.data[%s] < %s.data[%s]`, res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesFloat64",
						SeriesType: typesys.Float64Type,
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
						SeriesName: "SeriesBool",
						SeriesType: typesys.BoolType,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("b2 := int64(0)\nif %s.data[%s] { b2 = 1 }\n%s[%s] = %s.data[%s] <= b2", op2, op2Index, res, resIndex, op1, op1Index)}
						},
					},
					{
						SeriesName: "SeriesInt32",
						SeriesType: typesys.Int32Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = %s.data[%s] <= int64(%s.data[%s])`, res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesInt64",
						SeriesType: typesys.Int64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = %s.data[%s] <= %s.data[%s]`, res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesFloat64",
						SeriesType: typesys.Float64Type,
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
						SeriesName: "SeriesBool",
						SeriesType: typesys.BoolType,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("b2 := int64(0)\nif %s.data[%s] { b2 = 1 }\n%s[%s] = %s.data[%s] > b2", op2, op2Index, res, resIndex, op1, op1Index)}
						},
					},
					{
						SeriesName: "SeriesInt32",
						SeriesType: typesys.Int32Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = %s.data[%s] > int64(%s.data[%s])`, res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesInt64",
						SeriesType: typesys.Int64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = %s.data[%s] > %s.data[%s]`, res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesFloat64",
						SeriesType: typesys.Float64Type,
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
						SeriesName: "SeriesBool",
						SeriesType: typesys.BoolType,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("b2 := int64(0)\nif %s.data[%s] { b2 = 1 }\n%s[%s] = %s.data[%s] >= b2", op2, op2Index, res, resIndex, op1, op1Index)}
						},
					},
					{
						SeriesName: "SeriesInt32",
						SeriesType: typesys.Int32Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = %s.data[%s] >= int64(%s.data[%s])`, res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesInt64",
						SeriesType: typesys.Int64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = %s.data[%s] >= %s.data[%s]`, res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesFloat64",
						SeriesType: typesys.Float64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = float64(%s.data[%s]) >= %s.data[%s]`, res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
				},
			},
		},
	},

	"gdl_series_float64_ops.go": {
		SeriesName: "SeriesFloat64",
		SeriesType: typesys.Float64Type,
		Operations: map[string]Operation{
			"Mul": {
				OpCode: typesys.OP_BINARY_MUL,
				ApplyTo: []OperationApplyTo{
					{
						SeriesName: "SeriesBool",
						SeriesType: typesys.BoolType,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("if %s.data[%s] { %s[%s] = %s.data[%s] }", op2, op2Index, res, resIndex, op1, op1Index)}
						},
					},
					{
						SeriesName: "SeriesInt32",
						SeriesType: typesys.Int32Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] * float64(%s.data[%s])", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesInt64",
						SeriesType: typesys.Int64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] * float64(%s.data[%s])", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesFloat64",
						SeriesType: typesys.Float64Type,
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
						SeriesName: "SeriesBool",
						SeriesType: typesys.BoolType,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("b2 := float64(0)\nif %s.data[%s] { b2 = 1 }\n%s[%s] = %s.data[%s] / b2", op2, op2Index, res, resIndex, op1, op1Index)}
						},
					},
					{
						SeriesName: "SeriesInt32",
						SeriesType: typesys.Int32Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] / float64(%s.data[%s])", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesInt64",
						SeriesType: typesys.Int64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] / float64(%s.data[%s])", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesFloat64",
						SeriesType: typesys.Float64Type,
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
						SeriesName: "SeriesBool",
						SeriesType: typesys.BoolType,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("b2 := float64(0)\nif %s.data[%s] { b2 = 1 }\n%s[%s] = math.Mod(%s.data[%s], b2)", op2, op2Index, res, resIndex, op1, op1Index)}
						},
					},
					{
						SeriesName: "SeriesInt32",
						SeriesType: typesys.Int32Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = math.Mod(float64(%s.data[%s]), float64(%s.data[%s]))", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesInt64",
						SeriesType: typesys.Int64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = math.Mod(float64(%s.data[%s]), float64(%s.data[%s]))", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesFloat64",
						SeriesType: typesys.Float64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = math.Mod(float64(%s.data[%s]), float64(%s.data[%s]))", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
				},
			},

			"Pow": {
				OpCode: typesys.OP_BINARY_POW,
				ApplyTo: []OperationApplyTo{
					{
						SeriesName: "SeriesBool",
						SeriesType: typesys.BoolType,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("b2 := float64(0)\nif %s.data[%s] { b2 = 1 }\n%s[%s] = math.Pow(%s.data[%s], b2)", op2, op2Index, res, resIndex, op1, op1Index)}
						},
					},
					{
						SeriesName: "SeriesInt32",
						SeriesType: typesys.Int32Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = math.Pow(%s.data[%s], float64(%s.data[%s]))", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesInt64",
						SeriesType: typesys.Int64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = math.Pow(%s.data[%s], float64(%s.data[%s]))", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesFloat64",
						SeriesType: typesys.Float64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = math.Pow(%s.data[%s], float64(%s.data[%s]))", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
				},
			},

			"Add": {
				OpCode: typesys.OP_BINARY_ADD,
				ApplyTo: []OperationApplyTo{
					{
						SeriesName: "SeriesBool",
						SeriesType: typesys.BoolType,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("b2 := float64(0)\nif %s.data[%s] { b2 = 1 }\n%s[%s] = %s.data[%s] + b2", op2, op2Index, res, resIndex, op1, op1Index)}
						},
					},
					{
						SeriesName: "SeriesInt32",
						SeriesType: typesys.Int32Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] + float64(%s.data[%s])", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesInt64",
						SeriesType: typesys.Int64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] + float64(%s.data[%s])", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesFloat64",
						SeriesType: typesys.Float64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] + %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesString",
						SeriesType: typesys.StringType,
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
						SeriesName: "SeriesBool",
						SeriesType: typesys.BoolType,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("b2 := float64(0)\nif %s.data[%s] { b2 = 1 }\n%s[%s] = %s.data[%s] - b2", op2, op2Index, res, resIndex, op1, op1Index)}
						},
					},
					{
						SeriesName: "SeriesInt32",
						SeriesType: typesys.Int32Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] - float64(%s.data[%s])", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesInt64",
						SeriesType: typesys.Int64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] - float64(%s.data[%s])", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesFloat64",
						SeriesType: typesys.Float64Type,
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
						SeriesName: "SeriesInt32",
						SeriesType: typesys.Int32Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = %s.data[%s] == float64(%s.data[%s])`, res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesInt64",
						SeriesType: typesys.Int64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = %s.data[%s] == float64(%s.data[%s])`, res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesFloat64",
						SeriesType: typesys.Float64Type,
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
						SeriesName: "SeriesInt32",
						SeriesType: typesys.Int32Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = %s.data[%s] != float64(%s.data[%s])`, res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesInt64",
						SeriesType: typesys.Int64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = %s.data[%s] != float64(%s.data[%s])`, res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesFloat64",
						SeriesType: typesys.Float64Type,
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
						SeriesName: "SeriesBool",
						SeriesType: typesys.BoolType,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("b2 := float64(0)\nif %s.data[%s] { b2 = 1 }\n%s[%s] = %s.data[%s] < b2", op2, op2Index, res, resIndex, op1, op1Index)}
						},
					},
					{
						SeriesName: "SeriesInt32",
						SeriesType: typesys.Int32Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = %s.data[%s] < float64(%s.data[%s])`, res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesInt64",
						SeriesType: typesys.Int64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = %s.data[%s] < float64(%s.data[%s])`, res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesFloat64",
						SeriesType: typesys.Float64Type,
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
						SeriesName: "SeriesBool",
						SeriesType: typesys.BoolType,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("b2 := float64(0)\nif %s.data[%s] { b2 = 1 }\n%s[%s] = %s.data[%s] <= b2", op2, op2Index, res, resIndex, op1, op1Index)}
						},
					},
					{
						SeriesName: "SeriesInt32",
						SeriesType: typesys.Int32Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = %s.data[%s] <= float64(%s.data[%s])`, res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesInt64",
						SeriesType: typesys.Int64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = %s.data[%s] <= float64(%s.data[%s])`, res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesFloat64",
						SeriesType: typesys.Float64Type,
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
						SeriesName: "SeriesBool",
						SeriesType: typesys.BoolType,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("b2 := float64(0)\nif %s.data[%s] { b2 = 1 }\n%s[%s] = %s.data[%s] > b2", op2, op2Index, res, resIndex, op1, op1Index)}
						},
					},
					{
						SeriesName: "SeriesInt32",
						SeriesType: typesys.Int32Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = %s.data[%s] > float64(%s.data[%s])`, res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesInt64",
						SeriesType: typesys.Int64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = %s.data[%s] > float64(%s.data[%s])`, res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesFloat64",
						SeriesType: typesys.Float64Type,
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
						SeriesName: "SeriesBool",
						SeriesType: typesys.BoolType,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("b2 := float64(0)\nif %s.data[%s] { b2 = 1 }\n%s[%s] = %s.data[%s] >= b2", op2, op2Index, res, resIndex, op1, op1Index)}
						},
					},
					{
						SeriesName: "SeriesInt32",
						SeriesType: typesys.Int32Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = %s.data[%s] >= float64(%s.data[%s])`, res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesInt64",
						SeriesType: typesys.Int64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = %s.data[%s] >= float64(%s.data[%s])`, res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesFloat64",
						SeriesType: typesys.Float64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = %s.data[%s] >= %s.data[%s]`, res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
				},
			},
		},
	},

	"gdl_series_string_ops.go": {
		SeriesName: "SeriesString",
		SeriesType: typesys.StringType,
		Operations: map[string]Operation{
			"Add": {
				OpCode: typesys.OP_BINARY_ADD,
				ApplyTo: []OperationApplyTo{
					{
						SeriesName: "SeriesBool",
						SeriesType: typesys.BoolType,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.pool.Put(*%s.data[%s] + boolToString(%s.data[%s]))", res, resIndex, op1, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesInt32",
						SeriesType: typesys.Int32Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.pool.Put(*%s.data[%s] + intToString(int64(%s.data[%s])))", res, resIndex, op1, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesInt64",
						SeriesType: typesys.Int64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.pool.Put(*%s.data[%s] + intToString(%s.data[%s]))", res, resIndex, op1, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesFloat64",
						SeriesType: typesys.Float64Type,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.pool.Put(*%s.data[%s] + floatToString(%s.data[%s]))", res, resIndex, op1, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesString",
						SeriesType: typesys.StringType,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.pool.Put(*%s.data[%s] + *%s.data[%s])", res, resIndex, op1, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesTime",
						SeriesType: typesys.TimeType,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.pool.Put(*%s.data[%s] + %s.data[%s].String())", res, resIndex, op1, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesDuration",
						SeriesType: typesys.DurationType,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.pool.Put(*%s.data[%s] + %s.data[%s].String())", res, resIndex, op1, op1, op1Index, op2, op2Index)}
						},
					},
				},
			},

			"Eq": {
				OpCode: typesys.OP_BINARY_EQ,
				ApplyTo: []OperationApplyTo{
					{
						SeriesName: "SeriesString",
						SeriesType: typesys.StringType,
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
						SeriesName: "SeriesString",
						SeriesType: typesys.StringType,
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
						SeriesName: "SeriesString",
						SeriesType: typesys.StringType,
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
						SeriesName: "SeriesString",
						SeriesType: typesys.StringType,
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
						SeriesName: "SeriesString",
						SeriesType: typesys.StringType,
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
						SeriesName: "SeriesString",
						SeriesType: typesys.StringType,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = *%s.data[%s] >= *%s.data[%s]`, res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
				},
			},
		},
	},

	"gdl_series_time_ops.go": {
		SeriesName: "SeriesTime",
		SeriesType: typesys.TimeType,
		Operations: map[string]Operation{
			"Add": {
				OpCode: typesys.OP_BINARY_ADD,
				ApplyTo: []OperationApplyTo{
					{
						SeriesName: "SeriesString",
						SeriesType: typesys.StringType,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.pool.Put(%s.data[%s].String() + *%s.data[%s])", res, resIndex, op2, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesTime",
						SeriesType: typesys.TimeType,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s].AddDate(%s.data[%s].Year(), int(%s.data[%s].Month()), %s.data[%s].Day())", res, resIndex, op1, op1Index, op2, op2Index, op2, op2Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesDuration",
						SeriesType: typesys.DurationType,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s].Add(%s.data[%s])", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
				},
			},

			"Sub": {
				OpCode: typesys.OP_BINARY_SUB,
				ApplyTo: []OperationApplyTo{
					{
						SeriesName: "SeriesTime",
						SeriesType: typesys.TimeType,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s].Sub(%s.data[%s])", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesDuration",
						SeriesType: typesys.DurationType,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s].Add(-%s.data[%s])", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
				},
			},
		},
	},

	"gdl_series_duration_ops.go": {
		SeriesName: "SeriesDuration",
		SeriesType: typesys.DurationType,
		Operations: map[string]Operation{
			"Add": {
				OpCode: typesys.OP_BINARY_ADD,
				ApplyTo: []OperationApplyTo{
					{
						SeriesName: "SeriesString",
						SeriesType: typesys.StringType,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.pool.Put(%s.data[%s].String() + *%s.data[%s])", res, resIndex, op2, op1, op1Index, op2, op2Index)}
						},
					},
					{
						SeriesName: "SeriesDuration",
						SeriesType: typesys.DurationType,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] + %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
				},
			},

			"Sub": {
				OpCode: typesys.OP_BINARY_SUB,
				ApplyTo: []OperationApplyTo{
					{
						SeriesName: "SeriesDuration",
						SeriesType: typesys.DurationType,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] - %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
						},
					},
				},
			},
		},
	},
}
