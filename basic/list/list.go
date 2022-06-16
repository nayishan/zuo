package list

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode
	var next *ListNode
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil {
		return nil
	}
	var leftTail *ListNode
	changeHead := false
	cur := head
	var pre *ListNode
	var next *ListNode
	index := 1
	for index < left {
		leftTail = cur
		cur = cur.Next
		index++
	}

	if leftTail == nil {
		changeHead = true // index == 1
	}
	betweenTail := cur
	pre = cur
	for index <= right {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
		index++
	}
	betweenHead := pre
	rightHead := cur

	if changeHead {
		betweenTail.Next = rightHead
		head = betweenHead
	} else {
		leftTail.Next = betweenHead
		betweenTail.Next = rightHead
	}
	return head
}
