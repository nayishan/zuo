package main

import "math"

//所有的可能性是(M+1)^K
//注意不要剪枝。一定要走到K次
func right1(N, M, K int) float64 {
	if N < 1 || M < 1 || K < 1 {
		return 0
	}
	var all int64
	var kill int64
	all = int64(math.Pow(float64(M+1), float64(K)))
	kill = process1(K, N, M)
	return float64(kill / all)
}

//返回砍了count之后，怪物血量小于0的次数
func process1(times int, hp int, demage int) int64 {
	if times == 0 {
		if hp <= 0 {
			return 1
		} else {
			return 0
		}
	}
	if hp <= 0 {
		return int64(math.Pow(float64(demage)+1, float64(times)))
	}
	var ans int64
	for i := 0; i <= demage; i++ {
		ans += process1(times-1, hp-i, demage)
	}
	return ans

}

func dp1(K, N, M int) float64 {
	if N < 1 || M < 1 || K < 1 {
		return 0
	}
	all := int64(math.Pow(float64(N+1), float64(K)))
	dp := make([][]int64, N+1)
	for i := 0; i <= N+1; i++ {
		dp[i] = make([]int64, K+1)
	}
	dp[0][0] = 1
	for times := 0; times <= N; times++ {
		dp[times][0] = int64(math.Pow(float64(M)+1, float64(times)))
		for hp := 0; hp <= N; hp++ {
			var ways int64
			for i := 0; i < M; i++ {
				if hp-i <= 0 {
					ways += int64(math.Pow(float64(M)+1, float64(times)-1))
				} else {
					ways += dp[times-1][hp-i]
				}
			}
			dp[times][hp] = ways
		}
	}
	return float64(dp[K][N] / all)

}

func dp2(K, N, M int) float64 {
	if N < 1 || M < 1 || K < 1 {
		return 0
	}
	all := int64(math.Pow(float64(N+1), float64(K)))
	dp := make([][]int64, N+1)
	for i := 0; i <= N+1; i++ {
		dp[i] = make([]int64, K+1)
	}
	dp[0][0] = 1
	for times := 0; times <= N; times++ {
		dp[times][0] = int64(math.Pow(float64(M)+1, float64(times)))
		for hp := 0; hp <= N; hp++ {
			dp[times][hp] = dp[times][hp-1] + dp[times-1][hp]
			if hp-1-M >= 0 {
				dp[times][hp] -= dp[times-1][hp-1-M]
			} else {
				dp[times][hp] -= int64(math.Pow(float64(M)+1, float64(times)-1))

			}
		}
	}
	return float64(dp[K][N] / all)

}
