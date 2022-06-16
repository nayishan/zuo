package main

import "fmt"

func xuanze(a []int) {
	if a == nil || len(a) == 1 {
		return
	}
	//0 ~ n-1
	//1 ~n-1
	//2~n-1
	N := len(a)
	for i := 0; i < N; i++ {
		// 0~n-1 min
		// 1~n-1 min
		minIndex := i
		for j := i + 1; j < N; j++ {
			if a[minIndex] > a[j] {
				minIndex = j
			}
		}
		swap(a, i, minIndex)
	}
	fmt.Println(a)
}
func swap(a []int, m int, n int) {
	temp := a[m]
	a[m] = a[n]
	a[n] = temp
}
func main() {
	a := []int{3, 6, 5, 4}
	xuanze(a)
	b := []int{1, 3, 4, 2, 9, 6, 5, 7}
	xuanze(b)
}
