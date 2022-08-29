package main

import (
	"fmt"
)

func pop(s *[]int) int {
	ans := (*s)[0]
	*s = (*s)[1:]
	return ans
}

func push(s *[]int, elem int) {
	N := len(*s)
	cur := 0
	if N > 0 {
		post := (*s)[0]
		(*s)[0] = elem
		cur = post
		for i := 1; i < N; i++ {
			post = (*s)[i]
			(*s)[i] = cur
			cur = post
		}
	} else {
		cur = elem
	}

	*s = append(*s, cur)
}

// 将栈底返回，其它元素落下
func process(s *[]int) int {
	result := pop(s)
	if len(*s) == 0 {
		return result
	} else {
		last := process(s)
		push(s, result)
		return last
	}
}

func reverse(s *[]int) {
	if len(*s) == 0 {
		return
	}
	result := process(s)
	reverse(s)
	push(s, result)
}

func main() {
	s := []int{1, 2, 3}
	reverse(&s)
	fmt.Println(s)
}
