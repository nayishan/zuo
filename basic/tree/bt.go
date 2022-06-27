package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Push(nums *[]*TreeNode, val *TreeNode) {

	*nums = append(*nums, val)
}

func Pop(nums *[]*TreeNode) *TreeNode {
	N := len(*nums)
	ans := (*nums)[N-1]

	*nums = (*nums)[:N-1]
	return ans

}

func IsEmpty(nums *[]*TreeNode) bool {
	if len(*nums) == 0 {
		return true
	} else {
		return false
	}
}

//1.添加root
//2.弹出打印
//3.有右添加右，有左添加左
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	stack := make([]*TreeNode, 0)
	ans := make([]int, 0)
	temp := root

	Push(&stack, temp)
	for !IsEmpty(&stack) {
		temp = Pop(&stack)
		ans = append(ans, temp.Val)
		if temp.Right != nil {
			Push(&stack, temp.Right)
		}
		if temp.Left != nil {
			Push(&stack, temp.Left)
		}
	}
	return ans
}

//添加root 以及其所有的左子，左子的左子...
//弹出打印，弹出的点有右子，添加右子和右子的所有左子，左子的左子....
func inorderTraversal(root *TreeNode) []int {

	if root == nil {
		return nil
	}

	stack := make([]*TreeNode, 0)
	ans := make([]int, 0)
	temp := root

	for temp != nil {
		Push(&stack, temp)
		temp = temp.Left
	}

	for !IsEmpty(&stack) {
		temp = Pop(&stack)
		ans = append(ans, temp.Val)
		if temp.Right != nil {
			temp = temp.Right
			for temp != nil {
				Push(&stack, temp)
				temp = temp.Left
			}
		}
	}
	return ans
}

//准备两个栈，添加root进栈1,
//弹出root进栈2,如果root有左子添加左子到栈1，如果有右子添加右子到栈1
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	stack1 := make([]*TreeNode, 0)
	stack2 := make([]*TreeNode, 0)
	ans := make([]int, 0)

	temp := root

	Push(&stack1, temp)

	for !IsEmpty(&stack1) {
		temp = Pop(&stack1)
		Push(&stack2, temp)
		if temp.Left != nil {
			Push(&stack1, temp.Left)
		}
		if temp.Right != nil {
			Push(&stack1, temp.Right)
		}
	}

	for !IsEmpty(&stack2) {
		temp = Pop(&stack2)
		ans = append(ans, temp.Val)
	}
	return ans
}

type SignedTreeNode struct {
	Node *TreeNode
	End  bool
}

func add(nums *[]*SignedTreeNode, val *SignedTreeNode) {

	*nums = append(*nums, val)
}
func poll(nums *[]*SignedTreeNode) *SignedTreeNode {

	ans := (*nums)[0]
	*nums = (*nums)[1:]
	return ans

}

func FifoIsEmpty(nums *[]*SignedTreeNode) bool {
	if len(*nums) == 0 {
		return true
	} else {
		return false
	}
}
func setTailEnd(nums *[]*SignedTreeNode) {
	N := len(*nums)

	(*nums)[N-1].End = true
}
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	ans := make([][]int, 0)
	levelans := make([]int, 0)
	fifo := make([]*SignedTreeNode, 0)
	temp := root
	add(&fifo, &SignedTreeNode{temp, true})
	for !FifoIsEmpty(&fifo) {
		signedTemp := poll(&fifo)

		if signedTemp.Node.Left != nil {
			add(&fifo, &SignedTreeNode{signedTemp.Node.Left, false})
		}
		if signedTemp.Node.Right != nil {
			add(&fifo, &SignedTreeNode{signedTemp.Node.Right, false})
		}
		if signedTemp.End && !FifoIsEmpty(&fifo) {
			setTailEnd(&fifo)
		}
		levelans = append(levelans, signedTemp.Node.Val)

		if signedTemp.End {
			ans = append(ans, levelans)
			levelans = []int{}
		}
	}

	return ans
}
