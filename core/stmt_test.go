package preludiocore

import (
	"preludiocompiler"
	"testing"
)

func init() {
	be = new(ByteEater).InitVM()
}

func TestStmt(t *testing.T) {
	var bytecode []byte
	var source string

	// BOOL
	source = `true`
	bytecode = preludiocompiler.CompileSource(source)
	be.RunBytecode(bytecode)

	if be.__currentResult == nil {
		t.Error("Expected result, got nil")
	} else if be.__currentResult.isBoolScalar() == false {
		t.Error("Expected bool scalar, got", be.__currentResult)
	} else if b, err := be.__currentResult.getBoolScalar(); err != nil || b != true {
		t.Error("Expected true, got", b, err)
	}

	// INT
	source = `1 + 2`
	bytecode = preludiocompiler.CompileSource(source)
	be.RunBytecode(bytecode)

	if be.__currentResult == nil {
		t.Error("Expected result, got nil")
	} else if be.__currentResult.isIntegerScalar() == false {
		t.Error("Expected integer scalar, got", be.__currentResult)
	} else if i, err := be.__currentResult.getIntegerScalar(); err != nil || i != 3 {
		t.Error("Expected 3, got", i, err)
	}

	// LONG EXPRESSIONS
	source = `1 + 2 * 3 - 4 + 5 * 6`
	bytecode = preludiocompiler.CompileSource(source)
	be.RunBytecode(bytecode)

	if be.__currentResult == nil {
		t.Error("Expected result, got nil")
	} else if be.__currentResult.isIntegerScalar() == false {
		t.Error("Expected integer scalar, got", be.__currentResult)
	} else if i, err := be.__currentResult.getIntegerScalar(); err != nil || i != 23 {
		t.Error("Expected 23, got", i, err)
	}

	// source = `1 + (2 * 3) - 4 + 5 * 6`
	// bytecode = preludiocompiler.CompileSource(source)
	// be.RunBytecode(bytecode)

	// if be.__currentResult == nil {
	// 	t.Error("Expected result, got nil")
	// } else if be.__currentResult.isIntegerScalar() == false {
	// 	t.Error("Expected integer scalar, got", be.__currentResult)
	// } else if i, err := be.__currentResult.getIntegerScalar(); err != nil || i != 23 {
	// 	t.Error("Expected 23, got", i, err)
	// }

}
