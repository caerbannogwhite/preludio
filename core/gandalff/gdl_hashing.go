package gandalff

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
