package main

import (
	"fmt"
	"zuo/lib/linklist"
)

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
func min(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}
func right(arr []int, num int) int {
	if arr == nil {
		return 0
	}
	if len(arr) == 0 {
		return 0
	}
	if num < 0 {
		return 0
	}
	ans := 0
	N := len(arr)
	for L := 0; L < N; L++ {
		for R := L; R < N; R++ {
			maxTemp := arr[L]
			minTemp := arr[L]
			for i := L + 1; i <= R; i++ {
				maxTemp = max(maxTemp, arr[i])
				minTemp = min(minTemp, arr[i])
			}
			if maxTemp-minTemp <= num {
				ans++
			}
		}
	}
	return ans
}

//当[L..R] 满足<=num 它的子序列也满足
//当[L..R] 不满足 <=num R往右扩和L往左扩也不满足
func right2(arr []int, num int) int {
	if arr == nil {
		return 0
	}
	if len(arr) == 0 {
		return 0
	}
	if num < 0 {
		return 0
	}
	N := len(arr)
	ans := 0
	R := 0
	qMax := linklist.Acl_fifo_new()
	qMin := linklist.Acl_fifo_new()
	for L := 0; L < N; L++ {
		//[L..R)
		for R < N {
			for (qMax.Acl_size() != 0) && arr[qMax.Acl_tail().(int)] <= arr[R] {
				qMax.Acl_pop_back()
			}
			qMax.Acl_push_back(R)
			for (qMin.Acl_size() != 0) && arr[qMin.Acl_tail().(int)] >= arr[R] {
				qMin.Acl_pop_back()
			}

			qMin.Acl_push_back(R)

			if arr[qMax.Acl_head().(int)]-arr[qMin.Acl_head().(int)] > num {
				break
			} else {
				R++
			}
		}

		ans += R - L
		//此处去掉可能即将到来的L
		//L只可能出现在第一的位置，因为如果它要是小的话，它在上面的逻辑中，已经被pop出去了
		if qMax.Acl_head() == L {
			qMax.Acl_pop_front()
		}
		if qMin.Acl_head() == L {
			qMin.Acl_pop_front()
		}

	}

	return ans

}

func main() {
	arr := []int{4, 7, 5, 6, 10, 2, 8}
	num := 3
	fmt.Println(right(arr, num))
	fmt.Println(right2(arr, num))
}
