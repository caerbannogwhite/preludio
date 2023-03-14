package preludio

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io/fs"
	"os"
	"strconv"
	"strings"

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
	// parameters
	__printWarnings  bool
	__isCLI          bool
	__printToStdout  bool
	__outputSnippets bool
	__verbose        bool
	__debugLevel     int
	__inputPath      string

	// internal
	__panicMode             bool
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
	vm.__printWarnings = flag
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

func (vm *ByteEater) SetOutputSnippets(flag bool) *ByteEater {
	vm.__outputSnippets = flag
	return vm
}

func (vm *ByteEater) SetVerbose(flag bool) *ByteEater {
	vm.__verbose = flag
	return vm
}

func (vm *ByteEater) SetDebugLevel(level int) *ByteEater {
	vm.__debugLevel = level
	return vm
}

func (vm *ByteEater) InitVM() *ByteEater {
	vm.__currentDataFrameNames = map[string]bool{}
	vm.__globalNameSpace = map[string]*__p_intern__{}
	vm.__pipelineNameSpace = map[string]*__p_intern__{}
	vm.__currentDataFrame = nil

	return vm
}

type LOG_TYPE uint8

const (
	LOG_INFO    LOG_TYPE = 0
	LOG_WARNING LOG_TYPE = 1
	LOG_ERROR   LOG_TYPE = 2
	LOG_DEBUG   LOG_TYPE = 3
)

// Run Preludio Bytecode from byte array
func (vm *ByteEater) RunBytecode(bytecode []byte) {

	// clean vm state
	vm.__symbolTable = make([]string, 0)
	vm.__stack = make([]__p_intern__, 0)

	// set a new output for the new computation
	vm.__output = PreludioOutput{Log: make([]LogEnty, 0)}

	bytemark := bytecode[0:4]
	__symbolTableSize := binary.BigEndian.Uint32(bytecode[4:8])

	if vm.__debugLevel > 15 {
		vm.printDebug(15, "", "", "")
		vm.printDebug(15, "BYTECODE INFO", "", "")
		vm.printDebug(15, "====================", "", "")
		vm.printDebug(15, "SIZE", fmt.Sprintf("%d", len(bytecode)), "")
		vm.printDebug(15, "BYTE MARK", fmt.Sprintf("%x %x %x %x", bytemark[0], bytemark[1], bytemark[2], bytemark[3]), "")
		vm.printDebug(15, "SYMBOL TABLE SIZE", fmt.Sprintf("%d", __symbolTableSize), "")
		vm.printDebug(15, "", "", "")
		vm.printDebug(15, "STRING SYMBOLS", "", "")
		vm.printDebug(15, "====================", "", "")
	}

	offset := uint32(8)
	for i := uint32(0); i < __symbolTableSize; i++ {
		l := binary.BigEndian.Uint32(bytecode[offset : offset+4])
		offset += 4

		v := string(bytecode[offset : offset+l])
		vm.__symbolTable = append(vm.__symbolTable, v)
		offset += l
	}

	if vm.__debugLevel > 15 {
		for _, symbol := range vm.__symbolTable {
			vm.printDebug(15, "", symbol, "")
		}

		vm.printDebug(15, "", "", "")
		vm.printDebug(15, "INSTRUCTIONS", "", "")
		vm.printDebug(15, "====================", "", "")
	}

	vm.RunPrqlInstructions(bytecode, offset)
}

// Run Preludio bytecode from a binary file located
// at __inputPath with SetInputPath
// TO DEPRECATE (?)
func (vm *ByteEater) RunFileBytecode() {
	var err error
	var file *os.File
	var stats fs.FileInfo

	// clean vm state
	vm.__symbolTable = make([]string, 0)
	vm.__stack = make([]__p_intern__, 0)

	// set a new output for the new computation
	vm.__output = PreludioOutput{Log: make([]LogEnty, 0)}

	file, err = os.Open(vm.__inputPath)
	if err != nil {
		vm.setPanicMode(err.Error())
	}
	defer file.Close()

	stats, err = file.Stat()
	if err != nil {
		vm.setPanicMode(err.Error())
	}

	var size int64 = stats.Size()
	bytecode := make([]byte, size)

	bufr := bufio.NewReader(file)
	_, err = bufr.Read(bytecode)
	if err != nil {
		vm.setPanicMode(err.Error())
	}

	bytemark := bytecode[0:4]
	__symbolTableSize := binary.BigEndian.Uint32(bytecode[4:8])

	if vm.__debugLevel > 15 {
		vm.printDebug(15, "", "", "")
		vm.printDebug(15, "BYTECODE INFO", "", "")
		vm.printDebug(15, "====================", "", "")
		vm.printDebug(15, "SIZE", fmt.Sprintf("%d", size), "")
		vm.printDebug(15, "BYTE MARK", fmt.Sprintf("%x %x %x %x", bytemark[0], bytemark[1], bytemark[2], bytemark[3]), "")
		vm.printDebug(15, "SYMBOL TABLE SIZE", fmt.Sprintf("%d", __symbolTableSize), "")
		vm.printDebug(15, "", "", "")
		vm.printDebug(15, "STRING SYMBOLS", "", "")
		vm.printDebug(15, "====================", "", "")
	}

	offset := uint32(8)
	for i := uint32(0); i < __symbolTableSize; i++ {
		l := binary.BigEndian.Uint32(bytecode[offset : offset+4])
		offset += 4

		v := string(bytecode[offset : offset+l])
		vm.__symbolTable = append(vm.__symbolTable, v)
		offset += l
	}

	if vm.__debugLevel > 15 {
		for _, symbol := range vm.__symbolTable {
			vm.printDebug(15, "", symbol, "")
		}

		vm.printDebug(15, "", "", "")
		vm.printDebug(15, "INSTRUCTIONS", "", "")
		vm.printDebug(15, "====================", "", "")
	}

	vm.RunPrqlInstructions(bytecode, offset)
}

func (vm *ByteEater) setPanicMode(msg string) {
	vm.__panicMode = true
	vm.printError(msg)
}

func (vm *ByteEater) stackIsEmpty() bool {
	return len(vm.__stack) == 0
}

func (vm *ByteEater) stackPush(e *__p_intern__) {
	vm.__stack = append(vm.__stack, *e)
}

func (vm *ByteEater) stackPop() *__p_intern__ {
	e := vm.__stack[len(vm.__stack)-1]
	vm.__stack = vm.__stack[:len(vm.__stack)-1]
	return &e
}

func (vm *ByteEater) stackLast() *__p_intern__ {
	return &vm.__stack[len(vm.__stack)-1]
}

func (vm *ByteEater) loadResults() {
	vm.__output.Data = make([][]Columnar, 0)
	for !vm.stackIsEmpty() && vm.stackLast().tag != PRELUDIO_INTERNAL_TAG_BEGIN_FRAME {

		vm.__output.Data = append(vm.__output.Data, make([]Columnar, 0))

		internal := vm.stackPop()
		internal.toResult(&vm.__output.Data[len(vm.__output.Data)-1], vm.__outputSnippets)
	}
}

func (vm *ByteEater) GetOutput() *PreludioOutput {
	return &vm.__output
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
			vm.printDebug(10, "OP_START_PIPELINE", "", "")

			// Insert BEGIN FRAME
			vm.stackPush(newPInternBeginFrame())

		case OP_END_PIPELINE:
			vm.printDebug(10, "OP_END_PIPELINE", "", "")

			vm.loadResults()

			// Extract BEGIN FRAME
			vm.stackPop()

		case OP_ASSIGN_STMT:
			vm.printDebug(10, "OP_ASSIGN_STMT", "", "")

		case OP_START_FUNC_CALL:
			vm.printDebug(10, "OP_START_FUNC_CALL", "", "")

		case OP_MAKE_FUNC_CALL:
			funcName := vm.__symbolTable[binary.BigEndian.Uint32(param2)]
			vm.printDebug(10, "OP_MAKE_FUNC_CALL", "", funcName)

			switch funcName {

			// Standard library functions build-ins
			case "derive":
				PreludioFunc_Derive("derive", vm)
			case "describe":
				PreludioFunc_Describe("describe", vm)
			case "from":
				PreludioFunc_From("from", vm)
			case "writeCSV":
				PreludioFunc_WriteCsv("writeCSV", vm)
			case "readCSV":
				PreludioFunc_ReadCsv("readCSV", vm)
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
						vm.setPanicMode(fmt.Sprintf("variable '%s' not callable.", funcName))
					}
				} else {
					vm.setPanicMode(fmt.Sprintf("variable '%s' not defined.", funcName))
				}
			}

			vm.__funcNumParams = 0

		case OP_START_LIST:
			vm.printDebug(10, "OP_START_LIST", "", "")

			vm.__listElementCounters = append(vm.__listElementCounters, 0)

		case OP_END_LIST:
			vm.printDebug(10, "OP_END_LIST", "", "")

			stackLen := len(vm.__stack)
			listLen := vm.__listElementCounters[len(vm.__listElementCounters)-1]

			listCopy := make([]__p_intern__, listLen)
			copy(listCopy, vm.__stack[stackLen-listLen:])
			vm.__stack = vm.__stack[:stackLen-listLen]

			vm.stackPush(newPInternTerm(__p_list__(listCopy)))

			vm.__listElementCounters = vm.__listElementCounters[:len(vm.__listElementCounters)-1]

		case OP_ADD_FUNC_PARAM:
			vm.printDebug(10, "OP_ADD_FUNC_PARAM", "", "")

		case OP_ADD_EXPR_TERM:
			vm.printDebug(10, "OP_ADD_EXPR_TERM", "", "")

		///////////////////////////////////////////////////////////////////////
		///////////					PUSH NAMED PARAM
		///////////
		///////////	Set the last element on the stack as a named
		///////////	parameter.
		case OP_PUSH_NAMED_PARAM:
			paramName := vm.__symbolTable[binary.BigEndian.Uint32(param2)]
			vm.printDebug(10, "OP_PUSH_NAMED_PARAM", "", paramName)

			vm.stackLast().setParamName(paramName)

		///////////////////////////////////////////////////////////////////////
		///////////					PUSH ASSIGN IDENT
		///////////
		///////////	Set the last element on the stack as an assigned
		///////////	expression.
		case OP_PUSH_ASSIGN_IDENT:
			ident := vm.__symbolTable[binary.BigEndian.Uint32(param2)]
			vm.printDebug(10, "OP_PUSH_ASSIGN_IDENT", "", ident)

			vm.stackLast().setAssignment(ident)

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
				vm.stackPush(newPInternTerm([]bool{val}))

			case TERM_INTEGER:
				termType = "INTEGER"
				termVal = vm.__symbolTable[binary.BigEndian.Uint32(param2)]
				val, _ := strconv.ParseInt(termVal, 10, 64)
				vm.stackPush(newPInternTerm([]int{int(val)}))

			case TERM_FLOAT:
				termType = "FLOAT"
				termVal = vm.__symbolTable[binary.BigEndian.Uint32(param2)]
				val, _ := strconv.ParseFloat(termVal, 64)
				vm.stackPush(newPInternTerm([]float64{val}))

			case TERM_STRING:
				termType = "STRING"
				termVal = vm.__symbolTable[binary.BigEndian.Uint32(param2)]
				vm.stackPush(newPInternTerm([]string{termVal}))

			case TERM_SYMBOL:
				termType = "SYMBOL"
				termVal = vm.__symbolTable[binary.BigEndian.Uint32(param2)]
				vm.stackPush(newPInternTerm(__p_symbol__(termVal)))

			default:
				vm.setPanicMode(fmt.Sprintf("ByteEater: unknown term code %d.", param1))
			}

			vm.printDebug(10, "OP_PUSH_TERM", termType, termVal)

		case OP_END_CHUNCK:
			vm.printDebug(10, "OP_END_CHUNCK", "", "")

			vm.__funcNumParams += 1
			if len(vm.__listElementCounters) > 0 {
				vm.__listElementCounters[len(vm.__listElementCounters)-1]++
			}

		case OP_GOTO:
			vm.printDebug(10, "OP_GOTO", "", "")

		///////////////////////////////////////////////////////////////////////
		///////////				ARITHMETIC OPERATIONS
		case OP_BINARY_MUL:
			vm.printDebug(10, "OP_BINARY_MUL", "", "")

			op2 := vm.stackPop()
			vm.stackLast().appendOperand(OP_BINARY_MUL, op2)

		case OP_BINARY_DIV:
			vm.printDebug(10, "OP_BINARY_DIV", "", "")
			op2 := vm.stackPop()

			vm.stackPop().appendOperand(OP_BINARY_DIV, op2)

		case OP_BINARY_MOD:
			vm.printDebug(10, "OP_BINARY_MOD", "", "")

			op2 := vm.stackPop()
			vm.stackLast().appendOperand(OP_BINARY_MOD, op2)

		case OP_BINARY_ADD:
			vm.printDebug(10, "OP_BINARY_ADD", "", "")

			op2 := vm.stackPop()
			vm.stackLast().appendOperand(OP_BINARY_ADD, op2)

		case OP_BINARY_SUB:
			vm.printDebug(10, "OP_BINARY_SUB", "", "")

			op2 := vm.stackPop()
			vm.stackLast().appendOperand(OP_BINARY_SUB, op2)

		case OP_BINARY_POW:
			vm.printDebug(10, "OP_BINARY_POW", "", "")

			op2 := vm.stackPop()
			vm.stackLast().appendOperand(OP_BINARY_POW, op2)

		///////////////////////////////////////////////////////////////////////
		///////////				LOGICAL OPERATIONS

		case OP_BINARY_EQ:
			vm.printDebug(10, "OP_BINARY_EQ", "", "")

			op2 := vm.stackPop()
			vm.stackLast().appendOperand(OP_BINARY_EQ, op2)

		case OP_BINARY_NE:
			vm.printDebug(10, "OP_BINARY_NE", "", "")

			op2 := vm.stackPop()
			vm.stackLast().appendOperand(OP_BINARY_NE, op2)

		case OP_BINARY_GE:
			vm.printDebug(10, "OP_BINARY_GE", "", "")

			op2 := vm.stackPop()
			vm.stackLast().appendOperand(OP_BINARY_GE, op2)

		case OP_BINARY_LE:
			vm.printDebug(10, "OP_BINARY_LE", "", "")

			op2 := vm.stackPop()
			vm.stackLast().appendOperand(OP_BINARY_LE, op2)

		case OP_BINARY_GT:
			vm.printDebug(10, "OP_BINARY_GT", "", "")

			op2 := vm.stackPop()
			vm.stackLast().appendOperand(OP_BINARY_GT, op2)

		case OP_BINARY_LT:
			vm.printDebug(10, "OP_BINARY_LT", "", "")

			op2 := vm.stackPop()
			vm.stackLast().appendOperand(OP_BINARY_LT, op2)

		case OP_BINARY_AND:
			vm.printDebug(10, "OP_BINARY_AND", "", "")

			op2 := vm.stackPop()
			vm.stackLast().appendOperand(OP_BINARY_AND, op2)

		case OP_BINARY_OR:
			vm.printDebug(10, "OP_BINARY_OR", "", "")

			op2 := vm.stackPop()
			vm.stackLast().appendOperand(OP_BINARY_OR, op2)

		///////////////////////////////////////////////////////////////////////
		///////////				OTHER OPERATIONS

		case OP_BINARY_COALESCE:
			vm.printDebug(10, "OP_BINARY_COALESCE", "", "")

		case OP_BINARY_MODEL:
			vm.printDebug(10, "OP_BINARY_MODEL", "", "")

		///////////////////////////////////////////////////////////////////////
		///////////				UNARY OPERATIONS

		case OP_UNARY_SUB:
			vm.printDebug(10, "OP_UNARY_SUB", "", "")

		case OP_UNARY_ADD:
			vm.printDebug(10, "OP_UNARY_ADD", "", "")

		case OP_UNARY_NOT:
			vm.printDebug(10, "OP_UNARY_NOT", "", "")

		///////////////////////////////////////////////////////////////////////
		///////////				NO OPERATION

		case NO_OP:
			vm.printDebug(10, "NO_OP", "", "")
		}

		if vm.__panicMode {
			for !vm.stackIsEmpty() {
				vm.stackPop()
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
		t1 := *vm.stackLast()
		switch t1.tag {
		case PRELUDIO_INTERNAL_TAG_EXPRESSION:
			positionalParams = append([]*__p_intern__{&t1}, positionalParams...)
			vm.stackPop()

		case PRELUDIO_INTERNAL_TAG_NAMED_PARAM:
			// Name of parameter is in the given list of names
			if _, ok := (*namedParams)[t1.name]; ok {
				(*namedParams)[t1.name] = &t1
			} else {
				vm.printWarning(fmt.Sprintf("%s does not know a parameter named '%s', the value will be ignored.", funcName, t1.name))
			}
			vm.stackPop()

		case PRELUDIO_INTERNAL_TAG_ASSIGNMENT:
			if acceptingAssignments {
				assignments[t1.name] = &t1
			} else {
				vm.printWarning(fmt.Sprintf("%s does not accept assignements, the value of '%s' will be ignored.", funcName, t1.name))
			}
			vm.stackPop()

		case PRELUDIO_INTERNAL_TAG_BEGIN_FRAME:
			break LOOP1
		}
	}

	if solve {
		for _, p := range positionalParams {
			if err := solveExpr(vm, p); err != nil {
				return positionalParams, assignments, err
			}
		}

		if namedParams != nil {
			for _, p := range *namedParams {
				if err := solveExpr(vm, p); err != nil {
					return positionalParams, assignments, err
				}
			}
		}

		if acceptingAssignments {
			for _, p := range assignments {
				if err := solveExpr(vm, p); err != nil {
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
	df, _ := vm.stackLast().getDataframe()
	vm.__currentDataFrame = &df

	vm.__currentDataFrameNames = map[string]bool{}
	for _, name := range df.Names() {
		vm.__currentDataFrameNames[name] = true
	}
}

func truncate(s string, n int) string {
	if len(s) > n {
		return s[:n-3] + "..."
	}
	return s
}

func (vm *ByteEater) printDebug(level uint8, opname, param1, param2 string) {
	msg := fmt.Sprintf("[ 🐛 ]  %-20s | %-20s | %-20s", truncate(opname, 20), truncate(param1, 20), param2)
	vm.__output.Log = append(vm.__output.Log, LogEnty{LogType: LOG_DEBUG, Level: level, Message: msg})

	if vm.__printToStdout && vm.__debugLevel > int(level) {
		fmt.Print(msg)
	}
}

func (vm *ByteEater) printInfo(level uint8, msg string) {
	msg = fmt.Sprintf("[ ℹ️ ]  %s", msg)
	vm.__output.Log = append(vm.__output.Log, LogEnty{LogType: LOG_INFO, Level: level, Message: msg})

	if vm.__printToStdout {
		fmt.Print(msg)
	}
}

func (vm *ByteEater) printWarning(msg string) {
	msg = fmt.Sprintf("[ ⚠️ ]  %s", msg)
	vm.__output.Log = append(vm.__output.Log, LogEnty{LogType: LOG_WARNING, Message: msg})

	if vm.__printToStdout {
		fmt.Print(msg)
	}
}

func (vm *ByteEater) printError(msg string) {
	msg = fmt.Sprintf("[ ☠️ ]  %s", msg)
	vm.__output.Log = append(vm.__output.Log, LogEnty{LogType: LOG_ERROR, Message: msg})

	if vm.__printToStdout {
		fmt.Print(msg)
	}
}
