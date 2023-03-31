package typesys

type BOOL []bool
type INT16 []int16
type INT32 []int
type INT64 []int64
type FLOAT32 []float32
type FLOAT64 []float64
type STRING []string

type BaseType uint8

const (
	NullType BaseType = iota
	BoolType
	Int16Type
	Int32Type
	Int64Type
	Float32Type
	Float64Type
	StringType
	AnyType
	ErrorType
	NonBaseType
)

func (bt BaseType) ToString() string {
	switch bt {
	case NullType:
		return "Null"
	case BoolType:
		return "Bool"
	case Int16Type:
		return "Int16"
	case Int32Type:
		return "Int32"
	case Int64Type:
		return "Int64"
	case Float32Type:
		return "Float32"
	case Float64Type:
		return "Float64"
	case StringType:
		return "String"
	case AnyType:
		return "Any"
	case ErrorType:
		return "Error"
	case NonBaseType:
		return "Nonbase"
	}
	return "Unknown"
}

func (bt BaseType) ToGoType() string {
	switch bt {
	case NullType:
		return "nil"
	case BoolType:
		return "[]bool"
	case Int16Type:
		return "[]int16"
	case Int32Type:
		return "[]int"
	case Int64Type:
		return "[]int64"
	case Float32Type:
		return "[]float32"
	case Float64Type:
		return "[]float64"
	case StringType:
		return "[]string"
	case AnyType:
		return "interface{}"
	case ErrorType:
		return "error"
	case NonBaseType:
		return "interface{}"
	}
	return "Unknown"
}

func GoToPreludioTypeString(t interface{}) string {
	switch t.(type) {
	case []bool:
		return "Bool"
	case []int16:
		return "Int16"
	case []int:
		return "Int32"
	case []int64:
		return "Int64"
	case []float32:
		return "Float32"
	case []float64:
		return "Float64"
	case []string:
		return "String"
	}
	return "Unknown"
}

type Primitive struct {
	name   string
	base   BaseType
	size   int
	schema Schema
}

func (p *Primitive) GetName() string {
	return p.name
}

func (p *Primitive) IsBaseType() bool {
	return p.schema.primitives == nil || len(p.schema.primitives) == 0
}

func (op OPCODE) GetBinaryOpResultSize(lop, rop Primitive) int {
	// If one of the operands is null, the result is null
	if lop.size == 0 || rop.size == 0 {
		return -1
	}

	// If one of the operands is a scalar, the result is a vector
	var size int
	if lop.size == 1 {
		size = rop.size
	} else if rop.size == 1 {
		size = lop.size
	} else if lop.size == rop.size {
		size = lop.size
	} else {
		return -1
	}

	return size
}

func (op OPCODE) CommuteOperands(lop, rop Primitive) (Primitive, Primitive) {
	if op.IsCommutative() && lop.base > rop.base {
		return rop, lop
	}
	return lop, rop
}

func (op OPCODE) IsCommutative() bool {
	switch op {
	case OP_BINARY_ADD, OP_BINARY_MUL, OP_BINARY_AND, OP_BINARY_OR, OP_BINARY_XOR:
		return true
	}
	return false
}

func (op OPCODE) IsBinaryOp() bool {
	switch op {
	case OP_BINARY_ADD, OP_BINARY_SUB, OP_BINARY_MUL, OP_BINARY_DIV, OP_BINARY_MOD,
		OP_BINARY_AND, OP_BINARY_OR, OP_BINARY_XOR, OP_BINARY_LSHIFT, OP_BINARY_RSHIFT:
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
		switch lop.base {
		case BoolType:
			switch rop.base {
			case BoolType, Int32Type:
				return Primitive{base: Int32Type, size: size}
			case Int16Type:
				return Primitive{base: Int16Type, size: size}
			case Int64Type:
				return Primitive{base: Int64Type, size: size}
			case Float32Type:
				return Primitive{base: Float32Type, size: size}
			case Float64Type:
				return Primitive{base: Float64Type, size: size}
			case StringType:
				return Primitive{base: StringType, size: size}
			default:
				return Primitive{base: ErrorType}
			}

		case Int16Type:
			switch rop.base {
			case Int16Type:
				return Primitive{base: Int16Type, size: size}
			case Int32Type:
				return Primitive{base: Int32Type, size: size}
			case Int64Type:
				return Primitive{base: Int64Type, size: size}
			case Float32Type:
				return Primitive{base: Float32Type, size: size}
			case Float64Type:
				return Primitive{base: Float64Type, size: size}
			case StringType:
				return Primitive{base: StringType, size: size}
			default:
				return Primitive{base: ErrorType}
			}

		case Int32Type:
			switch rop.base {
			case Int32Type:
				return Primitive{base: Int32Type, size: size}
			case Int64Type:
				return Primitive{base: Int64Type, size: size}
			case Float32Type:
				return Primitive{base: Float32Type, size: size}
			case Float64Type:
				return Primitive{base: Float64Type, size: size}
			case StringType:
				return Primitive{base: StringType, size: size}
			default:
				return Primitive{base: ErrorType}
			}

		case Int64Type:
			switch rop.base {
			case Int64Type:
				return Primitive{base: Int64Type, size: size}
			case Float32Type:
				return Primitive{base: Float32Type, size: size}
			case Float64Type:
				return Primitive{base: Float64Type, size: size}
			case StringType:
				return Primitive{base: StringType, size: size}
			default:
				return Primitive{base: ErrorType}
			}

		case Float32Type:
			switch rop.base {
			case Float32Type:
				return Primitive{base: Float32Type, size: size}
			case Float64Type:
				return Primitive{base: Float64Type, size: size}
			case StringType:
				return Primitive{base: StringType, size: size}
			default:
				return Primitive{base: ErrorType}
			}

		case Float64Type:
			switch rop.base {
			case Float64Type:
				return Primitive{base: Float64Type, size: size}
			case StringType:
				return Primitive{base: StringType, size: size}
			default:
				return Primitive{base: ErrorType}
			}

		default:
			return Primitive{base: ErrorType}
		}

	/////////////////////////////////////////////////////////////////////////////////////
	///////////////////				BINARY DIVISION

	case OP_BINARY_DIV:
		switch lop.base {
		case BoolType:
			switch rop.base {
			case BoolType, Int32Type:
				return Primitive{base: Int32Type, size: size}
			case Int16Type:
				return Primitive{base: Int16Type, size: size}
			case Int64Type:
				return Primitive{base: Int64Type, size: size}
			case Float32Type:
				return Primitive{base: Float32Type, size: size}
			case Float64Type:
				return Primitive{base: Float64Type, size: size}
			default:
				return Primitive{base: ErrorType}
			}

		case Int16Type:
			switch rop.base {
			case BoolType, Int16Type:
				return Primitive{base: Int16Type, size: size}
			case Int32Type:
				return Primitive{base: Int32Type, size: size}
			case Int64Type:
				return Primitive{base: Int64Type, size: size}
			case Float32Type:
				return Primitive{base: Float32Type, size: size}
			case Float64Type:
				return Primitive{base: Float64Type, size: size}
			default:
				return Primitive{base: ErrorType}
			}

		case Int32Type:
			switch rop.base {
			case BoolType, Int16Type, Int32Type:
				return Primitive{base: Int32Type, size: size}
			case Int64Type:
				return Primitive{base: Int64Type, size: size}
			case Float32Type:
				return Primitive{base: Float32Type, size: size}
			case Float64Type:
				return Primitive{base: Float64Type, size: size}
			default:
				return Primitive{base: ErrorType}
			}

		case Int64Type:
			switch rop.base {
			case BoolType, Int16Type, Int32Type, Int64Type:
				return Primitive{base: Int64Type, size: size}
			case Float32Type:
				return Primitive{base: Float32Type, size: size}
			case Float64Type:
				return Primitive{base: Float64Type, size: size}
			default:
				return Primitive{base: ErrorType}
			}

		case Float32Type:
			switch rop.base {
			case BoolType, Int16Type, Int32Type, Int64Type, Float32Type:
				return Primitive{base: Float32Type, size: size}
			case Float64Type:
				return Primitive{base: Float64Type, size: size}
			default:
				return Primitive{base: ErrorType}
			}

		case Float64Type:
			switch rop.base {
			case BoolType, Int16Type, Int32Type, Int64Type, Float32Type, Float64Type:
				return Primitive{base: Float64Type, size: size}
			default:
				return Primitive{base: ErrorType}
			}

		default:
			return Primitive{base: ErrorType}
		}

	/////////////////////////////////////////////////////////////////////////////////////
	///////////////////				BINARY ADDITION

	case OP_BINARY_ADD:
		switch lop.base {
		case BoolType:
			switch rop.base {
			case BoolType, Int32Type:
				return Primitive{base: Int32Type, size: size}
			case Int16Type:
				return Primitive{base: Int16Type, size: size}
			case Int64Type:
				return Primitive{base: Int64Type, size: size}
			case Float32Type:
				return Primitive{base: Float32Type, size: size}
			case Float64Type:
				return Primitive{base: Float64Type, size: size}
			case StringType:
				return Primitive{base: StringType, size: size}
			default:
				return Primitive{base: ErrorType}
			}

		case Int16Type:
			switch rop.base {
			case Int16Type:
				return Primitive{base: Int16Type, size: size}
			case Int32Type:
				return Primitive{base: Int32Type, size: size}
			case Int64Type:
				return Primitive{base: Int64Type, size: size}
			case Float32Type:
				return Primitive{base: Float32Type, size: size}
			case Float64Type:
				return Primitive{base: Float64Type, size: size}
			case StringType:
				return Primitive{base: StringType, size: size}
			default:
				return Primitive{base: ErrorType}
			}

		case Int32Type:
			switch rop.base {
			case Int32Type:
				return Primitive{base: Int32Type, size: size}
			case Int64Type:
				return Primitive{base: Int64Type, size: size}
			case Float32Type:
				return Primitive{base: Float32Type, size: size}
			case Float64Type:
				return Primitive{base: Float64Type, size: size}
			case StringType:
				return Primitive{base: StringType, size: size}
			default:
				return Primitive{base: ErrorType}
			}

		case Int64Type:
			switch rop.base {
			case Int64Type:
				return Primitive{base: Int64Type, size: size}
			case Float32Type:
				return Primitive{base: Float32Type, size: size}
			case Float64Type:
				return Primitive{base: Float64Type, size: size}
			case StringType:
				return Primitive{base: StringType, size: size}
			default:
				return Primitive{base: ErrorType}
			}

		case Float32Type:
			switch rop.base {
			case Float32Type:
				return Primitive{base: Float32Type, size: size}
			case Float64Type:
				return Primitive{base: Float64Type, size: size}
			case StringType:
				return Primitive{base: StringType, size: size}
			default:
				return Primitive{base: ErrorType}
			}

		case Float64Type:
			switch rop.base {
			case Float64Type:
				return Primitive{base: Float64Type, size: size}
			case StringType:
				return Primitive{base: StringType, size: size}
			default:
				return Primitive{base: ErrorType}
			}

		case StringType:
			switch rop.base {
			case StringType:
				return Primitive{base: StringType, size: size}
			default:
				return Primitive{base: ErrorType}
			}

		default:
			return Primitive{base: ErrorType}
		}

	/////////////////////////////////////////////////////////////////////////////////////
	///////////////////				BINARY SUBTRACTION

	case OP_BINARY_SUB:
		switch lop.base {
		case BoolType:
			switch rop.base {
			case BoolType, Int32Type:
				return Primitive{base: Int32Type, size: size}
			case Int16Type:
				return Primitive{base: Int16Type, size: size}
			case Int64Type:
				return Primitive{base: Int64Type, size: size}
			case Float32Type:
				return Primitive{base: Float32Type, size: size}
			case Float64Type:
				return Primitive{base: Float64Type, size: size}
			default:
				return Primitive{base: ErrorType}
			}

		case Int16Type:
			switch rop.base {
			case BoolType, Int16Type:
				return Primitive{base: Int16Type, size: size}
			case Int32Type:
				return Primitive{base: Int32Type, size: size}
			case Int64Type:
				return Primitive{base: Int64Type, size: size}
			case Float32Type:
				return Primitive{base: Float32Type, size: size}
			case Float64Type:
				return Primitive{base: Float64Type, size: size}
			default:
				return Primitive{base: ErrorType}
			}

		case Int32Type:
			switch rop.base {
			case BoolType, Int16Type, Int32Type:
				return Primitive{base: Int32Type, size: size}
			case Int64Type:
				return Primitive{base: Int64Type, size: size}
			case Float32Type:
				return Primitive{base: Float32Type, size: size}
			case Float64Type:
				return Primitive{base: Float64Type, size: size}
			default:
				return Primitive{base: ErrorType}
			}

		case Int64Type:
			switch rop.base {
			case BoolType, Int16Type, Int32Type, Int64Type:
				return Primitive{base: Int64Type, size: size}
			case Float32Type:
				return Primitive{base: Float32Type, size: size}
			case Float64Type:
				return Primitive{base: Float64Type, size: size}
			default:
				return Primitive{base: ErrorType}
			}

		case Float32Type:
			switch rop.base {
			case BoolType, Int16Type, Int32Type, Int64Type, Float32Type:
				return Primitive{base: Float32Type, size: size}
			case Float64Type:
				return Primitive{base: Float64Type, size: size}
			default:
				return Primitive{base: ErrorType}
			}

		case Float64Type:
			switch rop.base {
			case BoolType, Int16Type, Int32Type, Int64Type, Float32Type, Float64Type:
				return Primitive{base: Float64Type, size: size}
			default:
				return Primitive{base: ErrorType}
			}

		default:
			return Primitive{base: ErrorType}
		}
	}

	return Primitive{base: ErrorType}
}

type Schema struct {
	primitives []Primitive
}

func (s *Schema) IsEqual(o Schema) bool {
	if len(s.primitives) != len(o.primitives) {
		return false
	}
	for i, p := range s.primitives {
		if p.name != o.primitives[i].name {
			return false
		}
		if p.base != o.primitives[i].base {
			return false
		}
		if !p.schema.IsEqual(o.primitives[i].schema) {
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
