package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type AVLinfo struct {
	IsBanlanced bool
	Height      int
}

func isAVL(root *TreeNode) bool {
	return processAVL(root).IsBanlanced
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//1.总结需要的变量
//2.左树返回
//3.右树返回
//4.自己计算并返回
func processAVL(root *TreeNode) AVLinfo {
	if root == nil {
		return AVLinfo{true, 0}
	}
	leftInfo := processAVL(root.Left)
	rightInfo := processAVL(root.Right)
	var myInfo AVLinfo
	if !leftInfo.IsBanlanced {
		myInfo.IsBanlanced = false

	}
	if !rightInfo.IsBanlanced {
		myInfo.IsBanlanced = false
	}
	if leftInfo.Height-rightInfo.Height > 1 {
		myInfo.IsBanlanced = false
	}

	myInfo.Height = max(leftInfo.Height, rightInfo.Height)
	return myInfo
}

type SBTinfo struct {
	Min   int
	Max   int
	IsSBT bool
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func isValidBST(root *TreeNode) bool {
	return processSBT(root).IsSBT

}

//1. 总结需要的变量
//2. 左树返回
//3. 右树返回
//4.base case 无法返回实体值，只能返回空值，所以在计算自己的变量时
//需要对空值进行处理
func processSBT(root *TreeNode) *SBTinfo {
	if root == nil {
		return nil
	}
	leftInfo := processSBT(root.Left)
	rightInfo := processSBT(root.Right)

	myInfo := SBTinfo{
		Min:   root.Val,
		Max:   root.Val,
		IsSBT: true,
	}
	if leftInfo != nil {
		myInfo.Max = max(leftInfo.Max, myInfo.Max)
		myInfo.Min = min(leftInfo.Min, myInfo.Min)
	}
	if rightInfo != nil {
		myInfo.Max = max(rightInfo.Max, myInfo.Max)
		myInfo.Min = min(rightInfo.Min, myInfo.Min)

	}
	if leftInfo != nil && !leftInfo.IsSBT {
		myInfo.IsSBT = false

	}
	if rightInfo != nil && !rightInfo.IsSBT {
		myInfo.IsSBT = false
	}
	if leftInfo != nil && leftInfo.Max >= root.Val {
		myInfo.IsSBT = false
	}
	if rightInfo != nil && rightInfo.Min <= root.Val {
		myInfo.IsSBT = false
	}

	return &myInfo
}

type maxInfo struct {
	Height  int
	MaxPath int
}

func MaxPath(root *TreeNode) int {
	return processMax(root).MaxPath
}

//1. 总结需要的变量
//2. 左树返回
//3. 右树返回
//4. 三种情况需要考虑，左树的最大距离（不经过x）右树的最大距离（不经过x）
//左树高度 + 右树高度 +1（经过x)
func processMax(root *TreeNode) maxInfo {
	if root == nil {
		return maxInfo{0, 0}
	}
	leftInfo := processMax(root.Left)
	rightInfo := processMax(root.Right)

	var myInfo maxInfo
	m1 := leftInfo.MaxPath
	m2 := rightInfo.MaxPath
	m3 := leftInfo.Height + rightInfo.Height + 1
	myInfo.Height = max(leftInfo.Height, rightInfo.Height) + 1
	myInfo.MaxPath = max(m1, max(m2, m3))
	return myInfo
}

type CBTInfo struct {
	Size   int
	IsCBT  bool
	Height int
}

func IsCBT(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return processCBT(root).IsCBT
}
func processCBT(root *TreeNode) CBTInfo {
	if root == nil {
		return CBTInfo{0, true, 0}
	}
	leftInfo := processCBT(root.Left)
	rightInfo := processCBT(root.Right)

	size := 1
	isCBT := true
	height := 0

	size += leftInfo.Size + rightInfo.Size
	height = max(leftInfo.Height, rightInfo.Height) + 1

	if !leftInfo.IsCBT || !rightInfo.IsCBT {
		isCBT = false
	}
	if leftInfo.Height != rightInfo.Height || leftInfo.Height != rightInfo.Height+1 {
		isCBT = false
	}
	if leftInfo.Height == rightInfo.Height {

		if (1<<leftInfo.Height - 1) != leftInfo.Size {
			isCBT = false
		}
	}
	if leftInfo.Height == rightInfo.Height+1 {
		if (1<<rightInfo.Height - 1) != rightInfo.Size {
			isCBT = false
		}
	}
	return CBTInfo{size, isCBT, height}

}

type FBTInfo struct {
	Size   int
	Height int
}

func IsFBT(root *TreeNode) bool {
	if root == nil {
		return true
	}
	temp := processFBT(root)
	if 1<<temp.Height-1 != temp.Size {
		return false
	} else {
		return true
	}
}
func processFBT(root *TreeNode) FBTInfo {
	if root == nil {
		return FBTInfo{0, 0}
	}
	leftInfo := processFBT(root.Left)
	rightInfo := processFBT(root.Right)

	var myInfo FBTInfo
	myInfo.Height = max(leftInfo.Height, rightInfo.Height) + 1
	myInfo.Size = leftInfo.Size + rightInfo.Size + 1
	return myInfo
}

//当节点的全部数量等于是二叉搜索树的节点数量时，它是二叉搜索树
type SubSBTInfo struct {
	Size    int //节点的全部数量
	Min     int
	Max     int
	SubSize int //是二叉搜索树的节点的数量
}

func MaxSubSBT(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return processSubSBT(root).SubSize

}

func processSubSBT(root *TreeNode) *SubSBTInfo {
	if root == nil {
		return nil
	}

	leftInfo := processSubSBT(root.Left)
	rightInfo := processSubSBT(root.Right)

	Size := 1
	Min := root.Val
	Max := root.Val
	leftSubSize := 0

	if leftInfo != nil {
		Size += leftInfo.Size
		Min = min(leftInfo.Min, Min)
		Max = max(leftInfo.Max, Max)
		leftSubSize = leftInfo.SubSize
	}
	rightSubSize := 0
	if rightInfo != nil {
		Size += rightInfo.Size
		Min = min(rightInfo.Min, Min)
		Max = max(rightInfo.Max, Max)
		rightSubSize = rightInfo.SubSize

	}

	SubSize := max(leftSubSize, rightSubSize)
	leftIsSBT := true
	rightIsSBT := true
	myIsSBT := true
	if leftInfo != nil {
		if leftInfo.Size != leftInfo.SubSize {
			leftIsSBT = false
		}
		if !leftIsSBT || leftInfo.Max > root.Val {
			myIsSBT = false
		}
	}
	if rightInfo != nil {
		if rightInfo.Size != rightInfo.SubSize {
			rightIsSBT = false
		}
		if !rightIsSBT || rightInfo.Max > root.Val {
			myIsSBT = false
		}
	}
	if myIsSBT {
		SubSize = Size
	}
	return &SubSBTInfo{Size, Min, Max, SubSize}
}

type AncestorInfo struct {
	FindA    bool
	FindB    bool
	Ancestor *TreeNode
}

func lowestAncestor(root *TreeNode, a, b *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	return processAncestor(root, a, b).Ancestor
}
func processAncestor(root *TreeNode, a, b *TreeNode) AncestorInfo {
	if root == nil {
		return AncestorInfo{false, false, nil}
	}
	leftInfo := processAncestor(root.Left, a, b)
	rightInfo := processAncestor(root.Right, a, b)
	findA := false
	findB := false
	var ancestor *TreeNode
	ancestor = nil
	if leftInfo.FindA || rightInfo.FindA || root == a {
		findA = true
	}
	if leftInfo.FindB || rightInfo.FindB || root == b {
		findB = true
	}
	if leftInfo.Ancestor != nil {
		ancestor = leftInfo.Ancestor
	} else if rightInfo.Ancestor != nil {
		ancestor = rightInfo.Ancestor
	} else {
		if findA && findB {
			ancestor = root
		}
	}
	return AncestorInfo{findA, findB, ancestor}
}

type happyInfo struct {
	yes int
	no  int
}
type MWaysTreeNode struct {
	Happy int
	Child *[]*MWaysTreeNode
}

func maxHappy(root *MWaysTreeNode) int {
	if root == nil {
		return 0
	}
	return max(processHappy(root).yes, processHappy(root).no)
}
func processHappy(root *MWaysTreeNode) happyInfo {
	if root == nil {
		return happyInfo{0, 0}
	}

	rootYes := root.Happy
	rootNo := 0
	for i := 0; i < len(*root.Child); i++ {
		rootYes += processHappy((*root.Child)[i]).no
		rootNo += max(processHappy((*root.Child)[i]).no, processHappy((*root.Child)[i]).yes)
	}
	return happyInfo{rootYes, rootNo}
}
