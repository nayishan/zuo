package main

import "fmt"
import "math/rand"
import "time"

func find(a []int, num int) int {
	if a == nil {
		return -1
	}
	if len(a) == 0 {
		return -1
	}
	L := 0
	R := len(a) - 1
	mid := 0
	index := -1
	for L <= R {
		mid = (L + R) / 2
		if a[mid] >= num {
			if a[mid] == num {
				index = mid
			}
			R = mid - 1
		} else {
			L = mid + 1
		}
	}
	return index

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
func lenRandValueRand(maxLen int, maxValue int) []int {
	rand.Seed(time.Now().UnixNano())
	arrayLen := rand.Intn(maxLen)
	array := make([]int, arrayLen)
	for i := 0; i < arrayLen; i++ {
		array[i] = rand.Intn(maxValue)
	}
	return array
}
func vFind(a []int, num int) int {
	if a == nil {
		return -1
	}
	if len(a) == 0 {
		return -1
	}
	N := len(a)
	for i := 0; i < N; i++ {
		if a[i] == num {
			return i
		}
	}
	return -1
}

func main() {
	maxLen := 40
	maxValue := 100
	tryTimes := 100000
	for i := 0; i < tryTimes; i++ {
		arr := lenRandValueRand(maxLen, maxValue)
		insertSort(&arr)
		flag1 := find(arr, 20)
		flag2 := vFind(arr, 20)
		if flag1 != flag2 {
			fmt.Println("LeftestIndex is wrong!!!!")
			fmt.Println(arr)
			fmt.Println("flag1", flag1, "flag2", flag2)
			break
		}
	}
}
