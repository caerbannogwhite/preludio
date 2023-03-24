package preludiocore

import (
	"os"
	"preludiocompiler"
	"testing"

	"github.com/go-gota/gota/dataframe"
)

func init() {
	be = new(ByteEater).InitVM()
}

func Test_Expressions(t *testing.T) {
	var bytecode []byte

	// BOOL

	bytecode = preludiocompiler.CompileSource(`true`)
	be.RunBytecode(bytecode)

	if be.__currentResult == nil {
		t.Error("Expected result, got nil")
	} else if be.__currentResult.isBoolScalar() == false {
		t.Error("Expected bool scalar, got", be.__currentResult)
	} else if b, err := be.__currentResult.getBoolScalar(); err != nil || b != true {
		t.Error("Expected true, got", b, err)
	}

	bytecode = preludiocompiler.CompileSource(`false`)
	be.RunBytecode(bytecode)

	if be.__currentResult == nil {
		t.Error("Expected result, got nil")
	} else if be.__currentResult.isBoolScalar() == false {
		t.Error("Expected bool scalar, got", be.__currentResult)
	} else if b, err := be.__currentResult.getBoolScalar(); err != nil || b != false {
		t.Error("Expected false, got", b, err)
	}

	// INT

	bytecode = preludiocompiler.CompileSource(`1 + 2`)
	be.RunBytecode(bytecode)

	if be.__currentResult == nil {
		t.Error("Expected result, got nil")
	} else if be.__currentResult.isIntegerScalar() == false {
		t.Error("Expected integer scalar, got", be.__currentResult)
	} else if i, err := be.__currentResult.getIntegerScalar(); err != nil || i != 3 {
		t.Error("Expected 3, got", i, err)
	}

	bytecode = preludiocompiler.CompileSource(`1 * 5`)
	be.RunBytecode(bytecode)

	if be.__currentResult == nil {
		t.Error("Expected result, got nil")
	} else if be.__currentResult.isIntegerScalar() == false {
		t.Error("Expected integer scalar, got", be.__currentResult)
	} else if i, err := be.__currentResult.getIntegerScalar(); err != nil || i != 5 {
		t.Error("Expected 5, got", i, err)
	}

	bytecode = preludiocompiler.CompileSource(`1 - 2`)
	be.RunBytecode(bytecode)

	if be.__currentResult == nil {
		t.Error("Expected result, got nil")
	} else if be.__currentResult.isIntegerScalar() == false {
		t.Error("Expected integer scalar, got", be.__currentResult)
	} else if i, err := be.__currentResult.getIntegerScalar(); err != nil || i != -1 {
		t.Error("Expected -1, got", i, err)
	}

	// LONG EXPRESSIONS

	bytecode = preludiocompiler.CompileSource(`1 + 2 * 3 - 4 + 5 * 6`)
	be.RunBytecode(bytecode)

	if be.__currentResult == nil {
		t.Error("Expected result, got nil")
	} else if be.__currentResult.isIntegerScalar() == false {
		t.Error("Expected integer scalar, got", be.__currentResult)
	} else if i, err := be.__currentResult.getIntegerScalar(); err != nil || i != 33 {
		t.Error("Expected 33, got", i, err)
	}

	// bytecode = preludiocompiler.CompileSource(`(1 + 2) * (3 - 4) + 5 * 6`)
	// be.RunBytecode(bytecode)

	// if be.__currentResult == nil {
	// 	t.Error("Expected result, got nil")
	// } else if be.__currentResult.isIntegerScalar() == false {
	// 	t.Error("Expected integer scalar, got", be.__currentResult)
	// } else if i, err := be.__currentResult.getIntegerScalar(); err != nil || i != 23 {
	// 	t.Error("Expected 23, got", i, err)
	// }

}

func Test_Function_readCSV(t *testing.T) {
	var bytecode []byte
	var source string
	var err error
	var df dataframe.DataFrame

	// CSV, comma delimiter, no header
	content := `true,hello,.43403,0
false,world,3e-2,4294
true,,0.000000001,-324
false,this is a string,4E4,3245
false,"hello again",0.000000000001,0`

	err = os.WriteFile("csvtest00_read_comma.csv", []byte(content), 0644)
	if err != nil {
		t.Error("Error writing test file", err)
	}
	defer os.Remove("csvtest00_read_comma.csv")

	source = `readCSV "csvtest00_read_comma.csv" delimiter: "," header: false`

	bytecode = preludiocompiler.CompileSource(source)
	be.RunBytecode(bytecode)

	if be.__currentResult == nil {
		t.Error("Expected result, got nil")
	} else if be.__currentResult.isDataframe() == false {
		t.Error("Expected dataframe, got", be.__currentResult)
	} else if df, err = be.__currentResult.getDataframe(); err == nil {
		records := df.Records()

		if len(records) != 5 {
			t.Error("Expected 5 records, got", len(records))
		}

		if records[0][0] != "true" {
			t.Error("Expected \"true\", got", records[0][0])
		}
		if records[0][1] != "hello" {
			t.Error("Expected \"hello\", got", records[0][1])
		}
		if records[3][1] != "this is a string" {
			t.Error("Expected \"this is a string\", got", records[3][1])
		}
	} else {
		t.Error("Expected no error, got", err)
	}

	// CSV, semicolon delimiter, no header
	content = `true;hello;.43403;0
false;world;3e-2;4294
true;;0.000000001;-324
false;this is a string;4E4;3245
false;"hello again";0.000000000001;0`

	err = os.WriteFile("csvtest01_read_semicolon.csv", []byte(content), 0644)
	if err != nil {
		t.Error("Error writing test file", err)
	}
	defer os.Remove("csvtest01_read_semicolon.csv")

	source = `readCSV "csvtest01_read_semicolon.csv" delimiter: ";" header: false`

	bytecode = preludiocompiler.CompileSource(source)
	be.RunBytecode(bytecode)

	if be.__currentResult == nil {
		t.Error("Expected result, got nil", be.__output.Log)
	} else if be.__currentResult.isDataframe() == false {
		t.Error("Expected dataframe, got", be.__currentResult)
	} else if df, err = be.__currentResult.getDataframe(); err == nil {
		records := df.Records()

		if len(records) != 5 {
			t.Error("Expected 5 records, got", len(records))
		}

		if records[0][0] != "true" {
			t.Error("Expected \"true\", got", records[0][0])
		}
		if records[0][1] != "hello" {
			t.Error("Expected \"hello\", got", records[0][1])
		}
		if records[3][1] != "this is a string" {
			t.Error("Expected \"this is a string\", got", records[3][1])
		}
	} else {
		t.Error("Expected no error, got", err)
	}

	// CSV, tab delimiter, header
	content = `A bool	something	a numeric value	an integer value
true	hello	.43403	0
false	world	3e-2	4294
true	0.000000001	-324	-1
false	this is a string	4E4	3245
false	"hello again"	0.000000000001	0`

	err = os.WriteFile("csvtest02_read_tab_header.csv", []byte(content), 0644)
	if err != nil {
		t.Error("Error writing test file", err)
	}
	defer os.Remove("csvtest02_read_tab_header.csv")

	source = `readCSV "csvtest02_read_tab_header.csv" delimiter: "\t" header: true`

	bytecode = preludiocompiler.CompileSource(source)
	be.RunBytecode(bytecode)

	if be.__currentResult == nil {
		t.Error("Expected result, got nil")
	} else if be.__currentResult.isDataframe() == false {
		t.Error("Expected dataframe, got", be.__currentResult)
	} else if df, err = be.__currentResult.getDataframe(); err != nil {
		records := df.Records()

		if len(records) != 4 {
			t.Error("Expected 4 records, got", len(records))
		}

		if records[0][0] != "true" {
			t.Error("Expected \"true\", got", records[0][0])
		}
		if records[0][1] != "hello" {
			t.Error("Expected \"hello\", got", records[0][1])
		}
		if records[3][1] != "this is a string" {
			t.Error("Expected \"this is a string\", got", records[3][1])
		}
	}
}
