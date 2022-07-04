package main

import (
	"container/heap"
	"fmt"
	"math"
	"zuo/lib/duishuqi"
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
func lessMoney1(a []int) int {
	if a == nil {
		return 0
	}
	if len(a) == 0 {
		return 0
	}
	b := IntHeap(a)
	heap.Init(&b)
	ansTemp := make([]int, 0)
	for b.Len() != 0 {
		val1 := heap.Pop(&b)
		if b.Len() != 0 {
			val2 := heap.Pop(&b)
			ansTemp = append(ansTemp, val1.(int)+val2.(int))
			heap.Push(&b, val1.(int)+val2.(int))
			// fmt.Println("lessMoney1", b)
		}
	}
	// fmt.Println(ansTemp)
	ans := 0
	for i := range ansTemp {
		ans += ansTemp[i]
	}
	return ans
}

func lessMoney2(a []int) int {
	if a == nil {
		return 0
	}
	if len(a) == 0 {
		return 0
	}
	pre := 0
	ans := processLessMoney(a, pre)
	// for
	return ans
}
func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func processLessMoney(a []int, pre int) int {
	if len(a) == 1 {
		// fmt.Println("len(a)==1", pre)
		return pre
	}
	ans := math.MaxInt
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			nexts := removeAndAdd(a, i, j)
			temp := pre + a[i] + a[j]
			next := processLessMoney(nexts, temp)
			ans = min(ans, next)
		}
	}
	return ans
}
func removeAndAdd(a []int, i, j int) []int {
	N := len(a)
	ans := make([]int, N-2)
	index := 0
	for k := 0; k < len(ans); k++ {
		if index == i {
			index++
		}
		if index == j {
			index++
		}
		ans[k] = a[index]
		index++
	}
	ans = append(ans, a[i]+a[j])
	// fmt.Println("=====", i, j, a, ans)
	return ans

}

func main() {
	for i := 0; i < 1000; i++ {
		a := duishuqi.LenRandValueRand(7, 20)
		b := make([]int, len(a))
		copy(b, a)
		val1 := lessMoney1(a)
		val2 := lessMoney2(b)
		if val1 != val2 {
			fmt.Println("Oops!", "val1", val1, "val2", val2)
		}
	}
	fmt.Println("finish!")
}
