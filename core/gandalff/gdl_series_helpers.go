package gandalff

func __binVecInit(size int) []uint8 {
	return make([]uint8, (size+7)>>3)
}

func __binVecResize(v []uint8, size int) []uint8 {
	if size <= len(v)<<3 {
		return v
	}
	v2 := make([]uint8, (size+7)>>3)
	copy(v2, v)
	return v2
}

func __binVecFilterByIndices(srcVec []uint8, indices []int) []uint8 {
	dstVec := make([]uint8, (len(indices)+7)>>3)
	for dstIdx, srcIdx := range indices {
		if srcIdx%8 > dstIdx%8 {
			dstVec[dstIdx>>3] |= ((srcVec[srcIdx>>3] & (1 << uint(srcIdx%8))) >> uint(srcIdx%8-dstIdx%8))
		} else {
			dstVec[dstIdx>>3] |= ((srcVec[srcIdx>>3] & (1 << uint(srcIdx%8))) << uint(dstIdx%8-srcIdx%8))
		}
	}
	return dstVec
}

func __binVecCount(v []uint8) int {
	count := 0
	for _, x := range v {
		for x != 0 {
			count += int(x & 1)
			x >>= 1
		}
	}
	return count
}
