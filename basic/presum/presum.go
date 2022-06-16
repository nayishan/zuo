package main

import "fmt"

func rangeSum(a []int) []int {
	if a == nil {
		return nil
	}
	if len(a) == 0 {
		return nil
	}
	// help := make([]int, len(a))
	// help[0] = a[0]
	// for i := 1; i < len(a); i++ {
	// 	help[i] = a[i] + help[i-1]
	// }
	help := []int{}
	help = append(help, a[0])
	if len(a) == 1 {
		return help
	}
	for i := 1; i < len(a); i++ {
		help = append(help, a[i]+help[i-1])
	}
	return help

}
func preSum(a []int, l int, r int) int {
	if a == nil {
		return 0
	}
	if len(a) == 0 || l > len(a)-1 || r > len(a)-1 || l < 0 || r < 0 {
		return 0
	}
	sum := 0
	if l == 0 {
		sum = a[r]
	} else {
		sum = a[r] - a[l-1]
	}
	return sum
}
func main() {
	a := []int{3, 4, 2, 1, 6, 7, 8}
	help := rangeSum(a)
	fmt.Println(help)
	sum := preSum(help, 2, 6)
	fmt.Println(sum)
}
