package main

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func right(arr []int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	if len(arr)&1 == 0 {
		return process(arr, 0, (len(arr))/2, sum>>1)

	} else {
		return max(process(arr, 0, (len(arr)-1)/2, sum>>1), process(arr, 0, (len(arr)-1)/2+1, sum>>1))

	}

}

// index 是到了哪，picks 是取了多少
// -1返回无效解
func process(arr []int, index int, picks int, rest int) int {
	if index == len(arr) {
		if picks == 0 {
			return 0
		} else {
			return -1
		}
	} else {
		p1 := process(arr, index+1, picks-1, rest)
		next := 0
		if rest-arr[index] >= 0 {
			next = process(arr, index+1, picks-1, rest-arr[index])
		}
		p2 := 0
		if next != -1 {
			p2 = next + arr[index]
		}
		return max(p1, p2)

	}

}

func dp(arr []int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	sum = sum / 2
	N := len(arr)
	M := (len(arr) + 1) / 2
	dp := make([][][]int, N+1)
	for i := 0; i <= N; i++ {
		dp[i] = make([][]int, M+1)
		for j := 0; j <= M; j++ {
			dp[i][j] = make([]int, sum+1)

			for k := 0; k <= sum+1; k++ {
				dp[i][j][k] = -1
			}
		}
	}

	for rest := 0; rest <= sum; rest++ {
		dp[N][0][rest] = 0
	}

	for index := N - 1; index >= 0; index-- {
		for picks := 0; picks <= M; picks++ {
			for rest := 0; rest <= M; rest++ {
				p1 := dp[index+1][picks-1][rest]
				next := 0
				if rest-arr[index] >= 0 && picks >= 1 {
					next = dp[index+1][picks-1][rest-arr[index]]
				}
				p2 := 0
				if next != -1 {
					p2 = next + arr[index]
				}
				dp[index][picks][rest] = max(p1, p2)
			}
		}
	}
	if N&1 == 0 {
		return dp[0][N/2][sum]
	} else {
		return max(dp[0][N/2][sum], dp[0][N/2+1][sum])
	}

}
