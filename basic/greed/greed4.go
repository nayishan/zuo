package main

import (
	"container/heap"
	"fmt"
)

//正数数组cost
//正数数组profits
//正数K
//正数M
//cost 表示项目的花费
//profit 表示项目的收益
//K表示能串行的最多做K个项目
//M表示初始资金
//输出 最大的钱数
func findMaximizedCaptial1(K, W int, profit []int, cost []int) int {
	programA := make([]program, len(profit))
	for i := range profit {
		programA[i].profit = profit[i]
		programA[i].cost = cost[i]
	}
	heapA := IntHeap{programA, true}

	heapB := IntHeap{}
	heapB.small = false

	heap.Init(&heapA)
	k := 0
	for k <= K {
		k++
		for !heapA.IsEmpty() && (heapA.Peek().(program).cost <= W) {
			x := heap.Pop(&heapA)
			heap.Push(&heapB, x)
			fmt.Println("===", heapA.Peek().(program).cost, W, heapA, heapB)
		}
		fmt.Println(heapA.Peek().(program).cost, W, heapA, heapB)
		if !heapB.IsEmpty() {
			x := heap.Pop(&heapB)
			W += x.(program).profit
		} else {
			return W
		}
	}
	return W
}

type program struct {
	cost   int
	profit int
}
type IntHeap struct {
	programs []program
	//true是小根堆，falseG是大根堆
	small bool
}

func (h IntHeap) Len() int { return len(h.programs) }
func (h IntHeap) Less(i, j int) bool {
	if !h.small {
		return h.programs[i].cost > h.programs[j].cost
	} else {
		return h.programs[i].profit < h.programs[j].profit
	}
}
func (h IntHeap) Swap(i int, j int) {
	h.programs[i], h.programs[j] = h.programs[j], h.programs[i]
}

func (h *IntHeap) Push(x interface{}) {
	h.programs = append(h.programs, x.(program))
}

func (h *IntHeap) Pop() interface{} {
	old := h
	n := len(old.programs)
	x := old.programs[n-1]
	h.programs = old.programs[:n-1]
	return x
}
func (h *IntHeap) IsEmpty() bool {
	if h.Len() == 0 {
		return true
	} else {
		return false
	}
}
func (h *IntHeap) Peek() interface{} {
	old := h
	x := old.programs[0]
	return x
}
func main() {
	costA := []int{100, 2, 3, 300}
	profitsA := []int{20, 5, 7, 17}
	ans := findMaximizedCaptial1(10, 10, costA, profitsA)
	fmt.Println(ans)

}
