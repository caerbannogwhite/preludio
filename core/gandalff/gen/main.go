package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io/ioutil"
	"path/filepath"
	"typesys"
)

const (
	GOROUTINES           = 4
	RESULT_VAR_NAME      = "result"
	RESULT_SIZE_VAR_NAME = "resultSize"
	FINAL_RETURN_FMT     = "SeriesError{fmt.Sprintf(\"Cannot %s %%s and %%s\", s.Type().ToString(), other.Type().ToString())}"
)

type BuildInfo struct {
	OpCode        typesys.OPCODE
	Op1Nullable   bool
	Op1Scalar     bool
	Op2Nullable   bool
	Op2Scalar     bool
	Op1VarName    string
	Op1SeriesType string
	Op1InnerType  typesys.BaseType
	Op2VarName    string
	Op2SeriesType string
	Op2InnerType  typesys.BaseType
	MakeOperation MakeOperationType
}

func (bi BuildInfo) UpdateScalarInfo(Op1Scalar, Op2Scalar bool) BuildInfo {
	bi.Op1Scalar = Op1Scalar
	bi.Op2Scalar = Op2Scalar
	return bi
}

func (bi BuildInfo) UpdateNullableInfo(Op1Nullable, Op2Nullable bool) BuildInfo {
	bi.Op1Nullable = Op1Nullable
	bi.Op2Nullable = Op2Nullable
	return bi
}

// Generate the code to define the result inner array
func generateMakeResultStmt(info BuildInfo) []ast.Stmt {
	resInnerType := computeResInnerType(info.OpCode, info.Op1InnerType, info.Op2InnerType)

	if resInnerType == info.Op1InnerType {
		if info.Op1Scalar {
			if info.Op2Scalar {
				// TYPE1 | SCALAR | SCALAR
				return []ast.Stmt{
					&ast.AssignStmt{
						Lhs: []ast.Expr{
							&ast.Ident{Name: RESULT_VAR_NAME},
						},
						Tok: token.DEFINE,
						Rhs: []ast.Expr{
							&ast.Ident{Name: fmt.Sprintf("%s.data", info.Op1VarName)},
						},
					},
				}
			} else {
				// TYPE1 | SCALAR | SERIES
				return []ast.Stmt{
					&ast.AssignStmt{
						Lhs: []ast.Expr{
							&ast.Ident{Name: RESULT_SIZE_VAR_NAME},
						},
						Tok: token.DEFINE,
						Rhs: []ast.Expr{
							&ast.Ident{Name: fmt.Sprintf("len(%s.data)", info.Op2VarName)},
						},
					},
					&ast.AssignStmt{
						Lhs: []ast.Expr{
							&ast.Ident{Name: RESULT_VAR_NAME},
						},
						Tok: token.DEFINE,
						Rhs: []ast.Expr{
							&ast.CallExpr{
								Fun: &ast.Ident{Name: "make"},
								Args: []ast.Expr{
									&ast.Ident{Name: resInnerType.ToGoType()},
									&ast.Ident{Name: RESULT_SIZE_VAR_NAME},
								},
							},
						},
					},
				}
			}
		} else {
			if info.Op2Scalar {
				// TYPE1 | SERIES | SCALAR
				return []ast.Stmt{
					&ast.AssignStmt{
						Lhs: []ast.Expr{
							&ast.Ident{Name: RESULT_SIZE_VAR_NAME},
						},
						Tok: token.DEFINE,
						Rhs: []ast.Expr{
							&ast.Ident{Name: fmt.Sprintf("len(%s.data)", info.Op1VarName)},
						},
					},
					&ast.AssignStmt{
						Lhs: []ast.Expr{
							&ast.Ident{Name: RESULT_VAR_NAME},
						},
						Tok: token.DEFINE,
						Rhs: []ast.Expr{
							&ast.Ident{Name: fmt.Sprintf("%s.data", info.Op1VarName)},
						},
					},
				}
			} else {
				// TYPE1 | SERIES | SERIES
				return []ast.Stmt{
					&ast.AssignStmt{
						Lhs: []ast.Expr{
							&ast.Ident{Name: RESULT_SIZE_VAR_NAME},
						},
						Tok: token.DEFINE,
						Rhs: []ast.Expr{
							&ast.Ident{Name: fmt.Sprintf("len(%s.data)", info.Op1VarName)},
						},
					},
					&ast.AssignStmt{
						Lhs: []ast.Expr{
							&ast.Ident{Name: RESULT_VAR_NAME},
						},
						Tok: token.DEFINE,
						Rhs: []ast.Expr{
							&ast.Ident{Name: fmt.Sprintf("%s.data", info.Op1VarName)},
						},
					},
				}
			}
		}
	} else {
		if info.Op1Scalar {
			if info.Op2Scalar {
				// TYPE2 | SCALAR | SCALAR
				return []ast.Stmt{
					&ast.AssignStmt{
						Lhs: []ast.Expr{
							&ast.Ident{Name: RESULT_VAR_NAME},
						},
						Tok: token.DEFINE,
						Rhs: []ast.Expr{
							&ast.Ident{Name: fmt.Sprintf("%s.data", info.Op2VarName)},
						},
					},
				}
			} else {
				// TYPE2 | SCALAR | SERIES
				return []ast.Stmt{
					&ast.AssignStmt{
						Lhs: []ast.Expr{
							&ast.Ident{Name: RESULT_SIZE_VAR_NAME},
						},
						Tok: token.DEFINE,
						Rhs: []ast.Expr{
							&ast.Ident{Name: fmt.Sprintf("len(%s.data)", info.Op2VarName)},
						},
					},
					&ast.AssignStmt{
						Lhs: []ast.Expr{
							&ast.Ident{Name: RESULT_VAR_NAME},
						},
						Tok: token.DEFINE,
						Rhs: []ast.Expr{
							&ast.Ident{Name: fmt.Sprintf("%s.data", info.Op2VarName)},
						},
					},
				}
			}
		} else {
			if info.Op2Scalar {
				// TYPE2 | SERIES | SCALAR
				return []ast.Stmt{
					&ast.AssignStmt{
						Lhs: []ast.Expr{
							&ast.Ident{Name: RESULT_SIZE_VAR_NAME},
						},
						Tok: token.DEFINE,
						Rhs: []ast.Expr{
							&ast.Ident{Name: fmt.Sprintf("len(%s.data)", info.Op1VarName)},
						},
					},
					&ast.AssignStmt{
						Lhs: []ast.Expr{
							&ast.Ident{Name: RESULT_VAR_NAME},
						},
						Tok: token.DEFINE,
						Rhs: []ast.Expr{
							&ast.CallExpr{
								Fun: &ast.Ident{Name: "make"},
								Args: []ast.Expr{
									&ast.Ident{Name: resInnerType.ToGoType()},
									&ast.Ident{Name: fmt.Sprintf("len(%s.data)", info.Op1VarName)},
								},
							},
						},
					},
				}
			} else {
				// TYPE2 | SERIES | SERIES
				return []ast.Stmt{
					&ast.AssignStmt{
						Lhs: []ast.Expr{
							&ast.Ident{Name: RESULT_SIZE_VAR_NAME},
						},
						Tok: token.DEFINE,
						Rhs: []ast.Expr{
							&ast.Ident{Name: fmt.Sprintf("len(%s.data)", info.Op2VarName)},
						},
					},
					&ast.AssignStmt{
						Lhs: []ast.Expr{
							&ast.Ident{Name: RESULT_VAR_NAME},
						},
						Tok: token.DEFINE,
						Rhs: []ast.Expr{
							&ast.Ident{Name: fmt.Sprintf("%s.data", info.Op2VarName)},
						},
					},
				}
			}
		}
	}
}

// Generate the code to compute the operation
func generateOperationLoop(info BuildInfo) []ast.Stmt {

	statements := make([]ast.Stmt, 0)

	if info.Op1Scalar && info.Op2Scalar {
		statements = append(statements, &ast.ExprStmt{
			info.MakeOperation(RESULT_VAR_NAME, "0", info.Op1VarName, "0", info.Op2VarName, "0"),
		})
	} else {

		op1Index := "i"
		op2Index := "i"
		if info.Op1Scalar {
			op1Index = "0"
		}

		if info.Op2Scalar {
			op2Index = "0"
		}

		statements = append(statements, &ast.ForStmt{
			Init: &ast.AssignStmt{
				Lhs: []ast.Expr{
					&ast.Ident{Name: "i"},
				},
				Tok: token.DEFINE,
				Rhs: []ast.Expr{
					&ast.Ident{Name: "0"},
				},
			},
			Cond: &ast.BinaryExpr{
				X:  &ast.Ident{Name: "i"},
				Op: token.LSS,
				Y:  &ast.Ident{Name: RESULT_SIZE_VAR_NAME},
			},
			Post: &ast.IncDecStmt{
				X:   &ast.Ident{Name: "i"},
				Tok: token.INC,
			},
			Body: &ast.BlockStmt{
				List: []ast.Stmt{
					&ast.ExprStmt{
						info.MakeOperation(RESULT_VAR_NAME, "i", info.Op1VarName, op1Index, info.Op2VarName, op2Index),
					},
				},
			},
		})
	}

	return statements
}

func generateOperation(info BuildInfo) []ast.Stmt {
	resIsNullable := info.Op1Nullable || info.Op2Nullable
	resSeriesType := computeResSeriesType(info.OpCode, info.Op1InnerType, info.Op2InnerType)

	statements := make([]ast.Stmt, 0)

	// 1 - Generate the result inner data array
	statements = append(statements, generateMakeResultStmt(info)...)

	// 2 - Generate the loop to compute the operation
	statements = append(statements, generateOperationLoop(info)...)

	// ? - Generate the return statement with the result series
	statements = append(statements, &ast.ReturnStmt{
		Results: []ast.Expr{
			&ast.CompositeLit{
				Type: &ast.Ident{Name: resSeriesType},
				Elts: []ast.Expr{
					&ast.KeyValueExpr{
						Key:   &ast.Ident{Name: "isNullable"},
						Value: &ast.Ident{Name: fmt.Sprintf("%v", resIsNullable)},
					},
					&ast.KeyValueExpr{
						Key:   &ast.Ident{Name: "name"},
						Value: &ast.Ident{Name: fmt.Sprintf("%s.name", info.Op1VarName)},
					},
					&ast.KeyValueExpr{
						Key:   &ast.Ident{Name: "data"},
						Value: &ast.Ident{Name: RESULT_VAR_NAME},
					},
					&ast.KeyValueExpr{
						Key:   &ast.Ident{Name: "nullMask"},
						Value: &ast.Ident{Name: fmt.Sprintf("%s.nullMask", info.Op1VarName)},
					},
				},
			},
		},
	})

	return statements
}

// Generate the if statement to check the nullability of the operands
func generateNullabilityCheck(info BuildInfo) []ast.Stmt {
	return []ast.Stmt{
		&ast.IfStmt{
			Cond: ast.NewIdent(fmt.Sprintf("%s.isNullable", info.Op1VarName)),
			Body: &ast.BlockStmt{
				List: []ast.Stmt{
					&ast.IfStmt{
						Cond: ast.NewIdent(fmt.Sprintf("%s.isNullable", info.Op2VarName)),
						Body: &ast.BlockStmt{
							List: generateOperation(info.UpdateNullableInfo(true, true)),
						},
						Else: &ast.BlockStmt{
							List: generateOperation(info.UpdateNullableInfo(true, false)),
						},
					},
				},
			},
			Else: &ast.BlockStmt{
				List: []ast.Stmt{
					&ast.IfStmt{
						Cond: ast.NewIdent(fmt.Sprintf("%s.isNullable", info.Op2VarName)),
						Body: &ast.BlockStmt{
							List: generateOperation(info.UpdateNullableInfo(false, true)),
						},
						Else: &ast.BlockStmt{
							List: generateOperation(info.UpdateNullableInfo(false, false)),
						},
					},
				},
			},
		},
	}
}

// Generate the if statement to check the size of the series
func generateSizeCheck(info BuildInfo) ast.Stmt {
	return &ast.IfStmt{
		Cond: &ast.BinaryExpr{
			X:  &ast.Ident{Name: fmt.Sprintf("%s.Len()", info.Op1VarName)},
			Op: token.EQL,
			Y:  &ast.Ident{Name: "1"},
		},

		// CASE OP1_SIZE == 1
		Body: &ast.BlockStmt{
			List: []ast.Stmt{
				&ast.IfStmt{
					Cond: &ast.BinaryExpr{
						X:  &ast.Ident{Name: fmt.Sprintf("%s.Len()", info.Op2VarName)},
						Op: token.EQL,
						Y:  &ast.Ident{Name: "1"},
					},

					// CASE OP1_SIZE == 1 AND OP2_SIZE == 1
					Body: &ast.BlockStmt{
						List: generateNullabilityCheck(info.UpdateScalarInfo(true, true)),
					},

					// CASE OP1_SIZE == 1 AND OP2_SIZE != 1
					Else: &ast.BlockStmt{
						List: generateNullabilityCheck(info.UpdateScalarInfo(true, false)),
					},
				},
			},
		},

		// CASE OP1_SIZE != 1
		Else: &ast.BlockStmt{
			List: []ast.Stmt{
				&ast.IfStmt{
					Cond: &ast.BinaryExpr{
						X:  &ast.Ident{Name: fmt.Sprintf("%s.Len()", info.Op2VarName)},
						Op: token.EQL,
						Y:  &ast.Ident{Name: "1"},
					},

					// CASE OP1_SIZE != 1 AND OP2_SIZE == 1
					Body: &ast.BlockStmt{
						List: generateNullabilityCheck(info.UpdateScalarInfo(false, true)),
					},

					// CASE OP1_SIZE != 1 AND OP2_SIZE != 1
					Else: &ast.BlockStmt{
						List: generateNullabilityCheck(info.UpdateScalarInfo(false, false)),
					},
				},
			},
		},
	}
}

// Generate the switch statement to handle the different types of the second operand
func generateSwitchType(operation Operation, op1SeriesType string, op1InnerType typesys.BaseType, op1VarName, op2VarName string) ast.Stmt {

	bigSwitch := &ast.TypeSwitchStmt{
		Assign: &ast.AssignStmt{
			Lhs: []ast.Expr{ast.NewIdent("o")},
			Tok: token.DEFINE,
			Rhs: []ast.Expr{ast.NewIdent(fmt.Sprintf("%s.(type)", op2VarName))},
		},
		Body: &ast.BlockStmt{
			List: []ast.Stmt{},
		},
	}

	for _, op2 := range operation.ApplyTo {
		bigSwitch.Body.List = append(bigSwitch.Body.List,
			&ast.CaseClause{
				List: []ast.Expr{ast.NewIdent(op2.SeriesType)},
				Body: []ast.Stmt{
					generateSizeCheck(BuildInfo{
						OpCode:        operation.OpCode,
						Op1VarName:    op1VarName,
						Op1SeriesType: op1SeriesType,
						Op1InnerType:  op1InnerType,
						Op2VarName:    "o",
						Op2SeriesType: op2.SeriesType,
						Op2InnerType:  op2.InnerType,
						MakeOperation: op2.MakeOperation,
					}),
				},
			},
		)
	}
	return bigSwitch
}

func computeResSeriesType(opCode typesys.OPCODE, op1, op2 typesys.BaseType) string {
	switch computeResInnerType(opCode, op1, op2) {
	case typesys.BoolType:
		return "SeriesBool"
	case typesys.Int16Type:
		return "SeriesInt16"
	case typesys.Int32Type:
		return "SeriesInt32"
	case typesys.Int64Type:
		return "SeriesInt64"
	case typesys.Float32Type:
		return "SeriesFloat32"
	case typesys.Float64Type:
		return "SeriesFloat64"
	case typesys.StringType:
		return "SeriesString"
	}
	return "SeriesError"
}

func computeResInnerType(opCode typesys.OPCODE, op1, op2 typesys.BaseType) typesys.BaseType {
	return opCode.GetBinaryOpResultType(typesys.Primitive{Base: op1}, typesys.Primitive{Base: op2}).Base
}

func generateOperations() {
	for filename, info := range DATA {

		src, err := ioutil.ReadFile(filepath.Join("..", filename))
		if err != nil {
			panic(err)
		}

		// Parse the file.
		fset := token.NewFileSet()
		fast, err := parser.ParseFile(fset, filepath.Join("..", filename), src, parser.ParseComments)
		if err != nil {
			panic(err)
		}

		for i, decl := range fast.Decls {
			if funcDecl, ok := decl.(*ast.FuncDecl); ok {
				switch funcDecl.Name.Name {
				case "Mul":
					fast.Decls[i].(*ast.FuncDecl).Body.List = []ast.Stmt{
						generateSwitchType(info.Operations["Mul"], info.SeriesType, info.InnerType, "s", "other"),

						// default: return GDLErrorSeries
						&ast.ReturnStmt{
							Results: []ast.Expr{ast.NewIdent(fmt.Sprintf(FINAL_RETURN_FMT, "multiply"))},
						},
					}

				case "Div":
					fast.Decls[i].(*ast.FuncDecl).Body.List = []ast.Stmt{
						generateSwitchType(info.Operations["Div"], info.SeriesType, info.InnerType, "s", "other"),

						// default: return GDLErrorSeries
						&ast.ReturnStmt{
							Results: []ast.Expr{ast.NewIdent(fmt.Sprintf(FINAL_RETURN_FMT, "divide"))},
						},
					}

				case "Mod":
					fast.Decls[i].(*ast.FuncDecl).Body.List = []ast.Stmt{
						// generateSwitchType(info.Operations["Mod"], info.SeriesType, info.InnerType, "s", "other"),

						// default: return GDLErrorSeries
						&ast.ReturnStmt{
							Results: []ast.Expr{ast.NewIdent(fmt.Sprintf(FINAL_RETURN_FMT, "use modulo"))},
						},
					}

				case "Add":
					fast.Decls[i].(*ast.FuncDecl).Body.List = []ast.Stmt{
						generateSwitchType(info.Operations["Add"], info.SeriesType, info.InnerType, "s", "other"),

						// default: return GDLErrorSeries
						&ast.ReturnStmt{
							Results: []ast.Expr{ast.NewIdent(fmt.Sprintf(FINAL_RETURN_FMT, "sum"))},
						},
					}

				case "Sub":
					fast.Decls[i].(*ast.FuncDecl).Body.List = []ast.Stmt{
						generateSwitchType(info.Operations["Sub"], info.SeriesType, info.InnerType, "s", "other"),

						// default: return GDLErrorSeries
						&ast.ReturnStmt{
							Results: []ast.Expr{ast.NewIdent(fmt.Sprintf(FINAL_RETURN_FMT, "subtract"))},
						},
					}
				}
			}

			buf := new(bytes.Buffer)
			err = format.Node(buf, fset, fast)
			if err != nil {
				panic(err)
			}

			err = ioutil.WriteFile(filepath.Join("..", filename), buf.Bytes(), 0644)
			if err != nil {
				panic(err)
			}
		}
	}
}

func main() {
	generateOperations()
}
