package gandalff

import (
	"testing"
)

func Test_GDLSeries(t *testing.T) {

	s := NewGDLSeriesBool("test", true, []bool{true, false, true, false, true, false, true, false, true, false})

	s.Append(true)
	s.Filter([]bool{true, false, true, true, false, true, true, false, true, true}).Append(true)
}
