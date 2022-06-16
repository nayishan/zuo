package main

import (
	"math/rand"
	"time"
)

func quickSort(a []int) []int {
	if a == nil {
		return a
	}
	if len(a) < 2 {
		return a
	}

	process(&a, 0, len(a)-1)
	return a
}

func process(a *[]int, l int, r int) {
	if r <= l {
		return
	}

	l1, r1 := netherland(a, l, r)
	process(a, l, l1-1)
	process(a, r1+1, r)
}

func randIndex(l int, r int) int {
	rand.Seed(time.Now().UnixNano())
	N := r - l + 1
	index := rand.Intn(N) + l
	return index
}

func netherland(a *[]int, l int, r int) (int, int) {
	b := *a
	selected := randIndex(l, r)
	b[selected], b[r] = b[r], b[selected]

	index := l
	left, right := l-1, r
	for index < right {
		if b[index] < b[r] {
			left++
			b[index], b[left] = b[left], b[index]
			index++
		} else if b[index] > b[r] {
			right--
			b[index], b[right] = b[right], b[index]
		} else {
			index++
		}
	}
	b[right], b[r] = b[r], b[right]
	return left + 1, right
}
