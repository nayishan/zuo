package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	temp := head
	var ans *Node

	if head == nil {
		return nil
	}

	for temp != nil {
		next := temp.Next
		copyNode := &Node{temp.Val, nil, nil}
		temp.Next = copyNode
		copyNode.Next = next
		temp = next
	}

	tempHead := head
	tempEnd := head.Next
	for tempHead != nil {
		if tempHead.Random != nil {
			tempEnd.Random = tempHead.Random.Next
		} else {
			tempEnd.Random = nil
		}
		tempHead = tempHead.Next.Next
		if tempHead != nil {
			tempEnd = tempEnd.Next.Next
		}
	}

	tempHead = head
	tempEnd = head.Next
	var tempHeadNext *Node
	var tempEndNext *Node
	ans = tempEnd
	for tempHead != nil {
		tempHeadNext = tempHead.Next.Next
		if tempHeadNext != nil {
			tempEndNext = tempEnd.Next.Next
		} else {
			tempEndNext = nil
		}
		tempHead.Next = tempHeadNext
		tempEnd.Next = tempEndNext
		tempHead = tempHeadNext
		tempEnd = tempEndNext
	}
	return ans
}
