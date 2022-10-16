package main

import (
	"fmt"
	"zuo/lib/linklist"
	"zuo/lib/stack"
)

func min(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
func maxSubTimesMin(arr []int) int {
	if arr == nil {
		return 0
	}
	N := len(arr)
	ans := 0
	for start := 0; start < N; start++ {
		for end := start; end < N; end++ {
			sum := 0
			minValue := arr[start]
			for index := start; index <= end; index++ {
				minValue = min(minValue, arr[index])
				sum += arr[index]
			}
			ans = max(ans, sum*minValue)
		}
	}
	return ans
}

// 前缀和加动态窗口
func maxSubTImesMin2(arr []int) int {
	if arr == nil {
		return 0
	}
	N := len(arr)
	helper := make([]int, N)
	helper[0] = arr[0]
	for i := 1; i < N; i++ {
		helper[i] = arr[i] + helper[i-1]
	}
	ans := 0
	for start := 0; start < N; start++ {
		qMin := linklist.Acl_fifo_new()
		for end := start; end < N; end++ {
			startSum := 0
			if start-1 < 0 {
				startSum = 0

			} else {
				startSum = helper[start-1]
			}

			subSum := helper[end] - startSum
			for qMin.Acl_size() != 0 && arr[qMin.Acl_tail().(int)] >= arr[end] {
				qMin.Acl_pop_back()
			}
			qMin.Acl_push_back(end)

			ans = max(ans, subSum*arr[qMin.Acl_head().(int)])
		}
	}
	return ans
}

func maxSubTImesMin3(arr []int) int {
	if arr == nil {
		return 0
	}
	N := len(arr)
	helper := make([]int, N)
	helper[0] = arr[0]
	for i := 1; i < N; i++ {
		helper[i] = arr[i] + helper[i-1]
	}
	ans := 0
	stack := stack.StackNew()
	for i := 0; i < N; i++ {
		//等号的处理，弹出，对于目前的这个index这个结果肯定是错的。
		//但是我们能知道相等时，弹出的index和押入的index的左边界是相同的。
		//再继续，arr数组最后一个和该index相同的值的index的左边界和之前的左边界是相同的，
		//最后的这个index的右边界一定是正确的，此时的结果一定是正确的。
		//之前的错误，会被这个正确的结果替换掉。也就是最终的结果是正确。
		for !stack.IsEmpty() && arr[stack.Peek().(int)] >= arr[i] {
			popIndex := stack.Pop().(int)
			if stack.IsEmpty() {
				ans = max(ans, helper[i-1]*arr[popIndex])
			} else {
				ans = max(ans, (helper[i-1]-helper[stack.Peek().(int)])*arr[popIndex])
			}

		}
		stack.Push(i)
	}
	for !stack.IsEmpty() {
		popIndex := stack.Pop().(int)
		if stack.IsEmpty() {
			ans = max(ans, helper[N-1]*arr[popIndex])
		} else {
			ans = max(ans, (helper[N-1]-helper[stack.Peek().(int)])*arr[popIndex])
		}
	}
	return ans
}

func main() {
	arr := []int{3, 4, 3, 4, 2}
	fmt.Println(maxSubTimesMin(arr))
	fmt.Println(maxSubTImesMin3(arr))
	fmt.Println(maxSubTImesMin2(arr))
}
