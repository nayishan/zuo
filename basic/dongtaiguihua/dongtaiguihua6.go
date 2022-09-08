package main

import "fmt"

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
func longestsubstring(s1, s2 string) int {
	if len(s1) == 0 || len(s2) == 0 {
		return 0
	}
	return process1([]byte(s1), []byte(s2), len(s1)-1, len(s2)-1)
}

//返回str1[0..i] str2[0..j]最长公共子序列多长
func process1(str1, str2 []byte, i, j int) int {
	if i == 0 && j == 0 {
		if str1[0] == str2[0] {
			return 1
		} else {
			return 0
		}
	} else if i == 0 {
		if str1[0] == str2[j] {
			return 1
		} else {
			return process1(str1, str2, 0, j-1)
		}
	} else if j == 0 {
		if str1[i] == str2[0] {
			return 1
		} else {
			return process1(str1, str2, i-1, 0)
		}
	} else {
		//i没有
		p1 := process1(str1, str2, i-1, j)
		//j没有
		p2 := process1(str1, str2, i, j-1)
		//i，j都有
		p3 := 0
		if str1[i] == str2[j] {
			p3 = process1(str1, str2, i-1, j-1) + 1
		} else {
			p3 = 0
		}
		return max(p1, max(p2, p3))
	}

}
func longestsubstring2(s1, s2 string) int {
	if len(s1) == 0 || len(s2) == 0 {
		return 0
	}
	N := len(s1)
	M := len(s2)
	var dp [][]int
	for i := 0; i < N; i++ {
		row := make([]int, M)
		dp = append(dp, row)
	}
	// if i == 0 && j == 0 {
	// 	if str1[0] == str2[0] {
	// 		return 1
	// 	} else {
	// 		return 0
	// 	}
	// }
	str1 := []byte(s1)
	str2 := []byte(s2)
	if str1[0] == str2[0] {
		dp[0][0] = 1
	} else {
		dp[0][0] = 0
	}
	// else if i == 0 {
	// 		if str1[0] == str2[j] {
	// 			return 1
	// 		} else {
	// 			return process1(str1, str2, 0, j-1)
	// 		}
	// 	}
	for j := 1; j < M; j++ {
		if str1[0] == str2[j] {
			dp[0][j] = 1
		} else {
			dp[0][j] = dp[0][j-1]
		}
	}
	// else if j == 0 {
	// 	if str1[i] == str2[0] {
	// 		return 1
	// 	} else {
	// 		return process1(str1, str2, i-1, 0)
	// 	}
	// }
	for i := 1; i < N; i++ {
		if str2[0] == str1[i] {
			dp[i][0] = 1
		} else {
			dp[i][0] = dp[i-1][0]
		}
	}
	// else {
	// 		//i没有
	// 		p1 := process1(str1, str2, i-1, j)
	// 		//j没有
	// 		p2 := process1(str1, str2, i, j-1)
	// 		//i，j都有
	// 		p3 := 0
	// 		if str1[i] == str2[j] {
	// 			p3 = process1(str1, str2, i-1, j-1) + 1
	// 		} else {
	// 			p3 = 0
	// 		}
	// 		return max(p1, max(p2, p3))
	// 	}
	for i := 1; i < N; i++ {
		for j := 1; j < M; j++ {
			p1 := dp[i-1][j]
			p2 := dp[i][j-1]
			p3 := 0
			if str1[i] == str2[j] {
				p3 = dp[i-1][j-1] + 1
			}
			dp[i][j] = max(p1, max(p2, p3))
		}
	}
	return dp[N-1][M-1]

}

func main() {
	str1 := "abcde"
	str2 := "ace"
	ans1 := longestsubstring(str1, str2)
	fmt.Println(ans1)
	ans2 := longestsubstring2(str1, str2)
	fmt.Println(ans2)
}
