package main

import "fmt"

//给定一个string
//返回多少种转化方法
func number(str string) int {
	return process([]byte(str), 0)
}

//0 ~ index -1 无需过问
//index ... 有多少种方法
func process(str []byte, index int) int {
	//转化完成收集点数
	if index == len(str) {
		return 1
	}
	//前面转化错误
	if str[index] == '0' {
		return 0
	}
	p1 := process(str, index+1)
	p2 := 0
	if index+1 < len(str) && ((str[index]-'0')*10+str[index+1]-'0') < 27 {
		p2 = process(str, index+2)
	}
	return p1 + p2

}
func dp(str string) int {

	bStr := []byte(str)
	N := len(bStr)

	var dp []int
	for i := 0; i <= N; i++ {
		dp = append(dp, 0)
	}
	// if index == len(str) {
	// 	return 1
	// }
	dp[N] = 1
	// if str[index] == '0' {
	// 	return 0
	// }
	// p1 := process(str, index+1)
	// p2 := 0
	// if index+1 < len(str) && ((str[index]-'0')*10+str[index+1]-'0') < 27 {
	// 	p2 = process(str, index+2)
	// }
	for i := N - 1; i >= 0; i-- {
		if bStr[i] == '0' {
			dp[i] = 0
			continue
		}
		p1 := dp[i+1]
		p2 := 0
		if i+1 < N && ((bStr[i]-'0')*10+bStr[i+1]-'0') < 27 {
			p2 = dp[i+2]
		}
		dp[i] = p1 + p2

	}

	return dp[0]

}

func main() {
	str := "2132082"
	ans1 := number(str)
	fmt.Println(ans1)
	ans2 := dp(str)
	fmt.Println(ans2)
}
