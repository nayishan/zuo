package main

import "fmt"

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a >= b {
		return b
	} else {
		return a
	}
}

//返回先手获取的最好分数
func xianshou(arr []int, L, R int) int {
	if L == R {
		return arr[L]
	}
	p1 := arr[L] + houshou(arr, L+1, R)
	p2 := arr[R] + houshou(arr, L, R-1)
	return max(p1, p2)
}

//因为两边都很聪明，所以只能得到最优中最差的结果
func houshou(arr []int, L, R int) int {
	if L == R {
		return 0
	}
	p1 := xianshou(arr, L+1, R) //如果对手拿走了L
	p2 := xianshou(arr, L, R-1) //如果对手拿走了R
	return min(p1, p2)
}
func win(arr []int) int {
	if arr == nil {
		return 0
	}
	if len(arr) == 0 {
		return 0
	}
	first := xianshou(arr, 0, len(arr)-1)
	second := houshou(arr, 0, len(arr)-1)
	fmt.Println(first, second)
	return max(first, second)
}
func main() {
	arr := []int{5, 7, 4, 5, 8, 1, 6, 0, 3, 4, 6, 1, 7}
	fmt.Println(win(arr))
	fmt.Println(win2(arr))
	fmt.Println(win3(arr))
}

func win2(arr []int) int {
	if arr == nil {
		return 0
	}
	if len(arr) == 0 {
		return 0
	}
	var xianshoumap [][]int
	var houshoumap [][]int
	for i := 0; i < len(arr); i++ {
		col := make([]int, 0)
		for j := 0; j < len(arr); j++ {
			col = append(col, -1)
		}
		xianshoumap = append(xianshoumap, col)
	}
	for i := 0; i < len(arr); i++ {
		col := make([]int, 0)
		for j := 0; j < len(arr); j++ {
			col = append(col, -1)
		}
		houshoumap = append(houshoumap, col)
	}
	first := xianshou2(arr, 0, len(arr)-1, xianshoumap, houshoumap)
	second := houshou2(arr, 0, len(arr)-1, xianshoumap, houshoumap)
	fmt.Println(first, second)
	return max(first, second)

}

func xianshou2(arr []int, L, R int, xianshoumap, houshoumap [][]int) int {
	if xianshoumap[L][R] != -1 {
		return xianshoumap[L][R]
	} else {
		if L == R {
			xianshoumap[L][R] = arr[L]
			return xianshoumap[L][R]
		}
		p1 := arr[L] + houshou2(arr, L+1, R, xianshoumap, houshoumap)
		p2 := arr[R] + houshou2(arr, L, R-1, xianshoumap, houshoumap)
		xianshoumap[L][R] = max(p1, p2)
		return xianshoumap[L][R]
	}
}

func houshou2(arr []int, L, R int, xianshoumap, houshoumap [][]int) int {
	if houshoumap[L][R] != -1 {
		return houshoumap[L][R]
	} else {
		if L == R {
			houshoumap[L][R] = 0
			return houshoumap[L][R]
		} else {
			p1 := xianshou2(arr, L+1, R, xianshoumap, houshoumap) //如果对手拿走了L
			p2 := xianshou2(arr, L, R-1, xianshoumap, houshoumap) //如果对手拿走了R
			houshoumap[L][R] = min(p1, p2)
			return houshoumap[L][R]
		}
	}
}

func win3(arr []int) int {
	var xianshoumap [][]int
	var houshoumap [][]int
	for i := 0; i < len(arr); i++ {
		col := make([]int, len(arr))
		xianshoumap = append(xianshoumap, col)
	}
	for i := 0; i < len(arr); i++ {
		col := make([]int, len(arr))
		houshoumap = append(houshoumap, col)
	}
	//startCol = 0 是给定的L==R
	for startCol := 1; startCol < len(arr); startCol++ {
		L := 0
		R := startCol
		for R < len(arr) {
			// p1 := arr[L] + houshou(arr, L+1, R)
			// p2 := arr[R] + houshou(arr, L, R-1)
			xianshoumap[L][R] = max(arr[L]+houshoumap[L+1][R], arr[R]+houshoumap[L][R-1])
			// p1 := xianshou(arr, L+1, R) //如果对手拿走了L
			// p2 := xianshou(arr, L, R-1) //如果对手拿走了R
			houshoumap[L][R] = min(xianshoumap[L+1][R], xianshoumap[L][R-1])
			L++
			R++
		}
	}
	// first := xianshou(arr, 0, len(arr)-1)
	// second := houshou(arr, 0, len(arr)-1)
	// return min(p1, p2)
	return max(xianshoumap[0][len(arr)-1], houshoumap[0][len(arr)-1])

}
