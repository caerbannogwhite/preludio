package pseries

import "time"

type PSeriesType uint8

const (
	BoolType PSeriesType = iota
	IntType
	FloatType
	StringType
	TimeType
	DurationType
	ComplexType
	InterfaceType
)

func PSeriesTypeToString(t PSeriesType) string {
	switch t {
	case BoolType:
		return "bool"
	case IntType:
		return "int"
	case FloatType:
		return "float"
	case StringType:
		return "string"
	case TimeType:
		return "time"
	case DurationType:
		return "duration"
	case ComplexType:
		return "complex"
	case InterfaceType:
		return "interface"
	}
	return "unknown"
}

type PSeries interface {
	// Returns the length of the series.
	Len() int
	// Returns the name of the series.
	Name() string
	// Returns the type of the series.
	Type() PSeriesType
	// Returns the data of the series.
	Data() interface{}
	// Returns the data of the series as a slice of strings.
	StringData() []string
	// Returns the data of the series as a slice of bools.
	BoolData() []bool
	// Returns the data of the series as a slice of ints.
	IntData() []int
	// Returns the data of the series as a slice of floats.
	FloatData() []float64
	// Returns the data of the series as a slice of time.Time.
	TimeData() []time.Time
	// Returns the data of the series as a slice of time.Duration.
	DurationData() []time.Duration
	// Returns the data of the series as a slice of complex128.
	ComplexData() []complex128
	// Returns the data of the series as a slice of interface{}.
}

type PSeriesBool struct {
	name string
	data []bool
}

func NewPSeriesBool(name string, data []bool) PSeriesBool {
	return PSeriesBool{name: name, data: data}
}

func (s PSeriesBool) Len() int {
	return len(s.data)
}

func (s PSeriesBool) Name() string {
	return s.name
}

func (s PSeriesBool) Type() PSeriesType {
	return BoolType
}

func (s PSeriesBool) Data() interface{} {
	return s.data
}
