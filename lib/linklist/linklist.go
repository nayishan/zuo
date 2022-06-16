package linklist

import (
	"fmt"
)

// ListNode
type ListNode struct {
	Val  int
	Next *ListNode // 表示指向下一个结点
}

func InitListNode() ListNode {
	head := ListNode{
		Val:  -1,
		Next: nil,
	}
	return head
}

// 第一种插入方法：给链表插入一个结点，在单链表的最后加入，该方法简单
func InsertListNode(head *ListNode, newNode *ListNode) {
	// 1 先找到该链表的最后这个结点。
	// 2 创建一个辅助结点。
	temp := head
	for temp.Next != nil {

		temp = temp.Next // 让 temp 不断的指向下一个结点
	}
	// 3 将 newHeroNode 加入到链表的最后
	temp.Next = newNode
}

// 第二种插入方法：给链表插入一个结点，根据 no 的编号从小到大插入，该方法实用
func InsertListNode2(head *ListNode, newNode *ListNode) {
	// 1 找到适当的结点。
	// 2 创建一个辅助结点。
	temp := head
	flag := true
	// 让插入的结点的 no，和 temp 的下一个结点的 no 比较
	for {
		if temp.Next == nil { // 说明到链表的最后
			break
		} else if temp.Next.Val >= newNode.Val {
			// 说明 newHeroNode 就应该插入到 temp 后面
			break
		} else if temp.Next.Val == newNode.Val {
			// 链表中已经有这个 no,不用插入.
			flag = false
			break
		}
		temp = temp.Next
	}
	if !flag {
		fmt.Println("对不起，已经存在no=", newNode.Val)
		return
	} else {
		newNode.Next = temp.Next
		temp.Next = newNode
	}
}

// 删除一个结点
func DelListNode(head *ListNode, id int) {
	temp := head
	flag := false
	// 找到要删除结点的 no，和 temp 的下一个结点的 no 比较
	for {
		if temp.Next == nil { // 说明到链表的最后
			break
		} else if temp.Next.Val == id {
			// 找到了
			flag = true
			break
		}
		temp = temp.Next
	}
	if flag { // 找到,删除
		temp.Next = temp.Next.Next
	} else {
		fmt.Println("sorry, 要删除的id不存在")
	}
}

// 显示链表的所有结点信息
func PrintListNode(head *ListNode) {
	// 创建一个辅助结点
	temp := head
	// 先判断该链表是不是一个空的链表
	if temp.Next == nil {
		fmt.Println("空空如也。。。。")
		return
	}
	// 遍历这个链表
	for {
		fmt.Printf("[%d ]==>", temp.Next.Val)
		temp = temp.Next
		if temp.Next == nil {
			break
		}
	}
	fmt.Println()
}
