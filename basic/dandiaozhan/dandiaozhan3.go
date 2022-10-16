package main

import (
	"fmt"
	"zuo/lib/stack"
)

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func maxSqare(arr []int) int {
	if arr == nil {
		return 0
	}
	N := len(arr)
	stack := stack.StackNew()
	ans := 0
	for i := 0; i < N; i++ {
		for !stack.IsEmpty() && arr[stack.Peek().(int)] >= arr[i] {
			popIndex := stack.Pop().(int)
			leftLessIndex := -1
			if !stack.IsEmpty() {
				leftLessIndex = stack.Peek().(int)
			}
			ans = max(ans, arr[popIndex]*((i-1)-(leftLessIndex+1)+1))
		}
		stack.Push(i)
	}
	for !stack.IsEmpty() {
		popIndex := stack.Pop().(int)
		leftLessIndex := -1
		if !stack.IsEmpty() {
			leftLessIndex = stack.Peek().(int)
		}
		ans = max(ans, arr[popIndex]*((N-1)-(leftLessIndex+1)+1))
	}
	return ans
}

func main() {
	arr := []int{3, 2, 4, 2, 5}
	fmt.Println(maxSqare(arr))
}
