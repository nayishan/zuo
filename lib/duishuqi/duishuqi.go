package duishuqi

import (
	"fmt"
	"math/rand"
	"time"
)

func LenRandValueRand(maxLen int, maxValue int) []int {
	rand.Seed(time.Now().UnixNano())
	arrayLen := rand.Intn(maxLen)
	array := make([]int, arrayLen)
	for i := 0; i < arrayLen; i++ {
		array[i] = rand.Intn(maxValue)
	}
	return array
}

func LenRandValueRand2(maxLen int, maxValue int, minValue int) []int {
	rand.Seed(time.Now().UnixNano())
	arrayLen := rand.Intn(maxLen)
	array := make([]int, arrayLen)
	for i := 0; i < arrayLen; i++ {
		randNum := rand.Intn(maxValue - minValue)
		array[i] = randNum + minValue
	}
	return array
}

func swap(a *[]int, m int, n int) {
	temp := (*a)[m]
	(*a)[m] = (*a)[n]
	(*a)[n] = temp
}

func InsertSort(a *[]int) {
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

func WuXuBuDeng(maxLen int, maxValue int) []int {
	rand.Seed(time.Now().UnixNano())
	arrayLen := rand.Intn(maxLen) + 1
	array := make([]int, arrayLen)
	for i := 0; i < arrayLen; i++ {
		for {
			array[i] = rand.Intn(maxValue)
			if i == 0 {
				break
			} else {
				if array[i] != array[i-1] {
					break
				}
			}
		}
	}
	return array
}
func init() {
	fmt.Println("duishuqi init")
}
