package main

import "fmt"

func ways1(n int) int {
	if n < 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return process(1, n)
}

//前面的数是pre，还有rest 需要裂开，裂开的方法是
func process(pre int, rest int) int {
	if rest == 0 {
		return 1
	}
	if pre > rest {
		return 0
	}
	ways := 0
	for first := pre; first <= rest; first++ {
		ways += process(first, rest-first)
	}
	return ways
}

func dp1(n int) int {
	if n < 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
	}
	//第0行无效
	//

	for pre := 1; pre <= n; pre++ {
		dp[pre][0] = 1
	}
	//其余对角线以下是0
	for pre := 1; pre < n; pre++ {
		for rest := 1; rest < pre; rest++ {
			dp[pre][rest] = 0
		}
	}
	//对角线是1
	for pre := 1; pre <= n; pre++ {
		dp[pre][pre] = 1
	}
	for pre := n - 1; pre >= 1; pre-- {
		for rest := pre + 1; rest <= n; rest++ {
			ways := 0
			for first := pre; first <= rest; first++ {
				ways += dp[first][rest-first]
			}
			dp[pre][rest] = ways

		}

	}
	return dp[1][n]
}
func dp2(n int) int {
	if n < 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
	}
	//第0行无效
	//

	for pre := 1; pre <= n; pre++ {
		dp[pre][0] = 1
	}
	//其余对角线以下是0
	for pre := 1; pre < n; pre++ {
		for rest := 1; rest < pre; rest++ {
			dp[pre][rest] = 0
		}
	}
	//对角线是1
	for pre := 1; pre <= n; pre++ {
		dp[pre][pre] = 1
	}
	for pre := n - 1; pre >= 1; pre-- {
		for rest := pre + 1; rest <= n; rest++ {
			dp[pre][rest] = dp[pre+1][rest]
			//rest - pre 一定大于等于0
			dp[pre][rest] += dp[pre][rest-pre]
		}

	}
	return dp[1][n]
}
func main() {
	fmt.Println(dp1(13))
	fmt.Println(ways1(13))
	fmt.Println(dp2(13))
}
