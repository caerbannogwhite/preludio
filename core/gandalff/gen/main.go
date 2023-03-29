package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/token"
)

const OP1_SIZE = "op1.size"
const OP2_SIZE = "op2.size"

const VECTORIZATION = 4

var data = []struct{ Operation, Op1Type, Op2Type, ResultType string }{
	{"*", "int", "int", "int"},
	{"*", "int", "float64", "float64"},
}

type BuildInfo struct {
	Operation     func(op1, op2, res, index string) ast.Stmt
	Op1           string
	Op2           string
	Res           string
	Op1Nullable   bool
	Op2Nullable   bool
	Op1Scalar     bool
	Op2Scalar     bool
	Vectorization int
}

func generateOperation(info BuildInfo) []ast.Stmt {
	var upperBound *ast.Ident
	var mainLoop ast.Stmt
	var increment ast.Stmt
	var operations []ast.Stmt
	var endStmt ast.Stmt

	statements := make([]ast.Stmt, 0)

	if info.Op1Scalar && info.Op2Scalar {	// scalar * scalar
		if info.Op1Nullable {
			statements = append(statements, &ast.IfStmt{
				Cond: &ast.BinaryExpr{
					X:  &ast.Ident{Name: info.Op1 + ".Valid"},
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

	statements = append(statements, &ast.AssignStmt{
		Lhs: []ast.Expr{&ast.Ident{Name: "result"}},
		Tok: token.DEFINE,
		Rhs: []ast.Expr{&ast.Ident{Name: "make([]" + result + ", " + op1 + ".size)"}},
	})

	if vectorization == 1 {
		upperBound = &ast.Ident{Name: op1 + ".size"}
		increment = &ast.ExprStmt{
			X: &ast.UnaryExpr{
				X:  &ast.Ident{Name: "i"},
				Op: token.INC,
			}}

		operations = info.Operation(info.Op1, info.Op2, info.Res, "i")

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
			Rhs: []ast.Expr{&ast.Ident{Name: op1 + ".size / " + fmt.Sprint(vectorization) + " * " + fmt.Sprint(vectorization)}},
		})

		upperBound = &ast.Ident{Name: "upperBound"}

		increment = &ast.AssignStmt{
			Lhs: []ast.Expr{&ast.Ident{Name: "i"}},
			Tok: token.ADD_ASSIGN,
			Rhs: []ast.Expr{&ast.Ident{Name: fmt.Sprint(vectorization)}},
		}

		operations = make([]ast.Stmt, 0)
		for i := 0; i < vectorization; i++ {
			operations = append(operations, info.Operation(info.Op1, info.Op2, info.Res, fmt.Sprint(i)))
		}

		mainLoop = &ast.IfStmt{
			Cond: &ast.BinaryExpr{
				X:  &ast.Ident{Name: op1 + ".size"},
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
				Y:  &ast.Ident{Name: op1 + ".size"},
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
								X:  &ast.Ident{Name: op1 + ".data[i]"},
								Op: token.MUL,
								Y:  &ast.Ident{Name: op2 + ".data[i]"},
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
			Cond: ast.NewIdent(fmt.Sprintf("%s.isNullable", info.Op1)),
			Body: &ast.BlockStmt{
				List: []ast.Stmt{
					&ast.IfStmt{
						Cond: ast.NewIdent(fmt.Sprintf("%s.isNullable", info.Op2)),
						Body: &ast.BlockStmt{
							List: generateOperation(BuildInfo{Op1: info.Op1, Op2: info.Op2, Res: info.Res, Vectorization: info.Vectorization,
								Op1Nullable: true, Op2Nullable: true, Op1Scalar: info.Op1Scalar, Op2Scalar: info.Op2Scalar,
							}),
						},
						Else: &ast.BlockStmt{
							List: generateOperation(BuildInfo{Op1: info.Op1, Op2: info.Op2, Res: info.Res, Vectorization: info.Vectorization,
								Op1Nullable: true, Op2Nullable: false, Op1Scalar: info.Op1Scalar, Op2Scalar: info.Op2Scalar,
							}),
						},
					},
				},
			},
			Else: &ast.BlockStmt{
				List: []ast.Stmt{
					&ast.IfStmt{
						Cond: ast.NewIdent(fmt.Sprintf("%s.isNullable", info.Op2)),
						Body: &ast.BlockStmt{
							List: generateOperation(BuildInfo{Op1: info.Op1, Op2: info.Op2, Res: info.Res, Vectorization: info.Vectorization,
								Op1Nullable: false, Op2Nullable: true, Op1Scalar: info.Op1Scalar, Op2Scalar: info.Op2Scalar,
							}),
						},
						Else: &ast.BlockStmt{
							List: generateOperation(BuildInfo{Op1: info.Op1, Op2: info.Op2, Res: info.Res, Vectorization: info.Vectorization,
								Op1Nullable: false, Op2Nullable: false, Op1Scalar: info.Op1Scalar, Op2Scalar: info.Op2Scalar,
							}),
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
			X:  &ast.Ident{Name: fmt.Sprintf("%s.Len()", info.Op1)},
			Op: token.EQL,
			Y:  &ast.Ident{Name: "1"},
		},

		// CASE OP1_SIZE == 1
		Body: &ast.BlockStmt{
			List: []ast.Stmt{
				&ast.IfStmt{
					Cond: &ast.BinaryExpr{
						X:  &ast.Ident{Name: fmt.Sprintf("%s.Len()", info.Op2)},
						Op: token.EQL,
						Y:  &ast.Ident{Name: "1"},
					},

					// CASE OP1_SIZE == 1 AND OP2_SIZE == 1
					Body: &ast.BlockStmt{
						List: generateNullabilityCheck(BuildInfo{Op1: info.Op1, Op2: info.Op2, Res: info.Res, Vectorization: info.Vectorization,
							Op1Nullable: info.Op1Nullable, Op2Nullable: info.Op2Nullable, Op1Scalar: true, Op2Scalar: true,
						}),
					},

					// CASE OP1_SIZE == 1 AND OP2_SIZE != 1
					Else: &ast.BlockStmt{
						List: generateNullabilityCheck(BuildInfo{Op1: info.Op1, Op2: info.Op2, Res: info.Res, Vectorization: info.Vectorization,
							Op1Nullable: info.Op1Nullable, Op2Nullable: info.Op2Nullable, Op1Scalar: true, Op2Scalar: false,
						}),
					},
				},
			},
		},

		// CASE OP1_SIZE != 1
		Else: &ast.BlockStmt{
			List: []ast.Stmt{
				&ast.IfStmt{
					Cond: &ast.BinaryExpr{
						X:  &ast.Ident{Name: OP2_SIZE},
						Op: token.EQL,
						Y:  &ast.Ident{Name: "1"},
					},

					// CASE OP1_SIZE != 1 AND OP2_SIZE == 1
					Body: &ast.BlockStmt{
						List: generateNullabilityCheck(BuildInfo{Op1: info.Op1, Op2: info.Op2, Res: info.Res, Vectorization: info.Vectorization,
							Op1Nullable: info.Op1Nullable, Op2Nullable: info.Op2Nullable, Op1Scalar: false, Op2Scalar: true,
						}),
					},

					// CASE OP1_SIZE != 1 AND OP2_SIZE != 1
					Else: &ast.BlockStmt{
						List: generateOperation(BuildInfo{Op1: info.Op1, Op2: info.Op2, Res: info.Res, Vectorization: info.Vectorization,
							Op1Nullable: info.Op1Nullable, Op2Nullable: info.Op2Nullable, Op1Scalar: false, Op2Scalar: false,
						}),
					},
				},
			},
		},
	}
}

func generateSwitchType(op1, op2 string) ast.Stmt {
	types1 := []string{"[]bool", "[]int", "[]float64", "[]string"}
	types2 := []string{"[]bool", "[]int", "[]float64", "[]string"}

	bigSwitch := &ast.SwitchStmt{
		Body: &ast.BlockStmt{
			List: []ast.Stmt{},
		},
	}

	for _, t1 := range types1 {
		for _, t2 := range types2 {
			bigSwitch.Body.List = append(bigSwitch.Body.List, &ast.CaseClause{
				List: []ast.Expr{ast.NewIdent(t1)},
				Body: []ast.Stmt{
					&ast.SwitchStmt{
						Tag: ast.NewIdent(op2),
						Body: &ast.BlockStmt{
							List: []ast.Stmt{
								&ast.CaseClause{
									List: []ast.Expr{ast.NewIdent(t2)},
									Body: []ast.Stmt{
										&ast.AssignStmt{
											Lhs: []ast.Expr{ast.NewIdent("result")},
											Tok: token.DEFINE,
											Rhs: []ast.Expr{ast.NewIdent("make([]int, len(op1))")},
										},
										generateSizeCheck(op1, op2),
									},
								},
							},
						},
					},
				},
			})
		}
	}
	return bigSwitch
}

func main() {

	fmt.Println(data)

	res := generateSwitchType("op1", "op2")

	outputFileSet := token.NewFileSet()
	outputGoFile := &ast.File{
		Name: ast.NewIdent("testgen"),
		Decls: []ast.Decl{
			&ast.FuncDecl{
				Name: ast.NewIdent("main"),
				Type: &ast.FuncType{
					Params:  &ast.FieldList{},
					Results: &ast.FieldList{},
				},

				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						res,
					},
				},
			},
		},
	}

	buf := new(bytes.Buffer)
	err := format.Node(buf, outputFileSet, outputGoFile)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(buf.String())

	// mult := ast.SwitchStmt{
	// 	Body: ast.BlockStmt([]ast.CaseClause{ast.CaseClause{}}),
	// }

}
