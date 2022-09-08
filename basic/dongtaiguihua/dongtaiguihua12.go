package main

func coinWay(arr []int, aim int) int {
	return process(arr, 0, aim)
}

func process(arr []int, index int, rest int) int {
	if index == len(arr) {
		if rest == 0 {
			return 1
		} else {
			return 0
		}
	}
	ways := 0
	for zhang := 0; zhang*arr[index] <= rest; zhang++ {
		ways += process(arr, index+1, rest-zhang*arr[index])
	}
	return ways
}

func coinWay2(arr []int, aim int) int {
	N := len(arr)

	dp := make([][]int, N+1)
	for i := 0; i < N+1; i++ {
		dp[i] = make([]int, aim+1)
	}

	dp[N][0] = 1

	for index := N - 1; index >= 0; index++ {
		for rest := 0; rest <= aim; rest++ {
			ways := 0
			for zhang := 0; zhang*arr[index] <= rest; zhang++ {
				ways += dp[index+1][rest-(zhang*arr[index])]
			}
			dp[index][rest] = ways

		}
	}
	return dp[0][aim]
}
func coinWay3(arr []int, aim int) int {
	N := len(arr)

	dp := make([][]int, N+1)
	for i := 0; i < N+1; i++ {
		dp[i] = make([]int, aim+1)
	}

	dp[N][0] = 1

	for index := N - 1; index >= 0; index++ {
		for rest := 0; rest <= aim; rest++ {
			next := 0
			if rest-arr[index] >= 0 {
				next = dp[index][rest-arr[index]]
			}
			dp[index][rest] = dp[index+1][rest] + next

		}
	}
	return dp[0][aim]
}
