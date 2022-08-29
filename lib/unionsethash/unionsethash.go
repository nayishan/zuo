package unionsethash

import "fmt"

type Coordinate struct {
	X int
	Y int
}
type Node struct {
	V interface{}
}

type UnionSet struct {
	Nodes   map[Node]Node
	Parents map[Node]Node
	SizeMap map[Node]int
	size    int
}

func init() {
	fmt.Println("unionset start")
}
func Init(a []interface{}) *UnionSet {
	N := len(a)
	var union UnionSet
	union.Nodes = make(map[Node]Node, 0)
	union.Parents = make(map[Node]Node, 0)
	union.SizeMap = make(map[Node]int, 0)
	for i := 0; i < N; i++ {
		temp := Node{a[i]}
		union.Nodes[temp] = temp
		union.Parents[temp] = temp
		union.SizeMap[temp] = 1
		union.size++
	}
	return &union
}

func InitCoordinate(a [][]interface{}) *UnionSet {
	row := len(a)
	col := len(a[0])
	var union UnionSet
	union.Nodes = make(map[Node]Node)
	union.Parents = make(map[Node]Node)
	union.SizeMap = make(map[Node]int)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			temp := Node{Coordinate{i, j}}
			union.Nodes[temp] = temp
			union.Parents[temp] = temp
			union.SizeMap[temp] = 1
			if a[i][j] == '1' {
				union.size++
			}
		}
	}
	return &union
}

func (u *UnionSet) Union(a, b interface{}) {
	xA := u.FindFather(a)
	xB := u.FindFather(b)
	parentA := Node{xA}
	parentB := Node{xB}
	if parentA == parentB {
		return
	}
	u.size--
	sizeA := u.SizeMap[parentA]
	sizeB := u.SizeMap[parentB]
	if sizeA < sizeB {
		u.Parents[parentA] = parentB
		delete(u.SizeMap, parentA)
		u.SizeMap[parentB] = sizeA + sizeB
	} else {
		u.Parents[parentB] = parentA
		delete(u.SizeMap, parentB)
		u.SizeMap[parentA] = sizeA + sizeB
	}
}

//给一个节点，返回到往上不能再往上的节点
func (u *UnionSet) FindFather(a interface{}) interface{} {
	nodeA := Node{a}
	stack := make([]Node, 0)
	stack = append(stack, nodeA)
	for nodeA != u.Parents[nodeA] {
		nodeA = u.Parents[nodeA]
		stack = append(stack, nodeA)
	}
	//这里是优化，目标是把路径的所有节点指alpha
	for i := len(stack) - 1; i >= 0; i-- {
		u.Parents[stack[i]] = nodeA
	}
	return nodeA.V
}

func (u *UnionSet) IsSameSet(a, b interface{}) bool {
	return u.FindFather(a) == u.FindFather(b)
}
func (u *UnionSet) Sets() int {
	return u.size
}
