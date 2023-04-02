package typesys

import (
	"testing"
)

func TestTypeSystem(t *testing.T) {

	var res Primitive

	res = OP_BINARY_MUL.GetBinaryOpResultType(
		Primitive{Base: Int32Type},
		Primitive{Base: Int32Type},
	)

	if res.Base != Int32Type {
		t.Errorf("Expected Int32Type, got %v", res.Base)
	}

	res = OP_BINARY_MUL.GetBinaryOpResultType(
		Primitive{Base: Int32Type},
		Primitive{Base: Float64Type},
	)

	if res.Base != Float64Type {
		t.Errorf("Expected Float64Type, got %v", res.Base)
	}

	res = OP_BINARY_MUL.GetBinaryOpResultType(
		Primitive{Base: Float64Type},
		Primitive{Base: Int32Type},
	)

	if res.Base != Float64Type {
		t.Errorf("Expected Float64Type, got %v", res.Base)
	}
}
