package main

import "math"

func minCoin(arr []int, aim int) int {
	return process(arr, 0, aim)
}

func min(a, b int) int {
	if a >= b {
		return b
	} else {
		return a
	}
}

func process(arr []int, index int, rest int) int {
	if rest < 0 {
		return math.MaxInt
	}
	if index == len(arr) {
		if rest == 0 {
			return 1
		} else {
			return math.MaxInt
		}
	} else {
		ans := math.MaxInt
		for zhang := 0; zhang*arr[index] <= rest; zhang++ {
			next := process(arr, index+1, rest-zhang*arr[index])
			if next != math.MaxInt {
				ans = min(ans, zhang+next)
			}
		}
		return ans
	}
}

func dp1(arr []int, aim int) int {
	if aim == 0 {
		return 0
	}
	N := len(arr)
	dp := make([][]int, N+1)
	for i := 0; i < N+1; i++ {
		dp[i] = make([]int, aim+1)
	}
	dp[N][0] = 0
	for j := 1; j <= aim; j++ {
		dp[N][j] = math.MaxInt
	}
	for index := N - 1; index >= 0; index-- {
		for rest := 0; rest <= aim; rest++ {
			ans := math.MaxInt
			for zhang := 0; zhang*arr[index] <= rest; zhang++ {
				next := process(arr, index+1, rest-zhang*arr[index])
				if next != math.MaxInt {
					ans = min(ans, zhang+next)
				}
			}
			dp[index][rest] = ans
		}
	}
	return dp[0][aim]
}
func dp2(arr []int, aim int) int {
	if aim == 0 {
		return 0
	}
	N := len(arr)
	dp := make([][]int, N+1)
	for i := 0; i < N+1; i++ {
		dp[i] = make([]int, aim+1)
	}
	dp[N][0] = 0
	for j := 1; j <= aim; j++ {
		dp[N][j] = math.MaxInt
	}
	for index := N - 1; index >= 0; index-- {
		for rest := 0; rest <= aim; rest++ {
			dp[index][rest] = dp[index+1][rest]
			//两个都是为了放置越界
			if rest-arr[index] >= 0 && dp[index][rest-arr[index]] != math.MaxInt {
				dp[index][rest] = min(dp[index][rest-arr[index]]+1, dp[index+1][rest])
			}
		}
	}
	return dp[0][aim]
}
