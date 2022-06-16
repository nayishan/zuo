package t

type MyCircularQueue struct {
	size    int
	begin   int
	end     int
	maxSize int
	data    []int
}

func Constructor(k int) MyCircularQueue {
	var t MyCircularQueue
	t.data = make([]int, k)
	t.size = 0
	t.begin = 0
	t.end = 0
	t.maxSize = k
	return t
}

func (t *MyCircularQueue) next(index int) int {
	if index >= t.maxSize-1 {
		return 0
	} else {
		index++
		return index
	}
}

func (t *MyCircularQueue) pre(index int) int {
	if index <= 0 {
		return t.maxSize - 1
	} else {
		index--
		return index
	}
}

func (t *MyCircularQueue) EnQueue(value int) bool {
	ret := true
	if t.size >= t.maxSize {
		ret = false
		return ret
	}
	t.size++
	t.data[t.end] = value
	t.end = t.next(t.end)
	return ret
}

func (t *MyCircularQueue) DeQueue() bool {
	ret := true
	if t.size == 0 {
		ret = false
		return ret
	}
	t.size--
	t.begin = t.next(t.begin)
	return ret
}

func (t *MyCircularQueue) Front() int {
	if t.size == 0 {
		return -1
	}
	ret := t.data[t.begin]
	return ret
}

func (t *MyCircularQueue) Rear() int {
	if t.size == 0 {
		return -1
	}
	index := t.pre(t.end)
	ret := t.data[index]
	return ret

}

func (t *MyCircularQueue) IsEmpty() bool {
	if t.size == 0 {
		return true
	} else {
		return false
	}

}

func (t *MyCircularQueue) IsFull() bool {
	if t.size == t.maxSize {
		return true
	} else {
		return false
	}
}
