package main

import "fmt"

//当前来到row, col位置，还剩rest步，走完rest步之后，来到x,y位置，方法数多少种
func p(row int, col int, rest int, x int, y int) int {
	if rest == 0 {
		if row == x && col == y {
			return 1
		} else {
			return 0
		}
	}
	if row < 0 || row > 9 || col < 0 || col > 8 || rest < 0 {
		return 0

	}
	//8种蹦法
	return p(row+2, col-1, rest-1, x, y) +
		p(row+2, col+1, rest-1, x, y) +
		p(row-2, col-1, rest-1, x, y) +
		p(row-2, col+1, rest-1, x, y) +
		p(row+1, col-2, rest-1, x, y) +
		p(row+1, col+2, rest-1, x, y) +
		p(row-1, col-2, rest-1, x, y) +
		p(row-1, col+2, rest-1, x, y)

}
func way3(x int, y int, k int) int {
	return p(0, 0, k, x, y)
}

func HorseJump(targetX, targetY, curX, curY, rest int) int {
	if curX < 0 || curX > 9 || curY < 0 || curY > 8 || rest < 0 {
		return 0
	}

	if rest == 0 {
		if curX == targetX && curY == targetY {
			return 1
		} else {
			return 0
		}
	}
	//八个方向
	ans := 0
	for _, direction := range [][]int{{+2, +1}, {+1, +2}, {-1, +2}, {-2, +1}, {-2, -1}, {-1, -2}, {+1, -2}, {+2, -1}} {
		ans += HorseJump(targetX, targetY, curX+direction[0], curY+direction[1], rest-1)
	}
	return ans
}

func HorseJumpForDp(targetX, targetY, k int) int {
	dp := make([][][]int, 10)
	for i := range dp {
		dp[i] = make([][]int, 9)
		for w := range dp[i] {
			dp[i][w] = make([]int, k+1)
		}
	}

	dp[0][0][0] = 1

	for level := 1; level <= k; level++ {
		for i := 0; i < 10; i++ {
			for j := 0; j < 9; j++ {
				ans := 0
				for _, direction := range [][]int{{+2, +1}, {+1, +2}, {-1, +2}, {-2, +1}, {-2, -1}, {-1, -2}, {+1, -2}, {+2, -1}} {
					if i+direction[0] < 0 || i+direction[0] > 9 || j+direction[1] < 0 || j+direction[1] > 8 {
						continue
					}
					ans += dp[i+direction[0]][j+direction[1]][level-1]
				}
				dp[i][j][level] = ans
			}
		}
	}
	return dp[targetX][targetY][k]
}

func main() {
	x, y, k := 6, 8, 10
	fmt.Println(way3(x, y, k))
	fmt.Println(HorseJump(6, 8, 0, 0, 10))
	fmt.Println(HorseJumpForDp(6, 8, 10))
}
