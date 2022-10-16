package main

import "zuo/lib/stack"

func sumSubArrayMins(arr []int) int {
	left := nearLessEqualLeft(arr)
	right := nearLessRight(arr)
	ans := 0
	N := len(arr)
	for i := 0; i < N; i++ {
		start := i - left[i]
		end := right[i] - i
		ans += start * end * arr[i]
	}
	return ans
}

func nearLessEqualLeft(arr []int) []int {
	N := len(arr)
	left := make([]int, N)
	stack := stack.StackNew()
	//倒序入栈，目的是右面的相同的值的index可以止步于左边相同值的index
	for i := N - 1; i >= 0; i-- {
		//等于的判定为了止步于。
		for !stack.IsEmpty() && arr[stack.Peek().(int)] >= arr[i] {
			left[stack.Pop().(int)] = i
		}
		stack.Push(i)
	}
	for !stack.IsEmpty() {
		left[stack.Pop().(int)] = -1
	}
	return left
}

func nearLessRight(arr []int) []int {
	N := len(arr)
	right := make([]int, N)
	stack := stack.StackNew()
	for i := 0; i < N; i++ {
		//相同的值的index拥有共同的结尾。
		for !stack.IsEmpty() && arr[stack.Peek().(int)] > arr[i] {
			right[stack.Pop().(int)] = i
		}
	}
	for !stack.IsEmpty() {
		right[stack.Pop().(int)] = -1
	}
	return right

}
