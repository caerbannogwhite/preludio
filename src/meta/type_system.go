package preludiometa

import (
	"fmt"
	"time"
)

type BOOL []bool
type INT16 []int16
type INT32 []int32
type INT64 []int64
type FLOAT32 []float32
type FLOAT64 []float64
type STRING []string

type BaseType uint8
type BaseTypeCard struct {
	Base BaseType
	Card int
}

const (
	NullType BaseType = iota
	BoolType
	IntType
	// Int32Type
	Int64Type
	Float32Type
	Float64Type
	StringType
	TimeType
	DurationType
	AnyType
	ErrorType
	NonBaseType
)

func (bt BaseType) ToString() string {
	switch bt {
	case NullType:
		return "NA"
	case BoolType:
		return "Bool"
	case IntType:
		return "Int"
	// case Int32Type:
	// 	return "Int32"
	case Int64Type:
		return "Int64"
	case Float32Type:
		return "Float32"
	case Float64Type:
		return "Float64"
	case StringType:
		return "String"
	case TimeType:
		return "Time"
	case DurationType:
		return "Duration"
	case AnyType:
		return "Any"
	case ErrorType:
		return "Error"
	case NonBaseType:
		return "Nonbase"
	default:
		return "Unknown"
	}
}

func (bt BaseType) ToGoType() string {
	switch bt {
	case NullType:
		return "nil"
	case BoolType:
		return "[]bool"
	case IntType:
		return "[]int"
	// case Int32Type:
	// 	return "[]int32"
	case Int64Type:
		return "[]int64"
	case Float32Type:
		return "[]float32"
	case Float64Type:
		return "[]float64"
	case StringType:
		return "[]string"
	case TimeType:
		return "[]time.Time"
	case DurationType:
		return "[]time.Duration"
	case AnyType:
		return "interface{}"
	case ErrorType:
		return "error"
	case NonBaseType:
		return "interface{}"
	default:
		return "Unknown"
	}
}

func (bt BaseType) CanCoerceTo(other BaseType) bool {
	switch bt {
	case NullType:
		return true

	case BoolType:
		switch other {
		case BoolType, IntType, Int64Type, Float32Type, Float64Type, StringType:
			return true
		default:
			return false
		}

	case IntType:
		switch other {
		case Int64Type, Float32Type, Float64Type, StringType:
			return true
		default:
			return false
		}

	case Int64Type:
		switch other {
		case Int64Type, Float32Type, Float64Type, StringType:
			return true
		default:
			return false
		}

	case Float32Type:
		switch other {
		case Float32Type, Float64Type, StringType:
			return true
		default:
			return false
		}

	case Float64Type:
		switch other {
		case Float64Type, StringType:
			return true
		default:
			return false
		}

	case StringType:
		switch other {
		case StringType:
			return true
		default:
			return false
		}

	case AnyType:
		return true

	default:
		return false
	}
}

func (btc BaseTypeCard) ToString() string {
	return fmt.Sprintf("%s[%d]", btc.Base.ToString(), btc.Card)
}

func GoToPreludioTypeString(t interface{}) string {
	switch t.(type) {
	case []bool:
		return "Bool"
	case []int:
		return "Int"
	// case []int32:
	// 	return "Int32"
	case []int64:
		return "Int64"
	case []float32:
		return "Float32"
	case []float64:
		return "Float64"
	case []string:
		return "String"
	case []time.Time:
		return "Time"
	case []time.Duration:
		return "Duration"
	default:
		return "Unknown"
	}
}

type Primitive struct {
	Nullable bool
	Name     string
	Base     BaseType
	Size     int
	Schema   Schema
}

func (p *Primitive) GetName() string {
	return p.Name
}

func (p *Primitive) IsBaseType() bool {
	return p.Schema.primitives == nil || len(p.Schema.primitives) == 0
}

func (op OPCODE) GetBinaryOpResultSize(lop, rop Primitive) int {
	// If one of the operands is null, the result is null
	if lop.Size == 0 || rop.Size == 0 {
		return -1
	}

	// If one of the operands is a scalar, the result is a vector
	var size int
	if lop.Size == 1 {
		size = rop.Size
	} else if rop.Size == 1 {
		size = lop.Size
	} else if lop.Size == rop.Size {
		size = lop.Size
	} else {
		return -1
	}

	return size
}

func (op OPCODE) CommuteOperands(lop, rop Primitive) (Primitive, Primitive) {
	if op.IsCommutative() && lop.Base > rop.Base {
		return rop, lop
	}
	return lop, rop
}

func (op OPCODE) IsCommutative() bool {
	switch op {
	case OP_BINARY_ADD, OP_BINARY_MUL, OP_BINARY_AND, OP_BINARY_OR, OP_BINARY_XOR, OP_BINARY_EQ, OP_BINARY_NE:
		return true
	}
	return false
}

func (op OPCODE) IsBinaryOp() bool {
	switch op {
	case OP_BINARY_ADD, OP_BINARY_SUB, OP_BINARY_MUL, OP_BINARY_DIV, OP_BINARY_MOD, OP_BINARY_EXP,
		OP_BINARY_AND, OP_BINARY_OR, OP_BINARY_XOR, OP_BINARY_LSHIFT, OP_BINARY_RSHIFT,
		OP_BINARY_EQ, OP_BINARY_NE, OP_BINARY_LT, OP_BINARY_LE, OP_BINARY_GT, OP_BINARY_GE:
		return true
	}
	return false
}

func (op OPCODE) IsUnaryOp() bool {
	switch op {
	case OP_UNARY_ADD, OP_UNARY_SUB, OP_UNARY_NOT:
		return true
	}
	return false
}

func (op OPCODE) GetBinaryOpResultType(lop, rop Primitive) Primitive {

	lop, rop = op.CommuteOperands(lop, rop)
	size := op.GetBinaryOpResultSize(lop, rop)

	switch op {

	/////////////////////////////////////////////////////////////////////////////////////
	///////////////////				BINARY MULTIPLICATION

	case OP_BINARY_MUL:
		switch lop.Base {
		case NullType:
			switch rop.Base {
			case NullType, BoolType, IntType, Int64Type, Float32Type, Float64Type:
				return Primitive{Base: NullType, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case BoolType:
			switch rop.Base {
			case IntType:
				return Primitive{Base: IntType, Size: size}
			case BoolType, Int64Type:
				return Primitive{Base: Int64Type, Size: size}
			case Float32Type:
				return Primitive{Base: Float32Type, Size: size}
			case Float64Type:
				return Primitive{Base: Float64Type, Size: size}
			case StringType:
				return Primitive{Base: StringType, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case IntType:
			switch rop.Base {
			case IntType:
				return Primitive{Base: IntType, Size: size}
			case Int64Type:
				return Primitive{Base: Int64Type, Size: size}
			case Float32Type:
				return Primitive{Base: Float32Type, Size: size}
			case Float64Type:
				return Primitive{Base: Float64Type, Size: size}
			case StringType:
				return Primitive{Base: StringType, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case Int64Type:
			switch rop.Base {
			case Int64Type:
				return Primitive{Base: Int64Type, Size: size}
			case Float32Type:
				return Primitive{Base: Float32Type, Size: size}
			case Float64Type:
				return Primitive{Base: Float64Type, Size: size}
			case StringType:
				return Primitive{Base: StringType, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case Float32Type:
			switch rop.Base {
			case Float32Type:
				return Primitive{Base: Float32Type, Size: size}
			case Float64Type:
				return Primitive{Base: Float64Type, Size: size}
			case StringType:
				return Primitive{Base: StringType, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case Float64Type:
			switch rop.Base {
			case Float64Type:
				return Primitive{Base: Float64Type, Size: size}
			case StringType:
				return Primitive{Base: StringType, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		default:
			return Primitive{Base: ErrorType}
		}

	/////////////////////////////////////////////////////////////////////////////////////
	///////////////////				BINARY DIVISION

	case OP_BINARY_DIV:
		switch lop.Base {
		case NullType:
			switch rop.Base {
			case NullType, BoolType, IntType, Int64Type, Float32Type, Float64Type:
				return Primitive{Base: NullType, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case BoolType:
			switch rop.Base {
			case Float32Type:
				return Primitive{Base: Float32Type, Size: size}
			case BoolType, IntType, Int64Type, Float64Type, StringType:
				return Primitive{Base: Float64Type, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case IntType:
			switch rop.Base {
			case Float32Type:
				return Primitive{Base: Float32Type, Size: size}
			case BoolType, IntType, Int64Type, Float64Type, StringType:
				return Primitive{Base: Float64Type, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case Int64Type:
			switch rop.Base {
			case Float32Type:
				return Primitive{Base: Float32Type, Size: size}
			case BoolType, IntType, Int64Type, Float64Type, StringType:
				return Primitive{Base: Float64Type, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case Float32Type:
			switch rop.Base {
			case BoolType, IntType, Int64Type, Float32Type:
				return Primitive{Base: Float32Type, Size: size}
			case Float64Type:
				return Primitive{Base: Float64Type, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case Float64Type:
			switch rop.Base {
			case BoolType, IntType, Int64Type, Float32Type, Float64Type:
				return Primitive{Base: Float64Type, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		default:
			return Primitive{Base: ErrorType}
		}

	/////////////////////////////////////////////////////////////////////////////////////
	///////////////////				BINARY MODULUS
	case OP_BINARY_MOD:
		switch lop.Base {
		case NullType:
			switch rop.Base {
			case NullType, BoolType, IntType, Int64Type, Float32Type, Float64Type:
				return Primitive{Base: NullType, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case BoolType:
			switch rop.Base {
			case BoolType, IntType, Int64Type, Float64Type:
				return Primitive{Base: Float64Type, Size: size}
			case Float32Type:
				return Primitive{Base: Float32Type, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case IntType:
			switch rop.Base {
			case BoolType, IntType, Int64Type, Float64Type:
				return Primitive{Base: Float64Type, Size: size}
			case Float32Type:
				return Primitive{Base: Float32Type, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case Int64Type:
			switch rop.Base {
			case BoolType, IntType, Int64Type, Float64Type:
				return Primitive{Base: Float64Type, Size: size}
			case Float32Type:
				return Primitive{Base: Float32Type, Size: size}
			}

		case Float32Type:
			switch rop.Base {
			case BoolType, IntType, Int64Type, Float32Type, Float64Type:
				return Primitive{Base: Float32Type, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case Float64Type:
			switch rop.Base {
			case BoolType, IntType, Int64Type, Float32Type, Float64Type:
				return Primitive{Base: Float64Type, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		default:
			return Primitive{Base: ErrorType}
		}

	/////////////////////////////////////////////////////////////////////////////////////
	///////////////////				BINARY POWER
	case OP_BINARY_EXP:
		switch lop.Base {
		case NullType:
			switch rop.Base {
			case NullType, BoolType, IntType, Int64Type, Float32Type, Float64Type:
				return Primitive{Base: NullType, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case BoolType:
			switch rop.Base {
			case BoolType, IntType, Int64Type:
				return Primitive{Base: Int64Type, Size: size}
			case Float32Type, Float64Type:
				return Primitive{Base: Float64Type, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case IntType:
			switch rop.Base {
			case BoolType, IntType, Int64Type:
				return Primitive{Base: Int64Type, Size: size}
			case Float32Type, Float64Type:
				return Primitive{Base: Float64Type, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case Int64Type:
			switch rop.Base {
			case BoolType, IntType, Int64Type:
				return Primitive{Base: Int64Type, Size: size}
			case Float32Type, Float64Type:
				return Primitive{Base: Float64Type, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case Float32Type:
			switch rop.Base {
			case BoolType, IntType, Int64Type, Float32Type, Float64Type:
				return Primitive{Base: Float64Type, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case Float64Type:
			switch rop.Base {
			case BoolType, IntType, Int64Type, Float32Type, Float64Type:
				return Primitive{Base: Float64Type, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		default:
			return Primitive{Base: ErrorType}
		}

	/////////////////////////////////////////////////////////////////////////////////////
	///////////////////				BINARY ADDITION

	case OP_BINARY_ADD:
		switch lop.Base {
		case NullType:
			switch rop.Base {
			case NullType, BoolType, IntType, Int64Type, Float32Type, Float64Type, TimeType, DurationType:
				return Primitive{Base: NullType, Size: size}
			case StringType:
				return Primitive{Base: StringType, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case BoolType:
			switch rop.Base {
			case IntType:
				return Primitive{Base: IntType, Size: size}
			case BoolType, Int64Type:
				return Primitive{Base: Int64Type, Size: size}
			case Float32Type:
				return Primitive{Base: Float32Type, Size: size}
			case Float64Type:
				return Primitive{Base: Float64Type, Size: size}
			case StringType:
				return Primitive{Base: StringType, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case IntType:
			switch rop.Base {
			case IntType:
				return Primitive{Base: IntType, Size: size}
			case Int64Type:
				return Primitive{Base: Int64Type, Size: size}
			case Float32Type:
				return Primitive{Base: Float32Type, Size: size}
			case Float64Type:
				return Primitive{Base: Float64Type, Size: size}
			case StringType:
				return Primitive{Base: StringType, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case Int64Type:
			switch rop.Base {
			case Int64Type:
				return Primitive{Base: Int64Type, Size: size}
			case Float32Type:
				return Primitive{Base: Float32Type, Size: size}
			case Float64Type:
				return Primitive{Base: Float64Type, Size: size}
			case StringType:
				return Primitive{Base: StringType, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case Float32Type:
			switch rop.Base {
			case Float32Type:
				return Primitive{Base: Float32Type, Size: size}
			case Float64Type:
				return Primitive{Base: Float64Type, Size: size}
			case StringType:
				return Primitive{Base: StringType, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case Float64Type:
			switch rop.Base {
			case Float64Type:
				return Primitive{Base: Float64Type, Size: size}
			case StringType:
				return Primitive{Base: StringType, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case TimeType:
			switch rop.Base {
			case StringType:
				return Primitive{Base: StringType, Size: size}
			case TimeType:
				return Primitive{Base: TimeType, Size: size}
			case DurationType:
				return Primitive{Base: TimeType, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case DurationType:
			switch rop.Base {
			case StringType:
				return Primitive{Base: StringType, Size: size}
			case TimeType:
				return Primitive{Base: TimeType, Size: size}
			case DurationType:
				return Primitive{Base: DurationType, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case StringType:
			switch rop.Base {
			case StringType, TimeType, DurationType:
				return Primitive{Base: StringType, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		default:
			return Primitive{Base: ErrorType}
		}

	/////////////////////////////////////////////////////////////////////////////////////
	///////////////////				BINARY SUBTRACTION

	case OP_BINARY_SUB:
		switch lop.Base {
		case NullType:
			switch rop.Base {
			case NullType, BoolType, IntType, Int64Type, Float32Type, Float64Type, TimeType, DurationType:
				return Primitive{Base: NullType, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case BoolType:
			switch rop.Base {
			case IntType:
				return Primitive{Base: IntType, Size: size}
			case BoolType, Int64Type:
				return Primitive{Base: Int64Type, Size: size}
			case Float32Type:
				return Primitive{Base: Float32Type, Size: size}
			case Float64Type:
				return Primitive{Base: Float64Type, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case IntType:
			switch rop.Base {
			case BoolType, IntType:
				return Primitive{Base: IntType, Size: size}
			case Int64Type:
				return Primitive{Base: Int64Type, Size: size}
			case Float32Type:
				return Primitive{Base: Float32Type, Size: size}
			case Float64Type:
				return Primitive{Base: Float64Type, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case Int64Type:
			switch rop.Base {
			case BoolType, IntType, Int64Type:
				return Primitive{Base: Int64Type, Size: size}
			case Float32Type:
				return Primitive{Base: Float32Type, Size: size}
			case Float64Type:
				return Primitive{Base: Float64Type, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case Float32Type:
			switch rop.Base {
			case BoolType, IntType, Int64Type, Float32Type:
				return Primitive{Base: Float32Type, Size: size}
			case Float64Type:
				return Primitive{Base: Float64Type, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case Float64Type:
			switch rop.Base {
			case BoolType, IntType, Int64Type, Float32Type, Float64Type:
				return Primitive{Base: Float64Type, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case TimeType:
			switch rop.Base {
			case TimeType:
				return Primitive{Base: DurationType, Size: size}
			case DurationType:
				return Primitive{Base: TimeType, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case DurationType:
			switch rop.Base {
			case DurationType:
				return Primitive{Base: DurationType, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		default:
			return Primitive{Base: ErrorType}
		}

	/////////////////////////////////////////////////////////////////////////////////////
	///////////////////				BINARY EQUALITY

	case OP_BINARY_EQ:
		switch lop.Base {
		case NullType:
			return Primitive{Base: NullType, Size: size}

		case BoolType:
			switch rop.Base {
			case BoolType:
				return Primitive{Base: BoolType, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case IntType:
			switch rop.Base {
			case IntType, Int64Type, Float32Type, Float64Type:
				return Primitive{Base: BoolType, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case Int64Type:
			switch rop.Base {
			case Int64Type, Float32Type, Float64Type:
				return Primitive{Base: BoolType, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case Float32Type:
			switch rop.Base {
			case Float32Type, Float64Type:
				return Primitive{Base: BoolType, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case Float64Type:
			switch rop.Base {
			case Float64Type:
				return Primitive{Base: BoolType, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case StringType:
			switch rop.Base {
			case StringType:
				return Primitive{Base: BoolType, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case TimeType:
			switch rop.Base {
			case TimeType:
				return Primitive{Base: BoolType, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case DurationType:
			switch rop.Base {
			case DurationType:
				return Primitive{Base: BoolType, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		default:
			return Primitive{Base: ErrorType}
		}

	/////////////////////////////////////////////////////////////////////////////////////
	///////////////////				BINARY INEQUALITY

	case OP_BINARY_NE:
		switch lop.Base {
		case NullType:
			return Primitive{Base: NullType, Size: size}

		case BoolType:
			switch rop.Base {
			case BoolType:
				return Primitive{Base: BoolType, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case IntType:
			switch rop.Base {
			case IntType, Int64Type, Float32Type, Float64Type:
				return Primitive{Base: BoolType, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case Int64Type:
			switch rop.Base {
			case Int64Type, Float32Type, Float64Type:
				return Primitive{Base: BoolType, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case Float32Type:
			switch rop.Base {
			case Float32Type, Float64Type:
				return Primitive{Base: BoolType, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case Float64Type:
			switch rop.Base {
			case Float64Type:
				return Primitive{Base: BoolType, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case StringType:
			switch rop.Base {
			case StringType:
				return Primitive{Base: BoolType, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case TimeType:
			switch rop.Base {
			case TimeType:
				return Primitive{Base: BoolType, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case DurationType:
			switch rop.Base {
			case DurationType:
				return Primitive{Base: BoolType, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		default:
			return Primitive{Base: ErrorType}
		}

	/////////////////////////////////////////////////////////////////////////////////////
	///////////////////				BINARY LESS THAN

	case OP_BINARY_LT:
		switch lop.Base {
		case NullType:
			return Primitive{Base: NullType, Size: size}

		case BoolType:
			switch rop.Base {
			case BoolType, IntType, Int64Type, Float32Type, Float64Type:
				return Primitive{Base: BoolType, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case IntType:
			switch rop.Base {
			case BoolType, IntType, Int64Type, Float32Type, Float64Type:
				return Primitive{Base: BoolType, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case Int64Type:
			switch rop.Base {
			case BoolType, IntType, Int64Type, Float32Type, Float64Type:
				return Primitive{Base: BoolType, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case Float32Type:
			switch rop.Base {
			case BoolType, IntType, Int64Type, Float32Type, Float64Type:
				return Primitive{Base: BoolType, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case Float64Type:
			switch rop.Base {
			case BoolType, IntType, Int64Type, Float32Type, Float64Type:
				return Primitive{Base: BoolType, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case StringType:
			switch rop.Base {
			case StringType:
				return Primitive{Base: BoolType, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case TimeType:
			switch rop.Base {
			case TimeType:
				return Primitive{Base: BoolType, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case DurationType:
			switch rop.Base {
			case DurationType:
				return Primitive{Base: BoolType, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		default:
			return Primitive{Base: ErrorType}
		}

	/////////////////////////////////////////////////////////////////////////////////////
	///////////////////				BINARY LESS THAN OR EQUAL TO

	case OP_BINARY_LE:
		switch lop.Base {
		case NullType:
			return Primitive{Base: NullType, Size: size}

		case BoolType:
			switch rop.Base {
			case BoolType, IntType, Int64Type, Float32Type, Float64Type:
				return Primitive{Base: BoolType, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case IntType:
			switch rop.Base {
			case BoolType, IntType, Int64Type, Float32Type, Float64Type:
				return Primitive{Base: BoolType, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case Int64Type:
			switch rop.Base {
			case BoolType, IntType, Int64Type, Float32Type, Float64Type:
				return Primitive{Base: BoolType, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case Float32Type:
			switch rop.Base {
			case BoolType, IntType, Int64Type, Float32Type, Float64Type:
				return Primitive{Base: BoolType, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case Float64Type:
			switch rop.Base {
			case BoolType, IntType, Int64Type, Float32Type, Float64Type:
				return Primitive{Base: BoolType, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case StringType:
			switch rop.Base {
			case StringType:
				return Primitive{Base: BoolType, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case TimeType:
			switch rop.Base {
			case TimeType:
				return Primitive{Base: BoolType, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case DurationType:
			switch rop.Base {
			case DurationType:
				return Primitive{Base: BoolType, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		default:
			return Primitive{Base: ErrorType}
		}

	/////////////////////////////////////////////////////////////////////////////////////
	///////////////////				BINARY GREATER THAN

	case OP_BINARY_GT:
		switch lop.Base {
		case NullType:
			return Primitive{Base: NullType, Size: size}

		case BoolType:
			switch rop.Base {
			case BoolType, IntType, Int64Type, Float32Type, Float64Type:
				return Primitive{Base: BoolType, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case IntType:
			switch rop.Base {
			case BoolType, IntType, Int64Type, Float32Type, Float64Type:
				return Primitive{Base: BoolType, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case Int64Type:
			switch rop.Base {
			case BoolType, IntType, Int64Type, Float32Type, Float64Type:
				return Primitive{Base: BoolType, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case Float32Type:
			switch rop.Base {
			case BoolType, IntType, Int64Type, Float32Type, Float64Type:
				return Primitive{Base: BoolType, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case Float64Type:
			switch rop.Base {
			case BoolType, IntType, Int64Type, Float32Type, Float64Type:
				return Primitive{Base: BoolType, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case StringType:
			switch rop.Base {
			case StringType:
				return Primitive{Base: BoolType, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case TimeType:
			switch rop.Base {
			case TimeType:
				return Primitive{Base: BoolType, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case DurationType:
			switch rop.Base {
			case DurationType:
				return Primitive{Base: BoolType, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		default:
			return Primitive{Base: ErrorType}
		}

	/////////////////////////////////////////////////////////////////////////////////////
	///////////////////				BINARY GREATER THAN OR EQUAL TO

	case OP_BINARY_GE:
		switch lop.Base {
		case NullType:
			return Primitive{Base: NullType, Size: size}

		case BoolType:
			switch rop.Base {
			case BoolType, IntType, Int64Type, Float32Type, Float64Type:
				return Primitive{Base: BoolType, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case IntType:
			switch rop.Base {
			case BoolType, IntType, Int64Type, Float32Type, Float64Type:
				return Primitive{Base: BoolType, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case Int64Type:
			switch rop.Base {
			case BoolType, IntType, Int64Type, Float32Type, Float64Type:
				return Primitive{Base: BoolType, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case Float32Type:
			switch rop.Base {
			case BoolType, IntType, Int64Type, Float32Type, Float64Type:
				return Primitive{Base: BoolType, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case Float64Type:
			switch rop.Base {
			case BoolType, IntType, Int64Type, Float32Type, Float64Type:
				return Primitive{Base: BoolType, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case StringType:
			switch rop.Base {
			case StringType:
				return Primitive{Base: BoolType, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case TimeType:
			switch rop.Base {
			case TimeType:
				return Primitive{Base: BoolType, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case DurationType:
			switch rop.Base {
			case DurationType:
				return Primitive{Base: BoolType, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		default:
			return Primitive{Base: ErrorType}
		}

	/////////////////////////////////////////////////////////////////////////////////////
	///////////////////				BINARY AND
	case OP_BINARY_AND:
		switch lop.Base {
		case NullType:
			switch rop.Base {
			case NullType, BoolType:
				return Primitive{Base: NullType, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case BoolType:
			switch rop.Base {
			case BoolType:
				return Primitive{Base: BoolType, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}
		default:
			return Primitive{Base: ErrorType}
		}

	/////////////////////////////////////////////////////////////////////////////////////
	///////////////////				BINARY OR
	case OP_BINARY_OR:
		switch lop.Base {
		case NullType:
			switch rop.Base {
			case NullType:
				return Primitive{Base: NullType, Size: size}
			case BoolType:
				return Primitive{Base: BoolType, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case BoolType:
			switch rop.Base {
			case BoolType:
				return Primitive{Base: BoolType, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}
		default:
			return Primitive{Base: ErrorType}
		}
	}

	return Primitive{Base: ErrorType}
}

type Schema struct {
	primitives []Primitive
}

func (s *Schema) IsEqual(o Schema) bool {
	if len(s.primitives) != len(o.primitives) {
		return false
	}
	for i, p := range s.primitives {
		if p.Name != o.primitives[i].Name {
			return false
		}
		if p.Base != o.primitives[i].Base {
			return false
		}
		if !p.Schema.IsEqual(o.primitives[i].Schema) {
			return false
		}
	}
	return true
}

func InitSchema() Schema {
	return Schema{primitives: make([]Primitive, 0)}
}

func (s *Schema) AddPrimitive(p Primitive) {
	s.primitives = append(s.primitives, p)
}

func (s *Schema) GetDataTypes() []BaseType {
	types := make([]BaseType, len(s.primitives))
	for i, p := range s.primitives {
		types[i] = p.Base
	}
	return types
}
