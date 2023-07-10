package gandalff

import (
	"fmt"
	"testing"
)

func checkEqSlice(a, b interface{}, t *testing.T, msg string) bool {
	switch a.(type) {
	case []bool:
		return checkEqSliceBool(a.([]bool), b.([]bool), t, msg)
	case []int32:
		return checkEqSliceInt32(a.([]int32), b.([]int32), t, msg)
	case []int64:
		return checkEqSliceInt64(a.([]int64), b.([]int64), t, msg)
	case []float64:
		return checkEqSliceFloat64(a.([]float64), b.([]float64), t, msg)
	case []string:
		return checkEqSliceString(a.([]string), b.([]string), t, msg)
	default:
		fmt.Printf("checkEqSlice: type %T not supported\n", a)
		return false
	}
}

func checkEqSliceBool(a, b []bool, t *testing.T, msg string) bool {
	if len(a) != len(b) {
		return false
	}

	if t == nil {
		for i, x := range a {
			if x != b[i] {
				fmt.Printf("    %s: %4d - expected '%v', got '%v'\n", msg, i, b[i], a[i])
				return false
			}
		}
	} else {
		for i, x := range a {
			if x != b[i] {
				t.Errorf("%s: %4d - expected '%v', got '%v'", msg, i, b[i], a[i])
				return false
			}
		}
	}
	return true
}

func checkEqSliceInt32(a, b []int32, t *testing.T, msg string) bool {
	if len(a) != len(b) {
		return false
	}

	if t == nil {
		for i, x := range a {
			if x != b[i] {
				fmt.Printf("    %s: %4d - expected '%v', got '%v'\n", msg, i, b[i], a[i])
				return false
			}
		}
	} else {
		for i, x := range a {
			if x != b[i] {
				t.Errorf("%s: %4d - expected '%v', got '%v'", msg, i, b[i], a[i])
				return false
			}
		}
	}
	return true
}

func checkEqSliceInt64(a, b []int64, t *testing.T, msg string) bool {
	if len(a) != len(b) {
		return false
	}

	if t == nil {
		for i, x := range a {
			if x != b[i] {
				fmt.Printf("    %s: %4d - expected '%v', got '%v'\n", msg, i, b[i], a[i])
				return false
			}
		}
	} else {
		for i, x := range a {
			if x != b[i] {
				t.Errorf("%s: %4d - expected '%v', got '%v'", msg, i, b[i], a[i])
				return false
			}
		}
	}
	return true
}

func checkEqSliceFloat64(a, b []float64, t *testing.T, msg string) bool {
	if len(a) != len(b) {
		return false
	}

	if t == nil {
		for i, x := range a {
			if x != b[i] {
				fmt.Printf("    %s: %4d - expected '%v', got '%v'\n", msg, i, b[i], a[i])
				return false
			}
		}
	} else {
		for i, x := range a {
			if x != b[i] {
				t.Errorf("%s: %4d - expected '%v', got '%v'", msg, i, b[i], a[i])
				return false
			}
		}
	}
	return true
}

func checkEqSliceString(a, b []string, t *testing.T, msg string) bool {
	if len(a) != len(b) {
		return false
	}

	if t == nil {
		for i, x := range a {
			if x != b[i] {
				fmt.Printf("    %s: %4d - expected '%v', got '%v'\n", msg, i, b[i], a[i])
				return false
			}
		}
	} else {
		for i, x := range a {
			if x != b[i] {
				t.Errorf("%s: %4d - expected '%v', got '%v'", msg, i, b[i], a[i])
				return false
			}
		}
	}
	return true
}
