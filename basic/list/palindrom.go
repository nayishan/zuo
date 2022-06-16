package palindrom

type ListNode struct {
	Val  int
	Next *ListNode
}

func IsPalindrom1(head *ListNode) bool {
	if nil == head {
		return false
	}
	stack := make([]int, 0)
	tempHead := head
	for tempHead != nil {
		stack = append(stack, tempHead.Val)
		tempHead = tempHead.Next
	}
	N := len(stack)
	for head != nil {
		if stack[N-1] != head.Val {
			return false
		}
		head = head.Next
		stack = stack[:N-1]
		N = N - 1
	}
	return true
}
