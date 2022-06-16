package main

import (
	"fmt"
	"math"
	/* "math" */)

//进位信息是两个数&  并且向左移一位
//异或是两个数的无进位相加
func add(a int, b int) int {
	jinwei := (a & b) << 1
	wujinwei := a ^ b
	for jinwei != 0 {
		temp := wujinwei
		wujinwei = jinwei ^ wujinwei
		jinwei = (jinwei & temp) << 1

	}
	return wujinwei
}

//求相反数
func negNum(a int) int {
	return add(^a, 1)
}

//a + b的相反数
func minus(a int, b int) int {
	return add(a, negNum(b))
}

func multi(a int, b int) int {
	ans := 0
	for b != 0 {
		if b&1 != 0 {
			ans += a
		}
		a = a << 1
		b = b >> 1
	}
	return ans
}
func isNeg(a int) bool {
	return a < 0
}
func div(a int, b int) int {

	x := 0
	y := 0
	if isNeg(a) {
		x = negNum(a)
	} else {
		x = a
	}
	if isNeg(b) {
		y = negNum(b)
	} else {
		y = b
	}
	ans := 0
	//x 右移找商的各个位的值
	for i := 31; i >= 0; i = minus(i, 1) {
		if (x >> i) >= y {
			ans |= (1 << i)
			x = minus(x, y<<i)
		}
	}
	if isNeg(a) != isNeg(b) {
		return negNum(ans)
	} else {
		return ans
	}
}

//注意math.MinInt32是没有对应的正值的，主函数就是处理这个边界问题
func divide(a int, b int) int {
	if a == math.MinInt32 && b == math.MaxInt32 {
		return 1
	} else if b == math.MinInt32 {
		return 0
	} else if a == math.MinInt32 {
		if b == negNum(1) { //leetcode 规定MinInt32/-1 == MaxInt32
			return math.MaxInt32
		} else {
			// (a+1)/b = c
			// c * b = e
			//a - e = m
			//m /b =  n
			//n + c
			//以上的所有步骤是为了防止溢出，并且计算出正确值的操作。
			c := div(add(a, 1), b)
			e := multi(c, b)
			m := minus(a, e)
			n := div(m, b)
			return add(n, c)
		}

	} else {
		return div(a, b)
	}
}
func int2byte(a int) []int {
	bytes := make([]int, 32)
	for i := 0; i < 32; i++ {
		if a&(1<<i) != 0 {
			bytes[31-i] = 1

		} else {
			bytes[31-i] = 0
		}
	}
	return bytes
}

func main() {
	fmt.Println(22 + 44)
	fmt.Println(add(22, 44))
	fmt.Println(44 - 22)
	fmt.Println(minus(44, 22))
	fmt.Println(6 * 7)
	fmt.Println(multi(6, 7))
	fmt.Println(int2byte(-2147483648))
	fmt.Println(divide(-2147483648, 1))
	fmt.Println(divide(-2147483648, -1))
}
