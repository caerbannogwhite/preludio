package preludiometa

func (op OPCODE) ToString() string {
	switch op {
	case OP_BINARY_MUL:
		return "*"
	case OP_BINARY_DIV:
		return "/"
	case OP_BINARY_MOD:
		return "%"
	case OP_BINARY_EXP:
		return "^"
	case OP_BINARY_ADD:
		return "+"
	case OP_BINARY_SUB:
		return "-"
	case OP_BINARY_AND:
		return "and"
	case OP_BINARY_OR:
		return "or"
	case OP_BINARY_EQ:
		return "=="
	case OP_BINARY_NE:
		return "!="
	case OP_BINARY_LT:
		return "<"
	case OP_BINARY_LE:
		return "<="
	case OP_BINARY_GT:
		return ">"
	case OP_BINARY_GE:
		return ">="

	case OP_UNARY_ADD:
		return "+"
	case OP_UNARY_SUB:
		return "-"
	case OP_UNARY_NOT:
		return "!"

	default:
		return "UNKNOWN OPERATOR"
	}
}

func (op OPCODE) ToCodeString() string {
	switch op {
	case OP_BINARY_MUL:
		return "  * <BINARY_MUL>"
	case OP_BINARY_DIV:
		return "  / <BINARY_DIV>"
	case OP_BINARY_MOD:
		return "  % <BINARY_MOD>"
	case OP_BINARY_EXP:
		return "  ^ <BINARY_EXP>"
	case OP_BINARY_ADD:
		return "  + <BINARY_ADD>"
	case OP_BINARY_SUB:
		return "  - <BINARY_SUB>"
	case OP_BINARY_AND:
		return "and <BINARY_AND>"
	case OP_BINARY_OR:
		return " or <BINARY_OR>"
	case OP_BINARY_EQ:
		return " == <BINARY_EQ>"
	case OP_BINARY_NE:
		return " != <BINARY_NE>"
	case OP_BINARY_LT:
		return "  < <BINARY_LT>"
	case OP_BINARY_LE:
		return " <= <BINARY_LE>"
	case OP_BINARY_GT:
		return "  > <BINARY_GT>"
	case OP_BINARY_GE:
		return " >= <BINARY_GE>"

	case OP_UNARY_ADD:
		return "  + <UNARY_ADD>"
	case OP_UNARY_SUB:
		return "  - <UNARY_SUB>"
	case OP_UNARY_NOT:
		return "  ! <UNARY_NOT>"

	default:
		return "UNKNOWN OPERATOR"
	}
}
