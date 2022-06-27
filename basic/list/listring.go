package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return false
	}

	slow := head.Next
	fast := head.Next.Next

	for fast != nil && fast.Next != nil && slow != fast {
		fast = fast.Next.Next
		slow = slow.Next
	}
	if fast == nil || fast.Next == nil {
		return false
	}

	return true

}

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return nil
	}

	slow := head.Next
	fast := head.Next.Next

	for fast != nil && fast.Next != nil && slow != fast {
		fast = fast.Next.Next
		slow = slow.Next
	}
	if fast == nil || fast.Next == nil {
		return nil
	}

	fast = head

	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}

func findEnd(head *ListNode) *ListNode {
	temp := head
	var pre *ListNode = nil
	for temp != nil {
		pre = temp
		temp = temp.Next
	}
	return pre
}

func Count(head *ListNode, entry *ListNode) int {

	tempHead := head
	ans := 0
	for tempHead != entry {
		ans++
		tempHead = tempHead.Next
	}
	return ans
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	tempA := headA
	tempB := headB
	entryA := detectCycle(tempA)
	entryB := detectCycle(tempB)

	if entryA != nil && entryB != nil {

		//1. out cycle
		tempA = headA
		tempB = headB
		indexA := Count(headA, entryA)
		indexB := Count(headB, entryB)
		indexDiff := 0
		tempA = headA
		tempB = headB
		if indexA > indexB {
			indexDiff = indexA - indexB
			for indexDiff >= 0 {
				indexDiff--
				tempA = tempA.Next
			}
			for tempA != entryA && tempB != entryB {
				if tempA == tempB {
					return tempA
				}
				tempA = tempA.Next
				tempB = tempB.Next
			}
		} else {
			indexDiff = indexB - indexA
			for indexDiff >= 0 {
				indexDiff--
				tempB = tempB.Next
			}
			for tempA != entryA && tempB != entryB {
				if tempA == tempB {
					return tempA
				}
				tempA = tempA.Next
				tempB = tempB.Next
			}
		}

		// 2. in cycle
		if entryA == entryB {
			return entryA
		}
		temp := entryA
		temp = temp.Next
		for temp != entryA {
			if temp == entryB {
				//return entryA is an answer too.
				return entryB
			}
		}

	}

	if entryA == nil && entryB == nil {
		tempA = headA
		tempEnd := findEnd(tempA)
		if tempEnd == nil {
			return nil
		}

		tempEnd.Next = tempB
		tempA = headA
		entryC := detectCycle(tempA)
		tempEnd.Next = nil
		if entryC == nil {
			return nil
		} else {
			return entryC
		}
	}

	return nil

}
