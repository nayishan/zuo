package main

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
func right(arr []int) int {
	num := 0

	for i := range arr {
		num += arr[i]
	}

	return process(arr, 0, num>>1)
}

//在[index...]范围上返回接近rest的不超过rest的值的值
func process(arr []int, index int, rest int) int {
	if index == len(arr) {
		return 0
	}

	p1 := process(arr, index+1, rest)
	p2 := 0
	if arr[index] < rest {
		p2 = process(arr, index+1, rest-arr[index]) + arr[index]
	}
	return max(p1, p2)
}

func dp(arr []int) int {
	N := len(arr)
	sum := 0
	for i := range arr {
		sum += arr[i]
	}
	sum = sum / 2
	dp := make([][]int, N+1)
	for i := 0; i < N; i++ {
		dp[i] = make([]int, sum+1)
	}

	for i := 0; i <= sum/2; i++ {
		dp[N][i] = 0
	}
	for index := N - 1; index >= 0; index-- {
		for rest := 0; rest <= sum/2; rest++ {
			p1 := dp[index+1][rest]
			p2 := 0
			if arr[index] < rest {
				p2 = dp[index+1][rest-arr[index]] + arr[index]
			}
			dp[index][rest] = max(p1, p2)
		}
	}
	return dp[0][sum/2]
}
