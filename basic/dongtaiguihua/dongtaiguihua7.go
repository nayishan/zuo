package main

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}

}
func longestPalindSubSeq(str string) int {
	if str == "" {
		return 0
	}
	return process([]byte(str), 0, len(str)-1)
}
func process(str []byte, L, R int) int {
	if L == R {
		return 1
	}
	if L+1 == R {
		if str[L] == str[R] {
			return 2
		} else {
			return 1
		}
	}
	p1 := process(str, L+1, R-1)
	p2 := process(str, L, R-1)
	p3 := process(str, L+1, R)
	p4 := 0

	if str[L] == str[R] {
		p4 = process(str, L-1, R) + 2
	}
	return max(max(p1, p4), max(p2, p3))
}

func longestPalindSubSeq2(str string) int {

	N := len(str)
	var dp [][]int
	for i := 0; i < N; i++ {
		row := make([]int, N)
		dp = append(dp, row)
	}
	// if L == R {
	// 	return 1
	// }
	// if L+1 == R {
	// 	if str[L] == str[R] {
	// 		return 2
	// 	} else {
	// 		return 1
	// 	}
	// }
	dp[N-1][N-1] = 1
	for i := 0; i < N-1; i++ {
		dp[i][i] = 1
		if str[i] == str[i+1] {
			dp[i][i+1] = 2
		} else {
			dp[i][i] = 1
		}
	}
	// p1 := process(str, L+1, R-1)
	// p2 := process(str, L, R-1)
	// p3 := process(str, L+1, R)
	// p4 := 0
	//
	// if str[L] == str[R] {
	// 	p4 = process(str, L-1, R) + 2
	// }
	for L := N - 1; L >= 0; L-- {
		for R := L + 2; R < N; R++ {
			p1 := dp[L+1][R-1]
			p2 := dp[L][R-1]
			p3 := dp[L+1][R]
			p4 := 0
			if str[L] == str[R] {
				p4 = dp[L-1][R-1] + 2
			}
			dp[L][R] = max(max(p1, p2), max(p3, p4))
		}
	}

	//return max(max(p1, p4), max(p2, p3))
	return dp[0][N-1]
}

func longestPalindSubSeq3(str string) int {

	N := len(str)
	var dp [][]int
	for i := 0; i < N; i++ {
		row := make([]int, N)
		dp = append(dp, row)
	}
	// if L == R {
	// 	return 1
	// }
	// if L+1 == R {
	// 	if str[L] == str[R] {
	// 		return 2
	// 	} else {
	// 		return 1
	// 	}
	// }
	dp[N-1][N-1] = 1
	for i := 0; i < N-1; i++ {
		dp[i][i] = 1
		if str[i] == str[i+1] {
			dp[i][i+1] = 2
		} else {
			dp[i][i] = 1
		}
	}
	// p1 := process(str, L+1, R-1)
	// p2 := process(str, L, R-1)
	// p3 := process(str, L+1, R)
	// p4 := 0
	//
	// if str[L] == str[R] {
	// 	p4 = process(str, L-1, R) + 2
	// }
	for L := N - 1; L >= 0; L-- {
		for R := L + 2; R < N; R++ {
			p2 := dp[L][R-1]
			p3 := dp[L+1][R]
			dp[L][R] = max(p2, p3)
			p4 := 0
			if str[L] == str[R] {
				p4 = dp[L-1][R-1] + 2
				dp[L][R] = max(dp[L][R], p4)
			}
		}
	}
	//return max(max(p1, p4), max(p2, p3))
	return dp[0][N-1]
}
