package stack

// import "fmt"

type stack []interface{}

// func init() stack {
// 	temp := make([]int, 0)
// 	fmt.Println("init stack")
// 	return temp
// }
func StackNew() *stack {
	temp := make(stack, 0)
	return &temp

}

//Push pop Looukup Size isEmpry

func (s *stack) Push(val interface{}) {
	*s = append(*s, val)
}
func (s *stack) Peek() interface{} {
	return (*s)[s.Size()-1]
}

func (s *stack) Pop() interface{} {
	N := len(*s)
	ans := (*s)[N-1]
	*s = (*s)[:N-1]
	return ans
}

func (s *stack) IsEmpty() bool {
	N := len(*s)
	if N == 0 {
		return true
	} else {
		return false
	}
}

func (s *stack) Size() int {
	return len(*s)
}
