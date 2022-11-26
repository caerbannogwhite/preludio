package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"os"
	"text/template"
)

const TERM_NULL uint16 = 0
const TERM_BOOL uint16 = 1
const TERM_NUMERIC uint16 = 2
const TERM_STRING uint16 = 3
const TERM_INTERVAL uint16 = 5
const TERM_RANGE uint16 = 6
const TERM_LIST uint16 = 7
const TERM_PIPELINE uint16 = 8
const TERM_IDENT uint16 = 10

const OP_BEGIN_PIPELINE uint16 = 0
const OP_END_PIPELINE uint16 = 1
const OP_ASSIGN_TABLE uint16 = 2

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
const OP_BINARY_MIN uint16 = 104

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
	InputPath   string
	SymbolTable []string
}

func main() {
	inputPath := os.Args[1]
	vm := PRQLVirtualMachine{InputPath: inputPath}

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

	// return bytes, err
}

func (vm *PRQLVirtualMachine) read_instruction() {

}

func prql_derive(vm *PRQLVirtualMachine) {

}

func prql_export(vm *PRQLVirtualMachine) {

	tmpl := `
	PROC EXPORT data=PG_SUMM_CSV(drop=lot4 lot5)
	outfile="&res_dir.\&out_prefix._summary_PG.csv" dbms=csv replace;
	putnames=YES;
	RUN;
	`
	fmt.Println(tmpl)
}

func prql_import(vm *PRQLVirtualMachine) {

	type Params struct {
		TableName string
		InputPath string
		Delimiter string
		SkipRows  int
	}

	params := Params{TableName: "test", InputPath: "cars.csv", Delimiter: ";", SkipRows: 4}

	tmpl, err := template.New("import").Parse(`
	DATA {{.TableName}};
	length lot product device $30.;
	INFILE {{.InputPath}} DELIMITER="{{.Delimiter}}" FIRSTOBS={{.SkipRows}} DSD;
	INPUT VOID $ PRODUCT $ LOT $ DEVICE AMT_ACT AMT_TOT AMT_3_8 TOTAL;
	RUN;
	`)

	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, params)
	if err != nil {
		panic(err)
	}

	fmt.Println(tmpl)
}

func prql_select(vm *PRQLVirtualMachine) {

	type Params struct {
		TableName string
		Columns   string
	}

	params := Params{TableName: "test"}

	tmpl, err := template.New("select").Parse(`
	DATA {{.TableName}};
	keep={{.Columns}};
	RUN;
	`)

	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, params)
	if err != nil {
		panic(err)
	}

	fmt.Println(tmpl)
}
