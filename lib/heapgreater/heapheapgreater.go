package main

import (
	"container/heap"
	"fmt"
)

type Custom struct {
	Id        int
	Buy       int
	EnterTime int
}

type zone struct {
	custom   []Custom
	indexMap map[Custom]int
	heapsize int
}

func (z zone) Len() int {
	return z.heapsize
}

func (z zone) Less(i, j int) bool {
	if z.custom[i].Buy != z.custom[j].Buy {
		return z.custom[i].Buy < z.custom[j].Buy
	} else {
		return z.custom[i].EnterTime < z.custom[j].EnterTime
	}
}

func (z zone) Swap(i, j int) {
	z.custom[i], z.custom[j] = z.custom[j], z.custom[i]
	z.indexMap[z.custom[i]] = i
	z.indexMap[z.custom[j]] = j
}

func (z *zone) Push(x interface{}) {
	if len(z.custom) == z.heapsize {
		(*z).custom = append((*z).custom, x.(Custom))
	} else {
		(*z).custom[z.heapsize] = x.(Custom)
	}
	z.indexMap[z.custom[z.heapsize]] = z.heapsize
	z.heapsize++
}

func (z *zone) Pop() interface{} {
	(*z).heapsize--
	x := (*z).custom[z.heapsize]
	delete(z.indexMap, z.custom[z.heapsize])
	return x
}
func (z *zone) contains(c Custom) bool {
	if _, ok := z.indexMap[c]; ok {
		return true
	} else {
		return false
	}
}
func (z *zone) getIndex(c Custom) int {
	return z.indexMap[c]
}
func (z *zone) isEmpty() bool {
	if z.heapsize == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	z := zone{
		custom:   []Custom{},
		indexMap: map[Custom]int{},
		heapsize: 0,
	}
	heap.Init(&z)
	heap.Push(&z, Custom{0, 0, 0})
	heap.Push(&z, Custom{0, 0, 1})
	for !z.isEmpty() {
		fmt.Println(heap.Pop(&z))

	}
}
