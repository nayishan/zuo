package dp1

func uniquePaths(row int, col int) int {
	dp := make([][]int, row)
	for n := 0; n < row; n++ {
		dp[n] = make([]int, col)
	}
	dp[0][0] = 1
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if i-1 >= 0 {
				dp[i][j] = dp[i-1][j] + 1
			}
			if j-1 >= 0 {
				dp[i][j] = dp[i][j-1] + 1
			}
		}
	}
	return dp[row-1][col-1]
}
