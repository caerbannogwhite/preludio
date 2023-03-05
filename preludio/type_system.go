package preludio

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
