package main

const PRQL_INTERNAL_DIMENSION_SCALAR = 0
const PRQL_INTERNAL_DIMENSION_SERIES = 1
const PRQL_INTERNAL_DIMENSION_TABLE = 2

const PRQL_INTERNAL_TAG_ERROR = 0
const PRQL_INTERNAL_TAG_TERM = 1
const PRQL_INTERNAL_TAG_PARAM_NAME = 2
const PRQL_INTERNAL_TAG_ASSING_IDENT = 3

type PRQLInternal struct {
	Dimension     int
	Tag           int
	ScalarType    int
	ScalarBool    bool
	ScalarNumeric float64
	ScalarString  string
	Name          string
	ErrorMsg      string
}

func NewPRQLInternalError(msg string) *PRQLInternal {
	return &PRQLInternal{Tag: PRQL_INTERNAL_TAG_ERROR, ErrorMsg: msg}
}

func NewPRQLInternalScalarBool(val bool) *PRQLInternal {
	return &PRQLInternal{Dimension: PRQL_INTERNAL_DIMENSION_SCALAR, Tag: PRQL_INTERNAL_TAG_TERM, ScalarBool: val}
}

func NewPRQLInternalScalarNumeric(val float64) *PRQLInternal {
	return &PRQLInternal{Dimension: PRQL_INTERNAL_DIMENSION_SCALAR, Tag: PRQL_INTERNAL_TAG_TERM, ScalarNumeric: val}
}

func NewPRQLInternalScalarString(val string) *PRQLInternal {
	return &PRQLInternal{Dimension: PRQL_INTERNAL_DIMENSION_SCALAR, Tag: PRQL_INTERNAL_TAG_TERM, ScalarString: val}
}

func NewPRQLInternalParamName(val string) *PRQLInternal {
	return &PRQLInternal{Tag: PRQL_INTERNAL_TAG_PARAM_NAME, Name: val}
}

func NewPRQLInternalAssignIdent(val string) *PRQLInternal {
	return &PRQLInternal{Tag: PRQL_INTERNAL_TAG_ASSING_IDENT, Name: val}
}
