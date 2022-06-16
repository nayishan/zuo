package main

func heapinsert(a *[]int, value int, heapsize int) {
	b := *a
	b[heapsize] = value
	heapsize += 1
	index := heapsize - 1
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
func heapify(a *[]int, index int, heapsize int) {
	b := *a

	for leftChildIndex(index) < heapsize {
		leftChild := b[leftChildIndex(index)]
		rightChild := b[index]

		if rightChildIndex(index) <= heapsize {
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

func heapSort(a []int) []int {
	if a == nil {
		return a
	}
	if len(a) < 2 {
		return a
	}
	//helper := make([]int, len(a))
	for i := len(a) - 1; i >= 0; i-- {
		heapify(&a, i, len(a))
	}
	heapSize := len(a)
	for heapSize > 0 {
		heapify(&a, 0, heapSize)
		heapSize--
		a[0], a[heapSize] = a[heapSize], a[0]
	}
	return a
}
