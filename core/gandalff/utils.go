package gandalff

import "testing"

func checkEqSliceInt32(a, b []int32, t *testing.T, msg string) bool {
	if len(a) != len(b) {
		return false
	}

	for i, x := range a {
		if x != b[i] {
			t.Errorf("%s: %4d - expected '%v', got '%v'", msg, i, b[i], a[i])
			return false
		}
	}
	return true
}

func checkEqSliceInt64(a, b []int64, t *testing.T, msg string) bool {
	if len(a) != len(b) {
		return false
	}

	for i, x := range a {
		if x != b[i] {
			t.Errorf("%s: %4d - expected '%v', got '%v'", msg, i, b[i], a[i])
			return false
		}
	}
	return true
}

func checkEqSliceFloat64(a, b []float64, t *testing.T, msg string) bool {
	if len(a) != len(b) {
		return false
	}

	for i, x := range a {
		if x != b[i] {
			t.Errorf("%s: %4d - expected '%v', got '%v'", msg, i, b[i], a[i])
			return false
		}
	}
	return true
}

func checkEqSliceString(a, b []string, t *testing.T, msg string) bool {
	if len(a) != len(b) {
		return false
	}

	for i, x := range a {
		if x != b[i] {
			t.Errorf("%s: %4d - expected '%v', got '%v'", msg, i, b[i], a[i])
			return false
		}
	}
	return true
}
