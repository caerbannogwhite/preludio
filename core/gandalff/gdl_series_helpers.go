package gandalff

func __initNullMask(size int) []uint8 {
	return make([]uint8, (size+7)>>3)
}
