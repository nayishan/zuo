package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func ListPartition(head *ListNode, value int) *ListNode {

	var bigHead *ListNode
	var bigEnd *ListNode
	var equalHead *ListNode
	var equalEnd *ListNode
	var smallHead *ListNode
	var smallEnd *ListNode
	var ans *ListNode
	temp := head

	for temp != nil {
		if temp.Val > value {
			if bigHead == nil {
				bigHead = temp
				bigEnd = temp
			} else {
				bigEnd.Next = temp
				bigEnd = bigEnd.Next
			}
		} else if temp.Val == value {
			if equalHead == nil {
				equalHead = temp
				equalEnd = temp
			} else {
				equalEnd.Next = temp
				equalEnd = equalEnd.Next
			}
		} else {
			if smallHead == nil {
				smallHead = temp
				smallEnd = temp
			} else {
				smallEnd.Next = temp
				smallEnd = smallEnd.Next
			}
		}

		temp = temp.Next
	}

	if smallHead == nil {
		if equalHead == nil {
			ans = bigHead
		} else {
			ans = equalHead
			equalEnd.Next = bigHead
		}
	} else {
		if equalHead == nil {
			smallEnd.Next = bigHead

		} else {
			smallEnd.Next = equalHead
			equalEnd.Next = smallHead
		}

	}
	if bigEnd != nil {
		bigEnd.Next = nil
	}
	return ans
}

func ListPartition2(head *ListNode, value int) *ListNode {

	var bigEqualHead *ListNode
	var bigEqualEnd *ListNode
	var smallHead *ListNode
	var smallEnd *ListNode
	var ans *ListNode
	temp := head

	for temp != nil {
		if temp.Val >= value {
			if bigEqualHead == nil {
				bigEqualHead = temp
				bigEqualEnd = temp
			} else {
				bigEqualEnd.Next = temp
				bigEqualEnd = bigEqualEnd.Next
			}
		} else {
			if smallHead == nil {
				smallHead = temp
				smallEnd = temp
			} else {
				smallEnd.Next = temp
				smallEnd = smallEnd.Next
			}
		}

		temp = temp.Next
	}

	if smallHead == nil {
		ans = bigEqualHead
	} else {
		ans = smallHead
		smallEnd.Next = bigEqualHead
	}
	if bigEqualEnd != nil {
		bigEqualEnd.Next = nil
	}
	return ans
}
