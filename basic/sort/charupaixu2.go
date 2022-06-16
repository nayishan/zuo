package main

import "fmt"

func swap(a []int, m int, n int) {
	temp := a[m]
	a[m] = a[n]
	a[n] = temp
}
func charu2(a []int) {
	if a == nil || len(a) == 1 {
		return
	}
	//0~1
	//0~2
	//0~3
	//0~n-1
	N := len(a)
	for end := 0; end < N; end++ {
		//0~0  1
		//0~1  2
		//0~n-1 n-1
		for pre := end - 1; pre >= 0 && a[pre] > a[pre+1]; pre-- {
			swap(a, pre, pre+1)
		}
	}
	fmt.Println(a)
}

func main() {
	a := []int{3, 6, 5, 4}
	charu2(a)
	b := []int{1, 3, 4, 2, 9, 6, 5, 7}
	charu2(b)
}
