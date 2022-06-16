package main

import (
	"container/heap"
	"fmt"
)

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

func quickSort(a [][]int) [][]int {
	if a == nil {
		return a
	}
	if len(a) < 2 {
		return a
	}
	process(&a, 0, len(a)-1)
	return a
}
func process(a *[][]int, l int, r int) {
	if r <= l {
		return
	}
	l1, r1 := netherland(a, l, r)
	process(a, l, l1-1)
	process(a, r1+1, r)
}

func netherland(a *[][]int, l int, r int) (int, int) {
	b := *a
	base := b[r][0]
	less := l - 1
	more := r
	index := l
	for index < more {
		if b[index][0] > base {
			more--
			b[index], b[more] = b[more], b[index]

		} else if b[index][0] < base {
			less++
			b[index], b[less] = b[less], b[index]
			index++

		} else {
			index++
		}
	}
	b[index], b[r] = b[r], b[index]
	return less + 1, more
}

func findMinArrowShots(points [][]int) int {
	helper := quickSort(points)
	// fmt.Println("healper:", helper, "len:", len(helper))
	endHeap := IntHeap{}
	heap.Init(&endHeap)
	shots := 0
	// index := 0
	for i := 0; i < len(helper); i++ {
		// index++
		canShots := false
		if endHeap.Len() > 0 {
			a := heap.Pop(&endHeap)
			if a.(int) >= helper[i][0] {
				heap.Push(&endHeap, a)
			} else {
				// fmt.Println("a:", a, "helper:", helper[i][0], "index:", index)
				canShots = true
				for endHeap.Len() > 0 {
					heap.Pop(&endHeap)
				}
			}
		}
		if canShots {
			shots += 1
		}
		heap.Push(&endHeap, helper[i][1])
	}
	shots += 1
	return shots
}

func main() {
	shots := findMinArrowShots([][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}})
	fmt.Println("shots:", shots)
	shots = findMinArrowShots([][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}})
	fmt.Println("shots:", shots)
	points := [][]int{{1, 9}, {7, 16}, {2, 5}, {7, 12}, {9, 11}, {2, 10}, {9, 16}, {3, 9}, {1, 3}}
	shots = findMinArrowShots(points)
	fmt.Println("shots:", shots)
}
