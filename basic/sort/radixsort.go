package radixsort

import (
	"math"
)

func arraySort(a []int) []int {
	if a == nil {
		return nil
	}
	if len(a) < 2 {
		return a
	}
	max := 1
	for i := 0; i < len(a); i++ {
		temp := maxBits(a[i])
		if temp > max {
			max = temp
		}
	}
	radixsort(&a, 0, len(a)-1, max)
	return a
}

func maxBits(a int) int {
	temp := 1
	for a/10 != 0 {
		a = a / 10
		temp++
	}
	return temp
}

func getBit(num int, bit int) int {
	temp := num / (int(math.Pow10(bit - 1)))
	temp = temp % 10
	return temp
}

func radixsort(a *[]int, l, r int, maxbits int) {
	helper := make([]int, r-l+1)

	for i := 1; i <= maxbits; i++ {
		index := 0
		bucket := make([][]int, 10)
		for n := 0; n < 10; n++ {
			bucket[n] = make([]int, 0)
		}
		for j := l; j <= r; j++ {
			temp := getBit((*a)[j], i)
			bucket[temp] = append(bucket[temp], (*a)[j])
		}
		for m := 0; m < 10; m++ {
			if len(bucket[m]) != 0 {
				copy(helper[index:], bucket[m])
				index += len(bucket[m])
			}
		}
		copy((*a)[l:], helper)
	}
}

func arraySort2(a []int) []int {
	if a == nil {
		return nil
	}
	if len(a) < 2 {
		return a
	}
	max := 1
	for i := 0; i < len(a); i++ {
		temp := maxBits(a[i])
		if temp > max {
			max = temp
		}
	}
	radixsort2(&a, 0, len(a)-1, max)
	return a
}

func radixsort2(a *[]int, l, r int, maxbits int) {
	presum := make([]int, 10)
	tempA := make([]int, r-l+1)

	for i := 1; i <= maxbits; i++ {
		helper := make([]int, 10)

		for j := l; j <= r; j++ {
			temp := getBit((*a)[j], i)
			helper[temp]++
		}

		//presum
		presum[0] = helper[0]
		for k := 1; k < 10; k++ {
			presum[k] = presum[k-1] + helper[k]
		}

		//for stability from right to left traverse
		for index := len(*a) - 1; index >= 0; index-- {
			temp := getBit((*a)[index], i)
			tempA[presum[temp]-1] = (*a)[index]
			presum[temp]--
		}
		copy((*a)[l:], tempA)
	}
}
