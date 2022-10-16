package main

import (
	"fmt"
	"zuo/lib/stack"
)

func nums(L int) int {
	return ((L + 1) * L) >> 1
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func numSubmat(arr [][]int) int {
	if arr == nil || len(arr[0]) == 0 {
		return 0
	}
	ans := 0
	N := len(arr)
	M := len(arr[0])
	height := make([]int, M)
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if arr[i][j] == 0 {
				height[j] = 0
			} else {
				height[j] = height[j] + 1
			}
		}
		ans += countFrimBottom(height)
	}
	return ans
}

func countFrimBottom(arr []int) int {
	stack := stack.StackNew()
	N := len(arr)
	ans := 0
	for i := 0; i < len(arr); i++ {
		for !stack.IsEmpty() && arr[stack.Peek().(int)] >= arr[i] {
			popIndex := stack.Pop().(int)
			if arr[popIndex] > arr[i] {
				leftLessIndex := -1
				if !stack.IsEmpty() {
					leftLessIndex = stack.Peek().(int)
				}
				bigHeight := arr[popIndex]
				smallHeight := arr[i]
				if leftLessIndex != -1 {
					smallHeight = max(smallHeight, arr[leftLessIndex])
				}
				L := ((i - 1) - (leftLessIndex + 1) + 1)
				ans += nums(L) * (bigHeight - smallHeight)
			}
		}
		stack.Push(i)
	}
	for !stack.IsEmpty() {
		popIndex := stack.Pop().(int)
		leftLessIndex := -1
		if !stack.IsEmpty() {
			leftLessIndex = stack.Peek().(int)
		}
		bigHeight := arr[popIndex]
		smallHeight := 0
		if leftLessIndex != -1 {
			smallHeight = max(smallHeight, arr[leftLessIndex])
		}
		L := ((N - 1) - (leftLessIndex + 1) + 1)
		fmt.Println(N, popIndex)
		ans += nums(L) * (bigHeight - smallHeight)

	}
	return ans
}

func main() {
	matrix := [][]int{
		{1, 1},
		{1, 1},
	}
	fmt.Println(numSubmat(matrix))

}
