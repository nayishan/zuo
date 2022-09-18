package main

import (
	"math"
)

func min(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

func minCoins(arr []int, aim int) int {
	return process(arr, 0, aim)
}
func process(arr []int, index int, rest int) int {
	if rest < 0 {
		return math.MaxInt
	}
	if index == len(arr) {
		if rest == 0 {
			return 0
		} else {
			return math.MaxInt
		}
	} else {
		p1 := process(arr, index+1, rest)
		p2 := process(arr, index+1, rest-arr[index])
		if p2 != math.MaxInt {
			p2++
		}
		return min(p1, p2)
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
		for j := 0; j < aim+1; j++ {
			dp[i][j] = math.MaxInt
		}
	}
	dp[N][0] = 0
	for index := N - 1; index >= 0; index++ {
		for rest := 0; rest <= aim; rest++ {
			p1 := dp[index+1][rest]
			p2 := math.MaxInt
			if rest-arr[index] > 0 {
				p2 = dp[index+1][rest-arr[index]]
			}
			if p2 != math.MaxInt {
				p2++
			}
			dp[index][rest] = min(p1, p2)
		}
	}
	return dp[0][aim]
}

//收集张数去重，将原数组改成货币值，货币张数两个数组，多叉树遍历
type Info struct {
	coins  []int
	zhangs []int
}

func getInfo(arr []int) Info {
	counts := make(map[int]int)
	for value := range arr {
		counts[value] += 1
	}
	N := len(counts)
	coins := make([]int, N)
	zhangs := make([]int, N)
	index := 0
	for k, v := range counts {
		coins[index] = k
		zhangs[index] = v
	}

	return Info{coins, zhangs}
}

func dp2(arr []int, aim int) int {
	if aim == 0 {
		return 0
	}
	info := getInfo(arr)
	coins := info.coins
	zhangs := info.zhangs
	N := len(coins)
	dp := make([][]int, N+1)
	for i := 0; i < N+1; i++ {
		dp[i] = make([]int, aim+1)
		for j := 0; j < aim+1; j++ {
			dp[i][j] = math.MaxInt
		}
	}
	dp[N][0] = 0
	for index := N - 1; index >= 0; index-- {
		for rest := 0; rest <= aim; rest++ {
			dp[index][rest] = dp[index+1][rest]
			for zhang := 1; zhang*coins[index] <= aim && zhang <= zhangs[index]; zhang++ {
				if rest-zhang*coins[index] >= 0 && dp[index+1][rest-zhang*coins[index]] != math.MaxInt {
					dp[index][rest] = min(dp[index][rest], zhang+dp[index+1][rest-zhang*coins[index]])
				}
			}
		}
	}
	return dp[0][aim]
}

func dp3() {
	//难暂时不做了
}
