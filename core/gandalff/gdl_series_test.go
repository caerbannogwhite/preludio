package gandalff

import (
	"testing"
	"typesys"
)

func Test_Series(t *testing.T) {

	s := NewSeries("test", typesys.BoolType, true, false, []bool{true, false, true, false, true, false, true, false, true, false}, nil)

	r := s.Append(true).
		Append([]NullableBool{{true, true}, {true, false}}).
		Filter([]bool{true, false, true, false, true, false, true, false, true, false, true, true, false})

	if e, ok := r.(SeriesError); ok {
		t.Errorf("Expected a series, got an error: %s", e.GetError())
	}

	if r.Len() != 7 {
		t.Errorf("Expected length of 7, got %d", r.Len())
	}
}
