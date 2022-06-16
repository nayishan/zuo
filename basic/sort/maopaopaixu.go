package main

import "fmt"

func swap(a []int, m int, n int) {
	temp := a[m]
	a[m] = a[n]
	a[n] = temp
}
func maopao(a []int) {
	if a == nil || len(a) == 1 {
		return
	}
	// 0~ n-1
	// 0 ~n-2
	//0 ~n-3
	N := len(a)
	for end := N - 1; end >= 0; end-- {
		//0 1, 1 2, 2 3, 4 5, ....,N-2 N-1
		for second := 1; second <= end; second++ {
			if a[second-1] > a[second] {
				swap(a, second, second-1)
			}
		}
	}
	fmt.Println(a)
}
func main() {
	a := []int{3, 6, 5, 4}
	maopao(a)
	b := []int{1, 3, 4, 2, 9, 6, 5, 7}
	maopao(b)
}
