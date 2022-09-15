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
func right(arr []int, w int) []int {
	if arr == nil || w < 1 || len(arr) < w {
		return nil
	}
	res := make([]int, 0)
	L := 0
	R := L + w - 1
	for R < len(arr) {
		temp := arr[L]
		for i := L + 1; i <= R; i++ {
			temp = max(temp, arr[i])
		}
		res = append(res, temp)
		L++
		R++
	}
	return res

}

func getMaxWindow(arr []int, w int) []int {
	if arr == nil || w < 1 || len(arr) < w {
		return nil
	}

	res := make([]int, 0)
	//窗口最大值的更新结构，里面的data是下标
	qMax := linklist.Acl_fifo_new()

	for R := 0; R < len(arr); R++ {
		for qMax.Acl_size() != 0 {
			if arr[qMax.Acl_tail().(int)] <= arr[R] {
				qMax.Acl_pop_back()
			} else {
				break
			}
		}
		qMax.Acl_push_back(R)

		if qMax.Acl_head() == R-w {
			qMax.Acl_pop_front()
		}
		if R >= w-1 {
			res = append(res, arr[qMax.Acl_head().(int)])
		}
	}
	return res

}

func main() {
	arr := []int{4, 3, 5, 4, 3, 3, 6, 7}
	w := 3
	fmt.Println(right(arr, w))
	fmt.Println(getMaxWindow(arr, w))

}
