package main

import (
	"fmt"
	"zuo/lib/linklist"
	"zuo/lib/stack"
)

// 如果arr= [3,1,2,3]
//     index 0 1 2 3
// ans =
//[
//  0:[-1,1]
//  1:[-1,-1]
//  2:[1,-1]
//  3:[2,-1]
//]
func GetNearLessNoRepeat(arr []int) [][2]int {
	N := len(arr)
	res := make([][2]int, N)
	stack := stack.StackNew()
	for i := 0; i < N; i++ {
		for !stack.IsEmpty() && arr[stack.Peek().(int)] > arr[i] {
			popIndex := stack.Pop().(int)
			leftLessIndex := -1
			if !stack.IsEmpty() {
				leftLessIndex = stack.Peek().(int)
			}
			res[popIndex][0] = leftLessIndex
			res[popIndex][1] = i
		}
		stack.Push(i)
	}
	for !stack.IsEmpty() {
		popIndex := stack.Pop().(int)
		leftLessIndex := -1
		if !stack.IsEmpty() {
			leftLessIndex = stack.Peek().(int)
		}
		res[popIndex][0] = leftLessIndex
		res[popIndex][1] = -1
	}
	return res
}

func GetNearLess(arr []int) [][2]int {
	N := len(arr)
	res := make([][2]int, N)
	stack := stack.StackNew()
	for i := 0; i < N; i++ {

		for !stack.IsEmpty() && arr[stack.Peek().(*(linklist.Acl_fifo)).Acl_tail().(int)] > arr[i] {
			popList := stack.Pop().(*(linklist.Acl_fifo))
			leftLessIndex := -1
			if !stack.IsEmpty() {
				leftLessIndex = stack.Peek().(*(linklist.Acl_fifo)).Acl_tail().(int)
			}

			for node := 0; node < popList.Acl_size(); node++ {
				res[popList.Acl_index(node).(int)][0] = leftLessIndex
				res[popList.Acl_index(node).(int)][1] = i
			}
		}

		if !stack.IsEmpty() && arr[stack.Peek().(*(linklist.Acl_fifo)).Acl_tail().(int)] == arr[i] {
			stack.Peek().(*(linklist.Acl_fifo)).Acl_push_back(i)
		} else {
			list := linklist.Acl_fifo_new()
			list.Acl_push_back(i)
			stack.Push(list)
		}

	}

	for !stack.IsEmpty() {

		popList := stack.Pop().(*(linklist.Acl_fifo))
		leftLessIndex := -1
		if !stack.IsEmpty() {
			leftLessIndex = stack.Peek().(*(linklist.Acl_fifo)).Acl_tail().(int)
		}

		for i := 0; i < popList.Acl_size(); i++ {
			res[popList.Acl_index(i).(int)][0] = leftLessIndex
			res[popList.Acl_index(i).(int)][1] = -1
		}
	}

	return res
}

func main() {
	arr := []int{1, 2, 3, 4}
	fmt.Println(GetNearLessNoRepeat(arr))
	fmt.Println(GetNearLess(arr))
	arr2 := []int{1, 1, 2, 3, 4}
	fmt.Println(GetNearLess(arr2))
}
