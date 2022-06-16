package main

import (
	"fmt"
	"math/rand"
	"time"
)

func lenRandValueRand(maxLen int, maxValue int) []int {
	rand.Seed(time.Now().UnixNano())
	arrayLen := rand.Intn(maxLen)
	array := make([]int, arrayLen)
	for i := 0; i < arrayLen; i++ {
		array[i] = rand.Intn(maxValue)
	}
	return array
}
func swap(a *[]int, m int, n int) {
	temp := (*a)[m]
	(*a)[m] = (*a)[n]
	(*a)[n] = temp
}

func insertSort(a *[]int) {
	if a == nil {
		return
	}
	if len(*a) < 2 {
		return
	}
	//0~0
	//0~1
	N := len(*a)
	for end := 1; end < N; end++ {
		//0~0 1
		//0~1 2
		for newIndex := end; newIndex > 0; newIndex-- {
			if (*a)[newIndex] < (*a)[newIndex-1] {
				swap(a, newIndex, newIndex-1)
			} else {
				break
			}
		}
	}
}

func maxInt(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func isSorted(a []int) bool {
	if a == nil {
		return true
	}
	if len(a) < 2 {
		return true
	}
	max := a[0]
	for i := 1; i < len(a); i++ {
		if max > a[i] {
			return false
		} else {
			max = maxInt(max, a[i])
		}
	}
	return true
}

func main() {
	maxLen := 40
	maxValue := 100
	tryTimes := 100000
	for i := 0; i < tryTimes; i++ {
		arr := lenRandValueRand(maxLen, maxValue)
		insertSort(&arr)
		if !isSorted(arr) {
			fmt.Println("insertSort is wrong!!!!")
			fmt.Println(arr)
		}
	}
}
