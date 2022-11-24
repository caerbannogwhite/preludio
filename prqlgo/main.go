package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"os"
)

const TERM_NULL uint16 = 0
const TERM_BOOL uint16 = 1
const TERM_NUMERIC uint16 = 2
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

type PRQLVirtualMachine struct {
	__debugLevel     int
	__verbosityLevel int
	InputPath        string
	SymbolTable      []string
	__stack          []*PRQLInternal
	__functions      map[string]PRQLFunction
	__funcNumParams  uint64
}

func main() {
	inputPath := os.Args[1]
	vm := PRQLVirtualMachine{InputPath: inputPath, __debugLevel: 20}

	vm.__functions = map[string]PRQLFunction{}
	vm.__functions["derive"] = prql_derive
	vm.__functions["from"] = prql_from
	vm.__functions["export_csv"] = prql_export_csv
	vm.__functions["import_csv"] = prql_import_csv
	vm.__functions["select"] = prql_select

	vm.read_prql_bytecode()
}

func (vm *PRQLVirtualMachine) read_prql_bytecode() {
	file, err := os.Open(vm.InputPath)

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
	symbolTableSize := binary.BigEndian.Uint64(bytes[8:16])

	fmt.Printf("SIZE:              %d\n", size)
	fmt.Printf("BYTE MARK:         %x %x %x %x\n", bytemark[0], bytemark[1], bytemark[2], bytemark[3])
	fmt.Printf("SYMBOL TABLE SIZE: %d\n\n", symbolTableSize)

	fmt.Printf("STRING SYMBOLS\n")
	fmt.Printf("==============\n")
	offset := uint64(16)
	for i := uint64(0); i < symbolTableSize; i++ {
		l := binary.BigEndian.Uint64(bytes[offset : offset+8])
		offset += 8

		v := string(bytes[offset : offset+l])
		vm.SymbolTable = append(vm.SymbolTable, v)
		fmt.Printf("%s\n", v)
		offset += l
	}

	vm.read_instructions(bytes, offset)

	// return bytes, err
}

func (vm *PRQLVirtualMachine) stackPush(e *PRQLInternal) {
	vm.__stack = append(vm.__stack, e)
}

func (vm *PRQLVirtualMachine) stackPop() *PRQLInternal {
	e := vm.__stack[len(vm.__stack)-1]
	vm.__stack = vm.__stack[:len(vm.__stack)-1]
	return e
}

func (vm *PRQLVirtualMachine) read_instructions(bytes []byte, offset uint64) {

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
				fmt.Printf("%30.30s | %30.30s | %30.30s | %50.50s \n", "OP_BEGIN_PIPELINE", "", "", "")
			}

		case OP_END_PIPELINE:
			if vm.__debugLevel > 10 {
				fmt.Printf("%30.30s | %30.30s | %30.30s | %50.50s \n", "OP_END_PIPELINE", "", "", "")
			}

		case OP_ASSIGN_STMT:
			if vm.__debugLevel > 10 {
				fmt.Printf("%30.30s | %30.30s | %30.30s | %50.50s \n", "OP_ASSIGN_STMT", "", "", "")
			}

		// case OP_BEGIN_FUNC_CALL:3
		case OP_MAKE_FUNC_CALL:
			funcName := vm.SymbolTable[binary.BigEndian.Uint64(param2)]
			if vm.__debugLevel > 10 {
				fmt.Printf("%30.30s | %30.30s | %30.30s | %50.50s \n", "OP_MAKE_FUNC_CALL", "", funcName, "")
			}

			if f, ok := vm.__functions[funcName]; ok {
				f(vm)
			} else {
				vm.stackPush(NewPRQLInternalError(fmt.Sprintf("function '%s' not found", funcName)))
			}

			vm.__funcNumParams = 0

		case OP_BEGIN_LIST:
			if vm.__debugLevel > 10 {
				fmt.Printf("%30.30s | %30.30s | %30.30s | %50.50s \n", "OP_BEGIN_LIST", "", "", "")
			}

		case OP_END_LIST:
			if vm.__debugLevel > 10 {
				fmt.Printf("%30.30s | %30.30s | %30.30s | %50.50s \n", "OP_END_LIST", "", "", "")
			}

		case OP_ADD_FUNC_PARAM:
			paramName := vm.SymbolTable[binary.BigEndian.Uint64(param2)]
			if vm.__debugLevel > 10 {
				fmt.Printf("%30.30s | %30.30s | %30.30s | %50.50s \n", "OP_ADD_FUNC_PARAM", "", paramName, "")
			}

			vm.stackPush(NewPRQLInternalParamName(paramName))

		case OP_ADD_EXPR_TERM:
			if vm.__debugLevel > 10 {
				fmt.Printf("%30.30s | %30.30s | %30.30s | %50.50s \n", "OP_ADD_EXPR_TERM", "", "", "")
			}

		case OP_PUSH_NAMED_PARAM:
			if vm.__debugLevel > 10 {
				fmt.Printf("%30.30s | %30.30s | %30.30s | %50.50s \n", "OP_PUSH_NAMED_PARAM", "", "", "")
			}

		case OP_PUSH_ASSIGN_IDENT:
			ident := vm.SymbolTable[binary.BigEndian.Uint64(param2)]
			if vm.__debugLevel > 10 {
				fmt.Printf("%30.30s | %30.30s | %30.30s | %50.50s \n", "OP_PUSH_ASSIGN_IDENT", "", ident, "")
			}

			vm.stackPush(NewPRQLInternalParamName(ident))

		case OP_PUSH_TERM:
			termType := ""
			termVal := ""

			switch param1 {
			case TERM_NULL:
				termType = "NULL"
			case TERM_BOOL:
				termType = "BOOL"

			case TERM_NUMERIC:
				termType = "NUMERIC"

			case TERM_STRING:
				termType = "STRING"
				termVal = vm.SymbolTable[binary.BigEndian.Uint64(param2)]
			case TERM_IDENT:
				termType = "IDENT"
				termVal = vm.SymbolTable[binary.BigEndian.Uint64(param2)]
			}

			if vm.__debugLevel > 10 {
				fmt.Printf("%30.30s | %30.30s | %30.30s | %50.50s \n", "OP_PUSH_TERM", termType, termVal, "")
			}

		case OP_END_FUNC_CALL_PARAM:
			if vm.__debugLevel > 10 {
				fmt.Printf("%30.30s | %30.30s | %30.30s | %50.50s \n", "OP_END_FUNC_CALL_PARAM", "", "", "")
			}

		case OP_GOTO:
			if vm.__debugLevel > 10 {
				fmt.Printf("%30.30s | %30.30s | %30.30s | %50.50s \n", "OP_GOTO", "", "", "")
			}
		}
	}
}
