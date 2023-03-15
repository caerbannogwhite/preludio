package preludiocore

type UserDefinedFunction func(*ByteEater)

// type PreludioBool bool
// type PreludioInt int
// type PreludioFloat float64
// type PreludioString string

type __p_symbol__ string

type __p_list__ []__p_intern__

func BoolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func BoolToFloat64(b bool) float64 {
	if b {
		return 1.0
	}
	return 0.0
}

// type __p_type__ interface {
// 	Sum(__p_type__) __p_type__
// 	Sub(__p_type__) __p_type__
// 	Mul(__p_type__) __p_type__
// 	Div(__p_type__) __p_type__
// 	Mod(__p_type__) __p_type__
// 	Pow(__p_type__) __p_type__
// 	And(__p_type__) __p_type__
// 	Or(__p_type__) __p_type__
// 	Not() __p_type__
// 	Equal(__p_type__) __p_type__
// 	NotEqual(__p_type__) __p_type__
// 	Greater(__p_type__) __p_type__
// 	GreaterEqual(__p_type__) __p_type__
// 	Less(__p_type__) __p_type__
// 	LessEqual(__p_type__) __p_type__
// 	Len() int
// 	ToString() string
// }
