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

func maxRectangle(matrix [][]int) int {
	if matrix == nil {
		return 0
	}
	//此处可以用行代替matrix
	N := len(matrix)
	M := len(matrix[0])
	helper := make([][]int, N)
	for i := 0; i < N; i++ {
		helper[i] = make([]int, M)
	}
	ans := 0
	for j := 0; j < M; j++ {
		helper[0][j] = matrix[0][j]
		ans = max(ans, maxSqare(helper[0]))
	}
	for i := 1; i < N; i++ {
		for j := 0; j < M; j++ {
			if matrix[i][j] == 1 {
				helper[i][j] = helper[i-1][j] + matrix[i][j]
			} else {
				helper[i][j] = 0
			}
		}
		ans = max(ans, maxSqare(helper[i]))
	}
	return ans
}

func main() {
	matrix := [][]int{
		{1, 1, 1, 1, 1},
		{1, 0, 1, 1, 1},
		{1, 1, 1, 0, 1},
		{1, 1, 1, 1, 1},
	}
	fmt.Println(maxRectangle(matrix))
}
