package preludiocore

import (
	"fmt"
	"preludiocompiler"
	"strconv"
	"strings"
)

func solveExpr(vm *ByteEater, i *__p_intern__) error {
	// TODO: check if this is possible and
	// if it's the case to raise an error
	if i == nil || i.expr == nil || len(i.expr) == 0 {
		return fmt.Errorf("invalid expression")
	}

	// Check if the expression is:
	//  - a symbol: resolve it
	//  - a list: recursively solve all the expressions
	if len(i.expr) == 1 {
		switch l := i.expr[0].(type) {
		case __p_symbol__:
			i.expr[0] = vm.symbolResolution(l)

		case __p_list__:
			for idx := range l {
				if err := solveExpr(vm, &l[idx]); err != nil {
					return err
				}
			}

		default:
		}
	}

	stack := make([]interface{}, 0)

	for len(i.expr) > 1 {

		// Load the stack until we find an operators
		var ok bool
		var op preludiocompiler.OPCODE
		for {
			if op, ok = isOperator(i.expr[0]); ok {
				i.expr = i.expr[1:len(i.expr)]
				break
			}

			stack = append(stack, i.expr[0])
			i.expr = i.expr[1:len(i.expr)]
		}

		var result interface{}

		// UNARY
		// if op, ok := isOperator(t2); ok {
		// 	i.expr = i.expr[2:len(i.expr)]

		// 	if s, ok := t1.(__p_symbol__); ok {
		// 		t1 = vm.symbolResolution(s)
		// 	}

		// 	switch op {
		// 	case preludiocompiler.OP_UNARY_ADD:
		// 	case preludiocompiler.OP_UNARY_SUB:
		// 	case preludiocompiler.OP_UNARY_NOT:
		// 	}
		// } else

		// BINARY
		{
			t2 := stack[len(stack)-1]
			t1 := stack[len(stack)-2]
			stack = stack[0 : len(stack)-2]

			// Symbo resolution
			if s, ok := t1.(__p_symbol__); ok {
				t1 = vm.symbolResolution(s)
			}

			if s, ok := t2.(__p_symbol__); ok {
				t2 = vm.symbolResolution(s)
			}

			switch op {

			///////////////////////////////////////////////////////////////////
			////////					MULTIPLICATION

			case preludiocompiler.OP_BINARY_MUL:
				switch val1 := t1.(type) {
				case []bool:
					switch val2 := t2.(type) {
					case []bool:
						var res []int
						if len(val1) == 1 {
							res = make([]int, len(val2))
							for j, b := range val2 {
								res[j] = BoolToInt(val1[0]) * BoolToInt(b)
							}
						} else if len(val2) == 1 {
							res = make([]int, len(val1))
							for j, b := range val1 {
								res[j] = BoolToInt(b) * BoolToInt(val2[0])
							}
						} else if len(val1) == len(val2) {
							res = make([]int, len(val2))
							for j, b := range val2 {
								res[j] = BoolToInt(val1[j]) * BoolToInt(b)
							}
						}
						result = res
					case []int:
						var res []int
						if len(val1) == 1 {
							if val1[0] {
								res = val2
							} else {
								res = make([]int, len(val2))
								for j := range val2 {
									res[j] = 0
								}
							}
						} else if len(val2) == 1 {
							res = make([]int, len(val1))
							for j, b := range val1 {
								res[j] = BoolToInt(b) * val2[0]
							}
						} else if len(val1) == len(val2) {
							res = make([]int, len(val2))
							for j, n := range val2 {
								res[j] = BoolToInt(val1[j]) * n
							}
						}
						result = res
					case []float64:
						var res []float64
						if len(val1) == 1 {
							if val1[0] {
								res = val2
							} else {
								res = make([]float64, len(val2))
								for j := range val2 {
									res[j] = 0.0
								}
							}
						} else if len(val2) == 1 {
							res = make([]float64, len(val1))
							for j, b := range val1 {
								res[j] = BoolToFloat64(b) * val2[0]
							}
						} else if len(val1) == len(val2) {
							res = make([]float64, len(val2))
							for j, n := range val2 {
								res[j] = BoolToFloat64(val1[j]) * n
							}
						}
						result = res
					case []string:
						var res []string
						if len(val1) == 1 {
							if val1[0] {
								res = val2
							} else {
								res = make([]string, len(val2))
								for j := range val2 {
									res[j] = ""
								}
							}
						} else if len(val2) == 1 {
							res = make([]string, len(val1))
							for j, b := range val1 {
								if b {
									res[j] = val2[j]
								} else {
									res[j] = ""
								}
							}
						} else if len(val1) == len(val2) {
							res = make([]string, len(val2))
							for j, b := range val1 {
								if b {
									res[j] = val2[j]
								} else {
									res[j] = ""
								}
							}
						}
						result = res

					default:
						return fmt.Errorf("binary \"*\" operator not implemented for %s and %s",
							preludiocompiler.GoToPreludioTypeString(val1),
							preludiocompiler.GoToPreludioTypeString(val2))
					}

				case []int:
					switch val2 := t2.(type) {
					case []bool:
						var res []int
						if len(val1) == 1 {
							res = make([]int, len(val2))
							for j, b := range val2 {
								res[j] = val1[0] * BoolToInt(b)
							}
						} else if len(val2) == 1 {
							res = make([]int, len(val1))
							for j, n := range val1 {
								res[j] = n * BoolToInt(val2[0])
							}
						} else if len(val1) == len(val2) {
							res = make([]int, len(val2))
							for j, b := range val2 {
								res[j] = val1[j] * BoolToInt(b)
							}
						}
						result = res
					case []int:
						var res []int
						if len(val1) == 1 {
							res = make([]int, len(val2))
							for j, n := range val2 {
								res[j] = val1[0] * n
							}
						} else if len(val2) == 1 {
							res = make([]int, len(val1))
							for j, n := range val1 {
								res[j] = n * val2[0]
							}
						} else if len(val1) == len(val2) {
							res = make([]int, len(val2))
							for j, n := range val2 {
								res[j] = val1[j] * n
							}
						}
						result = res
					case []float64:
						var res []float64
						if len(val1) == 1 {
							res = make([]float64, len(val2))
							for j, f := range val2 {
								res[j] = float64(val1[0]) * f
							}
						} else if len(val2) == 1 {
							res = make([]float64, len(val1))
							for j, n := range val1 {
								res[j] = float64(n) * val2[0]
							}
						} else if len(val1) == len(val2) {
							res = make([]float64, len(val2))
							for j, f := range val2 {
								res[j] = float64(val1[j]) * f
							}
						}
						result = res
					case []string:
						var res []string
						if len(val1) == 1 {
							res = make([]string, len(val2))
							for j, s := range val2 {
								res[j] = strings.Repeat(s, int(val1[0]))
							}
						} else if len(val2) == 1 {
							res = make([]string, len(val1))
							for j, n := range val1 {
								res[j] = strings.Repeat(val2[0], n)
							}
						} else if len(val1) == len(val2) {
							res = make([]string, len(val2))
							for j, s := range val2 {
								res[j] = strings.Repeat(s, val1[j])
							}
						}
						result = res
					default:
						return fmt.Errorf("binary \"*\" operator not implemented for %s and %s",
							preludiocompiler.GoToPreludioTypeString(val1),
							preludiocompiler.GoToPreludioTypeString(val2))
					}

				case []float64:
					switch val2 := t2.(type) {
					case []bool:
						var res []float64
						if len(val1) == 1 {
							res = make([]float64, len(val2))
							for j, b := range val2 {
								res[j] = val1[0] * BoolToFloat64(b)
							}
						} else if len(val2) == 1 {
							res = make([]float64, len(val1))
							for j, f := range val1 {
								res[j] = f * BoolToFloat64(val2[0])
							}
						} else if len(val1) == len(val2) {
							res = make([]float64, len(val2))
							for j, b := range val2 {
								res[j] = val1[j] * BoolToFloat64(b)
							}
						}
						result = res
					case []int:
						var res []float64
						if len(val1) == 1 {
							res = make([]float64, len(val2))
							for j, n := range val2 {
								res[j] = val1[0] * float64(n)
							}
						} else if len(val2) == 1 {
							res = make([]float64, len(val1))
							for j, f := range val1 {
								res[j] = f * float64(val2[0])
							}
						} else if len(val1) == len(val2) {
							res = make([]float64, len(val2))
							for j, n := range val2 {
								res[j] = val1[j] * float64(n)
							}
						}
						result = res
					case []float64:
						var res []float64
						if len(val1) == 1 {
							res = make([]float64, len(val2))
							for j, f := range val2 {
								res[j] = val1[0] * f
							}
						} else if len(val2) == 1 {
							res = make([]float64, len(val1))
							for j, f := range val1 {
								res[j] = f * val2[0]
							}
						} else if len(val1) == len(val2) {
							res = make([]float64, len(val2))
							for j, f := range val2 {
								res[j] = val1[j] * f
							}
						}
						result = res
					default:
						return fmt.Errorf("binary \"*\" operator not implemented for %s and %s",
							preludiocompiler.GoToPreludioTypeString(val1),
							preludiocompiler.GoToPreludioTypeString(val2))
					}

				case []string:
					switch val2 := t2.(type) {
					case []bool:
						var res []string
						if len(val1) == 1 {
							res = make([]string, len(val2))
							for j, b := range val2 {
								if b {
									res[j] = val1[0]
								} else {
									res[j] = ""
								}
							}
						} else if len(val2) == 1 {
							res = make([]string, len(val1))
							for j, s := range val1 {
								if val2[0] {
									res[j] = s
								} else {
									res[j] = ""
								}
							}
						} else if len(val1) == len(val2) {
							res = make([]string, len(val2))
							for j, s := range val1 {
								if val2[j] {
									res[j] = s
								} else {
									res[j] = ""
								}
							}
						}
						result = res
					case []int:
						var res []string
						if len(val1) == 1 {
							res = make([]string, len(val2))
							for j, n := range val2 {
								res[j] = strings.Repeat(val1[0], n)
							}
						} else if len(val2) == 1 {
							res = make([]string, len(val1))
							for j, n := range val2 {
								res[j] = strings.Repeat(val1[0], n)
							}
						} else if len(val1) == len(val2) {
							res = make([]string, len(val2))
							for j, s := range val1 {
								res[j] = strings.Repeat(s, val2[j])
							}
						}
						result = res
					default:
						return fmt.Errorf("binary \"*\" operator not implemented for %s and %s",
							preludiocompiler.GoToPreludioTypeString(val1),
							preludiocompiler.GoToPreludioTypeString(val2))
					}

				default:
					return fmt.Errorf("binary \"*\" operator not implemented for %s and %s",
						preludiocompiler.GoToPreludioTypeString(t1),
						preludiocompiler.GoToPreludioTypeString(t2))
				}

			///////////////////////////////////////////////////////////////////
			////////					DIVISION

			case preludiocompiler.OP_BINARY_DIV:
				switch val1 := t1.(type) {
				case []bool:
					switch val2 := t2.(type) {
					case []bool:
						var res []float64
						if len(val1) == 1 {
							res = make([]float64, len(val2))
							for j, b := range val2 {
								res[j] = BoolToFloat64(val1[0]) / BoolToFloat64(b)
							}
						} else if len(val2) == 1 {
							res = make([]float64, len(val1))
							for j, b := range val1 {
								res[j] = BoolToFloat64(b) / BoolToFloat64(val2[0])
							}
						} else if len(val1) == len(val2) {
							res = make([]float64, len(val2))
							for j, b := range val2 {
								res[j] = BoolToFloat64(val1[j]) / BoolToFloat64(b)
							}
						}
						result = res
					case []int:
						var res []float64
						if len(val1) == 1 {
							if val1[0] {
								res = make([]float64, len(val2))
								for j, n := range val2 {
									res[j] = 1 / float64(n)
								}
							} else {
								res = make([]float64, len(val2))
								for j := range val2 {
									res[j] = 0
								}
							}
						} else if len(val2) == 1 {
							res = make([]float64, len(val1))
							for j, b := range val1 {
								res[j] = BoolToFloat64(b) / float64(val2[0])
							}
						} else if len(val1) == len(val2) {
							res = make([]float64, len(val2))
							for j, n := range val2 {
								res[j] = BoolToFloat64(val1[j]) / float64(n)
							}
						}
						result = res
					case []float64:
						var res []float64
						if len(val1) == 1 {
							if val1[0] {
								res = val2
							} else {
								res = make([]float64, len(val2))
								for j := range val2 {
									res[j] = 0.0
								}
							}
						} else if len(val2) == 1 {
							res = make([]float64, len(val1))
							for j, b := range val1 {
								res[j] = BoolToFloat64(b) / val2[0]
							}
						} else if len(val1) == len(val2) {
							res = make([]float64, len(val2))
							for j, n := range val2 {
								res[j] = BoolToFloat64(val1[j]) / n
							}
						}
						result = res
					default:
						return fmt.Errorf("binary \"/\" operator not implemented for %s and %s",
							preludiocompiler.GoToPreludioTypeString(val1),
							preludiocompiler.GoToPreludioTypeString(val2))
					}

				case []int:
					switch val2 := t2.(type) {
					case []bool:
						var res []float64
						if len(val1) == 1 {
							res = make([]float64, len(val2))
							for j, b := range val2 {
								res[j] = float64(val1[0]) / BoolToFloat64(b)
							}
						} else if len(val2) == 1 {
							res = make([]float64, len(val1))
							for j, n := range val1 {
								res[j] = float64(n) / BoolToFloat64(val2[0])
							}
						} else if len(val1) == len(val2) {
							res = make([]float64, len(val2))
							for j, b := range val2 {
								res[j] = float64(val1[j]) / BoolToFloat64(b)
							}
						}
						result = res
					case []int:
						var res []float64
						if len(val1) == 1 {
							res = make([]float64, len(val2))
							for j, n := range val2 {
								res[j] = float64(val1[0]) / float64(n)
							}
						} else if len(val2) == 1 {
							res = make([]float64, len(val1))
							for j, n := range val1 {
								res[j] = float64(n) / float64(val2[0])
							}
						} else if len(val1) == len(val2) {
							res = make([]float64, len(val2))
							for j, n := range val2 {
								res[j] = float64(val1[j]) / float64(n)
							}
						}
						result = res
					case []float64:
						var res []float64
						if len(val1) == 1 {
							res = make([]float64, len(val2))
							for j, f := range val2 {
								res[j] = float64(val1[0]) / f
							}
						} else if len(val2) == 1 {
							res = make([]float64, len(val1))
							for j, n := range val1 {
								res[j] = float64(n) / val2[0]
							}
						} else if len(val1) == len(val2) {
							res = make([]float64, len(val2))
							for j, f := range val2 {
								res[j] = float64(val1[j]) / f
							}
						}
						result = res
					default:
						return fmt.Errorf("binary \"/\" operator not implemented for %s and %s",
							preludiocompiler.GoToPreludioTypeString(val1),
							preludiocompiler.GoToPreludioTypeString(val2))
					}

				case []float64:
					switch val2 := t2.(type) {
					case []bool:
						var res []float64
						if len(val1) == 1 {
							res = make([]float64, len(val2))
							for j, b := range val2 {
								res[j] = val1[0] / BoolToFloat64(b)
							}
						} else if len(val2) == 1 {
							res = make([]float64, len(val1))
							for j, f := range val1 {
								res[j] = f / BoolToFloat64(val2[0])
							}
						} else if len(val1) == len(val2) {
							res = make([]float64, len(val2))
							for j, b := range val2 {
								res[j] = val1[j] / BoolToFloat64(b)
							}
						}
						result = res
					case []int:
						var res []float64
						if len(val1) == 1 {
							res = make([]float64, len(val2))
							for j, n := range val2 {
								res[j] = val1[0] / float64(n)
							}
						} else if len(val2) == 1 {
							res = make([]float64, len(val1))
							for j, f := range val1 {
								res[j] = f / float64(val2[0])
							}
						} else if len(val1) == len(val2) {
							res = make([]float64, len(val2))
							for j, n := range val2 {
								res[j] = val1[j] / float64(n)
							}
						}
						result = res
					case []float64:
						var res []float64
						if len(val1) == 1 {
							res = make([]float64, len(val2))
							for j, f := range val2 {
								res[j] = val1[0] / f
							}
						} else if len(val2) == 1 {
							res = make([]float64, len(val1))
							for j, f := range val1 {
								res[j] = f / val2[0]
							}
						} else if len(val1) == len(val2) {
							res = make([]float64, len(val2))
							for j, f := range val2 {
								res[j] = val1[j] / f
							}
						}
						result = res
					default:
						return fmt.Errorf("binary \"/\" operator not implemented for %s and %s",
							preludiocompiler.GoToPreludioTypeString(val1),
							preludiocompiler.GoToPreludioTypeString(val2))
					}

				default:
					return fmt.Errorf("binary \"/\" operator not implemented for %s and %s",
						preludiocompiler.GoToPreludioTypeString(t1),
						preludiocompiler.GoToPreludioTypeString(t2))
				}

			///////////////////////////////////////////////////////////////////
			////////					MODULUS

			case preludiocompiler.OP_BINARY_MOD:
				switch val1 := t1.(type) {
				case []bool:
					switch val2 := t2.(type) {
					case []bool:
						var res []int
						if len(val1) == 1 {
							res = make([]int, len(val2))
							for j, b := range val2 {
								res[j] = BoolToInt(val1[0]) % BoolToInt(b)
							}
						} else if len(val2) == 1 {
							res = make([]int, len(val1))
							for j, b := range val1 {
								res[j] = BoolToInt(b) % BoolToInt(val2[0])
							}
						} else if len(val1) == len(val2) {
							res = make([]int, len(val2))
							for j, b := range val2 {
								res[j] = BoolToInt(val1[j]) % BoolToInt(b)
							}
						}
						result = res
					case []int:
						var res []int
						if len(val1) == 1 {
							if val1[0] {
								res = val2
							} else {
								res = make([]int, len(val2))
								for j := range val2 {
									res[j] = 0
								}
							}
						} else if len(val2) == 1 {
							res = make([]int, len(val1))
							for j, b := range val1 {
								res[j] = BoolToInt(b) % val2[0]
							}
						} else if len(val1) == len(val2) {
							res = make([]int, len(val2))
							for j, n := range val2 {
								res[j] = BoolToInt(val1[j]) % n
							}
						}
						result = res
					default:
						return fmt.Errorf("binary \"%%\" operator not implemented for %s and %s",
							preludiocompiler.GoToPreludioTypeString(val1),
							preludiocompiler.GoToPreludioTypeString(val2))
					}

				case []int:
					switch val2 := t2.(type) {
					case []bool:
						var res []int
						if len(val1) == 1 {
							res = make([]int, len(val2))
							for j, b := range val2 {
								res[j] = val1[0] % BoolToInt(b)
							}
						} else if len(val2) == 1 {
							res = make([]int, len(val1))
							for j, n := range val1 {
								res[j] = n % BoolToInt(val2[0])
							}
						} else if len(val1) == len(val2) {
							res = make([]int, len(val2))
							for j, b := range val2 {
								res[j] = val1[j] % BoolToInt(b)
							}
						}
						result = res
					case []int:
						var res []int
						if len(val1) == 1 {
							res = make([]int, len(val2))
							for j, n := range val2 {
								res[j] = val1[0] % n
							}
						} else if len(val2) == 1 {
							res = make([]int, len(val1))
							for j, n := range val1 {
								res[j] = n % val2[0]
							}
						} else if len(val1) == len(val2) {
							res = make([]int, len(val2))
							for j, n := range val2 {
								res[j] = val1[j] % n
							}
						}
						result = res
					default:
						return fmt.Errorf("binary \"%%\" operator not implemented for %s and %s",
							preludiocompiler.GoToPreludioTypeString(val1),
							preludiocompiler.GoToPreludioTypeString(val2))
					}

				default:
					return fmt.Errorf("binary \"%%\" operator not implemented for %s and %s",
						preludiocompiler.GoToPreludioTypeString(t1),
						preludiocompiler.GoToPreludioTypeString(t2))
				}

			///////////////////////////////////////////////////////////////////
			////////					ADDITION

			case preludiocompiler.OP_BINARY_ADD:
				switch val1 := t1.(type) {
				case []bool:
					switch val2 := t2.(type) {
					case []bool:
						var res []int
						if len(val1) == 1 {
							res = make([]int, len(val2))
							for j, b := range val2 {
								res[j] = BoolToInt(val1[0]) + BoolToInt(b)
							}
						} else if len(val2) == 1 {
							res = make([]int, len(val1))
							for j, b := range val1 {
								res[j] = BoolToInt(b) + BoolToInt(val2[0])
							}
						} else if len(val1) == len(val2) {
							res = make([]int, len(val2))
							for j, b := range val2 {
								res[j] = BoolToInt(val1[j]) + BoolToInt(b)
							}
						}
						result = res
					case []int:
						var res []int
						if len(val1) == 1 {
							res = make([]int, len(val2))
							for j, n := range val2 {
								res[j] = BoolToInt(val1[0]) + n
							}
						} else if len(val2) == 1 {
							res = make([]int, len(val1))
							for j, b := range val1 {
								res[j] = BoolToInt(b) + val2[0]
							}
						} else if len(val1) == len(val2) {
							res = make([]int, len(val2))
							for j, n := range val2 {
								res[j] = BoolToInt(val1[j]) + n
							}
						}
						result = res
					case []float64:
						var res []float64
						if len(val1) == 1 {
							res = make([]float64, len(val2))
							for j, f := range val2 {
								res[j] = BoolToFloat64(val1[0]) + f
							}
						} else if len(val2) == 1 {
							res = make([]float64, len(val1))
							for j, b := range val1 {
								res[j] = BoolToFloat64(b) + val2[0]
							}
						} else if len(val1) == len(val2) {
							res = make([]float64, len(val2))
							for j, n := range val2 {
								res[j] = BoolToFloat64(val1[j]) + n
							}
						}
						result = res
					case []string:
						var res []string
						if len(val1) == 1 {
							res = make([]string, len(val2))
							for j, s := range val2 {
								res[j] = fmt.Sprintf("%t%s", val1[0], s)
							}
						} else if len(val2) == 1 {
							res = make([]string, len(val1))
							for j, b := range val1 {
								res[j] = fmt.Sprintf("%t%s", b, val2[0])
							}
						} else if len(val1) == len(val2) {
							res = make([]string, len(val2))
							for j, b := range val1 {
								res[j] = fmt.Sprintf("%t%s", b, val2[j])
							}
						}
						result = res

					default:
						return fmt.Errorf("binary \"+\" operator not implemented for %s and %s",
							preludiocompiler.GoToPreludioTypeString(val1),
							preludiocompiler.GoToPreludioTypeString(val2))
					}

				case []int:
					switch val2 := t2.(type) {
					case []bool:
						var res []int
						if len(val1) == 1 {
							res = make([]int, len(val2))
							for j, b := range val2 {
								res[j] = val1[0] + BoolToInt(b)
							}
						} else if len(val2) == 1 {
							res = make([]int, len(val1))
							for j, n := range val1 {
								res[j] = n + BoolToInt(val2[0])
							}
						} else if len(val1) == len(val2) {
							res = make([]int, len(val2))
							for j, b := range val2 {
								res[j] = val1[j] + BoolToInt(b)
							}
						}
						result = res
					case []int:
						var res []int
						if len(val1) == 1 {
							res = make([]int, len(val2))
							for j, n := range val2 {
								res[j] = val1[0] + n
							}
						} else if len(val2) == 1 {
							res = make([]int, len(val1))
							for j, n := range val1 {
								res[j] = n + val2[0]
							}
						} else if len(val1) == len(val2) {
							res = make([]int, len(val2))
							for j, n := range val2 {
								res[j] = val1[j] + n
							}
						}
						result = res
					case []float64:
						var res []float64
						if len(val1) == 1 {
							res = make([]float64, len(val2))
							for j, f := range val2 {
								res[j] = float64(val1[0]) + f
							}
						} else if len(val2) == 1 {
							res = make([]float64, len(val1))
							for j, n := range val1 {
								res[j] = float64(n) + val2[0]
							}
						} else if len(val1) == len(val2) {
							res = make([]float64, len(val2))
							for j, f := range val2 {
								res[j] = float64(val1[j]) + f
							}
						}
						result = res
					case []string:
						var res []string
						if len(val1) == 1 {
							res = make([]string, len(val2))
							for j, s := range val2 {
								res[j] = fmt.Sprintf("%d%s", val1[0], s)
							}
						} else if len(val2) == 1 {
							res = make([]string, len(val1))
							for j, b := range val1 {
								res[j] = fmt.Sprintf("%d%s", b, val2[0])
							}
						} else if len(val1) == len(val2) {
							res = make([]string, len(val2))
							for j, b := range val1 {
								res[j] = fmt.Sprintf("%d%s", b, val2[j])
							}
						}
						result = res
					default:
						return fmt.Errorf("binary \"+\" operator not implemented for %s and %s",
							preludiocompiler.GoToPreludioTypeString(val1),
							preludiocompiler.GoToPreludioTypeString(val2))
					}

				case []float64:
					switch val2 := t2.(type) {
					case []bool:
						var res []float64
						if len(val1) == 1 {
							res = make([]float64, len(val2))
							for j, b := range val2 {
								res[j] = val1[0] + BoolToFloat64(b)
							}
						} else if len(val2) == 1 {
							res = make([]float64, len(val1))
							for j, f := range val1 {
								res[j] = f + BoolToFloat64(val2[0])
							}
						} else if len(val1) == len(val2) {
							res = make([]float64, len(val2))
							for j, b := range val2 {
								res[j] = val1[j] + BoolToFloat64(b)
							}
						}
						result = res
					case []int:
						var res []float64
						if len(val1) == 1 {
							res = make([]float64, len(val2))
							for j, n := range val2 {
								res[j] = val1[0] + float64(n)
							}
						} else if len(val2) == 1 {
							res = make([]float64, len(val1))
							for j, f := range val1 {
								res[j] = f + float64(val2[0])
							}
						} else if len(val1) == len(val2) {
							res = make([]float64, len(val2))
							for j, n := range val2 {
								res[j] = val1[j] + float64(n)
							}
						}
						result = res
					case []float64:
						var res []float64
						if len(val1) == 1 {
							res = make([]float64, len(val2))
							for j, f := range val2 {
								res[j] = val1[0] + f
							}
						} else if len(val2) == 1 {
							res = make([]float64, len(val1))
							for j, f := range val1 {
								res[j] = f + val2[0]
							}
						} else if len(val1) == len(val2) {
							res = make([]float64, len(val2))
							for j, f := range val2 {
								res[j] = val1[j] + f
							}
						}
						result = res
					case []string:
						var res []string
						if len(val1) == 1 {
							res = make([]string, len(val2))
							for j, f := range val2 {
								res[j] = fmt.Sprintf("%f%s", val1[0], f)
							}
						} else if len(val2) == 1 {
							res = make([]string, len(val1))
							for j, f := range val1 {
								res[j] = fmt.Sprintf("%f%s", f, val2[0])
							}
						} else if len(val1) == len(val2) {
							res = make([]string, len(val2))
							for j, f := range val1 {
								res[j] = fmt.Sprintf("%f%s", f, val2[j])
							}
						}
						result = res
					default:
						return fmt.Errorf("binary \"+\" operator not implemented for %s and %s",
							preludiocompiler.GoToPreludioTypeString(val1),
							preludiocompiler.GoToPreludioTypeString(val2))
					}

				case []string:
					switch val2 := t2.(type) {
					case []bool:
						var res []string
						if len(val1) == 1 {
							res = make([]string, len(val2))
							for j, b := range val2 {
								res[j] = fmt.Sprintf("%s%t", val1[0], b)
							}
						} else if len(val2) == 1 {
							res = make([]string, len(val1))
							for j, s := range val1 {
								res[j] = fmt.Sprintf("%s%t", s, val2[0])
							}
						} else if len(val1) == len(val2) {
							res = make([]string, len(val2))
							for j, s := range val1 {
								res[j] = fmt.Sprintf("%s%t", s, val2[j])
							}
						}
						result = res
					case []int:
						var res []string
						if len(val1) == 1 {
							res = make([]string, len(val2))
							for j, n := range val2 {
								res[j] = fmt.Sprintf("%s%d", val1[0], n)
							}
						} else if len(val2) == 1 {
							res = make([]string, len(val1))
							for j, s := range val1 {
								res[j] = fmt.Sprintf("%s%d", s, val2[0])
							}
						} else if len(val1) == len(val2) {
							res = make([]string, len(val2))
							for j, s := range val1 {
								res[j] = fmt.Sprintf("%s%d", s, val2[j])
							}
						}
						result = res
					case []float64:
						var res []string
						if len(val1) == 1 {
							res = make([]string, len(val2))
							for j, f := range val2 {
								res[j] = fmt.Sprintf("%s%f", val1[0], f)
							}
						} else if len(val2) == 1 {
							res = make([]string, len(val1))
							for j, s := range val1 {
								res[j] = fmt.Sprintf("%s%f", s, val2[0])
							}
						} else if len(val1) == len(val2) {
							res = make([]string, len(val2))
							for j, s := range val1 {
								res[j] = fmt.Sprintf("%s%f", s, val2[j])
							}
						}
						result = res
					case []string:
						var res []string
						if len(val1) == 1 {
							res = make([]string, len(val2))
							for j, s := range val2 {
								res[j] = fmt.Sprintf("%s%s", val1[0], s)
							}
						} else if len(val2) == 1 {
							res = make([]string, len(val1))
							for j, s := range val1 {
								res[j] = fmt.Sprintf("%s%s", s, val2[0])
							}
						} else if len(val1) == len(val2) {
							res = make([]string, len(val2))
							for j, s := range val1 {
								res[j] = fmt.Sprintf("%s%s", s, val2[j])
							}
						}
						result = res
					default:
						return fmt.Errorf("binary \"+\" operator not implemented for %s and %s",
							preludiocompiler.GoToPreludioTypeString(val1),
							preludiocompiler.GoToPreludioTypeString(val2))
					}

				default:
					return fmt.Errorf("binary \"+\" operator not implemented for %s and %s",
						preludiocompiler.GoToPreludioTypeString(t1),
						preludiocompiler.GoToPreludioTypeString(t2))
				}

			///////////////////////////////////////////////////////////////////
			////////					SUBTRACTION

			case preludiocompiler.OP_BINARY_SUB:
				switch val1 := t1.(type) {
				case []bool:
					switch val2 := t2.(type) {
					case []bool:
						var res []int
						if len(val1) == 1 {
							res = make([]int, len(val2))
							for j, b := range val2 {
								res[j] = BoolToInt(val1[0]) - BoolToInt(b)
							}
						} else if len(val2) == 1 {
							res = make([]int, len(val1))
							for j, b := range val1 {
								res[j] = BoolToInt(b) - BoolToInt(val2[0])
							}
						} else if len(val1) == len(val2) {
							res = make([]int, len(val2))
							for j, b := range val2 {
								res[j] = BoolToInt(val1[j]) - BoolToInt(b)
							}
						}
						result = res
					case []int:
						var res []int
						if len(val1) == 1 {
							res = make([]int, len(val2))
							for j, n := range val2 {
								res[j] = BoolToInt(val1[0]) - n
							}
						} else if len(val2) == 1 {
							res = make([]int, len(val1))
							for j, b := range val1 {
								res[j] = BoolToInt(b) - val2[0]
							}
						} else if len(val1) == len(val2) {
							res = make([]int, len(val2))
							for j, n := range val2 {
								res[j] = BoolToInt(val1[j]) - n
							}
						}
						result = res
					case []float64:
						var res []float64
						if len(val1) == 1 {
							res = make([]float64, len(val2))
							for j, f := range val2 {
								res[j] = BoolToFloat64(val1[0]) - f
							}
						} else if len(val2) == 1 {
							res = make([]float64, len(val1))
							for j, b := range val1 {
								res[j] = BoolToFloat64(b) - val2[0]
							}
						} else if len(val1) == len(val2) {
							res = make([]float64, len(val2))
							for j, n := range val2 {
								res[j] = BoolToFloat64(val1[j]) - n
							}
						}
						result = res
					default:
						return fmt.Errorf("binary \"-\" operator not implemented for %s and %s",
							preludiocompiler.GoToPreludioTypeString(val1),
							preludiocompiler.GoToPreludioTypeString(val2))
					}

				case []int:
					switch val2 := t2.(type) {
					case []bool:
						var res []int
						if len(val1) == 1 {
							res = make([]int, len(val2))
							for j, b := range val2 {
								res[j] = val1[0] - BoolToInt(b)
							}
						} else if len(val2) == 1 {
							res = make([]int, len(val1))
							for j, n := range val1 {
								res[j] = n - BoolToInt(val2[0])
							}
						} else if len(val1) == len(val2) {
							res = make([]int, len(val2))
							for j, b := range val2 {
								res[j] = val1[j] - BoolToInt(b)
							}
						}
						result = res
					case []int:
						var res []int
						if len(val1) == 1 {
							res = make([]int, len(val2))
							for j, n := range val2 {
								res[j] = val1[0] - n
							}
						} else if len(val2) == 1 {
							res = make([]int, len(val1))
							for j, n := range val1 {
								res[j] = n - val2[0]
							}
						} else if len(val1) == len(val2) {
							res = make([]int, len(val2))
							for j, n := range val2 {
								res[j] = val1[j] - n
							}
						}
						result = res
					case []float64:
						var res []float64
						if len(val1) == 1 {
							res = make([]float64, len(val2))
							for j, f := range val2 {
								res[j] = float64(val1[0]) - f
							}
						} else if len(val2) == 1 {
							res = make([]float64, len(val1))
							for j, n := range val1 {
								res[j] = float64(n) - val2[0]
							}
						} else if len(val1) == len(val2) {
							res = make([]float64, len(val2))
							for j, f := range val2 {
								res[j] = float64(val1[j]) - f
							}
						}
						result = res
					default:
						return fmt.Errorf("binary \"-\" operator not implemented for %s and %s",
							preludiocompiler.GoToPreludioTypeString(val1),
							preludiocompiler.GoToPreludioTypeString(val2))
					}

				case []float64:
					switch val2 := t2.(type) {
					case []bool:
						var res []float64
						if len(val1) == 1 {
							res = make([]float64, len(val2))
							for j, b := range val2 {
								res[j] = val1[0] - BoolToFloat64(b)
							}
						} else if len(val2) == 1 {
							res = make([]float64, len(val1))
							for j, f := range val1 {
								res[j] = f - BoolToFloat64(val2[0])
							}
						} else if len(val1) == len(val2) {
							res = make([]float64, len(val2))
							for j, b := range val2 {
								res[j] = val1[j] - BoolToFloat64(b)
							}
						}
						result = res
					case []int:
						var res []float64
						if len(val1) == 1 {
							res = make([]float64, len(val2))
							for j, n := range val2 {
								res[j] = val1[0] - float64(n)
							}
						} else if len(val2) == 1 {
							res = make([]float64, len(val1))
							for j, f := range val1 {
								res[j] = f - float64(val2[0])
							}
						} else if len(val1) == len(val2) {
							res = make([]float64, len(val2))
							for j, n := range val2 {
								res[j] = val1[j] - float64(n)
							}
						}
						result = res
					case []float64:
						var res []float64
						if len(val1) == 1 {
							res = make([]float64, len(val2))
							for j, f := range val2 {
								res[j] = val1[0] - f
							}
						} else if len(val2) == 1 {
							res = make([]float64, len(val1))
							for j, f := range val1 {
								res[j] = f - val2[0]
							}
						} else if len(val1) == len(val2) {
							res = make([]float64, len(val2))
							for j, f := range val2 {
								res[j] = val1[j] - f
							}
						}
						result = res
					default:
						return fmt.Errorf("binary \"-\" operator not implemented for %s and %s",
							preludiocompiler.GoToPreludioTypeString(val1),
							preludiocompiler.GoToPreludioTypeString(val2))
					}

				default:
					return fmt.Errorf("binary \"-\" operator not implemented for %s and %s",
						preludiocompiler.GoToPreludioTypeString(t1),
						preludiocompiler.GoToPreludioTypeString(t2))
				}

			///////////////////////////////////////////////////////////////////
			////////					EQUAL
			case preludiocompiler.OP_BINARY_EQ:
				switch val1 := t1.(type) {
				case []bool:
					switch val2 := t2.(type) {
					case []bool:
						var res []bool
						if len(val1) == 1 {
							res = make([]bool, len(val2))
							for j, b := range val2 {
								res[j] = val1[0] == b
							}
						} else if len(val2) == 1 {
							res = make([]bool, len(val1))
							for j, b := range val1 {
								res[j] = b == val2[0]
							}
						} else if len(val1) == len(val2) {
							res = make([]bool, len(val2))
							for j, b := range val2 {
								res[j] = val1[j] == b
							}
						}
						result = res
					case []int:
						var res []bool
						if len(val1) == 1 {
							res = make([]bool, len(val2))
							for j, n := range val2 {
								res[j] = BoolToInt(val1[0]) == n
							}
						} else if len(val2) == 1 {
							res = make([]bool, len(val1))
							for j, b := range val1 {
								res[j] = BoolToInt(b) == val2[0]
							}
						} else if len(val1) == len(val2) {
							res = make([]bool, len(val2))
							for j, n := range val2 {
								res[j] = BoolToInt(val1[j]) == n
							}
						}
						result = res
					case []float64:
						var res []bool
						if len(val1) == 1 {
							res = make([]bool, len(val2))
							for j, f := range val2 {
								res[j] = BoolToFloat64(val1[0]) == f
							}
						} else if len(val2) == 1 {
							res = make([]bool, len(val1))
							for j, b := range val1 {
								res[j] = BoolToFloat64(b) == val2[0]
							}
						} else if len(val1) == len(val2) {
							res = make([]bool, len(val2))
							for j, f := range val2 {
								res[j] = BoolToFloat64(val1[j]) == f
							}
						}
						result = res
					case []string:
						var res []bool
						if len(val1) == 1 {
							res = make([]bool, len(val2))
							for j, s := range val2 {
								res[j] = fmt.Sprintf("%t", val1[0]) == s
							}
						} else if len(val2) == 1 {
							res = make([]bool, len(val1))
							for j, b := range val1 {
								res[j] = fmt.Sprintf("%t", b) == val2[0]
							}
						} else if len(val1) == len(val2) {
							res = make([]bool, len(val2))
							for j, s := range val2 {
								res[j] = fmt.Sprintf("%t", val1[j]) == s
							}
						}
						result = res
					default:
						return fmt.Errorf("bynary \"==\" operator not implemented for %s and %s",
							preludiocompiler.GoToPreludioTypeString(val1),
							preludiocompiler.GoToPreludioTypeString(val2))
					}

				case []int:
					switch val2 := t2.(type) {
					case []bool:
						var res []bool
						if len(val1) == 1 {
							res = make([]bool, len(val2))
							for j, b := range val2 {
								res[j] = val1[0] == BoolToInt(b)
							}
						} else if len(val2) == 1 {
							res = make([]bool, len(val1))
							for j, n := range val1 {
								res[j] = n == BoolToInt(val2[0])
							}
						} else if len(val1) == len(val2) {
							res = make([]bool, len(val2))
							for j, b := range val2 {
								res[j] = val1[j] == BoolToInt(b)
							}
						}
						result = res
					case []int:
						var res []bool
						if len(val1) == 1 {
							res = make([]bool, len(val2))
							for j, n := range val2 {
								res[j] = val1[0] == n
							}
						} else if len(val2) == 1 {
							res = make([]bool, len(val1))
							for j, n := range val1 {
								res[j] = n == val2[0]
							}
						} else if len(val1) == len(val2) {
							res = make([]bool, len(val2))
							for j, n := range val2 {
								res[j] = val1[j] == n
							}
						}
						result = res
					case []float64:
						var res []bool
						if len(val1) == 1 {
							res = make([]bool, len(val2))
							for j, f := range val2 {
								res[j] = float64(val1[0]) == f
							}
						} else if len(val2) == 1 {
							res = make([]bool, len(val1))
							for j, n := range val1 {
								res[j] = float64(n) == val2[0]
							}
						} else if len(val1) == len(val2) {
							res = make([]bool, len(val2))
							for j, f := range val2 {
								res[j] = float64(val1[j]) == f
							}
						}
						result = res
					case []string:
						var res []bool
						if len(val1) == 1 {
							res = make([]bool, len(val2))
							for j, s := range val2 {
								res[j] = strconv.Itoa(val1[0]) == s
							}
						} else if len(val2) == 1 {
							res = make([]bool, len(val1))
							for j, n := range val1 {
								res[j] = strconv.Itoa(n) == val2[0]
							}
						} else if len(val1) == len(val2) {
							res = make([]bool, len(val2))
							for j, s := range val2 {
								res[j] = strconv.Itoa(val1[j]) == s
							}
						}
						result = res
					default:
						return fmt.Errorf("bynary \"==\" operator not implemented for %s and %s",
							preludiocompiler.GoToPreludioTypeString(val1),
							preludiocompiler.GoToPreludioTypeString(val2))
					}

				case []float64:
					switch val2 := t2.(type) {
					case []bool:
						var res []bool
						if len(val1) == 1 {
							res = make([]bool, len(val2))
							for j, b := range val2 {
								res[j] = val1[0] == BoolToFloat64(b)
							}
						} else if len(val2) == 1 {
							res = make([]bool, len(val1))
							for j, f := range val1 {
								res[j] = f == BoolToFloat64(val2[0])
							}
						} else if len(val1) == len(val2) {
							res = make([]bool, len(val2))
							for j, b := range val2 {
								res[j] = val1[j] == BoolToFloat64(b)
							}
						}
						result = res
					case []int:
						var res []bool
						if len(val1) == 1 {
							res = make([]bool, len(val2))
							for j, n := range val2 {
								res[j] = val1[0] == float64(n)
							}
						} else if len(val2) == 1 {
							res = make([]bool, len(val1))
							for j, f := range val1 {
								res[j] = f == float64(val2[0])
							}
						} else if len(val1) == len(val2) {
							res = make([]bool, len(val2))
							for j, n := range val2 {
								res[j] = val1[j] == float64(n)
							}
						}
						result = res
					case []float64:
						var res []bool
						if len(val1) == 1 {
							res = make([]bool, len(val2))
							for j, f := range val2 {
								res[j] = val1[0] == f
							}
						} else if len(val2) == 1 {
							res = make([]bool, len(val1))
							for j, f := range val1 {
								res[j] = f == val2[0]
							}
						} else if len(val1) == len(val2) {
							res = make([]bool, len(val2))
							for j, f := range val2 {
								res[j] = val1[j] == f
							}
						}
						result = res
					case []string:
						var res []bool
						if len(val1) == 1 {
							res = make([]bool, len(val2))
							for j, s := range val2 {
								res[j] = strconv.FormatFloat(val1[0], 'f', -1, 64) == s
							}
						}
						result = res
					default:
						return fmt.Errorf("bynary \"==\" operator not implemented for %s and %s",
							preludiocompiler.GoToPreludioTypeString(val1),
							preludiocompiler.GoToPreludioTypeString(val2))
					}

				case []string:
					switch val2 := t2.(type) {
					case []bool:
						var res []bool
						if len(val1) == 1 {
							res = make([]bool, len(val2))
							for j, b := range val2 {
								res[j] = val1[0] == strconv.FormatBool(b)
							}
						} else if len(val2) == 1 {
							res = make([]bool, len(val1))
							for j, s := range val1 {
								res[j] = s == strconv.FormatBool(val2[0])
							}
						} else if len(val1) == len(val2) {
							res = make([]bool, len(val2))
							for j, b := range val2 {
								res[j] = val1[j] == strconv.FormatBool(b)
							}
						}
						result = res
					case []int:
						var res []bool
						if len(val1) == 1 {
							res = make([]bool, len(val2))
							for j, n := range val2 {
								res[j] = val1[0] == strconv.Itoa(n)
							}
						} else if len(val2) == 1 {
							res = make([]bool, len(val1))
							for j, s := range val1 {
								res[j] = s == strconv.Itoa(val2[0])
							}
						} else if len(val1) == len(val2) {
							res = make([]bool, len(val2))
							for j, n := range val2 {
								res[j] = val1[j] == strconv.Itoa(n)
							}
						}
						result = res
					case []float64:
						var res []bool
						if len(val1) == 1 {
							res = make([]bool, len(val2))
							for j, f := range val2 {
								res[j] = val1[0] == strconv.FormatFloat(f, 'f', -1, 64)
							}
						} else if len(val2) == 1 {
							res = make([]bool, len(val1))
							for j, s := range val1 {
								res[j] = s == strconv.FormatFloat(val2[0], 'f', -1, 64)
							}
						} else if len(val1) == len(val2) {
							res = make([]bool, len(val2))
							for j, f := range val2 {
								res[j] = val1[j] == strconv.FormatFloat(f, 'f', -1, 64)
							}
						}
						result = res
					case []string:
						var res []bool
						if len(val1) == 1 {
							res = make([]bool, len(val2))
							for j, s := range val2 {
								res[j] = val1[0] == s
							}
						} else if len(val2) == 1 {
							res = make([]bool, len(val1))
							for j, s := range val1 {
								res[j] = s == val2[0]
							}
						} else if len(val1) == len(val2) {
							res = make([]bool, len(val2))
							for j, s := range val2 {
								res[j] = val1[j] == s
							}
						}
						result = res
					default:
						return fmt.Errorf("bynary \"==\" operator not implemented for %s and %s",
							preludiocompiler.GoToPreludioTypeString(val1),
							preludiocompiler.GoToPreludioTypeString(val2))
					}

				default:
					return fmt.Errorf("bynary \"==\" operator not implemented for %s and %s",
						preludiocompiler.GoToPreludioTypeString(t1),
						preludiocompiler.GoToPreludioTypeString(t2))
				}
			}
		}

		i.expr = append([]interface{}{result}, i.expr...)
	}
	return nil
}
