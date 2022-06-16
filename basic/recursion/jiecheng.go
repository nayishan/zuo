package main

import "fmt"

func jiecheng(a int) (b int) {
	if a == 0 {
		return 1
	}
	cur := 1
	res := 0
	for i := 1; i <= a; i++ {
		cur = cur * i
		res += cur
	}
	fmt.Println(res)
	return res
}
func main() {
	res := jiecheng(13)
	fmt.Println("res", res)
}
