package typesys

type BOOL []bool
type INT16 []int16
type INT32 []int32
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
		return "[]int32"
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
		switch lop.Base {
		case BoolType:
			switch rop.Base {
			case BoolType, Int32Type:
				return Primitive{Base: Int32Type, Size: size}
			case Int16Type:
				return Primitive{Base: Int16Type, Size: size}
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

		case Int16Type:
			switch rop.Base {
			case Int16Type:
				return Primitive{Base: Int16Type, Size: size}
			case Int32Type:
				return Primitive{Base: Int32Type, Size: size}
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

		case Int32Type:
			switch rop.Base {
			case Int32Type:
				return Primitive{Base: Int32Type, Size: size}
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
		case BoolType:
			switch rop.Base {
			case BoolType, Int32Type:
				return Primitive{Base: Int32Type, Size: size}
			case Int16Type:
				return Primitive{Base: Int16Type, Size: size}
			case Int64Type:
				return Primitive{Base: Int64Type, Size: size}
			case Float32Type:
				return Primitive{Base: Float32Type, Size: size}
			case Float64Type:
				return Primitive{Base: Float64Type, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case Int16Type:
			switch rop.Base {
			case BoolType, Int16Type:
				return Primitive{Base: Int16Type, Size: size}
			case Int32Type:
				return Primitive{Base: Int32Type, Size: size}
			case Int64Type:
				return Primitive{Base: Int64Type, Size: size}
			case Float32Type:
				return Primitive{Base: Float32Type, Size: size}
			case Float64Type:
				return Primitive{Base: Float64Type, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case Int32Type:
			switch rop.Base {
			case BoolType, Int16Type, Int32Type:
				return Primitive{Base: Int32Type, Size: size}
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
			case BoolType, Int16Type, Int32Type, Int64Type:
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
			case BoolType, Int16Type, Int32Type, Int64Type, Float32Type:
				return Primitive{Base: Float32Type, Size: size}
			case Float64Type:
				return Primitive{Base: Float64Type, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case Float64Type:
			switch rop.Base {
			case BoolType, Int16Type, Int32Type, Int64Type, Float32Type, Float64Type:
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
		case BoolType:
			switch rop.Base {
			case BoolType, Int32Type:
				return Primitive{Base: Int32Type, Size: size}
			case Int16Type:
				return Primitive{Base: Int16Type, Size: size}
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

		case Int16Type:
			switch rop.Base {
			case Int16Type:
				return Primitive{Base: Int16Type, Size: size}
			case Int32Type:
				return Primitive{Base: Int32Type, Size: size}
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

		case Int32Type:
			switch rop.Base {
			case Int32Type:
				return Primitive{Base: Int32Type, Size: size}
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

		case StringType:
			switch rop.Base {
			case StringType:
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
		case BoolType:
			switch rop.Base {
			case BoolType, Int32Type:
				return Primitive{Base: Int32Type, Size: size}
			case Int16Type:
				return Primitive{Base: Int16Type, Size: size}
			case Int64Type:
				return Primitive{Base: Int64Type, Size: size}
			case Float32Type:
				return Primitive{Base: Float32Type, Size: size}
			case Float64Type:
				return Primitive{Base: Float64Type, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case Int16Type:
			switch rop.Base {
			case BoolType, Int16Type:
				return Primitive{Base: Int16Type, Size: size}
			case Int32Type:
				return Primitive{Base: Int32Type, Size: size}
			case Int64Type:
				return Primitive{Base: Int64Type, Size: size}
			case Float32Type:
				return Primitive{Base: Float32Type, Size: size}
			case Float64Type:
				return Primitive{Base: Float64Type, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case Int32Type:
			switch rop.Base {
			case BoolType, Int16Type, Int32Type:
				return Primitive{Base: Int32Type, Size: size}
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
			case BoolType, Int16Type, Int32Type, Int64Type:
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
			case BoolType, Int16Type, Int32Type, Int64Type, Float32Type:
				return Primitive{Base: Float32Type, Size: size}
			case Float64Type:
				return Primitive{Base: Float64Type, Size: size}
			default:
				return Primitive{Base: ErrorType}
			}

		case Float64Type:
			switch rop.Base {
			case BoolType, Int16Type, Int32Type, Int64Type, Float32Type, Float64Type:
				return Primitive{Base: Float64Type, Size: size}
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
