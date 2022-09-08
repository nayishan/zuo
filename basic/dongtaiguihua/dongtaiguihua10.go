package main

import (
	"fmt"
	"math"
)

func min(a int, b int) int {
	if a >= b {
		return b
	} else {
		return a
	}
}
func ways(matrix [][]int) int {
	row := len(matrix)
	col := len(matrix[0])
	return process(matrix, 0, 0, row-1, col-1)
}

// 从x,y出发到a,b的最短路径和
//返回最短路径和
func process(matrix [][]int, x, y int, a, b int) int {
	//越界无效
	if x > a || y > b {
		return math.MaxInt
	}
	//到达位置返回的路径和是0
	if x == a && y == b {
		return 0
	}
	p1 := math.MaxInt
	p2 := math.MaxInt
	next1 := process(matrix, x+1, y, a, b)
	if next1 != math.MaxInt {
		p1 = next1 + matrix[x][y]

	}
	next2 := process(matrix, x, y+1, a, b)
	if next2 != math.MaxInt {
		p2 = next2 + matrix[x][y]
	}
	ans := min(p1, p2)
	return ans
}

func ways2(matrix [][]int) int {
	var dp [][]int
	row := len(matrix)
	col := len(matrix[0])
	dp = make([][]int, row)
	for i := 0; i < row; i++ {
		dp[i] = make([]int, col)
	}
	dp[row-1][col-1] = 0
	for i := row - 2; i >= 0; i-- {
		dp[i][col-1] = dp[i+1][col-1] + matrix[i][col-1]
	}
	for i := col - 2; i >= 0; i-- {
		dp[row-1][i] = dp[row-1][i+1] + matrix[row-1][i]
	}
	for i := row - 2; i >= 0; i-- {
		for j := col - 2; j >= 0; j-- {
			p1 := dp[i+1][j] + matrix[i][j]
			p2 := dp[i][j+1] + matrix[i][j]
			dp[i][j] = min(p1, p2)
		}
	}
	return dp[0][0]
}

func ways3(matrix [][]int) int {
	var dp [][]int
	row := len(matrix)
	col := len(matrix[0])
	dp = make([][]int, row)
	for i := 0; i < row; i++ {
		dp[i] = make([]int, col)
	}
	dp[0][0] = matrix[0][0]
	for i := 1; i < row; i++ {
		dp[i][0] = dp[i-1][0] + matrix[i][0]
	}
	for i := 1; i < col; i++ {
		dp[0][i] = dp[0][i-1] + matrix[0][i]
	}
	for i := 1; i < row; i++ {
		for j := 1; j < row; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + matrix[i][j]
		}
	}
	return dp[row-1][col-1]
}
func ways4(matrix [][]int) int {
	row := len(matrix)
	col := len(matrix[0])
	array := make([]int, col)
	array[0] = matrix[0][0]
	for i := 1; i < col; i++ {
		array[i] = array[i-1] + matrix[0][i]
	}
	for i := 1; i < row; i++ {
		array[0] += matrix[i][0]
		for j := 1; j < col; j++ {
			array[j] = min(array[j], array[j-1]) + matrix[i][j]
		}
	}
	return array[col-1]
}

func main() {
	var matrix [][]int

	col := 4
	matrix = make([][]int, 4)
	for i := 0; i < 4; i++ {
		matrix[i] = make([]int, col)
	}
	matrix = [][]int{
		{1, 3, 5, 9},
		{8, 1, 3, 4},
		{5, 0, 6, 1},
		{8, 8, 4, 0},
	}
	fmt.Println(ways(matrix))
	fmt.Println(ways2(matrix))
	fmt.Println(ways3(matrix))
	fmt.Println(ways4(matrix))

}
