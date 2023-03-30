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
)

const OP1_SIZE = "op1.size"
const OP2_SIZE = "op2.size"

const VECTORIZATION = 4

var data = []struct{ Operation, Op1Type, Op2Type, ResultType string }{
	{"*", "int", "int", "int"},
	{"*", "int", "float64", "float64"},
}

type BuildInfo struct {
	Op1Nullable   bool
	Op1Scalar     bool
	Op2Nullable   bool
	Op2Scalar     bool
	Vectorization int
	Op1           string
	Op1Type       string
	Op1InnerType  string
	Op2           string
	Op2Type       string
	Op2InnerType  string
	Res           string
	ResType       string
	ResInnerType  string
	Operation     func(op1, op2, res, index string) ast.Stmt
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
	}

	// statements = append(statements, &ast.AssignStmt{
	// 	Lhs: []ast.Expr{&ast.Ident{Name: "result"}},
	// 	Tok: token.DEFINE,
	// 	Rhs: []ast.Expr{&ast.Ident{Name: "make([]" + result + ", " + op1 + ".size)"}},
	// })

	if info.Vectorization == 1 {
		upperBound = &ast.Ident{Name: info.Op1 + ".Len()"}
		increment = &ast.ExprStmt{
			X: &ast.UnaryExpr{
				X:  &ast.Ident{Name: "i"},
				Op: token.INC,
			}}

		operations = append(operations, info.Operation(info.Op1, info.Op2, info.Res, "i"))

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
			Rhs: []ast.Expr{&ast.Ident{Name: info.Op1 + ".Len() / " + fmt.Sprint(info.Vectorization) + " * " + fmt.Sprint(info.Vectorization)}},
		})

		upperBound = &ast.Ident{Name: "upperBound"}

		increment = &ast.AssignStmt{
			Lhs: []ast.Expr{&ast.Ident{Name: "i"}},
			Tok: token.ADD_ASSIGN,
			Rhs: []ast.Expr{&ast.Ident{Name: fmt.Sprint(info.Vectorization)}},
		}

		operations = make([]ast.Stmt, 0)
		for i := 0; i < info.Vectorization; i++ {
			operations = append(operations, info.Operation(info.Op1, info.Op2, info.Res, fmt.Sprint(i)))
		}

		mainLoop = &ast.IfStmt{
			Cond: &ast.BinaryExpr{
				X:  &ast.Ident{Name: info.Op1 + ".Len()"},
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
				Y:  &ast.Ident{Name: info.Op1 + ".Len()"},
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
								X:  &ast.Ident{Name: info.Op1 + ".data[i]"},
								Op: token.MUL,
								Y:  &ast.Ident{Name: info.Op2 + ".data[i]"},
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
						Cond: ast.NewIdent(fmt.Sprintf("%s.isNullable", info.Op2)),
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
						X:  &ast.Ident{Name: fmt.Sprintf("%s.Len()", info.Op2)},
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
						List: generateNullabilityCheck(BuildInfo{Op1: info.Op1, Op2: info.Op2, Res: info.Res, Vectorization: info.Vectorization,
							Op1Nullable: info.Op1Nullable, Op2Nullable: info.Op2Nullable, Op1Scalar: false, Op2Scalar: false,
						}),
					},
				},
			},
		},
	}
}

func generateSwitchType(
	op1, op1Type, op1InnerType string,
	op2 string, op2Tpes, op2InnerTypes []string,
	vectorization int,
) ast.Stmt {

	bigSwitch := &ast.TypeSwitchStmt{
		Assign: &ast.AssignStmt{
			Lhs: []ast.Expr{ast.NewIdent("o")},
			Tok: token.DEFINE,
			Rhs: []ast.Expr{ast.NewIdent(fmt.Sprintf("%s.(type)", op2))},
		},
		Body: &ast.BlockStmt{
			List: []ast.Stmt{},
		},
	}

	for i, op2Type := range op2Tpes {
		bigSwitch.Body.List = append(bigSwitch.Body.List,
			&ast.CaseClause{
				List: []ast.Expr{ast.NewIdent(op2Type)},
				Body: []ast.Stmt{
					generateSizeCheck(BuildInfo{
						Op1:           op1,
						Op1Type:       op1Type,
						Op1InnerType:  op1InnerType,
						Op2:           "o",
						Op2Type:       op2Type,
						Op2InnerType:  op2InnerTypes[i],
						Vectorization: vectorization,
						ResType:       computeResType(op1Type, op2Type),
					}),
				},
			},
		)
	}
	return bigSwitch
}

func computeResType(op1Type, op2Type string) string {
	return ""
}

func main() {

	filenames := []string{
		"gdl_series_bool.go",
		"gdl_series_float64.go",
		"gdl_series_int32.go",
		"gdl_series_string.go",
	}

	for _, filename := range filenames {
		fast, err := parser.ParseFile(token.NewFileSet(), filepath.Join("..", filename), nil, parser.ParseComments)
		if err != nil {
			panic(err)
		}

		for i, decl := range fast.Decls {
			if funcDecl, ok := decl.(*ast.FuncDecl); ok {
				switch funcDecl.Name.Name {
				case "Mul":
					fmt.Println(funcDecl.Body.List)
					fast.Decls[i].(*ast.FuncDecl).Body.List = []ast.Stmt{
						generateSwitchType("s", "GDLSeriesInt32", "[]int",
							"other", []string{"GDLSeriesInt32", "GDLSeriesFloat64"}, []string{"[]int", "[]float64"}, 2),
					}
				case "Div":
					fmt.Println(funcDecl.Body.List)
				case "Mod":
					fmt.Println(funcDecl.Body.List)
				case "Add":
					fmt.Println(funcDecl.Body.List)
				case "Sub":
					fmt.Println(funcDecl.Body.List)
				}
			}

			buf := new(bytes.Buffer)
			err = format.Node(buf, token.NewFileSet(), fast)
			if err != nil {
				panic(err)
			}

			err = ioutil.WriteFile(filepath.Join("..", filename), buf.Bytes(), 0644)
			if err != nil {
				panic(err)
			}
		}
	}

	// fmt.Println(data)

	// res := generateSwitchType(
	// 	"s", "GDLSeriesInt32", "[]int",
	// 	"other", []string{"GDLSeriesInt32", "GDLSeriesFloat64"}, []string{"[]int", "[]float64"}, 2)

	// outputFileSet := token.NewFileSet()
	// outputGoFile := &ast.File{
	// 	Name: ast.NewIdent("testgen"),
	// 	Decls: []ast.Decl{
	// 		&ast.FuncDecl{
	// 			Name: ast.NewIdent("main"),
	// 			Type: &ast.FuncType{
	// 				Params:  &ast.FieldList{},
	// 				Results: &ast.FieldList{},
	// 			},

	// 			Body: &ast.BlockStmt{
	// 				List: []ast.Stmt{
	// 					res,
	// 				},
	// 			},
	// 		},
	// 	},
	// }

	// buf := new(bytes.Buffer)
	// err := format.Node(buf, outputFileSet, outputGoFile)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(buf.String())
}
