package main

type UserDefinedFunction func(*PreludioVM)

// type PreludioBool bool
// type PreludioInt int
// type PreludioFloat float64
// type PreludioString string

type PreludioSymbol string

type PreludioList []*PreludioInternal

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
