package main

func coinWay(arr []int, aim int) int {
	return process(arr, 0, aim)

}

func process(arr []int, index int, rest int) int {
	if rest < 0 {
		return 0
	}
	if index == len(arr) {
		if rest == 0 {
			return 1
		} else {
			return 0
		}
	} else {
		return process(arr, index+1, rest-arr[index]) + process(arr, index+1, rest)
	}
}

func coinWay2(arr []int, aim int) int {
	if aim == 0 {
		return 1
	}
	N := len(arr)
	dp := make([][]int, N+1)
	for i := 0; i < aim+1; i++ {
		dp[i] = make([]int, aim)
	}
	dp[N][1] = 1
	for index := N - 1; index >= 0; index++ {
		for rest := 0; rest <= aim; rest++ {
			next := 0
			if rest-arr[index] >= 0 {
				next = dp[index+1][rest-arr[index]]
			}
			dp[index][rest] = dp[index+1][rest] + next
		}
	}
	return dp[0][aim]
}
