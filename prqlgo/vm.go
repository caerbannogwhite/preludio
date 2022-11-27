package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"math"
	"os"
)

type OPCODE uint16

const (
	TERM_NULL     uint16 = 0
	TERM_BOOL     uint16 = 1
	TERM_INTEGER  uint16 = 2
	TERM_FLOAT    uint16 = 3
	TERM_STRING   uint16 = 4
	TERM_INTERVAL uint16 = 5
	TERM_RANGE    uint16 = 6
	TERM_LIST     uint16 = 7
	TERM_PIPELINE uint16 = 8
	TERM_SYMBOL   uint16 = 10
)

const (
	OP_BEGIN_PIPELINE OPCODE = 0
	OP_END_PIPELINE   OPCODE = 1
	OP_ASSIGN_STMT    OPCODE = 2
	// OP_BEGIN_FUNC_CALL OPCODE = 3;
	OP_MAKE_FUNC_CALL      OPCODE = 4
	OP_BEGIN_LIST          OPCODE = 5
	OP_END_LIST            OPCODE = 6
	OP_ADD_FUNC_PARAM      OPCODE = 7
	OP_ADD_EXPR_TERM       OPCODE = 8
	OP_PUSH_NAMED_PARAM    OPCODE = 9
	OP_PUSH_ASSIGN_IDENT   OPCODE = 10
	OP_PUSH_TERM           OPCODE = 11
	OP_END_FUNC_CALL_PARAM OPCODE = 12
	OP_GOTO                OPCODE = 50

	OP_BINARY_MUL OPCODE = 100
	OP_BINARY_DIV OPCODE = 101
	OP_BINARY_MOD OPCODE = 102
	OP_BINARY_ADD OPCODE = 103
	OP_BINARY_SUB OPCODE = 104
	OP_BINARY_POW OPCODE = 105

	OP_BINARY_EQ OPCODE = 110
	OP_BINARY_NE OPCODE = 111
	OP_BINARY_GE OPCODE = 112
	OP_BINARY_LE OPCODE = 113
	OP_BINARY_GT OPCODE = 114
	OP_BINARY_LT OPCODE = 115

	OP_BINARY_AND      OPCODE = 120
	OP_BINARY_OR       OPCODE = 121
	OP_BINARY_COALESCE OPCODE = 122
	OP_BINARY_MODEL    OPCODE = 123

	OP_UNARY_SUB OPCODE = 130
	OP_UNARY_ADD OPCODE = 131
	OP_UNARY_NOT OPCODE = 132
)

type PrqlVirtualMachineParams struct {
	PrintWarnings  bool
	DebugLevel     int
	VerbosityLevel int
	InputPath      string
}

type PrqlVirtualMachine struct {
	__printWarnings        bool
	__debugLevel           int
	__verbosityLevel       int
	__inputPath            string
	__symbolTable          []string
	__stack                []*PrqlInternal
	__currentTable         *PrqlInternal
	__userDefinedVariables map[string]PrqlInternal
	__funcNumParams        uint64
}

func NewPrqlVirtualMachine(params *PrqlVirtualMachineParams) *PrqlVirtualMachine {
	vm := PrqlVirtualMachine{
		__printWarnings: params.PrintWarnings,
		__inputPath:     params.InputPath,
		__debugLevel:    params.DebugLevel,
	}
	vm.__userDefinedVariables = map[string]PrqlInternal{}
	vm.__currentTable = nil

	return &vm
}

func (vm *PrqlVirtualMachine) ReadPrqlBytecode() {
	file, err := os.Open(vm.__inputPath)

	if err != nil {
		// return nil, err
	}
	defer file.Close()

	stats, statsErr := file.Stat()
	if statsErr != nil {
		// return nil, statsErr
	}

	var size int64 = stats.Size()
	bytes := make([]byte, size)

	bufr := bufio.NewReader(file)
	_, err = bufr.Read(bytes)

	bytemark := bytes[0:4]
	__symbolTableSize := binary.BigEndian.Uint64(bytes[8:16])

	if vm.__debugLevel > 5 {
		fmt.Println()
		fmt.Printf("BYTECODE INFO\n")
		fmt.Printf("=============\n")
		fmt.Printf("SIZE:              %d\n", size)
		fmt.Printf("BYTE MARK:         %x %x %x %x\n", bytemark[0], bytemark[1], bytemark[2], bytemark[3])
		fmt.Printf("SYMBOL TABLE SIZE: %d\n\n", __symbolTableSize)
		fmt.Printf("STRING SYMBOLS\n")
		fmt.Printf("==============\n")
	}

	offset := uint64(16)
	for i := uint64(0); i < __symbolTableSize; i++ {
		l := binary.BigEndian.Uint64(bytes[offset : offset+8])
		offset += 8

		v := string(bytes[offset : offset+l])
		vm.__symbolTable = append(vm.__symbolTable, v)
		offset += l
	}

	if vm.__debugLevel > 5 {
		for _, symbol := range vm.__symbolTable {
			fmt.Printf("%s\n", symbol)
		}

		fmt.Println()
		fmt.Printf("INSTRUCTIONS\n")
		fmt.Printf("============\n")
	}

	vm.ReadPrqlInstructions(bytes, offset)
}

func (vm *PrqlVirtualMachine) StackIsEmpty() bool {
	return len(vm.__stack) == 0
}

func (vm *PrqlVirtualMachine) StackPush(e *PrqlInternal) {
	vm.__stack = append(vm.__stack, e)
}

func (vm *PrqlVirtualMachine) StackPop() *PrqlInternal {
	e := vm.__stack[len(vm.__stack)-1]
	vm.__stack = vm.__stack[:len(vm.__stack)-1]
	return e
}

func (vm *PrqlVirtualMachine) StackLast() *PrqlInternal {
	return vm.__stack[len(vm.__stack)-1]
}

func (vm *PrqlVirtualMachine) ReadPrqlInstructions(bytes []byte, offset uint64) {

	var opCode OPCODE
	var param1 uint16
	var param2 []byte

	usize := uint64(len(bytes))

MAIN_LOOP:
	for offset < usize {
		opCode = OPCODE(binary.BigEndian.Uint16(bytes[offset : offset+2]))
		offset += 2
		param1 = binary.BigEndian.Uint16(bytes[offset : offset+2])
		offset += 2
		param2 = bytes[offset : offset+8]
		offset += 8

		switch opCode {
		case OP_BEGIN_PIPELINE:
			if vm.__debugLevel > 10 {
				fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "OP_BEGIN_PIPELINE", "", "", "")
			}

		case OP_END_PIPELINE:
			if vm.__debugLevel > 10 {
				fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "OP_END_PIPELINE", "", "", "")
			}

		case OP_ASSIGN_STMT:
			if vm.__debugLevel > 10 {
				fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "OP_ASSIGN_STMT", "", "", "")
			}

		// case OP_BEGIN_FUNC_CALL:3
		case OP_MAKE_FUNC_CALL:
			funcName := vm.__symbolTable[binary.BigEndian.Uint64(param2)]
			if vm.__debugLevel > 10 {
				fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "OP_MAKE_FUNC_CALL", "", funcName, "")
			}

			switch funcName {

			// Standard library functions build-ins
			case "derive":
				PrqlFunc_Derive("derive", vm)
			case "from":
				PrqlFunc_From("from", vm)
			case "exportCSV":
				PrqlFunc_ExportCsv("exportCSV", vm)
			case "importCSV":
				PrqlFunc_ImportCsv("importCSV", vm)
			case "select":
				PrqlFunc_Select("select", vm)

			// User defined functions
			default:
				if internal, ok := vm.__userDefinedVariables[funcName]; ok {
					switch value := internal.Expr.GetValue().(type) {
					case UserDefinedFunction:
						value(vm)
					default:
						vm.StackPush(NewPrqlInternalError(fmt.Sprintf("variable '%s' not callable.", funcName)))
					}
				} else {
					vm.StackPush(NewPrqlInternalError(fmt.Sprintf("name '%s' not found.", funcName)))
				}
			}

			vm.__funcNumParams = 0

			if vm.StackLast().Tag == PRQL_INTERNAL_TAG_ERROR {
				for !vm.StackIsEmpty() && vm.StackLast().Tag == PRQL_INTERNAL_TAG_ERROR {
					err := vm.StackPop()
					vm.PrintError(err.ErrorMsg)
				}
				break MAIN_LOOP
			}

		case OP_BEGIN_LIST:
			if vm.__debugLevel > 10 {
				fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "OP_BEGIN_LIST", "", "", "")
			}

		case OP_END_LIST:
			if vm.__debugLevel > 10 {
				fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "OP_END_LIST", "", "", "")
			}

		case OP_ADD_FUNC_PARAM:
			if vm.__debugLevel > 10 {
				fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "OP_ADD_FUNC_PARAM", "", "", "")
			}

		case OP_ADD_EXPR_TERM:
			if vm.__debugLevel > 10 {
				fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "OP_ADD_EXPR_TERM", "", "", "")
			}

		case OP_PUSH_NAMED_PARAM:
			paramName := vm.__symbolTable[binary.BigEndian.Uint64(param2)]
			if vm.__debugLevel > 10 {
				fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "OP_PUSH_NAMED_PARAM", "", paramName, "")
			}

			vm.StackPush(NewPrqlInternalParamName(paramName))

		case OP_PUSH_ASSIGN_IDENT:
			ident := vm.__symbolTable[binary.BigEndian.Uint64(param2)]
			if vm.__debugLevel > 10 {
				fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "OP_PUSH_ASSIGN_IDENT", "", ident, "")
			}

			vm.StackPush(NewPrqlInternalParamName(ident))

		case OP_PUSH_TERM:
			termType := ""
			termVal := ""

			switch param1 {
			// case TERM_NULL:
			// 	termType = "NULL"

			case TERM_BOOL:
				termType = "BOOL"
				termVal = "true"
				val := true
				if binary.BigEndian.Uint64(param2) == 0 {
					val = false
					termVal = "false"
				}
				vm.StackPush(NewPrqlInternalTerm(val))

			case TERM_INTEGER:
				termType = "INTEGER"
				val := int64(binary.LittleEndian.Uint64(param2))
				termVal = fmt.Sprintf("%d", val)
				vm.StackPush(NewPrqlInternalTerm(val))

			case TERM_FLOAT:
				termType = "FLOAT"
				val := math.Float64frombits(binary.LittleEndian.Uint64(param2))
				termVal = fmt.Sprintf("%f", val)
				vm.StackPush(NewPrqlInternalTerm(val))

			case TERM_STRING:
				termType = "STRING"
				termVal = vm.__symbolTable[binary.BigEndian.Uint64(param2)]
				vm.StackPush(NewPrqlInternalTerm(termVal))

			case TERM_SYMBOL:
				termType = "SYMBOL"
				termVal = vm.__symbolTable[binary.BigEndian.Uint64(param2)]
				vm.StackPush(NewPrqlInternalTerm(PrqlSymbol(termVal)))

			default:
				vm.StackPush(NewPrqlInternalError(fmt.Sprintf("PrlqVM: unknown term code %d.", param1)))

			}

			if vm.__debugLevel > 10 {
				fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "OP_PUSH_TERM", termType, termVal, "")
			}

		case OP_END_FUNC_CALL_PARAM:
			if vm.__debugLevel > 10 {
				fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "OP_END_FUNC_CALL_PARAM", "", "", "")
			}

			vm.__funcNumParams += 1

		case OP_GOTO:
			if vm.__debugLevel > 10 {
				fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "OP_GOTO", "", "", "")
			}

		case OP_BINARY_MUL:
			if vm.__debugLevel > 10 {
				fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "OP_BINARY_MUL", "", "", "")
			}

			op2 := vm.StackPop()
			vm.StackLast().Expr.Mul(op2.Expr)

		case OP_BINARY_DIV:
			if vm.__debugLevel > 10 {
				fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "OP_BINARY_DIV", "", "", "")
			}

			op2 := vm.StackPop()
			vm.StackLast().Expr.Div(op2.Expr)

		case OP_BINARY_MOD:
			if vm.__debugLevel > 10 {
				fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "OP_BINARY_MOD", "", "", "")
			}

			op2 := vm.StackPop()
			vm.StackLast().Expr.Mod(op2.Expr)

		case OP_BINARY_ADD:
			if vm.__debugLevel > 10 {
				fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "OP_BINARY_ADD", "", "", "")
			}

			op2 := vm.StackPop()
			vm.StackLast().Expr.Add(op2.Expr)

		case OP_BINARY_SUB:
			if vm.__debugLevel > 10 {
				fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "OP_BINARY_SUB", "", "", "")
			}

			op2 := vm.StackPop()
			vm.StackLast().Expr.Sub(op2.Expr)

		case OP_BINARY_POW:
			if vm.__debugLevel > 10 {
				fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "OP_BINARY_POW", "", "", "")
			}

			op2 := vm.StackPop()
			vm.StackLast().Expr.Pow(op2.Expr)

		case OP_BINARY_EQ:
		case OP_BINARY_NE:
		case OP_BINARY_GE:
		case OP_BINARY_LE:
		case OP_BINARY_GT:
		case OP_BINARY_LT:
		case OP_BINARY_AND:
		case OP_BINARY_OR:

		case OP_BINARY_COALESCE:
		case OP_BINARY_MODEL:

		case OP_UNARY_SUB:
			if vm.__debugLevel > 10 {
				fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "OP_UNARY_SUB", "", "", "")
			}

		case OP_UNARY_ADD:
		case OP_UNARY_NOT:

		}
	}
}

func (vm *PrqlVirtualMachine) GetFunctionParams(funcName string, positionalParamsNum uint64, namedParams *map[string]*PrqlExpr, acceptingAssignments bool) ([]*PrqlExpr, map[string]*PrqlExpr) {

	positionalParams := make([]*PrqlExpr, positionalParamsNum)
	var assignments map[string]*PrqlExpr
	if acceptingAssignments {
		assignments = map[string]*PrqlExpr{}
	}

	var t1, t2 PrqlInternal
	var positionalParamsIdx = uint64(0)
	for i := uint64(0); i < vm.__funcNumParams; i++ {
		t1 = *vm.StackPop()
		switch t1.Tag {
		case PRQL_INTERNAL_TAG_ERROR:
		case PRQL_INTERNAL_TAG_EXPRESSION:
			if positionalParamsIdx < positionalParamsNum {
				positionalParams[positionalParamsIdx] = t1.Expr
				positionalParamsIdx++
			} else {
				vm.PrintWarning(fmt.Sprintf("function %s expects exactly %d positional parametes, the remaining values will be ignored.", funcName, positionalParamsNum))
			}

		case PRQL_INTERNAL_TAG_PARAM_NAME:
			t2 = *vm.StackPop()

			// Name of parameter is in the given list of names
			if _, ok := (*namedParams)[t1.Name]; ok {
				(*namedParams)[t1.Name] = t2.Expr
			} else {
				vm.PrintWarning(fmt.Sprintf("function %s does not know a parameter named '%s', the value will be ignored.", funcName, t1.Name))
			}

		case PRQL_INTERNAL_TAG_ASSING_IDENT:
			t2 = *vm.StackPop()
			if acceptingAssignments {
				assignments[t1.Name] = t2.Expr
			} else {
				vm.PrintWarning(fmt.Sprintf("function %s does not accept assignements, the value of '%s' will be ignored.", funcName, t1.Name))
			}

		}
	}

	for _, p := range positionalParams {
		p.Solve()
	}

	for _, p := range *namedParams {
		p.Solve()
	}

	if acceptingAssignments {
		for _, p := range assignments {
			p.Solve()
		}
	}

	return positionalParams, assignments
}

// func (vm *PrqlVirtualMachine) IdentResolutionStrategy(ident string) *PrqlInternal {
// 	if vm.__currentTable != nil {
// 		switch t := vm.__currentTable.Value.(type) {
// 		case dataframe.DataFrame:
// 			names := t.Names()
// 			for _, name := range names {
// 				if name == ident {
// 					return NewPrqlInternalTerm(t.Col(name))
// 				}
// 			}
// 		default:
// 		}
// 	}

// 	if v, ok := vm.__userDefinedVariables[ident]; ok {
// 		return &v
// 	}

// 	return NewPrqlInternalError(fmt.Sprintf("name '%s' not found.", ident))
// }

func (vm *PrqlVirtualMachine) PrintWarning(msg string) {
	if vm.__printWarnings {
		fmt.Printf("[ ⚠️ Warning ⚠️ ] %s\n", msg)
	}
}

func (vm *PrqlVirtualMachine) PrintError(msg string) {
	fmt.Printf("[ ☠️ Error ☠️ ] %s\n", msg)
}
