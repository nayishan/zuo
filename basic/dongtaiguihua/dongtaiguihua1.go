package main

import "fmt"

//尝试法
func ways1(N, start, aim, K int) int {
	return process1(start, K, aim, N)
}

//机器人来到cur
//还有rest步需要走
//目标是aim
//位置总数
func process1(cur, rest, aim, N int) int {
	if rest == 0 {
		if cur == aim {
			return 1
		} else {
			return 0
		}
	}
	if cur == 1 {
		return process1(2, rest-1, aim, N)
	}
	if cur == N {
		return process1(N-1, rest, aim, N)
	}
	return process1(cur-1, rest-1, aim, N) + process1(cur+1, rest-1, aim, N)
}

func ways2(N, start, aim, K int) int {
	var dp [][]int
	//创建缓存表
	//设置初值为-1,表示未算过
	for i := 0; i < N+1; i++ {
		var row []int
		for j := 0; j < K+1; j++ {
			row = append(row, -1)
		}
		dp = append(dp, row)
	}

	return process2(start, K, aim, N, dp)
}
func process2(cur, rest, aim, N int, dp [][]int) int {
	//算过
	if dp[cur][rest] != -1 {
		return dp[cur][rest]
	}
	//没算过
	ans := 0
	if rest == 0 {
		if cur == aim {
			ans = 1
		} else {
			ans = 0
		}
	} else if cur == 1 {
		ans = process2(2, rest-1, aim, N, dp)
	} else if cur == N {
		ans = process2(N-1, rest-1, aim, N, dp)
	} else {
		ans = process2(cur+1, rest-1, aim, N, dp) + process2(cur-1, rest-1, aim, N, dp)
	}
	dp[cur][rest] = ans
	return ans
}

func ways3(N, start, aim, K int) int {

	var dp [][]int
	//创建缓存表
	//设置初值为-1,表示未算过
	for i := 0; i < N+1; i++ {
		row := make([]int, K+1)
		dp = append(dp, row)
	}
	//第0列
	dp[aim][0] = 1
	//第1列开始
	for rest := 1; rest < K+1; rest++ {
		//将第一行和最后一行单独计算
		dp[1][rest] = dp[2][rest-1]
		for cur := 2; cur < N; cur++ {
			dp[cur][rest] = dp[cur+1][rest-1] + dp[cur-1][rest-1]
		}
		dp[N][rest] = dp[N-1][rest-1]
	}
	return dp[start][K]

}

func main() {
	ans1 := ways1(4, 2, 4, 4)
	fmt.Println(ans1)
	ans2 := ways2(4, 2, 4, 4)
	fmt.Println(ans2)
	ans3 := ways3(4, 2, 4, 4)
	fmt.Println(ans3)
}
