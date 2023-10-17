package gandalff

import (
	"fmt"
	"math"
	"sort"
	"testing"
	"time"
)

func checkEqSlice(a, b interface{}, t *testing.T, msg string) bool {
	switch a.(type) {
	case []bool:
		return checkEqSliceBool(a.([]bool), b.([]bool), t, msg)
	case []int:
		return checkEqSliceInt(a.([]int), b.([]int), t, msg)
	case []int32:
		return checkEqSliceInt32(a.([]int32), b.([]int32), t, msg)
	case []int64:
		return checkEqSliceInt64(a.([]int64), b.([]int64), t, msg)
	case []float64:
		return checkEqSliceFloat64(a.([]float64), b.([]float64), t, msg)
	case []string:
		return checkEqSliceString(a.([]string), b.([]string), t, msg)
	case []time.Time:
		return checkEqSliceTime(a.([]time.Time), b.([]time.Time), t, msg)
	case []time.Duration:
		return checkEqSliceDuration(a.([]time.Duration), b.([]time.Duration), t, msg)
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
				if msg != "" {
					fmt.Printf("    %s: %4d - expected '%v', got '%v'\n", msg, i, b[i], a[i])
				} else {
					fmt.Printf("    checkEqSliceBool: %4d - expected '%v', got '%v'\n", i, b[i], a[i])
				}
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

func checkEqSliceInt(a, b []int, t *testing.T, msg string) bool {
	if len(a) != len(b) {
		return false
	}

	if t == nil {
		for i, x := range a {
			if x != b[i] {
				if msg != "" {
					fmt.Printf("    %s: %4d - expected '%v', got '%v'\n", msg, i, b[i], a[i])
				} else {
					fmt.Printf("    checkEqSliceInt: %4d - expected '%v', got '%v'\n", i, b[i], a[i])
				}
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

func checkEqSliceIntQuiet(a, b []int, t *testing.T, msg string) bool {
	if len(a) != len(b) {
		return false
	}

	for i, x := range a {
		if x != b[i] {
			return false
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
				if msg != "" {
					fmt.Printf("    %s: %4d - expected '%v', got '%v'\n", msg, i, b[i], a[i])
				} else {
					fmt.Printf("    checkEqSliceInt32: %4d - expected '%v', got '%v'\n", i, b[i], a[i])
				}
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
				if msg != "" {
					fmt.Printf("    %s: %4d - expected '%v', got '%v'\n", msg, i, b[i], a[i])
				} else {
					fmt.Printf("    checkEqSliceInt64: %4d - expected '%v', got '%v'\n", i, b[i], a[i])
				}
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
			if math.IsNaN(x) && math.IsNaN(b[i]) {
				continue
			}
			if x != b[i] {
				if msg != "" {
					fmt.Printf("    %s: %4d - expected '%v', got '%v'\n", msg, i, b[i], a[i])
				} else {
					fmt.Printf("    checkEqSliceFloat64: %4d - expected '%v', got '%v'\n", i, b[i], a[i])
				}
				return false
			}
		}
	} else {
		for i, x := range a {
			if math.IsNaN(x) && math.IsNaN(b[i]) {
				continue
			}
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
				if msg != "" {
					fmt.Printf("    %s: %4d - expected '%v', got '%v'\n", msg, i, b[i], a[i])
				} else {
					fmt.Printf("    checkEqSliceString: %4d - expected '%v', got '%v'\n", i, b[i], a[i])
				}
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

func checkEqSliceTime(a, b []time.Time, t *testing.T, msg string) bool {
	if len(a) != len(b) {
		return false
	}

	if t == nil {
		for i, x := range a {
			if x != b[i] {
				if msg != "" {
					fmt.Printf("    %s: %4d - expected '%v', got '%v'\n", msg, i, b[i], a[i])
				} else {
					fmt.Printf("    checkEqSliceTime: %4d - expected '%v', got '%v'\n", i, b[i], a[i])
				}
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

func checkEqSliceDuration(a, b []time.Duration, t *testing.T, msg string) bool {
	if len(a) != len(b) {
		return false
	}

	if t == nil {
		for i, x := range a {
			if x != b[i] {
				if msg != "" {
					fmt.Printf("    %s: %4d - expected '%v', got '%v'\n", msg, i, b[i], a[i])
				} else {
					fmt.Printf("    checkEqSliceDuration: %4d - expected '%v', got '%v'\n", i, b[i], a[i])
				}
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

func checkEqPartitionMap(a, b map[int64][]int, t *testing.T, msg string) bool {
	if len(a) != len(b) {
		return false
	}

	// check if the two maps represent the same partitioning
	// the keys can be different, but the values must be the same
	if t == nil {
		for _, v := range a {
			found := false
			vSorted := sort.IntSlice(v)
			for _, w := range b {
				if checkEqSliceIntQuiet(vSorted, sort.IntSlice(w), t, msg) {
					found = true
					break
				}
			}
			if !found {
				fmt.Printf("    %s: expected partition %v not found\n", msg, v)
				return false
			}
		}
	} else {
		for _, v := range a {
			found := false
			vSorted := sort.IntSlice(v)
			for _, w := range b {
				if checkEqSliceIntQuiet(vSorted, sort.IntSlice(w), t, msg) {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("%s: expected partition %v not found\n", msg, v)
				return false
			}
		}
	}

	return true
}
