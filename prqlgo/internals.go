package main

// const PRQL_INTERNAL_DIMENSION_SCALAR = 0
// const PRQL_INTERNAL_DIMENSION_SERIES = 1
// const PRQL_INTERNAL_DIMENSION_DATAFRAME = 2

const PRQL_INTERNAL_TAG_ERROR = 0
const PRQL_INTERNAL_TAG_TERM = 1
const PRQL_INTERNAL_TAG_PARAM_NAME = 2
const PRQL_INTERNAL_TAG_ASSING_IDENT = 3

type UserDefinedFunction func(*PrqlVirtualMachine)

type PrqlInternal struct {
	// Dimension int
	Tag int
	// Scalar    interface{}
	// Series    series.Series
	// DataFrame dataframe.DataFrame
	Value    interface{}
	Name     string
	ErrorMsg string
}

func NewPrqlInternalError(msg string) *PrqlInternal {
	return &PrqlInternal{Tag: PRQL_INTERNAL_TAG_ERROR, ErrorMsg: msg}
}

func NewPrqlInternalTerm(val interface{}) *PrqlInternal {
	return &PrqlInternal{Value: val}
}

func NewPrqlInternalParamName(val string) *PrqlInternal {
	return &PrqlInternal{Tag: PRQL_INTERNAL_TAG_PARAM_NAME, Name: val}
}

func NewPrqlInternalAssignIdent(val string) *PrqlInternal {
	return &PrqlInternal{Tag: PRQL_INTERNAL_TAG_ASSING_IDENT, Name: val}
}
