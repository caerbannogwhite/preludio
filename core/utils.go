package preludiocore

import (
	"fmt"
	"typesys"
)

func truncate(s string, n int) string {
	if len(s) > n {
		return s[:n-3] + "..."
	}
	return s
}

func operatorToString(op typesys.OPCODE) string {
	switch op {
	case typesys.OP_BINARY_MUL:
		return "*"
	case typesys.OP_BINARY_DIV:
		return "/"
	case typesys.OP_BINARY_MOD:
		return "%"
	case typesys.OP_BINARY_POW:
		return "**"
	case typesys.OP_BINARY_ADD:
		return "+"
	case typesys.OP_BINARY_SUB:
		return "-"
	case typesys.OP_BINARY_AND:
		return "and"
	case typesys.OP_BINARY_OR:
		return "or"
	case typesys.OP_BINARY_EQ:
		return "=="
	case typesys.OP_BINARY_NE:
		return "!="
	case typesys.OP_BINARY_LT:
		return "<"
	case typesys.OP_BINARY_LE:
		return "<="
	case typesys.OP_BINARY_GT:
		return ">"
	case typesys.OP_BINARY_GE:
		return ">="
	default:
		return ""
	}
}

func boolSliceEqual(a, b []bool) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if b[i] != v {
			return false
		}
	}
	return true
}

func int64SliceEqual(a, b []int64) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if b[i] != v {
			return false
		}
	}
	return true
}

func float64SliceEqual(a, b []float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if b[i] != v {
			return false
		}
	}
	return true
}

func stringSliceEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if b[i] != v {
			return false
		}
	}
	return true
}

func currentResultChecker(be *ByteEater, expected interface{}) error {
	switch v := expected.(type) {
	case bool:
		if be.__currentResult == nil {
			return fmt.Errorf("expected result, got nil")
		} else if !be.__currentResult.isBoolScalar() {
			return fmt.Errorf("expected bool scalar, got %T", be.__currentResult)
		} else if b, err := be.__currentResult.getBoolScalar(); err != nil || b != v {
			return fmt.Errorf("expected %t, got %t: %T", v, b, err)
		}

	case []bool:
		if be.__currentResult == nil {
			return fmt.Errorf("expected result, got nil")
		} else if !be.__currentResult.isBoolVector() {
			return fmt.Errorf("expected bool vector, got %T", be.__currentResult)
		} else if b, err := be.__currentResult.getBoolVector(); err != nil || !boolSliceEqual(b, v) {
			return fmt.Errorf("expected %v, got %v: %T", v, b, err)
		}

	case int64:
		if be.__currentResult == nil {
			return fmt.Errorf("expected result, got nil")
		} else if !be.__currentResult.isInt64Scalar() {
			return fmt.Errorf("expected int64 scalar, got %T", be.__currentResult)
		} else if i, err := be.__currentResult.getInt64Scalar(); err != nil || i != v {
			return fmt.Errorf("expected %d, got %d: %T", v, i, err)
		}

	case []int64:
		if be.__currentResult == nil {
			return fmt.Errorf("expected result, got nil")
		} else if !be.__currentResult.isInt64Vector() {
			return fmt.Errorf("expected int64 vector, got %T", be.__currentResult)
		} else if i, err := be.__currentResult.getInt64Vector(); err != nil || !int64SliceEqual(i, v) {
			return fmt.Errorf("expected %v, got %v: %T", v, i, err)
		}

	case float64:
		if be.__currentResult == nil {
			return fmt.Errorf("expected result, got nil")
		} else if !be.__currentResult.isFloat64Scalar() {
			return fmt.Errorf("expected float64 scalar, got %T", be.__currentResult)
		} else if f, err := be.__currentResult.getFloat64Scalar(); err != nil || f != v {
			return fmt.Errorf("expected %f, got %f: %T", v, f, err)
		}

	case []float64:
		if be.__currentResult == nil {
			return fmt.Errorf("expected result, got nil")
		} else if !be.__currentResult.isFloat64Vector() {
			return fmt.Errorf("expected float64 vector, got %T", be.__currentResult)
		} else if f, err := be.__currentResult.getFloat64Vector(); err != nil || !float64SliceEqual(f, v) {
			return fmt.Errorf("expected %v, got %v: %T", v, f, err)
		}

	case string:
		if be.__currentResult == nil {
			return fmt.Errorf("expected result, got nil")
		} else if !be.__currentResult.isStringScalar() {
			return fmt.Errorf("expected string scalar, got %T", be.__currentResult)
		} else if s, err := be.__currentResult.getStringScalar(); err != nil || s != v {
			return fmt.Errorf("expected %s, got %s: %T", v, s, err)
		}

	case []string:
		if be.__currentResult == nil {
			return fmt.Errorf("expected result, got nil")
		} else if !be.__currentResult.isStringVector() {
			return fmt.Errorf("expected string vector, got %T", be.__currentResult)
		} else if s, err := be.__currentResult.getStringVector(); err != nil || !stringSliceEqual(s, v) {
			return fmt.Errorf("expected %v, got %v: %T", v, s, err)
		}

	default:
		return fmt.Errorf("unknown type %T", v)
	}

	return nil
}
