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
	VECTORIZATION    = 4
	FINAL_RETURN_FMT = "GDLSeriesError{fmt.Sprintf(\"Cannot %s %%s and %%s\", s.Type().ToString(), other.Type().ToString())}"
)

type BuildInfo struct {
	Op1Nullable   bool
	Op1Scalar     bool
	Op2Nullable   bool
	Op2Scalar     bool
	Vectorization int
	Op1VarName    string
	Op1SeriesType string
	Op1InnerType  typesys.BaseType
	Op2VarName    string
	Op2SeriesType string
	Op2InnerType  typesys.BaseType
	ResVarName    string
	ResSeriesType string
	ResInnerType  typesys.BaseType
	Operation     func(op1, op2, res, index string) ast.Stmt
}

func (bi BuildInfo) UpdateScalarInfo(Op1Scalar, Op2Scalar bool) BuildInfo {
	bi.Op1Scalar = Op1Scalar
	bi.Op2Scalar = Op2Scalar
	return bi
}

func generateOperation(info BuildInfo) []ast.Stmt {
	var upperBound *ast.Ident
	var mainLoop ast.Stmt
	var increment ast.Stmt
	var operations []ast.Stmt
	var endStmt ast.Stmt

	statements := make([]ast.Stmt, 0)

	/////////////////////////////////////////////////////////////////////////////////////
	//							SCALAR * SCALAR
	if info.Op1Scalar && info.Op2Scalar {
		if info.Op1Nullable {
			statements = append(statements, &ast.IfStmt{
				Cond: &ast.BinaryExpr{
					X:  &ast.Ident{Name: info.Op1VarName + ".Valid"},
					Op: token.EQL,
					Y:  &ast.Ident{Name: "false"},
				},
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.ReturnStmt{
							Results: []ast.Expr{
								&ast.Ident{Name: "result"},
							},
						},
					},
				},
			})
		}
	}

	// statements = append(statements, &ast.AssignStmt{
	// 	Lhs: []ast.Expr{&ast.Ident{Name: "result"}},
	// 	Tok: token.DEFINE,
	// 	Rhs: []ast.Expr{&ast.Ident{Name: "make([]" + result + ", " + op1 + ".size)"}},
	// })

	if info.Vectorization == 1 {
		upperBound = &ast.Ident{Name: info.Op1VarName + ".Len()"}
		increment = &ast.ExprStmt{
			X: &ast.UnaryExpr{
				X:  &ast.Ident{Name: "i"},
				Op: token.INC,
			}}

		operations = append(operations, info.Operation(info.Op1VarName, info.Op2VarName, info.ResVarName, "i"))

		mainLoop = &ast.ForStmt{
			Init: &ast.AssignStmt{
				Lhs: []ast.Expr{&ast.Ident{Name: "i"}},
				Tok: token.DEFINE,
				Rhs: []ast.Expr{ast.NewIdent("0")},
			},
			Cond: &ast.BinaryExpr{
				X:  &ast.Ident{Name: "i"},
				Op: token.LSS,
				Y:  upperBound,
			},
			Post: increment,
			Body: &ast.BlockStmt{
				List: operations,
			},
		}

		endStmt = &ast.EmptyStmt{}
	} else {
		statements = append(statements, &ast.AssignStmt{
			Lhs: []ast.Expr{&ast.Ident{Name: "upperBound"}},
			Tok: token.DEFINE,
			Rhs: []ast.Expr{&ast.Ident{Name: info.Op1VarName + ".Len() / " + fmt.Sprint(info.Vectorization) + " * " + fmt.Sprint(info.Vectorization)}},
		})

		upperBound = &ast.Ident{Name: "upperBound"}

		increment = &ast.AssignStmt{
			Lhs: []ast.Expr{&ast.Ident{Name: "i"}},
			Tok: token.ADD_ASSIGN,
			Rhs: []ast.Expr{&ast.Ident{Name: fmt.Sprint(info.Vectorization)}},
		}

		operations = make([]ast.Stmt, 0)
		for i := 0; i < info.Vectorization; i++ {
			operations = append(operations, info.Operation(info.Op1VarName, info.Op2VarName, info.ResVarName, fmt.Sprint(i)))
		}

		mainLoop = &ast.IfStmt{
			Cond: &ast.BinaryExpr{
				X:  &ast.Ident{Name: info.Op1VarName + ".Len()"},
				Op: token.GTR,
				Y:  &ast.Ident{Name: "upperBound"},
			},
			Body: &ast.BlockStmt{
				List: []ast.Stmt{
					&ast.ForStmt{
						Init: &ast.AssignStmt{
							Lhs: []ast.Expr{&ast.Ident{Name: "i"}},
							Tok: token.DEFINE,
							Rhs: []ast.Expr{ast.NewIdent("0")},
						},
						Cond: &ast.BinaryExpr{
							X:  &ast.Ident{Name: "i"},
							Op: token.LSS,
							Y:  upperBound,
						},
						Post: increment,
						Body: &ast.BlockStmt{
							List: operations,
						},
					},
				},
			},
		}

		endStmt = &ast.ForStmt{
			Init: &ast.AssignStmt{
				Lhs: []ast.Expr{&ast.Ident{Name: "i"}},
				Tok: token.DEFINE,
				Rhs: []ast.Expr{&ast.Ident{Name: "upperBound"}},
			},
			Cond: &ast.BinaryExpr{
				X:  &ast.Ident{Name: "i"},
				Op: token.LSS,
				Y:  &ast.Ident{Name: info.Op1VarName + ".Len()"},
			},
			Post: &ast.ExprStmt{
				X: &ast.UnaryExpr{
					X:  &ast.Ident{Name: "i"},
					Op: token.INC,
				},
			},
			Body: &ast.BlockStmt{
				List: []ast.Stmt{
					&ast.AssignStmt{
						Lhs: []ast.Expr{&ast.Ident{Name: "result[i]"}},
						Tok: token.ASSIGN,
						Rhs: []ast.Expr{
							&ast.BinaryExpr{
								X:  &ast.Ident{Name: info.Op1VarName + ".data[i]"},
								Op: token.MUL,
								Y:  &ast.Ident{Name: info.Op2VarName + ".data[i]"},
							},
						},
					},
				},
			},
		}
	}

	return append(statements,
		mainLoop,
		endStmt,
	)
}

func generateNullabilityCheck(info BuildInfo) []ast.Stmt {
	return []ast.Stmt{
		&ast.IfStmt{
			Cond: ast.NewIdent(fmt.Sprintf("%s.isNullable", info.Op1VarName)),
			Body: &ast.BlockStmt{
				List: []ast.Stmt{
					&ast.IfStmt{
						Cond: ast.NewIdent(fmt.Sprintf("%s.isNullable", info.Op2VarName)),
						Body: &ast.BlockStmt{
							// 	List: generateOperation(BuildInfo{Op1: info.Op1, Op2: info.Op2, Res: info.Res, Vectorization: info.Vectorization,
							// 		Op1Nullable: true, Op2Nullable: true, Op1Scalar: info.Op1Scalar, Op2Scalar: info.Op2Scalar,
							// 	}),
						},
						Else: &ast.BlockStmt{
							// 	List: generateOperation(BuildInfo{Op1: info.Op1, Op2: info.Op2, Res: info.Res, Vectorization: info.Vectorization,
							// 		Op1Nullable: true, Op2Nullable: false, Op1Scalar: info.Op1Scalar, Op2Scalar: info.Op2Scalar,
							// 	}),
						},
					},
				},
			},
			Else: &ast.BlockStmt{
				List: []ast.Stmt{
					&ast.IfStmt{
						Cond: ast.NewIdent(fmt.Sprintf("%s.isNullable", info.Op2VarName)),
						Body: &ast.BlockStmt{
							// 	List: generateOperation(BuildInfo{Op1: info.Op1, Op2: info.Op2, Res: info.Res, Vectorization: info.Vectorization,
							// 		Op1Nullable: false, Op2Nullable: true, Op1Scalar: info.Op1Scalar, Op2Scalar: info.Op2Scalar,
							// 	}),
						},
						Else: &ast.BlockStmt{
							// 	List: generateOperation(BuildInfo{Op1: info.Op1, Op2: info.Op2, Res: info.Res, Vectorization: info.Vectorization,
							// 		Op1Nullable: false, Op2Nullable: false, Op1Scalar: info.Op1Scalar, Op2Scalar: info.Op2Scalar,
							// 	}),
						},
					},
				},
			},
		},
	}
}

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

func generateSwitchType(operation Operation, op1SeriesType string, op1InnerType typesys.BaseType, op1VarName, op2VarName string, vectorization int) ast.Stmt {

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
						Op1VarName:    op1VarName,
						Op1SeriesType: op1SeriesType,
						Op1InnerType:  op1InnerType,
						Op2VarName:    "o",
						Op2SeriesType: op2.SeriesType,
						Op2InnerType:  op2.InnerType,
						Vectorization: vectorization,
						ResSeriesType: computeResSeriesType(operation.OpCode, op1InnerType, op2.InnerType),
						ResInnerType:  computeResInnerType(operation.OpCode, op1InnerType, op2.InnerType),
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
		return "GDLSeriesBool"
	case typesys.Int16Type:
		return "GDLSeriesInt16"
	case typesys.Int32Type:
		return "GDLSeriesInt32"
	case typesys.Int64Type:
		return "GDLSeriesInt64"
	case typesys.Float32Type:
		return "GDLSeriesFloat32"
	case typesys.Float64Type:
		return "GDLSeriesFloat64"
	case typesys.StringType:
		return "GDLSeriesString"
	}
	return "GDLSeriesError"
}

func computeResInnerType(opCode typesys.OPCODE, op1, op2 typesys.BaseType) typesys.BaseType {
	return opCode.GetBinaryOpResultType(typesys.Primitive{Base: op1}, typesys.Primitive{Base: op1}).Base
}

func main() {

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
						generateSwitchType(info.Operations["Mul"], info.SeriesType, info.InnerType, "s", "other", VECTORIZATION),

						// default: return GDLErrorSeries
						&ast.ReturnStmt{
							Results: []ast.Expr{ast.NewIdent(fmt.Sprintf(FINAL_RETURN_FMT, "multiply"))},
						},
					}

				case "Div":
					fast.Decls[i].(*ast.FuncDecl).Body.List = []ast.Stmt{
						// generateSwitchType("s", "GDLSeriesInt32", "[]int",
						// 	"other", []string{"GDLSeriesInt32", "GDLSeriesFloat64"}, []string{"[]int", "[]float64"}, VECTORIZATION),

						// default: return GDLErrorSeries
						&ast.ReturnStmt{
							Results: []ast.Expr{ast.NewIdent(fmt.Sprintf(FINAL_RETURN_FMT, "divide"))},
						},
					}

				case "Mod":
					fast.Decls[i].(*ast.FuncDecl).Body.List = []ast.Stmt{
						// generateSwitchType("s", "GDLSeriesInt32", "[]int",
						// 	"other", []string{"GDLSeriesInt32", "GDLSeriesFloat64"}, []string{"[]int", "[]float64"}, VECTORIZATION),

						// default: return GDLErrorSeries
						&ast.ReturnStmt{
							Results: []ast.Expr{ast.NewIdent(fmt.Sprintf(FINAL_RETURN_FMT, "use modulo"))},
						},
					}

				case "Add":
					fast.Decls[i].(*ast.FuncDecl).Body.List = []ast.Stmt{
						// generateSwitchType("s", "GDLSeriesInt32", "[]int",
						// 	"other", []string{"GDLSeriesInt32", "GDLSeriesFloat64"}, []string{"[]int", "[]float64"}, VECTORIZATION),

						// default: return GDLErrorSeries
						&ast.ReturnStmt{
							Results: []ast.Expr{ast.NewIdent(fmt.Sprintf(FINAL_RETURN_FMT, "sum"))},
						},
					}

				case "Sub":
					fast.Decls[i].(*ast.FuncDecl).Body.List = []ast.Stmt{
						// generateSwitchType("s", "GDLSeriesInt32", "[]int",
						// 	"other", []string{"GDLSeriesInt32", "GDLSeriesFloat64"}, []string{"[]int", "[]float64"}, VECTORIZATION),

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
