package preludiocore

import (
	"bytefeeder"
	"os"
	"testing"

	"github.com/go-gota/gota/dataframe"
)

func init() {
	be = new(ByteEater).InitVM()
}

func Test_Expressions(t *testing.T) {
	var bytecode []byte

	// BOOL

	bytecode = bytefeeder.CompileSource(`true`)
	be.RunBytecode(bytecode)

	if be.__currentResult == nil {
		t.Error("Expected result, got nil")
	} else if be.__currentResult.isBoolScalar() == false {
		t.Error("Expected bool scalar, got", be.__currentResult)
	} else if b, err := be.__currentResult.getBoolScalar(); err != nil || b != true {
		t.Error("Expected true, got", b, err)
	}

	bytecode = bytefeeder.CompileSource(`false`)
	be.RunBytecode(bytecode)

	if be.__currentResult == nil {
		t.Error("Expected result, got nil")
	} else if be.__currentResult.isBoolScalar() == false {
		t.Error("Expected bool scalar, got", be.__currentResult)
	} else if b, err := be.__currentResult.getBoolScalar(); err != nil || b != false {
		t.Error("Expected false, got", b, err)
	}

	bytecode = bytefeeder.CompileSource(`true + false`)
	be.RunBytecode(bytecode)

	if be.__currentResult == nil {
		t.Error("Expected result, got nil")
	} else if be.__currentResult.isIntegerScalar() == false {
		t.Error("Expected integer scalar, got", be.__currentResult)
	} else if i, err := be.__currentResult.getIntegerScalar(); err != nil || i != 1 {
		t.Error("Expected 1, got", i, err)
	}

	// INT

	bytecode = bytefeeder.CompileSource(`1 * 5`)
	be.RunBytecode(bytecode)

	if be.__currentResult == nil {
		t.Error("Expected result, got nil")
	} else if be.__currentResult.isIntegerScalar() == false {
		t.Error("Expected integer scalar, got", be.__currentResult)
	} else if i, err := be.__currentResult.getIntegerScalar(); err != nil || i != 5 {
		t.Error("Expected 5, got", i, err)
	}

	bytecode = bytefeeder.CompileSource(`1 / 3`)
	be.RunBytecode(bytecode)

	if be.__currentResult == nil {
		t.Error("Expected result, got nil")
	} else if be.__currentResult.isFloatScalar() == false {
		t.Error("Expected float scalar, got", be.__currentResult)
	} else if f, err := be.__currentResult.getFloatScalar(); err != nil || f != 0.3333333333333333 {
		t.Error("Expected 0.3333333333333333, got", f, err)
	}

	bytecode = bytefeeder.CompileSource(`4682 % 427`)
	be.RunBytecode(bytecode)

	if be.__currentResult == nil {
		t.Error("Expected result, got nil")
	} else if be.__currentResult.isIntegerScalar() == false {
		t.Error("Expected integer scalar, got", be.__currentResult)
	} else if i, err := be.__currentResult.getIntegerScalar(); err != nil || i != 412 {
		t.Error("Expected 412, got", i, err)
	}

	bytecode = bytefeeder.CompileSource(`1 - 2`)
	be.RunBytecode(bytecode)

	if be.__currentResult == nil {
		t.Error("Expected result, got nil")
	} else if be.__currentResult.isIntegerScalar() == false {
		t.Error("Expected integer scalar, got", be.__currentResult)
	} else if i, err := be.__currentResult.getIntegerScalar(); err != nil || i != -1 {
		t.Error("Expected -1, got", i, err)
	}

	bytecode = bytefeeder.CompileSource(`1 + 2`)
	be.RunBytecode(bytecode)

	if be.__currentResult == nil {
		t.Error("Expected result, got nil")
	} else if be.__currentResult.isIntegerScalar() == false {
		t.Error("Expected integer scalar, got", be.__currentResult)
	} else if i, err := be.__currentResult.getIntegerScalar(); err != nil || i != 3 {
		t.Error("Expected 3, got", i, err)
	}

	// FLOAT

	bytecode = bytefeeder.CompileSource(`1.325235e-3 * 5`)
	be.RunBytecode(bytecode)

	if be.__currentResult == nil {
		t.Error("Expected result, got nil")
	} else if be.__currentResult.isFloatScalar() == false {
		t.Error("Expected float scalar, got", be.__currentResult)
	} else if f, err := be.__currentResult.getFloatScalar(); err != nil || f != 0.006626175 {
		t.Error("Expected 0.006626175, got", f, err)
	}

	bytecode = bytefeeder.CompileSource(`1.325235e-3 / 3`)
	be.RunBytecode(bytecode)

	if be.__currentResult == nil {
		t.Error("Expected result, got nil")
	} else if be.__currentResult.isFloatScalar() == false {
		t.Error("Expected float scalar, got", be.__currentResult)
	} else if f, err := be.__currentResult.getFloatScalar(); err != nil || f != 0.00044174499999999995 {
		t.Error("Expected 0.00044174499999999995, got", f, err)
	}

	// STRING

	bytecode = bytefeeder.CompileSource(`"hello" + "world"`)
	be.RunBytecode(bytecode)

	if be.__currentResult == nil {
		t.Error("Expected result, got nil")
	} else if be.__currentResult.isStringScalar() == false {
		t.Error("Expected string scalar, got", be.__currentResult)
	} else if s, err := be.__currentResult.getStringScalar(); err != nil || s != "helloworld" {
		t.Error("Expected helloworld, got", s, err)
	}

	bytecode = bytefeeder.CompileSource(`"hello" * 3`)
	be.RunBytecode(bytecode)

	if be.__currentResult == nil {
		t.Error("Expected result, got nil")
	} else if be.__currentResult.isStringScalar() == false {
		t.Error("Expected string scalar, got", be.__currentResult)
	} else if s, err := be.__currentResult.getStringScalar(); err != nil || s != "hellohellohello" {
		t.Error("Expected hellohellohello, got", s, err)
	}

	// LONG EXPRESSIONS

	bytecode = bytefeeder.CompileSource(`1 + 2 * 3 - 4 + 5 * 6`)
	be.RunBytecode(bytecode)

	if be.__currentResult == nil {
		t.Error("Expected result, got nil")
	} else if be.__currentResult.isIntegerScalar() == false {
		t.Error("Expected integer scalar, got", be.__currentResult)
	} else if i, err := be.__currentResult.getIntegerScalar(); err != nil || i != 33 {
		t.Error("Expected 33, got", i, err)
	}

	bytecode = bytefeeder.CompileSource(`1 + 2 * 3 - 4 + 5 * 6 % 7 + "hello"`)
	be.RunBytecode(bytecode)

	if be.__currentResult == nil {
		t.Error("Expected result, got nil")
	} else if be.__currentResult.isStringScalar() == false {
		t.Error("Expected string scalar, got", be.__currentResult)
	} else if s, err := be.__currentResult.getStringScalar(); err != nil || s != "5hello" {
		t.Error("Expected 5hello, got", s, err)
	}

	bytecode = bytefeeder.CompileSource(`3.4 + 2.3 * 3.2 - 4.1 + 5.0 * 6.9`)
	be.RunBytecode(bytecode)

	if be.__currentResult == nil {
		t.Error("Expected result, got nil")
	} else if be.__currentResult.isFloatScalar() == false {
		t.Error("Expected float scalar, got", be.__currentResult)
	} else if f, err := be.__currentResult.getFloatScalar(); err != nil || f != 41.16 {
		t.Error("Expected 41.16, got", f, err)
	}

	// bytecode = bytefeeder.CompileSource(`(1 + 2) * (3 - 4) + 5 * 6`)
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

	bytecode = bytefeeder.CompileSource(source)
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

	bytecode = bytefeeder.CompileSource(source)
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

	bytecode = bytefeeder.CompileSource(source)
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
