package preludiocore

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io/fs"
	"os"
	"strconv"
	"strings"
	"typesys"

	"gandalff"
)

// ByteEater is the name of the Preludio Virtual Machine
type ByteEater struct {
	// parameters
	__param_printWarnings       bool
	__param_isCLI               bool
	__param_printToStdout       bool
	__param_fullOutput          bool
	__param_verbose             bool
	__param_debugLevel          int
	__param_outputSnippetLength int
	__param_inputPath           string

	// internals
	__panicMode             bool
	__symbolTable           []string
	__stack                 []__p_intern__
	__globalNamespace       map[string]*__p_intern__
	__pipelineNameSpace     map[string]*__p_intern__
	__currentDataFrameNames map[string]bool
	__funcNumParams         int
	__listElementCounters   []int
	__output                typesys.PreludioOutput
	__stringPool            *gandalff.StringPool
	__currentDataFrame      gandalff.DataFrame
	__currentResult         *__p_intern__
}

func (vm *ByteEater) GetParamPrintWarning() bool {
	return vm.__param_printWarnings
}

func (vm *ByteEater) SetParamPrintWarning(flag bool) *ByteEater {
	vm.__param_printWarnings = flag
	return vm
}

func (vm *ByteEater) GetParamIsCLI() bool {
	return vm.__param_isCLI
}

func (vm *ByteEater) SetParamIsCLI(flag bool) *ByteEater {
	vm.__param_isCLI = flag
	return vm
}

func (vm *ByteEater) GetParamPrintToStdout() bool {
	return vm.__param_printToStdout
}

func (vm *ByteEater) SetParamPrintToStdout(flag bool) *ByteEater {
	vm.__param_printToStdout = flag
	return vm
}

func (vm *ByteEater) GetParamFullOutput() bool {
	return vm.__param_fullOutput
}

func (vm *ByteEater) SetParamFullOutput(flag bool) *ByteEater {
	vm.__param_fullOutput = flag
	return vm
}

func (vm *ByteEater) GetParamVerbose() bool {
	return vm.__param_verbose
}

func (vm *ByteEater) SetParamVerbose(flag bool) *ByteEater {
	vm.__param_verbose = flag
	return vm
}

func (vm *ByteEater) GetParamDebugLevel() int {
	return vm.__param_debugLevel
}

func (vm *ByteEater) SetParamDebugLevel(level int) *ByteEater {
	vm.__param_debugLevel = level
	return vm
}

func (vm *ByteEater) GetParamOutputSnippetLength() int {
	return vm.__param_outputSnippetLength
}

func (vm *ByteEater) SetParamOutputSnippetLength(length int) *ByteEater {
	vm.__param_outputSnippetLength = length
	return vm
}

func (vm *ByteEater) InitVM() *ByteEater {

	// set default values
	vm.__param_outputSnippetLength = 10

	vm.__currentDataFrameNames = map[string]bool{}
	vm.__globalNamespace = map[string]*__p_intern__{}
	vm.__pipelineNameSpace = map[string]*__p_intern__{}

	vm.__stringPool = gandalff.NewStringPool()

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
	vm.__panicMode = false
	vm.__symbolTable = make([]string, 0)
	vm.__stack = make([]__p_intern__, 0)
	vm.__currentDataFrame = nil
	vm.__currentResult = nil

	// set a new output for the new computation
	vm.__output = typesys.PreludioOutput{Log: make([]typesys.LogEnty, 0)}

	bytemark := bytecode[0:4]
	__symbolTableSize := binary.BigEndian.Uint32(bytecode[4:8])

	if vm.__param_debugLevel > 25 {
		vm.printDebug(25, "", "", "")
		vm.printDebug(25, "BYTECODE INFO", "", "")
		vm.printDebug(25, "====================", "", "")
		vm.printDebug(25, "SIZE", fmt.Sprintf("%d", len(bytecode)), "")
		vm.printDebug(25, "BYTE MARK", fmt.Sprintf("%x %x %x %x", bytemark[0], bytemark[1], bytemark[2], bytemark[3]), "")
		vm.printDebug(25, "SYMBOL TABLE SIZE", fmt.Sprintf("%d", __symbolTableSize), "")
		vm.printDebug(25, "", "", "")
		vm.printDebug(25, "STRING SYMBOLS", "", "")
		vm.printDebug(25, "====================", "", "")
	}

	offset := uint32(8)
	for i := uint32(0); i < __symbolTableSize; i++ {
		l := binary.BigEndian.Uint32(bytecode[offset : offset+4])
		offset += 4

		v := string(bytecode[offset : offset+l])
		vm.__symbolTable = append(vm.__symbolTable, v)
		offset += l
	}

	if vm.__param_debugLevel > 25 {
		for _, symbol := range vm.__symbolTable {
			vm.printDebug(25, "", symbol, "")
		}

		vm.printDebug(25, "", "", "")
		vm.printDebug(25, "INSTRUCTIONS", "", "")
		vm.printDebug(25, "====================", "", "")
	}

	vm.RunPrqlInstructions(bytecode, offset)
}

// Run Preludio bytecode from a binary file located
// at __param_inputPath with SetInputPath
func (vm *ByteEater) RunFileBytecode() {
	var err error
	var file *os.File
	var stats fs.FileInfo

	// clean vm state
	vm.__panicMode = false
	vm.__symbolTable = make([]string, 0)
	vm.__stack = make([]__p_intern__, 0)
	vm.__currentDataFrame = nil
	vm.__currentResult = nil

	// set a new output for the new computation
	vm.__output = typesys.PreludioOutput{Log: make([]typesys.LogEnty, 0)}

	file, err = os.Open(vm.__param_inputPath)
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

	if vm.__param_debugLevel > 25 {
		vm.printDebug(25, "", "", "")
		vm.printDebug(25, "BYTECODE INFO", "", "")
		vm.printDebug(25, "====================", "", "")
		vm.printDebug(25, "SIZE", fmt.Sprintf("%d", size), "")
		vm.printDebug(25, "BYTE MARK", fmt.Sprintf("%x %x %x %x", bytemark[0], bytemark[1], bytemark[2], bytemark[3]), "")
		vm.printDebug(25, "SYMBOL TABLE SIZE", fmt.Sprintf("%d", __symbolTableSize), "")
		vm.printDebug(25, "", "", "")
		vm.printDebug(25, "STRING SYMBOLS", "", "")
		vm.printDebug(25, "====================", "", "")
	}

	offset := uint32(8)
	for i := uint32(0); i < __symbolTableSize; i++ {
		l := binary.BigEndian.Uint32(bytecode[offset : offset+4])
		offset += 4

		v := string(bytecode[offset : offset+l])
		vm.__symbolTable = append(vm.__symbolTable, v)
		offset += l
	}

	if vm.__param_debugLevel > 25 {
		for _, symbol := range vm.__symbolTable {
			vm.printDebug(25, "", symbol, "")
		}

		vm.printDebug(25, "", "", "")
		vm.printDebug(25, "INSTRUCTIONS", "", "")
		vm.printDebug(25, "====================", "", "")
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
	// get the results from the stack until we find the begin frame
	results := make([]__p_intern__, 0)
	for !vm.stackIsEmpty() && vm.stackLast().tag != PRELUDIO_INTERNAL_TAG_BEGIN_FRAME {

		result := vm.stackPop()
		if err := solveExpr(vm, result); err != nil {
			vm.setPanicMode(err.Error())
			break
		}

		results = append(results, *result)
	}

	if len(results) > 1 {
		vm.__currentResult = newPInternTerm(__p_list__(results))
	} else if len(results) == 1 {
		vm.__currentResult = &results[0]
	}
}

func (vm *ByteEater) GetOutput() *typesys.PreludioOutput {

	vm.__output.Data = make([][]typesys.Columnar, 0)

	if vm.__currentResult != nil {
		if vm.__currentResult.isList() {
			list, err := vm.__currentResult.getList()
			if err != nil {
				vm.setPanicMode(err.Error())
				return &vm.__output
			}

			for _, result := range list {
				vm.__output.Data = append(vm.__output.Data, make([]typesys.Columnar, 0))
				result.toResult(&vm.__output.Data[len(vm.__output.Data)-1], vm.__param_fullOutput, vm.__param_outputSnippetLength)
			}
		} else {
			vm.__output.Data = append(vm.__output.Data, make([]typesys.Columnar, 0))
			vm.__currentResult.toResult(&vm.__output.Data[len(vm.__output.Data)-1], vm.__param_fullOutput, vm.__param_outputSnippetLength)
		}
	}

	return &vm.__output
}

func (vm *ByteEater) RunPrqlInstructions(bytes []byte, offset uint32) {

	opCode := typesys.OPCODE(0)
	param1 := typesys.PARAM1(0)
	param2 := []byte{0, 0, 0, 0}

	usize := uint32(len(bytes))

MAIN_LOOP:
	for offset < usize {

		opCode = typesys.OPCODE(bytes[offset])
		offset++
		param1 = typesys.PARAM1(bytes[offset])
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

		case typesys.OP_START_STMT:
			vm.printDebug(10, "OP_START_STMT", "", "")

			// Insert BEGIN FRAME
			vm.stackPush(newPInternBeginFrame())

		case typesys.OP_END_STMT:
			vm.printDebug(10, "OP_END_STMT", "", "")

			vm.loadResults()

			// Estract BEGIN FRAME
			vm.stackPop()

		case typesys.OP_VAR_DECL:
			varName := vm.__symbolTable[binary.BigEndian.Uint32(param2)]
			vm.printDebug(10, "OP_VAR_DECL", "", varName)

			if _, ok := vm.__globalNamespace[varName]; ok {
				vm.setPanicMode(fmt.Sprintf("Variable \"%s\" is already declared", varName))
			} else {
				vm.loadResults()
				if vm.__currentResult != nil {
					vm.__globalNamespace[varName] = vm.__currentResult
					vm.__currentResult = nil
				}
			}

		case typesys.OP_VAR_ASSIGN:
			varName := vm.__symbolTable[binary.BigEndian.Uint32(param2)]
			vm.printDebug(10, "OP_VAR_ASSIGN", "", varName)

			if _, ok := vm.__globalNamespace[varName]; ok {
				vm.loadResults()
				if vm.__currentResult != nil {
					vm.__globalNamespace[varName] = vm.__currentResult
					vm.__currentResult = nil
				}
			} else {
				vm.setPanicMode(fmt.Sprintf("Variable \"%s\" is not declared", varName))
			}

		case typesys.OP_START_FUNC_CALL:
			vm.printDebug(10, "OP_START_FUNC_CALL", "", "")

		///////////////////////////////////////////////////////////////////////
		///////////				PIPELINE OPERATIONS					///////////
		case typesys.OP_START_PIPELINE:
			vm.printDebug(10, "OP_START_PIPELINE", "", "")

			// Insert BEGIN FRAME
			vm.stackPush(newPInternBeginFrame())

		case typesys.OP_END_PIPELINE:
			vm.printDebug(10, "OP_END_PIPELINE", "", "")

			vm.loadResults()

			// Extract BEGIN FRAME
			vm.stackPop()

		case typesys.OP_MAKE_FUNC_CALL:
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
				PreludioFunc_WriteCSV("writeCSV", vm)
			case "readCSV":
				PreludioFunc_ReadCSV("readCSV", vm)
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
				if internal, ok := vm.__globalNamespace[funcName]; ok {
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

		case typesys.OP_START_LIST:
			vm.printDebug(10, "OP_START_LIST", "", "")

			vm.__listElementCounters = append(vm.__listElementCounters, 0)

		case typesys.OP_END_LIST:
			vm.printDebug(10, "OP_END_LIST", "", "")

			stackLen := len(vm.__stack)
			listLen := vm.__listElementCounters[len(vm.__listElementCounters)-1]

			listCopy := make([]__p_intern__, listLen)
			copy(listCopy, vm.__stack[stackLen-listLen:])
			vm.__stack = vm.__stack[:stackLen-listLen]

			vm.stackPush(newPInternTerm(__p_list__(listCopy)))

			vm.__listElementCounters = vm.__listElementCounters[:len(vm.__listElementCounters)-1]

		case typesys.OP_ADD_FUNC_PARAM:
			vm.printDebug(10, "OP_ADD_FUNC_PARAM", "", "")

		case typesys.OP_ADD_EXPR_TERM:
			vm.printDebug(10, "OP_ADD_EXPR_TERM", "", "")

		///////////////////////////////////////////////////////////////////////
		///////////					PUSH NAMED PARAM
		///////////
		///////////	Set the last element on the stack as a named
		///////////	parameter.
		case typesys.OP_PUSH_NAMED_PARAM:
			paramName := vm.__symbolTable[binary.BigEndian.Uint32(param2)]
			vm.printDebug(10, "OP_PUSH_NAMED_PARAM", "", paramName)

			vm.stackLast().setParamName(paramName)

		///////////////////////////////////////////////////////////////////////
		///////////					PUSH ASSIGN IDENT
		///////////
		///////////	Set the last element on the stack as an assigned
		///////////	expression.
		case typesys.OP_PUSH_ASSIGN_IDENT:
			ident := vm.__symbolTable[binary.BigEndian.Uint32(param2)]
			vm.printDebug(10, "OP_PUSH_ASSIGN_IDENT", "", ident)

			vm.stackLast().setAssignment(ident)

		case typesys.OP_PUSH_TERM:
			termType := ""
			termVal := ""

			switch param1 {
			// case typesys.TERM_NULL:
			// 	termType = "NULL"

			case typesys.TERM_BOOL:
				termType = "BOOL"
				termVal = "true"
				val := true
				if binary.BigEndian.Uint32(param2) == 0 {
					val = false
					termVal = "false"
				}
				vm.stackPush(newPInternTerm(gandalff.NewSeriesBool("", true, []bool{val})))

			case typesys.TERM_INTEGER:
				termType = "INTEGER"
				termVal = vm.__symbolTable[binary.BigEndian.Uint32(param2)]
				val, _ := strconv.ParseInt(termVal, 10, 64)
				vm.stackPush(newPInternTerm(gandalff.NewSeriesInt64("", true, false, []int64{val})))

			case typesys.TERM_FLOAT:
				termType = "FLOAT"
				termVal = vm.__symbolTable[binary.BigEndian.Uint32(param2)]
				val, _ := strconv.ParseFloat(termVal, 64)
				vm.stackPush(newPInternTerm(gandalff.NewSeriesFloat64("", true, false, []float64{val})))

			case typesys.TERM_STRING:
				termType = "STRING"
				termVal = vm.__symbolTable[binary.BigEndian.Uint32(param2)]
				vm.stackPush(newPInternTerm(gandalff.NewSeriesString("", true, []string{termVal}, vm.__stringPool)))

			case typesys.TERM_SYMBOL:
				termType = "SYMBOL"
				termVal = vm.__symbolTable[binary.BigEndian.Uint32(param2)]
				vm.stackPush(newPInternTerm(__p_symbol__(termVal)))

			default:
				vm.setPanicMode(fmt.Sprintf("ByteEater: unknown term code %d.", param1))
			}

			vm.printDebug(10, "OP_PUSH_TERM", termType, termVal)

		case typesys.OP_END_CHUNCK:
			vm.printDebug(10, "OP_END_CHUNCK", "", "")

			vm.__funcNumParams += 1
			if len(vm.__listElementCounters) > 0 {
				vm.__listElementCounters[len(vm.__listElementCounters)-1]++
			}

		case typesys.OP_GOTO:
			vm.printDebug(10, "OP_GOTO", "", "")

		///////////////////////////////////////////////////////////////////////
		///////////				ARITHMETIC OPERATIONS
		case typesys.OP_BINARY_MUL:
			vm.printDebug(10, "OP_BINARY_MUL", "", "")

			op2 := vm.stackPop()
			vm.stackLast().appendOperand(typesys.OP_BINARY_MUL, op2)

		case typesys.OP_BINARY_DIV:
			vm.printDebug(10, "OP_BINARY_DIV", "", "")

			op2 := vm.stackPop()
			vm.stackLast().appendOperand(typesys.OP_BINARY_DIV, op2)

		case typesys.OP_BINARY_MOD:
			vm.printDebug(10, "OP_BINARY_MOD", "", "")

			op2 := vm.stackPop()
			vm.stackLast().appendOperand(typesys.OP_BINARY_MOD, op2)

		case typesys.OP_BINARY_ADD:
			vm.printDebug(10, "OP_BINARY_ADD", "", "")

			op2 := vm.stackPop()
			vm.stackLast().appendOperand(typesys.OP_BINARY_ADD, op2)

		case typesys.OP_BINARY_SUB:
			vm.printDebug(10, "OP_BINARY_SUB", "", "")

			op2 := vm.stackPop()
			vm.stackLast().appendOperand(typesys.OP_BINARY_SUB, op2)

		case typesys.OP_BINARY_POW:
			vm.printDebug(10, "OP_BINARY_POW", "", "")

			op2 := vm.stackPop()
			vm.stackLast().appendOperand(typesys.OP_BINARY_POW, op2)

		///////////////////////////////////////////////////////////////////////
		///////////				LOGICAL OPERATIONS

		case typesys.OP_BINARY_EQ:
			vm.printDebug(10, "OP_BINARY_EQ", "", "")

			op2 := vm.stackPop()
			vm.stackLast().appendOperand(typesys.OP_BINARY_EQ, op2)

		case typesys.OP_BINARY_NE:
			vm.printDebug(10, "OP_BINARY_NE", "", "")

			op2 := vm.stackPop()
			vm.stackLast().appendOperand(typesys.OP_BINARY_NE, op2)

		case typesys.OP_BINARY_GE:
			vm.printDebug(10, "OP_BINARY_GE", "", "")

			op2 := vm.stackPop()
			vm.stackLast().appendOperand(typesys.OP_BINARY_GE, op2)

		case typesys.OP_BINARY_LE:
			vm.printDebug(10, "OP_BINARY_LE", "", "")

			op2 := vm.stackPop()
			vm.stackLast().appendOperand(typesys.OP_BINARY_LE, op2)

		case typesys.OP_BINARY_GT:
			vm.printDebug(10, "OP_BINARY_GT", "", "")

			op2 := vm.stackPop()
			vm.stackLast().appendOperand(typesys.OP_BINARY_GT, op2)

		case typesys.OP_BINARY_LT:
			vm.printDebug(10, "OP_BINARY_LT", "", "")

			op2 := vm.stackPop()
			vm.stackLast().appendOperand(typesys.OP_BINARY_LT, op2)

		case typesys.OP_BINARY_AND:
			vm.printDebug(10, "OP_BINARY_AND", "", "")

			op2 := vm.stackPop()
			vm.stackLast().appendOperand(typesys.OP_BINARY_AND, op2)

		case typesys.OP_BINARY_OR:
			vm.printDebug(10, "OP_BINARY_OR", "", "")

			op2 := vm.stackPop()
			vm.stackLast().appendOperand(typesys.OP_BINARY_OR, op2)

		///////////////////////////////////////////////////////////////////////
		///////////				OTHER OPERATIONS

		case typesys.OP_BINARY_COALESCE:
			vm.printDebug(10, "OP_BINARY_COALESCE", "", "")

		case typesys.OP_BINARY_MODEL:
			vm.printDebug(10, "OP_BINARY_MODEL", "", "")

		///////////////////////////////////////////////////////////////////////
		///////////				UNARY OPERATIONS

		case typesys.OP_UNARY_SUB:
			vm.printDebug(10, "OP_UNARY_SUB", "", "")

		case typesys.OP_UNARY_ADD:
			vm.printDebug(10, "OP_UNARY_ADD", "", "")

		case typesys.OP_UNARY_NOT:
			vm.printDebug(10, "OP_UNARY_NOT", "", "")

		///////////////////////////////////////////////////////////////////////
		///////////				NO OPERATION

		case typesys.NO_OP:
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

func (vm *ByteEater) symbolResolution(symbol __p_symbol__) interface{} {
	// 1 - Look at the current DataFrame
	if vm.__currentDataFrame != nil {
		if ok := vm.__currentDataFrameNames[string(symbol)]; ok {
			ser := vm.__currentDataFrame.Series(string(symbol))
			switch ser.Type() {
			case typesys.BoolType:
				val, _ := ser.Data().([]bool)
				return val
			case typesys.Int64Type:
				val, _ := ser.Data().([]int64)
				return val
			case typesys.Float64Type:
				return ser.Data().([]float64)
			case typesys.StringType:
				return ser.Data().([]string)
			}
		}
	}

	// 2 - Look at the global Namespace
	if val, ok := vm.__globalNamespace[string(symbol)]; ok {
		if len(val.expr) == 1 {
			return val.expr[0]
		}
		return nil
	}

	// 2 - Try to split the symbol into pieces
	pieces := strings.Split(string(symbol), ".")
	return pieces
}

// Set the last element inserted into the stack as
// the current DataFrame
func (vm *ByteEater) setCurrentDataFrame() {
	df, _ := vm.stackLast().getDataframe()
	vm.__currentDataFrame = df

	vm.__currentDataFrameNames = map[string]bool{}
	for _, name := range df.Names() {
		vm.__currentDataFrameNames[name] = true
	}
}

func (vm *ByteEater) printDebug(level uint8, opname, param1, param2 string) {
	msg := fmt.Sprintf("[ 🐛 ]  %-20s | %-20s | %-20s", truncate(opname, 20), truncate(param1, 20), param2)
	vm.__output.Log = append(vm.__output.Log, typesys.LogEnty{LogType: typesys.LOG_DEBUG, Level: level, Message: msg})

	if vm.__param_printToStdout && vm.__param_debugLevel > int(level) {
		fmt.Print(msg)
	}
}

func (vm *ByteEater) printInfo(level uint8, msg string) {
	msg = fmt.Sprintf("[ ℹ️ ]  %s", msg)
	vm.__output.Log = append(vm.__output.Log, typesys.LogEnty{LogType: typesys.LOG_INFO, Level: level, Message: msg})

	if vm.__param_printToStdout {
		fmt.Print(msg)
	}
}

func (vm *ByteEater) printWarning(msg string) {
	msg = fmt.Sprintf("[ ⚠️ ]  %s", msg)
	vm.__output.Log = append(vm.__output.Log, typesys.LogEnty{LogType: typesys.LOG_WARNING, Message: msg})

	if vm.__param_printToStdout {
		fmt.Print(msg)
	}
}

func (vm *ByteEater) printError(msg string) {
	msg = fmt.Sprintf("[ ☠️ ]  %s", msg)
	vm.__output.Log = append(vm.__output.Log, typesys.LogEnty{LogType: typesys.LOG_ERROR, Message: msg})

	if vm.__param_printToStdout {
		fmt.Print(msg)
	}
}
