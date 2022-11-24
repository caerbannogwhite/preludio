package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"math"
	"os"

	"github.com/go-gota/gota/dataframe"
)

const TERM_NULL uint16 = 0
const TERM_BOOL uint16 = 1
const TERM_FLOAT uint16 = 2
const TERM_STRING uint16 = 3

// const TERM_INTERVAL uint16 = 5;
// const TERM_RANGE uint16 = 6;
// const TERM_LIST uint16 = 7;
// const TERM_PIPELINE uint16 = 8;
const TERM_IDENT uint16 = 10

const OP_BEGIN_PIPELINE uint16 = 0
const OP_END_PIPELINE uint16 = 1
const OP_ASSIGN_STMT uint16 = 2

// const OP_BEGIN_FUNC_CALL uint16 = 3;
const OP_MAKE_FUNC_CALL uint16 = 4
const OP_BEGIN_LIST uint16 = 5
const OP_END_LIST uint16 = 6
const OP_ADD_FUNC_PARAM uint16 = 7
const OP_ADD_EXPR_TERM uint16 = 8
const OP_PUSH_NAMED_PARAM uint16 = 9
const OP_PUSH_ASSIGN_IDENT uint16 = 10
const OP_PUSH_TERM uint16 = 11
const OP_END_FUNC_CALL_PARAM uint16 = 12
const OP_GOTO uint16 = 50

const OP_BINARY_MUL uint16 = 100
const OP_BINARY_DIV uint16 = 101
const OP_BINARY_MOD uint16 = 102
const OP_BINARY_ADD uint16 = 103
const OP_BINARY_SUB uint16 = 104

// const OP_BINARY_EQ uint16 = 110;
// const OP_BINARY_NE uint16 = 111;
// const OP_BINARY_GE uint16 = 112;
// const OP_BINARY_LE uint16 = 113;
// const OP_BINARY_GT uint16 = 114;
// const OP_BINARY_LT uint16 = 115;

// const OP_BINARY_AND uint16 = 120;
// const OP_BINARY_OR uint16 = 121;
// const OP_BINARY_COALESCE uint16 = 122;

type PrqlVirtualMachineParams struct {
	DebugLevel     int
	VerbosityLevel int
	InputPath      string
}

type PrqlVirtualMachine struct {
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
	vm := PrqlVirtualMachine{__inputPath: params.InputPath, __debugLevel: params.DebugLevel}
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

func (vm *PrqlVirtualMachine) StackPush(e *PrqlInternal) {
	vm.__stack = append(vm.__stack, e)
}

func (vm *PrqlVirtualMachine) StackPop() *PrqlInternal {
	e := vm.__stack[len(vm.__stack)-1]
	vm.__stack = vm.__stack[:len(vm.__stack)-1]
	return e
}

func (vm *PrqlVirtualMachine) ReadPrqlInstructions(bytes []byte, offset uint64) {

	var opCode uint16
	var param1 uint16
	var param2 []byte

	usize := uint64(len(bytes))

	for offset < usize {
		opCode = binary.BigEndian.Uint16(bytes[offset : offset+2])
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

			// Standard library functions
			case "derive":
				PrqlFunc_Derive(vm)
			case "from":
				PrqlFunc_From(vm)
			case "export_csv":
				PrqlFunc_ExportCsv(vm)
			case "import_csv":
				PrqlFunc_ImportCsv(vm)
			case "select":
				PrqlFunc_Select(vm)

			// User defined functions
			default:
				if f, ok := vm.__userDefinedVariables[funcName]; ok {
					switch t := f.Value.(type) {
					case UserDefinedFunction:
						t(vm)
					default:
						vm.StackPush(NewPrqlInternalError(fmt.Sprintf("variable '%s' not callable.", funcName)))
					}
				} else {
					vm.StackPush(NewPrqlInternalError(fmt.Sprintf("name '%s' not found.", funcName)))
				}
			}

			vm.__funcNumParams = 0

		case OP_BEGIN_LIST:
			if vm.__debugLevel > 10 {
				fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "OP_BEGIN_LIST", "", "", "")
			}

		case OP_END_LIST:
			if vm.__debugLevel > 10 {
				fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "OP_END_LIST", "", "", "")
			}

		case OP_ADD_FUNC_PARAM:
			paramName := vm.__symbolTable[binary.BigEndian.Uint64(param2)]
			if vm.__debugLevel > 10 {
				fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "OP_ADD_FUNC_PARAM", "", paramName, "")
			}

			vm.StackPush(NewPrqlInternalParamName(paramName))

		case OP_ADD_EXPR_TERM:
			if vm.__debugLevel > 10 {
				fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "OP_ADD_EXPR_TERM", "", "", "")
			}

		case OP_PUSH_NAMED_PARAM:
			if vm.__debugLevel > 10 {
				fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "OP_PUSH_NAMED_PARAM", "", "", "")
			}

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
			// 	vm.StackPush(NewPrqlInternalTerm(nil))

			case TERM_BOOL:
				termType = "BOOL"
				termVal = "true"
				val := true
				if binary.BigEndian.Uint64(param2) == 0 {
					val = false
					termVal = "false"
				}

				vm.StackPush(NewPrqlInternalTerm(val))

			// case TERM_INTEGER:

			case TERM_FLOAT:
				termType = "FLOAT"
				val := math.Float64frombits(binary.LittleEndian.Uint64(param2))
				termVal = fmt.Sprintf("%f", val)

				vm.StackPush(NewPrqlInternalTerm(val))

			case TERM_STRING:
				termType = "STRING"
				termVal = vm.__symbolTable[binary.BigEndian.Uint64(param2)]

				vm.StackPush(NewPrqlInternalTerm(termVal))

			case TERM_IDENT:
				termType = "IDENT"
				termVal = vm.__symbolTable[binary.BigEndian.Uint64(param2)]

				vm.StackPush(vm.IdentResolutionStrategy(termVal))

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

		case OP_GOTO:
			if vm.__debugLevel > 10 {
				fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "OP_GOTO", "", "", "")
			}

		case OP_BINARY_MUL:
			if vm.__debugLevel > 10 {
				fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "OP_BINARY_MUL", "", "", "")
			}

		case OP_BINARY_DIV:
			if vm.__debugLevel > 10 {
				fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "OP_BINARY_DIV", "", "", "")
			}

		case OP_BINARY_MOD:
			if vm.__debugLevel > 10 {
				fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "OP_BINARY_MOD", "", "", "")
			}

		case OP_BINARY_ADD:
			if vm.__debugLevel > 10 {
				fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "OP_BINARY_ADD", "", "", "")
			}

		case OP_BINARY_SUB:
			if vm.__debugLevel > 10 {
				fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "OP_BINARY_SUB", "", "", "")
			}

		}
	}
}

func (vm *PrqlVirtualMachine) IdentResolutionStrategy(ident string) *PrqlInternal {
	if vm.__currentTable != nil {
		switch t := vm.__currentTable.Value.(type) {
		case dataframe.DataFrame:
			names := t.Names()
			for _, name := range names {
				if name == ident {
					return NewPrqlInternalTerm(t.Col(name))
				}
			}
		default:
		}
	}

	if v, ok := vm.__userDefinedVariables[ident]; ok {
		return &v
	}

	return NewPrqlInternalError(fmt.Sprintf("name '%s' not found.", ident))
}
