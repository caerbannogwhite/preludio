package preludiocore

import (
	"fmt"
	"math"
)

func truncate(s string, n int) string {
	if len(s) > n {
		return s[:n-3] + "..."
	}
	return s
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
		if math.IsNaN(b[i]) {
			if !math.IsNaN(v) {
				return false
			} else {
				continue
			}
		}
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

func checkCurrentResult(be *ByteEater, expected interface{}) error {
	if be.__currentResult == nil {
		return fmt.Errorf("expected result, got nil")
	}

	switch v := expected.(type) {
	case bool:
		if !be.__currentResult.isBoolScalar() {
			return fmt.Errorf("expected bool scalar, got %T", be.__currentResult)
		} else if b, err := be.__currentResult.getBoolScalar(); err != nil || b != v {
			return fmt.Errorf("expected %t, got %t: %T", v, b, err)
		}

	case []bool:
		if !be.__currentResult.isBoolVector() {
			return fmt.Errorf("expected bool vector, got %T", be.__currentResult)
		} else if b, err := be.__currentResult.getBoolVector(); err != nil || !boolSliceEqual(b, v) {
			return fmt.Errorf("expected %v, got %v: %T", v, b, err)
		}

	case int64:
		if !be.__currentResult.isInt64Scalar() {
			return fmt.Errorf("expected int64 scalar, got %T", be.__currentResult)
		} else if i, err := be.__currentResult.getInt64Scalar(); err != nil || i != v {
			return fmt.Errorf("expected %d, got %d: %T", v, i, err)
		}

	case []int64:
		if !be.__currentResult.isInt64Vector() {
			return fmt.Errorf("expected int64 vector, got %T", be.__currentResult)
		} else if i, err := be.__currentResult.getInt64Vector(); err != nil || !int64SliceEqual(i, v) {
			return fmt.Errorf("expected %v, got %v: %T", v, i, err)
		}

	case float64:
		if !be.__currentResult.isFloat64Scalar() {
			return fmt.Errorf("expected float64 scalar, got %T", be.__currentResult)
		} else if f, err := be.__currentResult.getFloat64Scalar(); err != nil || !float64SliceEqual([]float64{f}, []float64{v}) {
			return fmt.Errorf("expected %f, got %f: %T", v, f, err)
		}

	case []float64:
		if !be.__currentResult.isFloat64Vector() {
			return fmt.Errorf("expected float64 vector, got %T", be.__currentResult)
		} else if f, err := be.__currentResult.getFloat64Vector(); err != nil || !float64SliceEqual(f, v) {
			return fmt.Errorf("expected %v, got %v: %T", v, f, err)
		}

	case string:
		if !be.__currentResult.isStringScalar() {
			return fmt.Errorf("expected string scalar, got %T", be.__currentResult)
		} else if s, err := be.__currentResult.getStringScalar(); err != nil || s != v {
			return fmt.Errorf("expected %s, got %s: %T", v, s, err)
		}

	case []string:
		if !be.__currentResult.isStringVector() {
			return fmt.Errorf("expected string vector, got %T", be.__currentResult)
		} else if s, err := be.__currentResult.getStringVector(); err != nil || !stringSliceEqual(s, v) {
			return fmt.Errorf("expected %v, got %v: %T", v, s, err)
		}

	default:
		return fmt.Errorf("unknown type %T", v)
	}

	return nil
}

func checkExpression(be *ByteEater, operand *__p_intern__, expected interface{}) error {
	err := be.solveExpr(operand)

	switch expectedTyped := expected.(type) {
	case bool:
		if err != nil {
			return err
		} else if v, err := operand.getBoolScalar(); err == nil {
			if v != expectedTyped {
				return fmt.Errorf("expected %t, got %t", expected, v)
			}
		} else {
			return err
		}

	case []bool:
		if err != nil {
			return err
		} else if v, err := operand.getBoolVector(); err == nil {
			if !boolSliceEqual(v, expectedTyped) {
				return fmt.Errorf("expected %v, got %v", expected, v)
			}
		} else {
			return err
		}

	case int64:
		if err != nil {
			return err
		} else if v, err := operand.getInt64Scalar(); err == nil {
			if v != expectedTyped {
				return fmt.Errorf("expected %d, got %d", expected, v)
			}
		} else {
			return err
		}

	case []int64:
		if err != nil {
			return err
		} else if v, err := operand.getInt64Vector(); err == nil {
			if !int64SliceEqual(v, expectedTyped) {
				return fmt.Errorf("expected %v, got %v", expected, v)
			}
		} else {
			return err
		}

	case float64:
		if err != nil {
			return err
		} else if v, err := operand.getFloat64Scalar(); err == nil {
			if v != expectedTyped {
				return fmt.Errorf("expected %f, got %f", expected, v)
			}
		} else {
			return err
		}

	case []float64:
		if err != nil {
			return err
		} else if v, err := operand.getFloat64Vector(); err == nil {
			if !float64SliceEqual(v, expectedTyped) {
				return fmt.Errorf("expected %v, got %v", expected, v)
			}
		} else {
			return err
		}

	case string:
		if err != nil {
			return err
		} else if v, err := operand.getStringScalar(); err == nil {
			if v != expectedTyped {
				return fmt.Errorf("expected %s, got %s", expected, v)
			}
		} else {
			return err
		}

	case []string:
		if err != nil {
			return err
		} else if v, err := operand.getStringVector(); err == nil {
			if !stringSliceEqual(v, expectedTyped) {
				return fmt.Errorf("expected %v, got %v", expected, v)
			}
		} else {
			return err
		}

	case error:
		if err == nil {
			return fmt.Errorf("expected error, got %v", operand)
		} else if err.Error() != expectedTyped.Error() {
			return fmt.Errorf("expected error string \"%v\", got \"%v\"", expected, err)
		}

	default:
		return fmt.Errorf("unknown type %T", expectedTyped)
	}

	return nil
}
