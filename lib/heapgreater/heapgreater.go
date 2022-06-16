package heapgreater

type comparator1 interface {
	Less(index1, index2 int) bool
	// Swap(index1, index2 int)
	// Len() int
}

type heapgreater struct {
	HeapSize int //maxIndex +1
	Value    []interface{}
	IndexMap map[interface{}][]int
	comparator1
}

func (t *heapgreater) HeapInit() {
	t.Value = make([]interface{}, 0)
	t.HeapSize = 0
	t.IndexMap = make(map[interface{}][]int)
}

func (t *heapgreater) heapinsert(index int) {
	//index  =0 also in this base case
	for t.Less((index-1)/2, index) {
		t.Swap(index, (index-1)>>1)
		index = (index - 1) >> 1
	}
}

func leftChildIndex(index int) int {
	return 2*index + 1
}

func rightChildIndex(index int) int {
	return 2*index + 2
}

func (t *heapgreater) heapify(index int) {

	for leftChildIndex(index) < t.HeapSize {
		largest := index
		if t.Less(index, leftChildIndex(index)) {
			largest = leftChildIndex(index)
		}
		hasRight := false

		if rightChildIndex(index) < t.HeapSize {
			hasRight = true
		}

		if hasRight {
			if t.Less(largest, rightChildIndex(index)) {
				largest = rightChildIndex(index)
			}
		}
		if largest == index {
			break
		} else {
			t.Swap(index, largest)
			index = largest
		}
	}
}

func (t *heapgreater) addIndexMap(key1 interface{}, value1 int) {
	t.IndexMap[key1] = append(t.IndexMap[key1], value1)
}
func (t *heapgreater) removeIndexMap(key1 interface{}, value1 int) {
	value := t.IndexMap[key1]
	if len(value) > 1 {
		for i, v := range value {
			if v == value1 {
				value[i], value[len(value)-1] = value[len(value)-1], value[i]
				t.IndexMap[key1] = t.IndexMap[key1][:len(value)-1]
				break
			}
		}
	} else {
		delete(t.IndexMap, key1)
	}
}

func (t *heapgreater) swapIndexMap(key1, key2 interface{}, value1, value2 int) {
	v1 := t.IndexMap[key1]
	v2 := t.IndexMap[key2]

	if key1 != key2 {
		for i, v := range v1 {
			if v == value2 {
				v1[i] = value1
			}
		}
		for i, v := range v2 {
			if v == value1 {
				v2[i] = value2
			}
		}

	}

}

func (t *heapgreater) Add(value interface{}) {
	if len(t.Value) == t.HeapSize {
		t.Value = append(t.Value, value)
	} else {
		t.Value[t.HeapSize] = value
	}
	t.addIndexMap(t.Value[t.HeapSize], t.HeapSize)
	t.heapinsert(t.HeapSize)
	t.HeapSize++
}

func (t *heapgreater) Poll() interface{} {
	b := t.Value
	temp := b[0]
	t.Swap(0, t.HeapSize-1)

	t.removeIndexMap(b[t.HeapSize-1], t.HeapSize-1)
	t.HeapSize -= 1
	t.heapify(0)
	return temp
}

func (t *heapgreater) Peek() interface{} {
	b := t.Value
	return b[0]
}

func (t *heapgreater) IsEmpty() bool {
	if t.HeapSize == 0 {
		return true
	} else {
		return false
	}
}

func (t *heapgreater) Contains(value interface{}) bool {
	if _, ok := t.IndexMap[value]; ok {
		return true
	} else {
		return false
	}
}

func (t *heapgreater) Swap(index1, index2 int) {
	t.swapIndexMap(t.Value[index1], t.Value[index2], index2, index1)
	t.Value[index1], t.Value[index2] = t.Value[index2], t.Value[index1]
}
