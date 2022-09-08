package main

import "fmt"

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

//w,v重量和价值数组
//bag 背包重要
//返回在不超过最大重要的最大价值
func getMaxValue(w, v []int, bag int) int {
	return process(w, v, 0, bag)
}

//当前到了index包括index，index之后的货物可以自由选择。
//做的选择不超过最大承重
//返回最大价值
func process(w, v []int, index int, bag int) int {
	if bag < 0 {
		return -1
	}

	if index == len(w) {
		return 0
	}
	//不要当前的货物
	p1 := process(w, v, index+1, bag)
	p2 := 0
	next := process(w, v, index+1, bag-w[index])
	if next != -1 {
		p2 = v[index] + next
	}
	return max(p1, p2)
}

func dp(w, v []int, bag int) int {
	N := len(w)
	//index 0~N
	//bag 0~ bag
	var dp [][]int
	for i := 0; i <= N; i++ {
		col := make([]int, bag+1)
		dp = append(dp, col)
	}
	//index 是N 都为0
	// if index == len(w) {
	// 	return 0
	// }
	for index := N - 1; index >= 0; index-- {
		for rest := 0; rest <= bag; rest++ {

			// p1 := process(w, v, index+1, bag)
			// p2 := 0
			// next := process(w, v, index+1, bag-w[index])
			// if next != -1 {
			// 	p2 = v[index] + next
			// }
			// return max(p1, p2)
			p1 := dp[index+1][rest]
			p2 := 0
			next := -1
			if rest-w[index] >= 0 {
				next = dp[index+1][rest-w[index]]
			}
			if next != -1 {
				p2 = v[index] + next

			}
			dp[index][rest] = max(p1, p2)
		}
	}

	//return process(w, v, 0, bag)
	return dp[0][bag]
}

func main() {
	weight := []int{3, 2, 4, 7, 3, 1, 7}
	value := []int{5, 6, 3, 19, 12, 4, 2}
	bag := 15
	ans1 := getMaxValue(weight, value, bag)
	fmt.Println(ans1)
	ans2 := dp(weight, value, bag)
	fmt.Println(ans2)
}
