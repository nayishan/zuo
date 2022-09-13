package main

import "math"

func nQueen(n int) int {
	if n < 1 {
		return 0
	}
	record := make([]int, n)
	return process(0, record, n)
}

//当前来到index行，要在index上放皇后
//所有列都尝试，但是必须要保证与之前的所有列不打架
func process(i int, record []int, n int) int {
	if i == n {
		return 1
	}
	res := 0
	//挨个列尝试
	for j := 0; j < n; j++ {
		if isValid(record, i, j) {
			record[i] = j
			res += process(i+1, record, n)
		}
	}
	return res
}

func isValid(record []int, i, j int) bool {
	for k := 0; k < i; k++ {
		if j == record[k] || (math.Abs(float64(record[k]-j)) == math.Abs(float64(i-k))) {
			return false
		}
	}
	return true
}

//32以上超过了int的范围
func nQueen2(n int) int {
	if n < 1 || n > 32 {
		return 1
	}
	limit := 0
	if n == 32 {
		limit = -1
	} else {
		limit = 1<<n - 1
	}
	return process2(limit, 0, 0, 0)
}

// limit 当colLim == limit 时，发现了一种有效方法
// leftDiaLim 左下的限制
// colLim 列限制
//rightDiaLim 右下的限制
func process2(limit int, colLim int, leftDiaLim int, rightDiaLim int) int {
	if limit == colLim {
		return 1
	}
	//加工出所有可以放皇后的位置
	pos := limit & (^(colLim | leftDiaLim | rightDiaLim))
	mostRightOne := 0
	res := 0
	for pos != 0 {
		mostRightOne = pos & (^pos + 1)
		pos = pos - mostRightOne
		res += process2(limit, colLim|mostRightOne, (leftDiaLim|mostRightOne)<<1, (rightDiaLim|mostRightOne)>>1)
	}
	return res
}
