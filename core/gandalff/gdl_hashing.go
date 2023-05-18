package gandalff

import "unsafe"

type SeriesPart struct {
	part map[uint64][]int
}

func HashStringPtrVec(vec *[]*string) map[uint64][]int {
	m := make(map[uint64][]int, DEFAULT_HASH_MAP_INITIAL_CAPACITY)

	var h uint64
	for i, v := range *vec {
		h = *(*uint64)(unsafe.Pointer(unsafe.Pointer(v)))
		m[h] = append(m[h], i)
	}
	return m
}

func CombineStringPtrVec(oldMap map[uint64][]int, vec *[]*string) map[uint64][]int {
	newMap := make(map[uint64][]int, DEFAULT_HASH_MAP_INITIAL_CAPACITY)

	var newHash uint64
	for hash, indexes := range oldMap {
		for _, index := range indexes {
			newHash = *(*uint64)(unsafe.Pointer(unsafe.Pointer((*vec)[index]))) + HASH_MAGIC_NUMBER + (hash << 6) + (hash >> 2)
			newMap[newHash] = append(newMap[newHash], index)
		}
	}

	return newMap
}

type HMElem struct {
	Key   uint64
	Value *[]int
}

type CustomHM [][]HMElem

func NewCustomHM(size int) *CustomHM {
	hm := make(CustomHM, size)
	return &hm
}

func (hm *CustomHM) Get(key uint64) *[]int {
	index := int(key % uint64(len(*hm)))
	for _, elem := range (*hm)[index] {
		if elem.Key == key {
			return elem.Value
		}
	}
	return nil
}

func (hm *CustomHM) Put(key uint64, value *[]int) {
	index := int(key % uint64(len(*hm)))
	(*hm)[index] = append((*hm)[index], HMElem{key, value})
}
