package preludiometa

import (
	"testing"
)

func TestTypeSystem(t *testing.T) {

	var res Primitive

	res = OP_BINARY_MUL.GetBinaryOpResultType(
		Primitive{Base: IntType},
		Primitive{Base: IntType},
	)

	if res.Base != IntType {
		t.Errorf("Expected IntType, got %v", res.Base)
	}

	res = OP_BINARY_MUL.GetBinaryOpResultType(
		Primitive{Base: IntType},
		Primitive{Base: Float64Type},
	)

	if res.Base != Float64Type {
		t.Errorf("Expected Float64Type, got %v", res.Base)
	}

	res = OP_BINARY_MUL.GetBinaryOpResultType(
		Primitive{Base: Float64Type},
		Primitive{Base: IntType},
	)

	if res.Base != Float64Type {
		t.Errorf("Expected Float64Type, got %v", res.Base)
	}
}
