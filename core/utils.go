package preludiocore

import "typesys"

func truncate(s string, n int) string {
	if len(s) > n {
		return s[:n-3] + "..."
	}
	return s
}

func operatorToString(op typesys.OPCODE) string {
	switch op {
	case typesys.OP_BINARY_MUL:
		return "*"
	case typesys.OP_BINARY_DIV:
		return "/"
	case typesys.OP_BINARY_MOD:
		return "%"
	case typesys.OP_BINARY_POW:
		return "**"
	case typesys.OP_BINARY_ADD:
		return "+"
	case typesys.OP_BINARY_SUB:
		return "-"
	case typesys.OP_BINARY_AND:
		return "and"
	case typesys.OP_BINARY_OR:
		return "or"
	case typesys.OP_BINARY_EQ:
		return "=="
	case typesys.OP_BINARY_NE:
		return "!="
	case typesys.OP_BINARY_LT:
		return "<"
	case typesys.OP_BINARY_LE:
		return "<="
	case typesys.OP_BINARY_GT:
		return ">"
	case typesys.OP_BINARY_GE:
		return ">="
	default:
		return ""
	}
}
