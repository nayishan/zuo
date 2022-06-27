package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// func Push(nums *[]interface{}, val interface{}) {
//
// 	*nums = append(*nums, val)
// }
//
// func Pop(nums *[]interface{}) *interface{} {
// 	N := len(*nums)
// 	ans := (*nums)[N-1]
//
// 	*nums = (*nums)[:N-1]
// 	return &ans
//
// }
//
// func IsEmpty(nums *[]interface{}) bool {
// 	if len(*nums) == 0 {
// 		return true
// 	} else {
// 		return false
// 	}
// }
type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (t *Codec) serialize(root *TreeNode) *[]string {
	ans := make([]string, 0)
	preSerial(root, &ans)
	return &ans

}

// Deserializes your encoded data to tree.
func (t *Codec) deserialize(data *[]string) (ans *TreeNode) {
	if len(*data) == 0 {
		return ans
	}
	tempString := (*data)[0]
	*data = (*data)[1:]
	if tempString == "nil" {
		ans = nil
	} else {
		var tempNode TreeNode
		ans = &tempNode
		tempNode.Val, _ = strconv.Atoi(tempString)
		ans.Left = t.deserialize(data)
		ans.Right = t.deserialize(data)
	}
	return ans
}

func preSerial(root *TreeNode, ans *[]string) {
	if root == nil {
		*ans = append(*ans, "nil")
	} else {
		*ans = append(*ans, strconv.Itoa(root.Val))
		preSerial(root.Left, ans)
		preSerial(root.Right, ans)
	}
}

func preDeserial(ans *TreeNode, strings *[]string) {
	if len(*strings) == 0 {
		return
	}
	tempString := (*strings)[0]
	*strings = (*strings)[1:]

	if tempString == "nil" {
		ans = nil
	} else {
		var tempNode TreeNode
		ans = &tempNode
		tempNode.Val, _ = strconv.Atoi(tempString)
		preDeserial(ans.Left, strings)
		preDeserial(ans.Right, strings)
	}
}

func add(nums *[]*TreeNode, val *TreeNode) {

	*nums = append(*nums, val)
}
func poll(nums *[]*TreeNode) *TreeNode {

	ans := (*nums)[0]
	(*nums) = (*nums)[1:]
	return ans
}

func FifoIsEmpty(nums *[]*TreeNode) bool {
	if len(*nums) == 0 {
		return true
	} else {
		return false
	}
}

func levelSerial(root *TreeNode, ans *[]string) {

	if root == nil {
		return
	}
	fifo := make([]*TreeNode, 0)
	add(&fifo, root)
	*ans = append(*ans, strconv.Itoa(root.Val))
	for !FifoIsEmpty(&fifo) {
		temp := poll(&fifo)
		if temp.Left == nil {
			*ans = append(*ans, "nil")
		} else {
			*ans = append(*ans, strconv.Itoa(temp.Left.Val))
			add(&fifo, temp.Left)
		}
		if temp.Right == nil {
			*ans = append(*ans, "nil")
		} else {
			*ans = append(*ans, strconv.Itoa(temp.Right.Val))
			add(&fifo, temp.Right)
		}
	}
}
func levelDeSerial(strings *[]string) *TreeNode {
	if strings == nil || len(*strings) == 0 {
		return nil
	}
	fifo := make([]*TreeNode, 0)
	temp := (*strings)[0]
	*strings = (*strings)[1:]
	ans := TreeNode{
		Left:  nil,
		Right: nil,
	}
	ans.Val, _ = strconv.Atoi(temp)
	add(&fifo, &ans)
	for !FifoIsEmpty(&fifo) {
		temp := poll(&fifo)
		tempString := (*strings)[0]
		if tempString == "nil" {
			temp.Left = nil
		} else {
			temp.Left = &TreeNode{
				Left:  nil,
				Right: nil,
			}
			temp.Left.Val, _ = strconv.Atoi(tempString)
			add(&fifo, temp.Left)
		}
		*strings = (*strings)[1:]
		tempString = (*strings)[0]
		if tempString == "nil" {
			temp.Right = nil
		} else {
			temp.Right = &TreeNode{
				Left:  nil,
				Right: nil,
			}
			temp.Right.Val, _ = strconv.Atoi(tempString)
			add(&fifo, temp.Right)
		}
		*strings = (*strings)[1:]
	}
	return &ans
}

func main() {
	a := TreeNode{}
	c := TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}
	a.Val = 1
	a.Left = nil
	a.Right = &c
	strings := make([]string, 0)
	preSerial(&a, &strings)
	fmt.Println(strings)
	preDeserial(&a, &strings)
}

// many ways tree
type Node struct {
	Val   int
	Child *[]*Node
}

// two ways TreeNode
type NewTreeNode struct {
	Val   int
	Left  *NewTreeNode
	Right *NewTreeNode
}

func encode(root *Node) *NewTreeNode {
	if root == nil {
		return nil
	}
	ans := &NewTreeNode{
		Val:   root.Val,
		Left:  nil,
		Right: nil,
	}
	ans.Left = processEn(root)
	return ans
}

func processEn(root *Node) *NewTreeNode {
	var head *NewTreeNode
	var cur *NewTreeNode
	N := len(*root.Child)
	for i := 0; i < N; i++ {
		temp := &NewTreeNode{}
		temp.Val = (*root.Child)[i].Val
		if head == nil {
			head = temp
		} else {
			cur.Right = temp
		}
		cur = temp
		cur.Left = processEn((*root.Child)[i])
	}
	return head
}

func decode(root *NewTreeNode) *Node {
	if root == nil {
		return nil
	}

	ans := &Node{
		Val:   root.Val,
		Child: nil,
	}

	ans.Child = processDe(root.Left)
	return ans
}
func processDe(root *NewTreeNode) *[]*Node {
	ans := make([]*Node, 0)
	for root != nil {
		temp := &Node{
			Val:   root.Val,
			Child: processDe(root.Left),
		}
		ans = append(ans, temp)
		root = root.Right
	}
	return &ans

}

type NewNode struct {
	Val    int
	Left   *NewNode
	Right  *NewNode
	Parent *NewNode
}

func getSuccessorNode(node *NewNode) *NewNode {
	if node == nil {
		return nil
	}
	var ans *NewNode
	if node.Right == nil {
		for node.Parent != nil {
			if node.Parent.Right == node {
				node = node.Parent
			} else {
				ans = node.Parent
			}
		}
	} else {
		node = node.Right
		for node.Left != nil {
			node = node.Left
		}
		ans = node
	}
	return ans
}

func printAllNode(N int) {
	process(1, N, true)
}

func process(i, N int, down bool) {
	if i > N {
		return
	}
	process(i+1, N, true)
	if !down {
		fmt.Println("凸")
	} else {
		fmt.Println("凹")
	}
	process(i+1, N, false)
}

func isCBT(root *TreeNode) bool {
	if root == nil {
		return true
	}
	isLeaf := false
	fifo := make([]*TreeNode, 0)

	add(&fifo, root)
	for !FifoIsEmpty(&fifo) {
		temp := poll(&fifo)
		if temp.Right != nil && temp.Left == nil {
			return false
		}
		if isLeaf && (temp.Left != nil || temp.Right != nil) {
			return false
		}
		if temp.Left != nil {
			add(&fifo, temp.Left)
		}
		if temp.Right != nil {
			add(&fifo, temp.Right)
		}
		if temp.Left == nil || temp.Right == nil {
			isLeaf = true
		}
	}
	return true
}
