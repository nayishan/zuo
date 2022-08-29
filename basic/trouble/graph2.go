// https://leetcode.cn/problems/redundant-connection-ii/
package main

import "fmt"

type Node struct {
	V interface{}
}

type UnionSet struct {
	Nodes   map[Node]Node
	Parents map[Node]Node
	SizeMap map[Node]int
	size    int
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

func findRedundantDirectedConnection(edges [][]int) []int {
	var unionSet UnionSet
	unionSet.Nodes = make(map[Node]Node)
	unionSet.Parents = make(map[Node]Node)
	unionSet.SizeMap = make(map[Node]int)
	directedMap := make(map[Node]Node)
	ans1 := make([]int, 0)
	ans2 := make([]int, 0)

	for i := range edges {
		parent := Node{edges[i][0]}
		child := Node{edges[i][1]}
		//找重复点入度
		if _, ok := directedMap[child]; !ok {
			directedMap[child] = parent
			_, ok := unionSet.Nodes[child]
			if !ok {
				unionSet.Nodes[child] = child
				unionSet.Parents[child] = child
				unionSet.SizeMap[child] = 1
				unionSet.size++
			}
			_, ok = unionSet.Nodes[parent]
			if !ok {
				unionSet.Nodes[parent] = parent
				unionSet.Parents[parent] = parent
				unionSet.SizeMap[parent] = 1
				unionSet.size++
			}
		} else {
			ans1 = append(ans1, directedMap[child].V.(int))
			ans1 = append(ans1, child.V.(int))
			ans2 = append(ans2, parent.V.(int))
			ans2 = append(ans2, child.V.(int))
			fmt.Println(ans1, ans2)
		}

		//没有环就添加
		if !unionSet.IsSameSet(edges[i][0], edges[i][1]) {
			unionSet.Union(edges[i][0], edges[i][1])
		} else {
			//如果有多入度
			if len(ans1) != 0 && len(ans2) != 0 {
				//如果有环，且就是当前的edge引起的，就是它。
				// fmt.Println(parent.V, child.V)
				// if parent.V == ans2[0] && child.V == ans2[1] {
				// 	return ans2
				// }
				return ans2
				//如果不是当前的edge引起的，那么就是之前添加的parent重复的edge引起的
				// return ans1
			} else {
				ans := make([]int, 0)
				ans = append(ans, edges[i]...)
				return ans
			}
		}
	}
	//如果没有环,就是没有添加进的edge引起的。
	return ans2
}

func main() {
	edge1 := [][]int{{2, 1}, {3, 1}, {4, 2}, {1, 4}}
	edge2 := [][]int{{1, 2}, {1, 3}, {2, 3}}
	edge3 := [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 1}, {1, 5}}
	ans1 := findRedundantDirectedConnection(edge1)
	fmt.Println(ans1)
	ans2 := findRedundantDirectedConnection(edge2)
	fmt.Println(ans2)
	ans3 := findRedundantDirectedConnection(edge3)
	fmt.Println(ans3)
	edge4 := [][]int{{4, 2}, {1, 5}, {5, 2}, {4, 3}, {4, 1}}

	ans4 := findRedundantDirectedConnection(edge4)
	fmt.Println(ans4)

}
