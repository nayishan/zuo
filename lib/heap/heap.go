package heap

type heap struct {
	HeapSize int //maxIndex +1
	Data     []int
}

func (t *heap) HeapInit(cap int) {
	t.Data = make([]int, cap)
	t.HeapSize = 0

}

func (t *heap) heapinsert(value int) {
	b := t.Data
	if t.HeapSize == cap(b) {
		//扩容
		b = append(b, value)
	} else {
		//0 ~
		b[t.HeapSize] = value
	}
	t.HeapSize += 1
	index := t.HeapSize - 1
	//index  =0 also in this base case
	for b[index] > b[(index-1)/2] {
		b[index], b[(index-1)>>1] = b[(index-1)>>1], b[index]
		index = (index - 1) >> 1
	}
}

func max(num1, index1, num2, index2 int) (int, int) {
	if num1 >= num2 {
		return num1, index1
	} else {
		return num2, index2
	}
}

func leftChildIndex(index int) int {
	return 2*index + 1
}
func rightChildIndex(index int) int {
	return 2*index + 2
}
func (t *heap) heapify(index int) {
	b := t.Data

	for leftChildIndex(index)+1 < t.HeapSize {
		leftChild := b[leftChildIndex(index)]
		rightChild := b[index]

		if rightChildIndex(index)+1 <= t.HeapSize {
			rightChild = b[rightChildIndex(index)]
		}
		bigChild, childIndex := max(leftChild, leftChildIndex(index), rightChild, rightChildIndex(index))
		if b[index] < bigChild {
			b[index], b[childIndex] = b[childIndex], b[index]
			index = childIndex
		} else {
			break
		}
	}
}

func (t *heap) Add(value int) {
	t.heapinsert(value)
}

func (t *heap) Poll() int {
	temp := t.Data[0]
	t.Data[0], t.Data[t.HeapSize-1] = t.Data[t.HeapSize-1], t.Data[0]
	t.HeapSize -= 1
	t.heapify(0)
	return temp
}

func (t *heap) Peek() int {
	return t.Data[0]
}

func (t *heap) IsEmpty() bool {
	if t.HeapSize == 0 {
		return true
	} else {
		return false
	}
}
