package main

import (
	"fmt"
	"zuo/lib/duishuqi"
)

func mergeSort(a []int) int {
	pairs := 0
	if a == nil || len(a) < 2 {
		return pairs
	}
	process(&a, 0, len(a)-1, &pairs)
	return pairs
}

func process(a *[]int, l int, r int, pairs *int) {
	if l == r {
		return
	}
	mid := l + ((r - l) >> 1)
	process(a, l, mid, pairs)
	process(a, mid+1, r, pairs)
	merge(a, l, mid, r, pairs)
}

func merge(a *[]int, l int, mid int, r int, pairs *int) {
	tempL := l
	tempR := mid + 1
	helper := make([]int, r-l+1)
	tempH := 0

	for tempL <= mid && tempR <= r {
		if (*a)[tempL] > (*a)[tempR] {
			helper[tempH] = (*a)[tempL]
			*pairs += (r - tempR + 1)
			tempL++
		} else {
			helper[tempH] = (*a)[tempR]
			tempR++
		}
		tempH++
	}
	if tempR > r {
		copy(helper[tempH:], (*a)[tempL:])
	}
	if tempL > mid {
		copy(helper[tempH:], (*a)[tempR:])
	}
	copy((*a)[l:], helper)
}

func reverse(a []int) int {
	sum := 0
	for begin := 0; begin < len(a); begin++ {
		for end := begin + 1; end < len(a); end++ {
			if a[begin] > a[end] {
				sum += 1
			}
		}
	}
	return sum
}

func main() {
	for i := 0; i < 10000; i++ {
		a := duishuqi.LenRandValueRand(20, 100)
		sum2 := reverse(a)
		sum1 := mergeSort(a)
		if sum1 != sum2 {
			fmt.Println("wrong", "mergeSort", sum1, "reverse", sum2)
			fmt.Println(a)
			break
		}

	}
}
