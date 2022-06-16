package tree

import (
	"math"

	"golang.org/x/tools/go/analysis/passes/nilfunc"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	M := (p == nil)
	N := (q == nil)
	if M != N {
		return false

	}
	if M && N {
		return true
	}
	return (p.Val == q.Val) && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func isMirrorTree(p *TreeNode, q *TreeNode) bool {
	M := (p == nil)
	N := (q == nil)
	if M != N {
		return false

	}
	if M && N {
		return true
	}
	return (p.Val == q.Val) && isMirrorTree(p.Left, q.Right) && isMirrorTree(p.Right, q.Left)

}
func max(n1 int, n2 int) int {
	if n1 > n2 {
		return n1
	} else {
		return n2
	}
}
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1

}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if preorder == nil || inorder == nil || len(preorder) != len(inorder) {
		return nil
	}
	valueKey := make(map[int]int, len(inorder))
	for i := 0; i < len(inorder); i++ {
		valueKey[inorder[i]] = i
	}
	return f(preorder, 0, len(preorder)-1, inorder, 0, len(inorder)-1, valueKey)
}

func f(a []int, l1 int, r1 int, b []int, l2 int, r2 int, valueKey map[int]int) *TreeNode {
	if l1 > r1 {
		return nil
	}
	node := new(TreeNode)
	if l1 == r1 {
		node.Val = a[l1]
		return node
	}
	find := valueKey[a[0]]
	node.Left = f(a, l1+1, l1+find-l2, b, 0, find-l2-1, valueKey)
	node.Right = f(a, l1+find-l2+1, r1, b, find+1, r2, valueKey)
	return node
}

/*
func buildTree(preorder []int, inorder []int) *TreeNode {
	if preorder == nil || inorder == nil || len(preorder) != len(inorder) || len(preorder) == 0{
		return nil
	}
    find:= 0
	for find = 0; find < len(preorder); find++ {
		if inorder[find] == preorder[0]{
            break
        }
	}
    root := &TreeNode{Val: preorder[0]}
	root.Left = buildTree(preorder[1:find+1], inorder[:find])
	root.Right = buildTree(preorder[find+1:], inorder[find+1:])

    return root
}*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func pop(stack *[]*TreeNode) *TreeNode {
	temp := (*stack)[0]
	*stack = (*stack)[1:]
	return temp
}
func push(stack *[]*TreeNode, node *TreeNode) {
	*stack = append(*stack, node)
}
func levelOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	ans := make([]int, 0)
	fifoA := make([]*TreeNode, 0)
	fifoB := make([]*TreeNode, 0)
	flag := 0
	push(&fifoA, root)
	for len(fifoA) != 0 || len(fifoB) != 0 {
		if flag == 0 {
			for len(fifoA) != 0 {
				node := pop(&fifoA)
				ans = append(ans, node.Val)
				if node.Left != nil {
					push(&fifoB, node.Left)
				}
				if node.Right != nil {
					push(&fifoB, node.Right)
				}
			}
			flag = 1
		} else {
			for len(fifoB) != 0 {
				node := pop(&fifoB)
				ans = append(ans, node.Val)
				if node.Left != nil {
					push(&fifoA, node.Left)
				}
				if node.Right != nil {
					push(&fifoA, node.Right)
				}
			}
			flag = 0
		}
	}
	return ans
}

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	ans := make([]int, 0)
	fifoA := make([]*TreeNode, 0)
	fifoB := make([]*TreeNode, 0)
	size := make([]int, 0)

	flag := 0
	push(&fifoA, root)
	for len(fifoA) != 0 || len(fifoB) != 0 {
		if flag == 0 {
			count := 0
			for len(fifoA) != 0 {
				count++
				node := pop(&fifoA)
				ans = append(ans, node.Val)
				if node.Left != nil {
					push(&fifoB, node.Left)
				}
				if node.Right != nil {
					push(&fifoB, node.Right)
				}
			}
			flag = 1
			size = append(size, count)
		} else {
			count := 0
			for len(fifoB) != 0 {
				count++
				node := pop(&fifoB)
				ans = append(ans, node.Val)
				if node.Left != nil {
					push(&fifoA, node.Left)
				}
				if node.Right != nil {
					push(&fifoA, node.Right)
				}
			}
			flag = 0
			size = append(size, count)
		}
	}
	return pack(ans, size)

}
func pack(ans []int, size []int) [][]int {
	N := len(size)
	res := make([][]int, N)
	for i := N - 1; i >= 0; i-- {
		M := len(ans)
		res = append(res, ans[M-size[i]:])
		ans = ans[:M-size[i]]
	}
	return res
}

type info struct {
	balanced bool
	hight    int
}

func isBalanced(root *TreeNode) bool {
	return process(root).balanced
}

func Inscope(a int, b int) bool {
	if a == max(a, b) {
		return a-b < 2
	} else {
		return b-a < 2
	}
}
func max2(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func process(root *TreeNode) *info {
	if root == nil {
		return &info{balanced: true, hight: 0}
	}
	infoTemp := new(info)
	leftInfo := process(root.Left)
	rightInfo := process(root.Right)
	infoTemp.hight = max2(leftInfo.hight, rightInfo.hight) + 1
	if leftInfo.balanced && rightInfo.balanced && Inscope(leftInfo.hight, rightInfo.hight) {
		infoTemp.balanced = true
	}
	return infoTemp
}

type vnode struct {
	isBST bool
	min   int
	max   int
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return process2(root).isBST

}
func min(a int, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func process2(root *TreeNode) *vnode {
	if root == nil {
		return nil
	}
	tempInfo := vnode{isBST: true, min: root.Val, max: root.Val}
	leftInfo := process2(root.Left)
	rightInfo := process2(root.Right)
	if leftInfo != nil {
		tempInfo.min = min(leftInfo.min, tempInfo.min)
		tempInfo.max = max(leftInfo.max, tempInfo.max)
	}
	if rightInfo != nil {
		tempInfo.min = min(rightInfo.min, tempInfo.min)
		tempInfo.max = max(rightInfo.max, tempInfo.max)
	}
	if leftInfo != nil && !leftInfo.isBST {
		tempInfo.isBST = false
	}
	if rightInfo != nil && !rightInfo.isBST {
		tempInfo.isBST = false
	}
	if leftInfo != nil && leftInfo.max >= root.Val {
		tempInfo.isBST = false
	}
	if rightInfo != nil && rightInfo.min <= root.Val {
		tempInfo.isBST = false
	}

	return &tempInfo
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		if targetSum != 0 {
			return false
		} else {
			return true
		}
	}
	path := make([][]int, 0)
	isTarget := false
	process3(root, 0, targetSum, &isTarget, path)
	return isTarget
}
func process3(root *TreeNode, preSum int, targetSum int, isTarget *bool) {
	if root.Left == nil && root.Right == nil {
		if preSum+root.Val == targetSum {
			*isTarget = true
		}
	}
	preSum += root.Val
	if root.Left != nil {
		process3(root.Left, preSum, targetSum, isTarget)
	}
	if root.Right != nil {
		process3(root.Right, preSum, targetSum, isTarget)
	}
}

func CopyPath(path []int, val int, ans *[][]int) {
	path = append(path, val)
	*ans = append(*ans, path)
}
func process4(root *TreeNode, path []int, preSum int, targetSum int, ans *[][]int) {
	if root.Left == nil && root.Right == nil {
		if preSum+root.Val == targetSum {
			CopyPath(path, root.Val, ans)
		}
		return
	}
	path = append(path, root.Val)
	preSum += root.Val
	if root.Left != nil {
		process4(root.Left, path, preSum, targetSum, ans)
	}
	if root.Right != nil {
		process4(root.Right, path, preSum, targetSum, ans)
	}
	path = path[:len(path)-1]
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return nil
	}
	path := make([]int, 0)
	ans := make([][]int, 0)
	process4(root, path, 0, targetSum, &ans)
	return ans
}
