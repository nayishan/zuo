package main

import "zuo/lib/unionsethash"

func findCircleNum(isConnected [][]int) int {
	N := len(isConnected)
	citys := make([]int, N)
	for i := 0; i < N; i++ {
		citys[i] = i
	}
	interfaceCitys := make([]interface{}, len(citys))
	for i, d := range citys {
		interfaceCitys[i] = d
	}
	u := unionsethash.Init(interfaceCitys)
	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			if isConnected[i][j] == 1 {
				u.Union(i, j)
			}
		}
	}
	return u.Sets()
}

func numIslands(board [][]byte) int {
	//转化interface{}
	row := len(board)
	col := len(board[0])
	coordinate := make([][]interface{}, row)
	for j := 0; j < row; j++ {
		coordinate[j] = make([]interface{}, col)
	}
	for j := 0; j < row; j++ {
		for i := 0; i < col; i++ {
			if board[j][i] == '1' {
				coordinate[j][i] = '1'
			}

		}
	}

	u := unionsethash.InitCoordinate(coordinate)
	for i := 1; i < col; i++ {
		if coordinate[0][i] == '1' && coordinate[0][i-1] == '1' {
			u.Union(unionsethash.Coordinate{X: 0, Y: i}, unionsethash.Coordinate{X: 0, Y: i - 1})
		}
	}
	for j := 1; j < row; j++ {
		if coordinate[j][0] == '1' && coordinate[j-1][0] == '1' {
			u.Union(unionsethash.Coordinate{X: j, Y: 0}, unionsethash.Coordinate{X: j - 1, Y: 0})
		}
	}
	for i := 1; i < col; i++ {
		for j := 1; j < row; j++ {
			if coordinate[j][i] == '1' && coordinate[j-1][i] == '1' {
				u.Union(unionsethash.Coordinate{X: j, Y: i}, unionsethash.Coordinate{X: j - 1, Y: i})
			}
			if coordinate[j][i] == '1' && coordinate[j][i-1] == '1' {
				u.Union(unionsethash.Coordinate{X: j, Y: i}, unionsethash.Coordinate{X: j, Y: i - 1})
			}
		}
	}
	return u.Sets()
}
