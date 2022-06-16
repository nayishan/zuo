package listmid

type ListNode struct {
	Val  int
	Next *ListNode
}

func MidOrUpMidNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}
	slow := head.Next
	fast := head.Next.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func MidOrDownMidNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow := head.Next
	fast := head.Next
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func MidOrUpMidPreNode(head *ListNode) *ListNode {
	if nil == head || nil == head.Next {
		return nil
	}
	if nil == head.Next.Next {
		return head
	}
	slow := head
	fast := head.Next.Next
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func MidOrDownMidPreNode(head *ListNode) *ListNode {
	if nil == head || nil == head.Next {
		return nil
	}
	slow := head
	fast := head.Next
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow

}
