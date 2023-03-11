package preludio

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io/fs"
	"os"
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

type OPCODE uint8
type PARAM1 uint8

const (
	TERM_NULL     PARAM1 = 0
	TERM_BOOL     PARAM1 = 1
	TERM_INTEGER  PARAM1 = 2
	TERM_FLOAT    PARAM1 = 3
	TERM_STRING   PARAM1 = 4
	TERM_INTERVAL PARAM1 = 5
	TERM_RANGE    PARAM1 = 6
	TERM_LIST     PARAM1 = 7
	TERM_PIPELINE PARAM1 = 8
	TERM_SYMBOL   PARAM1 = 10
)

const (
	OP_START_PIPELINE    OPCODE = 0
	OP_END_PIPELINE      OPCODE = 1
	OP_ASSIGN_STMT       OPCODE = 2
	OP_START_FUNC_CALL   OPCODE = 3
	OP_MAKE_FUNC_CALL    OPCODE = 4
	OP_START_LIST        OPCODE = 5
	OP_END_LIST          OPCODE = 6
	OP_ADD_FUNC_PARAM    OPCODE = 7
	OP_ADD_EXPR_TERM     OPCODE = 8
	OP_PUSH_NAMED_PARAM  OPCODE = 9
	OP_PUSH_ASSIGN_IDENT OPCODE = 10
	OP_PUSH_TERM         OPCODE = 11
	OP_END_CHUNCK        OPCODE = 12
	OP_GOTO              OPCODE = 50

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

	OP_UNARY_ADD OPCODE = 130
	OP_UNARY_SUB OPCODE = 131
	OP_UNARY_NOT OPCODE = 132

	NO_OP = 255
)

// ByteEater is the name of the Preludio Virtual Machine
type ByteEater struct {
	__reportWarnings bool
	__isCLI          bool
	__printToStdout  bool
	__debugLevel     int
	__verbosityLevel int
	__inputPath      string

	__symbolTable           []string
	__stack                 []__p_intern__
	__currentDataFrame      *dataframe.DataFrame
	__currentDataFrameNames map[string]bool
	__globalNameSpace       map[string]*__p_intern__
	__pipelineNameSpace     map[string]*__p_intern__
	__funcNumParams         int
	__listElementCounters   []int
	__output                PreludioOutput
}

func (vm *ByteEater) SetPrintWarning(flag bool) *ByteEater {
	vm.__reportWarnings = flag
	return vm
}

func (vm *ByteEater) SetCLI(flag bool) *ByteEater {
	vm.__isCLI = flag
	return vm
}

func (vm *ByteEater) SetPrintToStdout(flag bool) *ByteEater {
	vm.__printToStdout = flag
	return vm
}

func (vm *ByteEater) SetDebugLevel(level int) *ByteEater {
	vm.__debugLevel = level
	return vm
}

func (vm *ByteEater) SetVerbosityLevel(level int) *ByteEater {
	vm.__verbosityLevel = level
	return vm
}

func (vm *ByteEater) InitVM() *ByteEater {
	vm.__currentDataFrameNames = map[string]bool{}
	vm.__globalNameSpace = map[string]*__p_intern__{}
	vm.__pipelineNameSpace = map[string]*__p_intern__{}
	vm.__currentDataFrame = nil

	return vm
}

type PreludioOutput struct {
	Log       []string   `json:"log"`
	Errors    []string   `json:"errors"`
	Schema    []string   `json:"schema"`
	DataFrame [][]string `json:"result"`
}

func (vm *ByteEater) GetLog() []string {
	return vm.__output.Log
}

func (vm *ByteEater) PrintLog() {
	for _, l := range vm.__output.Log {
		fmt.Println(l)
	}
}

// Run Preludio Bytecode from byte array
func (vm *ByteEater) RunBytecode(bytecode []byte) *PreludioOutput {

	// clean vm state
	vm.__symbolTable = make([]string, 0)
	vm.__stack = make([]__p_intern__, 0)

	// set a new output for the new computation
	vm.__output = PreludioOutput{Log: make([]string, 0), Errors: make([]string, 0)}

	bytemark := bytecode[0:4]
	__symbolTableSize := binary.BigEndian.Uint32(bytecode[4:8])

	if vm.__printToStdout && vm.__debugLevel > 5 {
		fmt.Println()
		fmt.Printf("BYTECODE INFO\n")
		fmt.Printf("=============\n")
		fmt.Printf("SIZE:              %d\n", len(bytecode))
		fmt.Printf("BYTE MARK:         %x %x %x %x\n", bytemark[0], bytemark[1], bytemark[2], bytemark[3])
		fmt.Printf("SYMBOL TABLE SIZE: %d\n\n", __symbolTableSize)
		fmt.Printf("STRING SYMBOLS\n")
		fmt.Printf("==============\n")
	}

	offset := uint32(8)
	for i := uint32(0); i < __symbolTableSize; i++ {
		l := binary.BigEndian.Uint32(bytecode[offset : offset+4])
		offset += 4

		v := string(bytecode[offset : offset+l])
		vm.__symbolTable = append(vm.__symbolTable, v)
		offset += l
	}

	if vm.__printToStdout && vm.__debugLevel > 5 {
		for _, symbol := range vm.__symbolTable {
			fmt.Printf("%s\n", symbol)
		}

		fmt.Println()
		fmt.Printf("INSTRUCTIONS\n")
		fmt.Printf("============\n")
	}

	vm.RunPrqlInstructions(bytecode, offset)

	return &vm.__output
}

// Run Preludio bytecode from a binary file located
// at __inputPath - SetInputPath
// TO DEPRECATE (?)
func (vm *ByteEater) RunFileBytecode() *PreludioOutput {
	var err error
	var file *os.File
	var stats fs.FileInfo

	// clean vm state
	vm.__symbolTable = make([]string, 0)
	vm.__stack = make([]__p_intern__, 0)

	// set a new output for the new computation
	vm.__output = PreludioOutput{Log: make([]string, 0), Errors: make([]string, 0)}

	file, err = os.Open(vm.__inputPath)
	if err != nil {
		vm.__output.Errors = append(vm.__output.Errors, err.Error())
		return &vm.__output
	}
	defer file.Close()

	stats, err = file.Stat()
	if err != nil {
		vm.__output.Errors = append(vm.__output.Errors, err.Error())
		return &vm.__output
	}

	var size int64 = stats.Size()
	bytecode := make([]byte, size)

	bufr := bufio.NewReader(file)
	_, err = bufr.Read(bytecode)
	if err != nil {
		vm.__output.Errors = append(vm.__output.Errors, err.Error())
		return &vm.__output
	}

	bytemark := bytecode[0:4]
	__symbolTableSize := binary.BigEndian.Uint32(bytecode[4:8])

	if vm.__printToStdout && vm.__debugLevel > 5 {
		fmt.Println()
		fmt.Printf("BYTECODE INFO\n")
		fmt.Printf("=============\n")
		fmt.Printf("SIZE:              %d\n", size)
		fmt.Printf("BYTE MARK:         %x %x %x %x\n", bytemark[0], bytemark[1], bytemark[2], bytemark[3])
		fmt.Printf("SYMBOL TABLE SIZE: %d\n\n", __symbolTableSize)
		fmt.Printf("STRING SYMBOLS\n")
		fmt.Printf("==============\n")
	}

	offset := uint32(8)
	for i := uint32(0); i < __symbolTableSize; i++ {
		l := binary.BigEndian.Uint32(bytecode[offset : offset+4])
		offset += 4

		v := string(bytecode[offset : offset+l])
		vm.__symbolTable = append(vm.__symbolTable, v)
		offset += l
	}

	if vm.__printToStdout && vm.__debugLevel > 5 {
		for _, symbol := range vm.__symbolTable {
			fmt.Printf("%s\n", symbol)
		}

		fmt.Println()
		fmt.Printf("INSTRUCTIONS\n")
		fmt.Printf("============\n")
	}

	vm.RunPrqlInstructions(bytecode, offset)

	return &vm.__output
}

func (vm *ByteEater) StackIsEmpty() bool {
	return len(vm.__stack) == 0
}

func (vm *ByteEater) StackPush(e *__p_intern__) {
	vm.__stack = append(vm.__stack, *e)
}

func (vm *ByteEater) StackPop() *__p_intern__ {
	e := vm.__stack[len(vm.__stack)-1]
	vm.__stack = vm.__stack[:len(vm.__stack)-1]
	return &e
}

func (vm *ByteEater) StackLast() *__p_intern__ {
	return &vm.__stack[len(vm.__stack)-1]
}

func (vm *ByteEater) LoadResult() {

	for !vm.StackIsEmpty() {
		vm.__output.Log = append(vm.__output.Log, fmt.Sprintf("%s", vm.StackPop().getValue()))
	}

	// if !vm.StackIsEmpty() {
	// 	switch res := vm.StackPop().getValue().(type) {
	// 	case dataframe.DataFrame:
	// 		vm.__output.DataFrame = make([][]string, res.Ncol())
	// 		for i, name := range res.Names() {
	// 			vm.__output.DataFrame[i] = make([]string, res.Nrow())
	// 			for j, val := range res.Col(name).Records() {
	// 				vm.__output.DataFrame[i][j] = val
	// 			}
	// 		}
	// 	default:
	// 		vm.__output.DataFrame = [][]string{{fmt.Sprintf("%s", res)}}
	// 	}
	// }

	// if vm.__printToStdout && vm.__debugLevel > 5 {
	// 	fmt.Println("STACK DUMP")
	// 	for !vm.StackIsEmpty() {
	// 		fmt.Printf("%s\n", vm.StackPop().getValue())
	// 	}
	// }
}

func (vm *ByteEater) RunPrqlInstructions(bytes []byte, offset uint32) {

	opCode := OPCODE(0)
	param1 := PARAM1(0)
	param2 := []byte{0, 0, 0, 0}

	usize := uint32(len(bytes))

MAIN_LOOP:
	for offset < usize {

		opCode = OPCODE(bytes[offset])
		offset++
		param1 = PARAM1(bytes[offset])
		offset++

		param2[0] = bytes[offset]
		offset++
		param2[1] = bytes[offset]
		offset++
		param2[2] = bytes[offset]
		offset++
		param2[3] = bytes[offset]
		offset++

		switch opCode {

		///////////////////////////////////////////////////////////////////////
		///////////				PIPELINE OPERATIONS					///////////
		case OP_START_PIPELINE:
			if vm.__printToStdout && vm.__debugLevel > 10 {
				fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "OP_START_PIPELINE", "", "", "")
			}

			// Insert BEGIN FRAME
			vm.StackPush(newPInternBeginFrame())

		case OP_END_PIPELINE:
			if vm.__printToStdout && vm.__debugLevel > 10 {
				fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "OP_END_PIPELINE", "", "", "")
			}

			// Extract BEGIN FRAME
			vm.StackPop()

		case OP_ASSIGN_STMT:
			if vm.__printToStdout && vm.__debugLevel > 10 {
				fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "OP_ASSIGN_STMT", "", "", "")
			}

		case OP_START_FUNC_CALL:
			if vm.__printToStdout && vm.__debugLevel > 10 {
				fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "OP_START_FUNC_CALL", "", "", "")
			}

		case OP_MAKE_FUNC_CALL:
			funcName := vm.__symbolTable[binary.BigEndian.Uint32(param2)]
			if vm.__printToStdout && vm.__debugLevel > 10 {
				fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "OP_MAKE_FUNC_CALL", "", funcName, "")
			}

			switch funcName {

			// Standard library functions build-ins
			case "derive":
				PreludioFunc_Derive("derive", vm)
			case "describe":
				PreludioFunc_Describe("describe", vm)
			case "from":
				PreludioFunc_From("from", vm)
			case "exportCSV":
				PreludioFunc_ExportCsv("exportCSV", vm)
			case "importCSV":
				PreludioFunc_ImportCsv("importCSV", vm)
			case "new":
				PreludioFunc_New("new", vm)
			case "select":
				PreludioFunc_Select("select", vm)
			case "sort":
				PreludioFunc_Sort("sort", vm)
			case "take":
				PreludioFunc_Take("take", vm)

			// Environment functions
			case "toCurrent":
				PreludioFunc_ToCurrent("toCurrent", vm)

			// Coerce functions
			case "asBool":
				PreludioFunc_AsBool("asBool", vm)
			case "asInteger":
				PreludioFunc_AsInteger("asInteger", vm)
			case "asFloat":
				PreludioFunc_AsFloat("asFloat", vm)
			case "asString":
				PreludioFunc_AsString("asString", vm)

			// String functions
			case "strReplace":
				PreludioFunc_StrReplace("strReplace", vm)

			// User defined functions
			default:
				if internal, ok := vm.__globalNameSpace[funcName]; ok {
					switch value := internal.getValue().(type) {
					case UserDefinedFunction:
						value(vm)
					default:
						vm.StackPush(newPInternError(fmt.Sprintf("variable '%s' not callable.", funcName)))
					}
				} else {
					vm.StackPush(newPInternError(fmt.Sprintf("variable '%s' not defined.", funcName)))
				}
			}

			vm.__funcNumParams = 0

		case OP_START_LIST:
			if vm.__printToStdout && vm.__debugLevel > 10 {
				fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "OP_START_LIST", "", "", "")
			}

			vm.__listElementCounters = append(vm.__listElementCounters, 0)

		case OP_END_LIST:
			if vm.__printToStdout && vm.__debugLevel > 10 {
				fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "OP_END_LIST", "", "", "")
			}

			stackLen := len(vm.__stack)
			listLen := vm.__listElementCounters[len(vm.__listElementCounters)-1]

			listCopy := make([]__p_intern__, listLen)
			copy(listCopy, vm.__stack[stackLen-listLen:])
			vm.__stack = vm.__stack[:stackLen-listLen]

			vm.StackPush(newPInternTerm(__p_list__(listCopy)))

			vm.__listElementCounters = vm.__listElementCounters[:len(vm.__listElementCounters)-1]

		case OP_ADD_FUNC_PARAM:
			if vm.__printToStdout && vm.__debugLevel > 10 {
				fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "OP_ADD_FUNC_PARAM", "", "", "")
			}

		case OP_ADD_EXPR_TERM:
			if vm.__printToStdout && vm.__debugLevel > 10 {
				fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "OP_ADD_EXPR_TERM", "", "", "")
			}

		///////////////////////////////////////////////////////////////////////
		///////////					PUSH NAMED PARAM
		///////////
		///////////	Set the last element on the stack as a named
		///////////	parameter.
		case OP_PUSH_NAMED_PARAM:
			paramName := vm.__symbolTable[binary.BigEndian.Uint32(param2)]
			if vm.__printToStdout && vm.__debugLevel > 10 {
				fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "OP_PUSH_NAMED_PARAM", "", paramName, "")
			}

			vm.StackLast().setParamName(paramName)

		///////////////////////////////////////////////////////////////////////
		///////////					PUSH ASSIGN IDENT
		///////////
		///////////	Set the last element on the stack as an assigned
		///////////	expression.
		case OP_PUSH_ASSIGN_IDENT:
			ident := vm.__symbolTable[binary.BigEndian.Uint32(param2)]
			if vm.__printToStdout && vm.__debugLevel > 10 {
				fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "OP_PUSH_ASSIGN_IDENT", "", ident, "")
			}

			vm.StackLast().setAssignment(ident)

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
				if binary.BigEndian.Uint32(param2) == 0 {
					val = false
					termVal = "false"
				}
				vm.StackPush(newPInternTerm([]bool{val}))

			case TERM_INTEGER:
				termType = "INTEGER"
				termVal = vm.__symbolTable[binary.BigEndian.Uint32(param2)]
				val, _ := strconv.ParseInt(termVal, 10, 64)
				vm.StackPush(newPInternTerm([]int{int(val)}))

			case TERM_FLOAT:
				termType = "FLOAT"
				termVal = vm.__symbolTable[binary.BigEndian.Uint32(param2)]
				val, _ := strconv.ParseFloat(termVal, 64)
				vm.StackPush(newPInternTerm([]float64{val}))

			case TERM_STRING:
				termType = "STRING"
				termVal = vm.__symbolTable[binary.BigEndian.Uint32(param2)]
				vm.StackPush(newPInternTerm([]string{termVal}))

			case TERM_SYMBOL:
				termType = "SYMBOL"
				termVal = vm.__symbolTable[binary.BigEndian.Uint32(param2)]
				vm.StackPush(newPInternTerm(__p_symbol__(termVal)))

			default:
				vm.StackPush(newPInternError(fmt.Sprintf("ByteEater: unknown term code %d.", param1)))

			}

			if vm.__printToStdout && vm.__debugLevel > 10 {
				fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "OP_PUSH_TERM", termType, termVal, "")
			}

		case OP_END_CHUNCK:
			if vm.__printToStdout && vm.__debugLevel > 10 {
				fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "OP_END_CHUNCK", "", "", "")
			}

			vm.__funcNumParams += 1
			if len(vm.__listElementCounters) > 0 {
				vm.__listElementCounters[len(vm.__listElementCounters)-1]++
			}

		case OP_GOTO:
			if vm.__printToStdout && vm.__debugLevel > 10 {
				fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "OP_GOTO", "", "", "")
			}

		///////////////////////////////////////////////////////////////////////
		///////////				ARITHMETIC OPERATIONS
		case OP_BINARY_MUL:
			if vm.__printToStdout && vm.__debugLevel > 10 {
				fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "OP_BINARY_MUL", "", "", "")
			}

			op2 := vm.StackPop()
			vm.StackLast().appendOperand(OP_BINARY_MUL, op2)

		case OP_BINARY_DIV:
			if vm.__printToStdout && vm.__debugLevel > 10 {
				fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "OP_BINARY_DIV", "", "", "")
			}

			op2 := vm.StackPop()
			vm.StackLast().appendOperand(OP_BINARY_DIV, op2)

		case OP_BINARY_MOD:
			if vm.__printToStdout && vm.__debugLevel > 10 {
				fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "OP_BINARY_MOD", "", "", "")
			}

			op2 := vm.StackPop()
			vm.StackLast().appendOperand(OP_BINARY_MOD, op2)

		case OP_BINARY_ADD:
			if vm.__printToStdout && vm.__debugLevel > 10 {
				fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "OP_BINARY_ADD", "", "", "")
			}

			op2 := vm.StackPop()
			vm.StackLast().appendOperand(OP_BINARY_ADD, op2)

		case OP_BINARY_SUB:
			if vm.__printToStdout && vm.__debugLevel > 10 {
				fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "OP_BINARY_SUB", "", "", "")
			}

			op2 := vm.StackPop()
			vm.StackLast().appendOperand(OP_BINARY_SUB, op2)

		case OP_BINARY_POW:
			if vm.__printToStdout && vm.__debugLevel > 10 {
				fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "OP_BINARY_POW", "", "", "")
			}

			op2 := vm.StackPop()
			vm.StackLast().appendOperand(OP_BINARY_POW, op2)

		///////////////////////////////////////////////////////////////////////
		///////////				LOGICAL OPERATIONS

		case OP_BINARY_EQ:
			if vm.__printToStdout && vm.__debugLevel > 10 {
				fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "OP_BINARY_EQ", "", "", "")
			}

			op2 := vm.StackPop()
			vm.StackLast().appendOperand(OP_BINARY_EQ, op2)

		case OP_BINARY_NE:
			if vm.__printToStdout && vm.__debugLevel > 10 {
				fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "OP_BINARY_NE", "", "", "")
			}

			op2 := vm.StackPop()
			vm.StackLast().appendOperand(OP_BINARY_NE, op2)

		case OP_BINARY_GE:
			if vm.__printToStdout && vm.__debugLevel > 10 {
				fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "OP_BINARY_GE", "", "", "")
			}

			op2 := vm.StackPop()
			vm.StackLast().appendOperand(OP_BINARY_GE, op2)

		case OP_BINARY_LE:
			if vm.__printToStdout && vm.__debugLevel > 10 {
				fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "OP_BINARY_LE", "", "", "")
			}

			op2 := vm.StackPop()
			vm.StackLast().appendOperand(OP_BINARY_LE, op2)

		case OP_BINARY_GT:
			if vm.__printToStdout && vm.__debugLevel > 10 {
				fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "OP_BINARY_GT", "", "", "")
			}

			op2 := vm.StackPop()
			vm.StackLast().appendOperand(OP_BINARY_GT, op2)

		case OP_BINARY_LT:
			if vm.__printToStdout && vm.__debugLevel > 10 {
				fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "OP_BINARY_LT", "", "", "")
			}

			op2 := vm.StackPop()
			vm.StackLast().appendOperand(OP_BINARY_LT, op2)

		case OP_BINARY_AND:
			if vm.__printToStdout && vm.__debugLevel > 10 {
				fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "OP_BINARY_AND", "", "", "")
			}

			op2 := vm.StackPop()
			vm.StackLast().appendOperand(OP_BINARY_AND, op2)

		case OP_BINARY_OR:
			if vm.__printToStdout && vm.__debugLevel > 10 {
				fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "OP_BINARY_OR", "", "", "")
			}

			op2 := vm.StackPop()
			vm.StackLast().appendOperand(OP_BINARY_OR, op2)

		///////////////////////////////////////////////////////////////////////
		///////////				OTHER OPERATIONS

		case OP_BINARY_COALESCE:
		case OP_BINARY_MODEL:

		///////////////////////////////////////////////////////////////////////
		///////////				UNARY OPERATIONS
		case OP_UNARY_SUB:
			if vm.__printToStdout && vm.__debugLevel > 10 {
				fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "OP_UNARY_SUB", "", "", "")
			}

		case OP_UNARY_ADD:
		case OP_UNARY_NOT:

		}

		if !vm.StackIsEmpty() && vm.StackLast().tag == PRELUDIO_INTERNAL_TAG_ERROR {
			for !vm.StackIsEmpty() && vm.StackLast().tag == PRELUDIO_INTERNAL_TAG_ERROR {
				err := vm.StackPop()
				vm.PrintError(err.errorMsg)
			}
			break MAIN_LOOP
		}

	}
}

// Process and get the parameters from the VM stack
// ready to be used for the function call.
//
// funcName: name of the caller function (for error reporting)
//
// namedParams: map with the names of the expecting named parameters (nil if none)
//
// acceptingAssignments:
//
// solve: if false, parameters expression won't be solved
func (vm *ByteEater) GetFunctionParams(funcName string, namedParams *map[string]*__p_intern__, acceptingAssignments bool, solve bool) ([]*__p_intern__, map[string]*__p_intern__, error) {

	var assignments map[string]*__p_intern__
	if acceptingAssignments {
		assignments = map[string]*__p_intern__{}
	}

	positionalParams := make([]*__p_intern__, 0)

LOOP1:
	for {
		t1 := *vm.StackLast()
		switch t1.tag {
		case PRELUDIO_INTERNAL_TAG_ERROR:
		case PRELUDIO_INTERNAL_TAG_EXPRESSION:
			positionalParams = append([]*__p_intern__{&t1}, positionalParams...)
			vm.StackPop()

		case PRELUDIO_INTERNAL_TAG_NAMED_PARAM:
			// Name of parameter is in the given list of names
			if _, ok := (*namedParams)[t1.name]; ok {
				(*namedParams)[t1.name] = &t1
			} else {
				vm.PrintWarning(fmt.Sprintf("%s does not know a parameter named '%s', the value will be ignored.", funcName, t1.name))
			}
			vm.StackPop()

		case PRELUDIO_INTERNAL_TAG_ASSIGNMENT:
			if acceptingAssignments {
				assignments[t1.name] = &t1
			} else {
				vm.PrintWarning(fmt.Sprintf("%s does not accept assignements, the value of '%s' will be ignored.", funcName, t1.name))
			}
			vm.StackPop()

		case PRELUDIO_INTERNAL_TAG_BEGIN_FRAME:
			break LOOP1
		}
	}

	if solve {
		for _, p := range positionalParams {
			if err := p.solve(vm); err != nil {
				return positionalParams, assignments, err
			}
		}

		if namedParams != nil {
			for _, p := range *namedParams {
				if err := p.solve(vm); err != nil {
					return positionalParams, assignments, err
				}
			}
		}

		if acceptingAssignments {
			for _, p := range assignments {
				if err := p.solve(vm); err != nil {
					return positionalParams, assignments, err
				}
			}
		}
	}

	return positionalParams, assignments, nil
}

func (vm *ByteEater) SymbolResolution(symbol __p_symbol__) interface{} {
	// 1 - Look at the current DataFrame
	if vm.__currentDataFrame != nil {
		if ok := vm.__currentDataFrameNames[string(symbol)]; ok {
			ser := vm.__currentDataFrame.Col(string(symbol))
			switch ser.Type() {
			case series.Bool:
				val, _ := ser.Bool()
				return val
			case series.Int:
				val, _ := ser.Int()
				return val
			case series.Float:
				return ser.Float()
			case series.String:
				return ser.Records()
			}
		}
	}

	// 2 - Try to split the symbol into pieces
	pieces := strings.Split(string(symbol), ".")
	return pieces
}

// Set the last element inserted into the stack as
// the current DataFrame
func (vm *ByteEater) SetCurrentDataFrame() {
	df, _ := vm.StackLast().getDataframe()
	vm.__currentDataFrame = &df

	vm.__currentDataFrameNames = map[string]bool{}
	for _, name := range df.Names() {
		vm.__currentDataFrameNames[name] = true
	}
}

var WARNING_STYLE = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#ffff00")).
	Bold(true)

var ERROR_STYLE = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#ff8787")).
	Bold(true)

func (vm *ByteEater) PrintWarning(msg string) {
	if vm.__reportWarnings {
		if vm.__isCLI {
			vm.__output.Log = append(vm.__output.Log, WARNING_STYLE.Render(fmt.Sprintf("[ ⚠️ Warning ] %s\n", msg)))
		} else {
			vm.__output.Log = append(vm.__output.Log, fmt.Sprintf("[ ⚠️ Warning ] %s\n", msg))
		}

		if vm.__printToStdout {
			fmt.Printf(WARNING_STYLE.Render(fmt.Sprintf("[ ⚠️ Warning ] %s\n", msg)))
		}
	}
}

func (vm *ByteEater) PrintError(msg string) {
	if vm.__isCLI {
		vm.__output.Log = append(vm.__output.Log, ERROR_STYLE.Render(fmt.Sprintf("[ ☠️ Error ] %s\n", msg)))
	} else {
		vm.__output.Log = append(vm.__output.Log, fmt.Sprintf("[ ☠️ Error ] %s\n", msg))
	}

	if vm.__printToStdout {
		fmt.Printf(ERROR_STYLE.Render(fmt.Sprintf("[ ☠️ Error ] %s\n", msg)))
	}
}
