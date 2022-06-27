package isPalindrome

type ListNode struct {
	Val  int
	Next *ListNode
}

func midDownmid(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func reversePartList(point *ListNode) *ListNode {
	now := point.Next
	if point.Next == nil {
		return point
	}
	pre := point
	var next *ListNode
	for now != nil {
		next = now.Next
		now.Next = pre
		pre = now
		now = next
	}
	return pre
}

func reverseBack(right *ListNode) {
	var pre *ListNode = nil
	var next *ListNode = nil
	now := right
	for now != nil {
		next = now.Next
		now.Next = pre
		pre = now
		now = next

	}
}

func isEqual(left, right *ListNode) bool {
	l := left
	r := right
	for l != nil && r != nil {
		if l.Val != r.Val {
			return false
		}
		l = l.Next
		r = r.Next
	}
	return true

}
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	midOrDownMid := midDownmid(head)
	right := reversePartList(midOrDownMid)
	midOrDownMid.Next = nil
	left := head
	ans := isEqual(left, right)
	reverseBack(right)
	return ans
}
