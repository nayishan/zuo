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
	// fmt.Println(b)
	heap.Init(&b)
	ansTemp := make([]int, 0)
	for b.Len() != 0 {
		val1 := heap.Pop(&b)
		if b.Len() != 0 {
			val2 := heap.Pop(&b)
			ansTemp = append(ansTemp, val1.(int)+val2.(int))
			heap.Push(&b, val1.(int)+val2.(int))
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
	processLessMoney(a, &pre)
	// for
	return pre
}
func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
func processLessMoney(a []int, pre *int) int {
	if a == nil {
		return *pre
	}
	if len(a) < 2 {
		return *pre
	}
	ans := math.MaxInt
	for i := 0; i < len(a)-1; i++ {
		for j := i + 1; j < len(a); j++ {
			fmt.Println("=====", i, j)
			nexts := removeAndAdd(a, i, j)
			temp := *pre + a[i] + a[j]
			ans = min(ans, processLessMoney(nexts, &temp))
		}
	}
	return ans
}
func removeAndAdd(a []int, i, j int) []int {
	N := len(a)
	ans := make([]int, N-2)
	findI := false
	findJ := false
	for k := 0; k < len(ans); k++ {
		if k == i {
			findI = true
		} else if k == j {
			findJ = true
		}
		if findJ && findI {
			if k+2 < len(ans) {
				ans[k] = a[k+2]
			}
		} else if findI || findJ {
			if k+1 < len(ans) {
				ans[k] = a[k+1]
			}
		} else {
			ans[k] = a[k]
		}
	}
	ans = append(ans, a[i]+a[j])
	return ans

}

func main() {
	// for i := 0; i < 1000; i++ {
	a := duishuqi.LenRandValueRand(20, 200)
	val1 := lessMoney1(a)
	val2 := lessMoney2(a)
	if val1 != val2 {
		fmt.Println("Oops!", "val1", val1, "val2", val2)
	}
	// }
	fmt.Println("finish!")
}
