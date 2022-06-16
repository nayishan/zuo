package stack

import "fmt"

var stack []int

func init() {
	stack = make([]int, 0)
	fmt.Println("init stack,len", len(stack), "cap", cap(stack))
}

//Push pop Looukup Size isEmpry

func Push(val int) {
	stack = append(stack, val)
}

func Pop() (int, bool) {
	N := len(stack)
	if N == 0 {
		return 0, false
	}
	ans := stack[N-1]
	stack = stack[:N-1]
	return ans, true
}

func IsEmpty() bool {
	N := len(stack)
	if N == 0 {
		return true
	} else {
		return false
	}
}

func Size() int {
	return len(stack)
}
