package preludiocompiler

type BOOL []bool
type INT32 []int
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

func (bt BaseType) String() string {
	switch bt {
	case NullType:
		return "null"
	case BoolType:
		return "bool"
	case Int16Type:
		return "int16"
	case Int32Type:
		return "int32"
	case Int64Type:
		return "int64"
	case Float32Type:
		return "float32"
	case Float64Type:
		return "float64"
	case StringType:
		return "string"
	case AnyType:
		return "any"
	case ErrorType:
		return "error"
	case NonBaseType:
		return "nonbase"
	}
	return "unknown"
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

func (p *Primitive) ApplyBinaryOp(op OPCODE, o Primitive) Primitive {
	switch op {

	/////////////////////////////////////////////////////////////////////////////////////
	///////////////////				BINARY MULTIPLICATION

	case OP_BINARY_MUL:

		if p.size == 0 || o.size == 0 {
			return Primitive{base: ErrorType}
		}

		switch p.base {
		case BoolType:
			switch o.base {
			case BoolType:
				if p.size == o.size {
					return Primitive{base: Int32Type, size: p.size}
				} else if p.size == 1 {
					return Primitive{base: Int32Type, size: o.size}
				} else if o.size == 1 {
					return Primitive{base: Int32Type, size: p.size}
				}

			case Int16Type:
				if p.size == o.size {
					return Primitive{base: Int16Type, size: p.size}
				} else if p.size == 1 {
					return Primitive{base: Int16Type, size: o.size}
				} else if o.size == 1 {
					return Primitive{base: Int16Type, size: p.size}
				}

			case Int32Type:
				if p.size == o.size {
					return Primitive{base: Int32Type, size: p.size}
				} else if p.size == 1 {
					return Primitive{base: Int32Type, size: o.size}
				} else if o.size == 1 {
					return Primitive{base: Int32Type, size: p.size}
				}

			case Int64Type:
				if p.size == o.size {
					return Primitive{base: Int64Type, size: p.size}
				} else if p.size == 1 {
					return Primitive{base: Int64Type, size: o.size}
				} else if o.size == 1 {
					return Primitive{base: Int64Type, size: p.size}
				}

			case Float32Type:
				if p.size == o.size {
					return Primitive{base: Float32Type, size: p.size}
				} else if p.size == 1 {
					return Primitive{base: Float32Type, size: o.size}
				} else if o.size == 1 {
					return Primitive{base: Float32Type, size: p.size}
				}

			case Float64Type:
				if p.size == o.size {
					return Primitive{base: Float64Type, size: p.size}
				} else if p.size == 1 {
					return Primitive{base: Float64Type, size: o.size}
				} else if o.size == 1 {
					return Primitive{base: Float64Type, size: p.size}
				}

			case StringType:
				if p.size == o.size {
					return Primitive{base: StringType, size: p.size}
				} else if p.size == 1 {
					return Primitive{base: StringType, size: o.size}
				} else if o.size == 1 {
					return Primitive{base: StringType, size: p.size}
				}

			case AnyType:
				return Primitive{base: ErrorType}

			case ErrorType:
				return Primitive{base: ErrorType}

			case NonBaseType:
				return Primitive{base: ErrorType}
			}

		case Int16Type:
			switch o.base {
			case BoolType:
				if p.size == o.size {
					return Primitive{base: Int16Type, size: p.size}
				} else if p.size == 1 {
					return Primitive{base: Int16Type, size: o.size}
				} else if o.size == 1 {
					return Primitive{base: Int16Type, size: p.size}
				}

			case Int16Type:
				if p.size == o.size {
					return Primitive{base: Int16Type, size: p.size}
				} else if p.size == 1 {
					return Primitive{base: Int16Type, size: o.size}
				} else if o.size == 1 {
					return Primitive{base: Int16Type, size: p.size}
				}

			case Int32Type:
				if p.size == o.size {
					return Primitive{base: Int32Type, size: p.size}
				} else if p.size == 1 {
					return Primitive{base: Int32Type, size: o.size}
				} else if o.size == 1 {
					return Primitive{base: Int32Type, size: p.size}
				}

			case Int64Type:
				if p.size == o.size {
					return Primitive{base: Int64Type, size: p.size}
				} else if p.size == 1 {
					return Primitive{base: Int64Type, size: o.size}
				} else if o.size == 1 {
					return Primitive{base: Int64Type, size: p.size}
				}

			case Float32Type:
				if p.size == o.size {
					return Primitive{base: Float32Type, size: p.size}
				} else if p.size == 1 {
					return Primitive{base: Float32Type, size: o.size}
				} else if o.size == 1 {
					return Primitive{base: Float32Type, size: p.size}
				}

			case Float64Type:
				if p.size == o.size {
					return Primitive{base: Float64Type, size: p.size}
				} else if p.size == 1 {
					return Primitive{base: Float64Type, size: o.size}
				} else if o.size == 1 {
					return Primitive{base: Float64Type, size: p.size}
				}

			case StringType:
				if p.size == o.size {
					return Primitive{base: StringType, size: p.size}
				} else if p.size == 1 {
					return Primitive{base: StringType, size: o.size}
				} else if o.size == 1 {
					return Primitive{base: StringType, size: p.size}
				}

			case AnyType:
				return Primitive{base: ErrorType}

			case ErrorType:
				return Primitive{base: ErrorType}

			case NonBaseType:
				return Primitive{base: ErrorType}
			}
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
