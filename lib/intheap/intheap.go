package intheap

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// type IntHeap struct {
// 	heap []int
// 	//true是小根堆，false是大根堆
// 	bool
// }
//
//
// func (h IntHeap) Len() int           { return len(h.heap) }
// func (h IntHeap) Less(i, j int) bool {
// 	if h.bool{
// 		return h.heap[i] < h.heap[j]
// 	}else{
// 		return h.heap[i] > h.heap[j]
// 	}
//
// } // 小根堆  > 大根堆
// func (h IntHeap) Swap(i, j int)      { h.heap[i], h.heap[j] = h.heap[j], h.heap[i] }
//
// func (h *IntHeap) Push(x interface{}) {
// 	h.heap = append(h.heap, x.(int))
// }
//
// func (h *IntHeap) Pop() interface{} {
// 	old := h
// 	n := len(old.heap)
// 	x := old.heap[n-1]
// 	h.heap = old.heap[0 : n-1]
// 	return x
//
// }
