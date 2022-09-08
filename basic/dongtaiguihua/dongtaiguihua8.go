package main

import "fmt"

func horseJump(k int, x int, y int) int {
	return process(0, 0, k, x, y)
}

//目标ab
//还剩rest步
//目前是xy
func process(x, y int, rest int, a, b int) int {
	//不是挨个方法判断越界，就可以随便蹦，被调的函数判断越界
	if x < 0 || y < 0 || x > 9 || y > 8 {
		return 0
	}
	if rest == 0 {
		if x == a && y == b {
			return 1
		} else {
			return 0
		}
	}
	//8种蹦法
	ans := process(x+2, y+1, rest-1, a, b)
	ans += process(x+1, y+2, rest-1, a, b)
	ans += process(x-1, y+2, rest-1, a, b)
	ans += process(x-2, y+1, rest-1, a, b)
	ans += process(x-2, y-1, rest-1, a, b)
	ans += process(x-1, y-2, rest-1, a, b)
	ans += process(x+2, y-1, rest-1, a, b)
	ans += process(x+1, y-2, rest-1, a, b)
	return ans
}

func main() {
	x := 7
	y := 7
	k := 10
	ans1 := horseJump(k, x, y)
	fmt.Println(ans1)
	ans2 := horseJump2(k, x, y)
	fmt.Println(ans2)
}
func pick(dp [10][9][]int, x, y, rest int) int {
	// if x < 0 || y < 0 || x > 9 || y > 8 {
	// 	return 0
	// }
	if x < 0 || y < 0 || x > 9 || y > 8 {
		return 0
	}
	return dp[x][y][rest]
}

func horseJump2(k int, a int, b int) int {
	// var dp [][][]int
	// var Y [][]int
	// for i := 0; i < 10; i++ {
	// 	for j := 0; j < 9; j++ {
	// 		step := make([]int, k+1)
	// 		Y = append(Y, step)
	// 	}
	// 	dp = append(dp, Y)
	// }
	var dp [10][9][]int
	for i := 0; i < 10; i++ {
		for j := 0; j < 9; j++ {
			dp[i][j] = make([]int, k+1)
		}
	}
	// if rest == 0 {
	// 	if x == a && y == b {
	// 		return 1
	// 	} else {
	// 		return 0
	// 	}
	// }
	dp[a][b][0] = 1
	for rest := 1; rest <= k; rest++ {
		for x := 0; x < 10; x++ {
			for y := 0; y < 9; y++ {
				// ans := process(x+2, y+1, rest-1, a, b)
				// ans += process(x+1, y+2, rest-1, a, b)
				// ans += process(x-1, y+2, rest-1, a, b)
				// ans += process(x-2, y+1, rest-1, a, b)
				// ans += process(x-2, y-1, rest-1, a, b)
				// ans += process(x-1, y-2, rest-1, a, b)
				// ans += process(x+2, y-1, rest-1, a, b)
				// ans += process(x+1, y-2, rest-1, a, b)
				ways := pick(dp, x+2, y+1, rest-1)
				ways += pick(dp, x+1, y+2, rest-1)
				ways += pick(dp, x-1, y+2, rest-1)
				ways += pick(dp, x-2, y+1, rest-1)
				ways += pick(dp, x-2, y-1, rest-1)
				ways += pick(dp, x-1, y-2, rest-1)
				ways += pick(dp, x+2, y-1, rest-1)
				ways += pick(dp, x+1, y-2, rest-1)
				dp[x][y][rest] = ways
			}
		}
	}
	// return process(0, 0, k, x, y)
	return dp[0][0][k]
}
