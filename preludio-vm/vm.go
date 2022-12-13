package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io/fs"
	"math"
	"os"
	"strings"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

type OPCODE uint16
type PARAM1 uint16

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
)

type PreludioVMParams struct {
	PrintWarnings  bool
	DebugLevel     int
	VerbosityLevel int
	InputPath      string
}

type PreludioVM struct {
	__printWarnings         bool
	__debugLevel            int
	__verbosityLevel        int
	__inputPath             string
	__symbolTable           []string
	__stack                 []*PreludioInternal
	__currentDataFrame      *dataframe.DataFrame
	__currentDataFrameNames map[string]bool
	__globalNameSpace       map[string]*PreludioInternal
	__pipelineNameSpace     map[string]*PreludioInternal
	__funcNumParams         int
	__listElementCounters   []int
}

func NewPreludioVM(params *PreludioVMParams) *PreludioVM {
	vm := PreludioVM{
		__printWarnings:  params.PrintWarnings,
		__inputPath:      params.InputPath,
		__debugLevel:     params.DebugLevel,
		__verbosityLevel: params.VerbosityLevel,
	}

	vm.__currentDataFrameNames = map[string]bool{}
	vm.__globalNameSpace = map[string]*PreludioInternal{}
	vm.__pipelineNameSpace = map[string]*PreludioInternal{}
	vm.__currentDataFrame = nil

	return &vm
}

func (vm *PreludioVM) ReadPreludioBytecode() {
	var err error
	var file *os.File
	var stats fs.FileInfo

	file, err = os.Open(vm.__inputPath)
	if err != nil {
		// return nil, err
	}
	defer file.Close()

	stats, err = file.Stat()
	if err != nil {
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

func (vm *PreludioVM) StackIsEmpty() bool {
	return len(vm.__stack) == 0
}

func (vm *PreludioVM) StackPush(e *PreludioInternal) {
	vm.__stack = append(vm.__stack, e)
}

func (vm *PreludioVM) StackPop() *PreludioInternal {
	e := vm.__stack[len(vm.__stack)-1]
	vm.__stack = vm.__stack[:len(vm.__stack)-1]
	return e
}

func (vm *PreludioVM) StackLast() *PreludioInternal {
	return vm.__stack[len(vm.__stack)-1]
}

func (vm *PreludioVM) ReadPrqlInstructions(bytes []byte, offset uint64) {

	var opCode OPCODE
	var param1 PARAM1
	var param2 []byte

	usize := uint64(len(bytes))

MAIN_LOOP:
	for offset < usize {
		opCode = OPCODE(binary.BigEndian.Uint16(bytes[offset : offset+2]))
		offset += 2
		param1 = PARAM1(binary.BigEndian.Uint16(bytes[offset : offset+2]))
		offset += 2
		param2 = bytes[offset : offset+8]
		offset += 8

		switch opCode {

		///////////////////////////////////////////////////////////////////////
		///////////				PIPELINE OPERATIONS					///////////
		case OP_START_PIPELINE:
			if vm.__debugLevel > 10 {
				fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "OP_START_PIPELINE", "", "", "")
			}

		case OP_END_PIPELINE:
			if vm.__debugLevel > 10 {
				fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "OP_END_PIPELINE", "", "", "")
			}

		case OP_ASSIGN_STMT:
			if vm.__debugLevel > 10 {
				fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "OP_ASSIGN_STMT", "", "", "")
			}

		case OP_START_FUNC_CALL:
			if vm.__debugLevel > 10 {
				fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "OP_START_FUNC_CALL", "", "", "")
			}

			vm.StackPush(NewPreludioInternalStartFunc())

		case OP_MAKE_FUNC_CALL:
			funcName := vm.__symbolTable[binary.BigEndian.Uint64(param2)]
			if vm.__debugLevel > 10 {
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

			// User defined functions
			default:
				if internal, ok := vm.__globalNameSpace[funcName]; ok {
					switch value := internal.GetValue().(type) {
					case UserDefinedFunction:
						value(vm)
					default:
						vm.StackPush(NewPreludioInternalError(fmt.Sprintf("variable '%s' not callable.", funcName)))
					}
				} else {
					vm.StackPush(NewPreludioInternalError(fmt.Sprintf("variable '%s' not defined.", funcName)))
				}
			}

			vm.__funcNumParams = 0

		case OP_START_LIST:
			if vm.__debugLevel > 10 {
				fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "OP_START_LIST", "", "", "")
			}

			vm.__listElementCounters = append(vm.__listElementCounters, 0)

		case OP_END_LIST:
			if vm.__debugLevel > 10 {
				fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "OP_END_LIST", "", "", "")
			}

			stackLen := len(vm.__stack)
			listLen := vm.__listElementCounters[len(vm.__listElementCounters)-1]
			list := make([]*PreludioInternal, listLen)
			for i := 0; i < listLen; i++ {
				list[i] = vm.__stack[stackLen-listLen+i]
			}

			vm.__stack = vm.__stack[:stackLen-listLen]
			vm.StackPush(NewPreludioInternalTerm(PreludioList(list)))

			vm.__listElementCounters = vm.__listElementCounters[:len(vm.__listElementCounters)-1]

		case OP_ADD_FUNC_PARAM:
			if vm.__debugLevel > 10 {
				fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "OP_ADD_FUNC_PARAM", "", "", "")
			}

		case OP_ADD_EXPR_TERM:
			if vm.__debugLevel > 10 {
				fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "OP_ADD_EXPR_TERM", "", "", "")
			}

		///////////////////////////////////////////////////////////////////////
		///////////					PUSH NAMED PARAM
		///////////
		///////////	Set the last element on the stack as a named
		///////////	parameter.
		case OP_PUSH_NAMED_PARAM:
			paramName := vm.__symbolTable[binary.BigEndian.Uint64(param2)]
			if vm.__debugLevel > 10 {
				fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "OP_PUSH_NAMED_PARAM", "", paramName, "")
			}

			vm.StackLast().SetParamName(paramName)

		///////////////////////////////////////////////////////////////////////
		///////////					PUSH ASSIGN IDENT
		///////////
		///////////	Set the last element on the stack as an assigned
		///////////	expression.
		case OP_PUSH_ASSIGN_IDENT:
			ident := vm.__symbolTable[binary.BigEndian.Uint64(param2)]
			if vm.__debugLevel > 10 {
				fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "OP_PUSH_ASSIGN_IDENT", "", ident, "")
			}

			vm.StackLast().SetAssignment(ident)

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
				vm.StackPush(NewPreludioInternalTerm([]bool{val}))

			case TERM_INTEGER:
				termType = "INTEGER"
				val := int(binary.LittleEndian.Uint32(param2))
				termVal = fmt.Sprintf("%d", val)
				vm.StackPush(NewPreludioInternalTerm([]int{val}))

			case TERM_FLOAT:
				termType = "FLOAT"
				val := math.Float64frombits(binary.LittleEndian.Uint64(param2))
				termVal = fmt.Sprintf("%f", val)
				vm.StackPush(NewPreludioInternalTerm([]float64{val}))

			case TERM_STRING:
				termType = "STRING"
				termVal = vm.__symbolTable[binary.BigEndian.Uint64(param2)]
				vm.StackPush(NewPreludioInternalTerm([]string{termVal}))

			case TERM_SYMBOL:
				termType = "SYMBOL"
				termVal = vm.__symbolTable[binary.BigEndian.Uint64(param2)]
				vm.StackPush(NewPreludioInternalTerm(PreludioSymbol(termVal)))

			default:
				vm.StackPush(NewPreludioInternalError(fmt.Sprintf("PrlqVM: unknown term code %d.", param1)))

			}

			if vm.__debugLevel > 10 {
				fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "OP_PUSH_TERM", termType, termVal, "")
			}

		case OP_END_CHUNCK:
			if vm.__debugLevel > 10 {
				fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "OP_END_CHUNCK", "", "", "")
			}

			vm.__funcNumParams += 1
			if len(vm.__listElementCounters) > 0 {
				vm.__listElementCounters[len(vm.__listElementCounters)-1]++
			}

		case OP_GOTO:
			if vm.__debugLevel > 10 {
				fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "OP_GOTO", "", "", "")
			}

		///////////////////////////////////////////////////////////////////////
		///////////				ARITHMETIC OPERATIONS
		case OP_BINARY_MUL:
			if vm.__debugLevel > 10 {
				fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "OP_BINARY_MUL", "", "", "")
			}

			op2 := vm.StackPop()
			vm.StackLast().Mul(op2)

		case OP_BINARY_DIV:
			if vm.__debugLevel > 10 {
				fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "OP_BINARY_DIV", "", "", "")
			}

			op2 := vm.StackPop()
			vm.StackLast().Div(op2)

		case OP_BINARY_MOD:
			if vm.__debugLevel > 10 {
				fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "OP_BINARY_MOD", "", "", "")
			}

			op2 := vm.StackPop()
			vm.StackLast().Mod(op2)

		case OP_BINARY_ADD:
			if vm.__debugLevel > 10 {
				fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "OP_BINARY_ADD", "", "", "")
			}

			op2 := vm.StackPop()
			vm.StackLast().Add(op2)

		case OP_BINARY_SUB:
			if vm.__debugLevel > 10 {
				fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "OP_BINARY_SUB", "", "", "")
			}

			op2 := vm.StackPop()
			vm.StackLast().Sub(op2)

		case OP_BINARY_POW:
			if vm.__debugLevel > 10 {
				fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "OP_BINARY_POW", "", "", "")
			}

			op2 := vm.StackPop()
			vm.StackLast().Pow(op2)

		///////////////////////////////////////////////////////////////////////
		///////////				LOGICAL OPERATIONS
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

		///////////////////////////////////////////////////////////////////////
		///////////				UNARY OPERATIONS
		case OP_UNARY_SUB:
			if vm.__debugLevel > 10 {
				fmt.Printf("%-30s | %-30s | %-30s | %-50s \n", "OP_UNARY_SUB", "", "", "")
			}

		case OP_UNARY_ADD:
		case OP_UNARY_NOT:

		}

		if !vm.StackIsEmpty() && vm.StackLast().Tag == PRELUDIO_INTERNAL_TAG_ERROR {
			for !vm.StackIsEmpty() && vm.StackLast().Tag == PRELUDIO_INTERNAL_TAG_ERROR {
				err := vm.StackPop()
				vm.PrintError(err.ErrorMsg)
			}
			break MAIN_LOOP
		}

	}
}

func (vm *PreludioVM) GetFunctionParams(funcName string, implicitParamsNum uint64, positionalParamsNum uint64, namedParams *map[string]*PreludioInternal, acceptingAssignments bool) ([]*PreludioInternal, map[string]*PreludioInternal, error) {

	var assignments map[string]*PreludioInternal
	if acceptingAssignments {
		assignments = map[string]*PreludioInternal{}
	}

	i := uint64(0)
	positionalParamsIdx := uint64(0)
	positionalParamsTot := implicitParamsNum + positionalParamsNum

	positionalParams := make([]*PreludioInternal, positionalParamsTot)

LOOP1:
	for {
		t1 := *vm.StackPop()
		switch t1.Tag {
		case PRELUDIO_INTERNAL_TAG_ERROR:
		case PRELUDIO_INTERNAL_TAG_EXPRESSION:
			if positionalParamsIdx < positionalParamsTot {
				positionalParams[positionalParamsTot-positionalParamsIdx-1] = &t1
				positionalParamsIdx++
			} else {
				vm.PrintWarning(fmt.Sprintf("function %s expects exactly %d positional parametes, the remaining values will be ignored.", funcName, positionalParamsNum))
			}

		case PRELUDIO_INTERNAL_TAG_NAMED_PARAM:
			// Name of parameter is in the given list of names
			if _, ok := (*namedParams)[t1.Name]; ok {
				(*namedParams)[t1.Name] = &t1
			} else {
				vm.PrintWarning(fmt.Sprintf("function %s does not know a parameter named '%s', the value will be ignored.", funcName, t1.Name))
			}

		case PRELUDIO_INTERNAL_TAG_ASSIGNMENT:
			if acceptingAssignments {
				assignments[t1.Name] = &t1
			} else {
				vm.PrintWarning(fmt.Sprintf("function %s does not accept assignements, the value of '%s' will be ignored.", funcName, t1.Name))
			}

		case PRELUDIO_INTERNAL_TAG_START_FUNC:
			break LOOP1
		}
	}

	// implicit parameters: pushed into the stack before the
	// function was called
	for ; i < implicitParamsNum; i++ {
		t1 := *vm.StackPop()
		switch t1.Tag {
		case PRELUDIO_INTERNAL_TAG_ERROR:
		case PRELUDIO_INTERNAL_TAG_EXPRESSION:
			if positionalParamsIdx < positionalParamsTot {
				positionalParams[positionalParamsTot-positionalParamsIdx-1] = &t1
				positionalParamsIdx++
			} else {
				vm.PrintWarning(fmt.Sprintf("function %s expects exactly %d positional parametes, the remaining values will be ignored.", funcName, positionalParamsNum))
			}

			// case PRELUDIO_INTERNAL_TAG_NAMED_PARAM:
			// 	// Name of parameter is in the given list of names
			// 	if _, ok := (*namedParams)[t1.Name]; ok {
			// 		(*namedParams)[t1.Name] = t1.Expr
			// 	} else {
			// 		vm.PrintWarning(fmt.Sprintf("function %s does not know a parameter named '%s', the value will be ignored.", funcName, t1.Name))
			// 	}

			// case PRELUDIO_INTERNAL_TAG_ASSIGNMENT:
			// 	if acceptingAssignments {
			// 		assignments[t1.Name] = t1.Expr
			// 	} else {
			// 		vm.PrintWarning(fmt.Sprintf("function %s does not accept assignements, the value of '%s' will be ignored.", funcName, t1.Name))
			// 	}
		}
	}

	for _, p := range positionalParams {
		if err := p.Solve(vm); err != nil {
			return positionalParams, assignments, err
		}
	}

	if namedParams != nil {
		for _, p := range *namedParams {
			if err := p.Solve(vm); err != nil {
				return positionalParams, assignments, err
			}
		}
	}

	if acceptingAssignments {
		for _, p := range assignments {
			if err := p.Solve(vm); err != nil {
				return positionalParams, assignments, err
			}
		}
	}

	return positionalParams, assignments, nil
}

func (vm *PreludioVM) SymbolResolution(symbol PreludioSymbol) interface{} {
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
func (vm *PreludioVM) SetCurrentDataFrame() {
	df, _ := vm.StackLast().GetDataframe()
	vm.__currentDataFrame = &df

	vm.__currentDataFrameNames = map[string]bool{}
	for _, name := range df.Names() {
		vm.__currentDataFrameNames[name] = true
	}
}

func (vm *PreludioVM) PrintWarning(msg string) {
	if vm.__printWarnings {
		fmt.Printf("[ ⚠️ Warning ⚠️ ] %s\n", msg)
	}
}

func (vm *PreludioVM) PrintError(msg string) {
	fmt.Printf("[ ☠️  Error  ☠️ ] %s\n", msg)
}
