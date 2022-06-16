package addlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func size(head *ListNode) int {

	len := 0
	if head == nil {
		return len
	}
	for head != nil {
		head = head.Next
		len++
	}
	return len
}

func addTwoNumbers(list1 *ListNode, list2 *ListNode) *ListNode {
	len1 := size(list1)
	len2 := size(list2)
	longList := list1
	if len2 > len1 {
		longList = list2
	}
	newList := longList
	preList := longList
	val1 := 0
	val2 := 0
	overFlow := 0
	rest := 0
	for longList != nil {
		if list1 == nil {
			val1 = 0
		} else {
			val1 = list1.Val
			list1 = list1.Next
		}
		if list2 == nil {
			val2 = 0
		} else {
			val2 = list2.Val
			list2 = list2.Next
		}
		rest = (val1 + val2 + overFlow) % 10
		overFlow = (val1 + val2 + overFlow) / 10
		longList.Val = rest
		preList = longList
		longList = longList.Next
	}
	if overFlow != 0 {
		newNode := new(ListNode)
		newNode.Val = overFlow
		newNode.Next = nil
		preList.Next = newNode
	}
	return newList
}
