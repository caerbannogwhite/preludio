package gandalff

import (
	"testing"
	"typesys"
)

func Test_GDLSeries(t *testing.T) {

	s := NewGDLSeries("test", typesys.BoolType, true, false, []bool{true, false, true, false, true, false, true, false, true, false}, nil)

	r := s.Append(true).
		Append([]NullableBool{{true, true}, {true, false}}).
		FilterByMask([]bool{true, false, true, false, true, false, true, false, true, false, true, true, false})

	if e, ok := r.(GDLSeriesError); ok {
		t.Errorf("Expected a series, got an error: %s", e.Error())
	}

	if r.Len() != 7 {
		t.Errorf("Expected length of 7, got %d", r.Len())
	}
}
