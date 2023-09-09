package typesys

type OPCODE uint8

const (
	OP_START_STMT OPCODE = iota
	OP_END_STMT
	OP_START_PIPELINE
	OP_END_PIPELINE
	OP_START_FUNC_CALL
	OP_MAKE_FUNC_CALL
	OP_START_LIST
	OP_END_LIST
	OP_ADD_FUNC_PARAM
	OP_ADD_EXPR_TERM
	OP_PUSH_NAMED_PARAM
	OP_PUSH_ASSIGN_IDENT
	OP_PUSH_TERM
	OP_END_CHUNCK
	OP_VAR_DECL
	OP_VAR_ASSIGN
	OP_GOTO

	OP_BINARY_MUL
	OP_BINARY_DIV
	OP_BINARY_MOD
	OP_BINARY_ADD
	OP_BINARY_SUB
	OP_BINARY_POW

	OP_BINARY_EQ
	OP_BINARY_NE
	OP_BINARY_GE
	OP_BINARY_LE
	OP_BINARY_GT
	OP_BINARY_LT

	OP_BINARY_AND
	OP_BINARY_OR
	OP_BINARY_XOR
	OP_BINARY_COALESCE
	OP_BINARY_MODEL

	OP_BINARY_LSHIFT
	OP_BINARY_RSHIFT

	OP_UNARY_ADD
	OP_UNARY_SUB
	OP_UNARY_NOT

	NO_OP = 255
)

type PARAM1 uint8

const (
	TERM_NULL PARAM1 = iota
	TERM_BOOL
	TERM_INTEGER
	TERM_FLOAT
	TERM_STRING
	TERM_INTERVAL
	TERM_RANGE
	TERM_LIST
	TERM_PIPELINE
	TERM_SYMBOL
)

type LOG_TYPE uint8

const (
	LOG_INFO LOG_TYPE = iota
	LOG_WARNING
	LOG_ERROR
	LOG_DEBUG
)

type LogEnty struct {
	LogType LOG_TYPE `json:"logType"`
	Level   uint8    `json:"level"`
	Message string   `json:"message"`
}

type Columnar struct {
	Name         string   `json:"name"`
	Type         string   `json:"type"`
	ActualLength int      `json:"actualLength"` // actual length of the column
	Data         []string `json:"data"`
	Nulls        []bool   `json:"nulls"`
}

type PreludioOutput struct {
	Log  []LogEnty    `json:"log"`
	Data [][]Columnar `json:"data"`
}
