package groupreverse

type ListNode struct {
	val  int
	Next *ListNode
}

// start->1->2->3->end k=5 return end
func getKGroupEnd(start *ListNode, k int) *ListNode {
	end := start
	for index := k - 1; index > 0; index-- {
		if end != nil {
			end = end.Next
		} else {
			break
		}
	}
	return end
}

//s->1->2->3->end->end+1  reverse end->3->2->1->s->end+1
func reverse(start *ListNode, end *ListNode) {
	end = end.Next
	var pre *ListNode = nil
	cur := start
	var next *ListNode = nil
	for cur != end {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	start.Next = end
}

func reverseKGroup(start *ListNode, k int) *ListNode {
	head := start
	tempCurEnd := getKGroupEnd(head, k)
	if tempCurEnd == nil {
		return head
	}
	reverse(head, tempCurEnd)
	head = tempCurEnd

	var tempCurHead *ListNode = nil
	tempPreEnd := start
	for tempPreEnd.Next != nil {
		tempCurHead = tempPreEnd.Next
		tempCurEnd = getKGroupEnd(tempCurHead, k)
		if tempCurEnd == nil {
			return head
		}
		reverse(tempCurHead, tempCurEnd)
		tempPreEnd.Next = tempCurEnd
		tempPreEnd = tempCurHead
	}
	return head
}
