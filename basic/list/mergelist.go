package mergelist

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(head1 *ListNode, head2 *ListNode) *ListNode {
	if head1 == nil && head2 == nil {
		return nil
	}
	if head1 == nil {
		return head2
	}
	if head2 == nil {
		return head1
	}
	var head *ListNode
	var cur2 *ListNode
	if head1.Val > head2.Val {
		head = head2
		cur2 = head1
	} else {
		head = head1
		cur2 = head2
	}
	cur1 := head.Next
	cur := head

	for cur1 != nil && cur2 != nil {
		if cur1.Val < cur2.Val {
			cur.Next = cur1
			cur1 = cur1.Next
		} else {
			cur.Next = cur2
			cur2 = cur2.Next
		}
		cur = cur.Next
	}
	if cur1 == nil {
		cur.Next = cur2
	} else {
		cur.Next = cur1
	}
	return head
}
