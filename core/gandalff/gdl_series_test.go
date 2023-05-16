package gandalff

import (
	"fmt"
	"strings"
	"testing"
	"typesys"
	"unsafe"
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

func TestXxx(t *testing.T) {

	data := `name,age,weight,junior
Alice C,29,75.0,F
John Doe,30,80.5,true
Bob,31,85.0,T
Jane H,25,60.0,false
Mary,28,70.0,false
Oliver,32,90.0,true
Ursula,27,65.0,f
Charlie,33,95.0,t
`

	// Create a new dataframe from the CSV data.
	df := NewBaseDataFrame().FromCSV().
		SetReader(strings.NewReader(data)).
		SetDelimiter(',').
		SetHeader(true).
		SetGuessDataTypeLen(3).
		Read()

	names := df.Series("name").(GDLSeriesString).__getDataPtr()

	m := make(map[uint64][]int)
	for i, v := range *names {
		m[*(*uint64)(unsafe.Pointer(unsafe.Pointer(v)))] = append(m[*(*uint64)(unsafe.Pointer(unsafe.Pointer(v)))], i)
	}

	fmt.Println(m)
}
