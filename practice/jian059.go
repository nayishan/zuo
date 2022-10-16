package main

import (
	"fmt"
	"zuo/lib/linklist"
)

func maxSlidingWindow(nums []int, k int) []int {
	if nums == nil || k >= len(nums) {
		return nil
	}
	N := len(nums)
	R := 0

	qmax := linklist.Acl_fifo_new()
	ans := make([]int, N-k+1)
	for L := 0; L < N; L++ {
		for R <= L+k-1 && R < N {

			for qmax.Acl_size() != 0 && nums[qmax.Acl_tail().(int)] <= nums[R] {
				temp := qmax.Acl_pop_back()
				fmt.Println("pop tail ", nums[temp.(int)])
			}

			fmt.Println("push", R)
			qmax.Acl_push_back(R)

			if R == L+k-1 {
				ans[L] = nums[qmax.Acl_head().(int)]
				fmt.Println("add ans R", R, "ans[L]", ans[L])
			}
			R++
		}

		if qmax.Acl_head().(int) == L {
			fmt.Println("pop head", L)
			qmax.Acl_pop_front()
		}
	}

	return ans

}
func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	fmt.Println(maxSlidingWindow(nums, k))
}
