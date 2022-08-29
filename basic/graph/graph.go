package main

import (
	"container/heap"
	"fmt"
	"math"
)

type Node struct {
	Val int
	//入度 直接指向该点的线段个数
	In int
	//出度 从该点出发直接指向别的点的线段的个数
	Out int
	//出度指向的点
	Nexts *[]*Node
	//出度的边
	Edges *[]*Edge
}

type Edge struct {
	Weight int
	//from node
	Form *Node
	//to node
	To *Node
}

type Graph struct {
	//key nodeval value nodeptr
	Nodes map[int]*Node
	//key edgeptr
	Edges map[Edge]struct{}
}

func InitGraph() *Graph {
	var graph Graph
	graph.Nodes = make(map[int]*Node)
	graph.Edges = make(map[Edge]struct{})
	return &graph
}

func CreateGraph(matrix [][]int) *Graph {
	graph := InitGraph()
	for i := range matrix {
		weight := matrix[i][0]
		fromVal := matrix[i][1]
		toVal := matrix[i][2]
		_, fromOk := graph.Nodes[fromVal]
		if !fromOk {
			var from Node
			fromNexts := make([]*Node, 0)
			fromEdges := make([]*Edge, 0)
			from.Nexts = &fromNexts
			from.Edges = &fromEdges
			graph.Nodes[fromVal] = &from
			from.Val = fromVal
		}
		_, toOk := graph.Nodes[toVal]
		if !toOk {
			var to Node
			toNexts := make([]*Node, 0)
			toEdges := make([]*Edge, 0)
			to.Nexts = &toNexts
			to.Edges = &toEdges
			graph.Nodes[toVal] = &to
			to.Val = toVal
		}
		// from to 是局部变量需要重新赋值出来
		from := graph.Nodes[fromVal]
		to := graph.Nodes[toVal]
		from.Out++
		to.In++
		*(from.Nexts) = append(*(from.Nexts), to)
		var edge Edge
		edge.Weight = weight
		edge.Form = from
		edge.To = to
		*(from.Edges) = append(*(from.Edges), &edge)
		graph.Edges[edge] = struct{}{}
	}
	return graph
}

func bfs(start *Node) []int {
	if start == nil {
		return nil
	}
	ans := make([]int, 0)
	fifo := make([]*Node, 0)
	set := make(map[Node]struct{})
	fifo = append(fifo, start)
	set[*start] = struct{}{}
	for len(fifo) != 0 {
		temp := fifo[0]
		if len(fifo) == 0 {
			fifo = []*Node{}
		} else {
			fifo = fifo[1:]
		}
		ans = append(ans, temp.Val)
		for i := 0; i < len(*temp.Nexts); i++ {
			_, ok := set[*(*temp.Nexts)[i]]
			if !ok {
				set[*(*temp.Nexts)[i]] = struct{}{}
				fifo = append(fifo, (*temp.Nexts)[i])
			}
		}
	}
	return ans
}

func dfs1(start *Node) []int {
	if start == nil {
		return nil
	}
	set := make(map[*Node]struct{})
	ans := process1(start, &set)
	return ans
}
func process1(start *Node, set *map[*Node]struct{}) []int {
	ans := make([]int, 0)
	_, ok := (*set)[start]
	if ok {
		return nil
	} else {
		//此处一定要递归之前处理。
		ans = append(ans, start.Val)
		(*set)[start] = struct{}{}
	}
	for i := 0; i < len(*start.Nexts); i++ {
		next := process1((*start.Nexts)[i], set)
		if next != nil {
			ans = append(ans, next...)
		}
	}
	return ans
}
func dfs2(start *Node) []int {
	if start == nil {
		return nil
	}
	set := make(map[*Node]struct{})
	stack := make([]*Node, 0)
	ans := make([]int, 0)
	stack = append(stack, start)
	set[start] = struct{}{}
	for len(stack) != 0 {
		pop := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans = append(ans, pop.Val)
		for i := range *pop.Nexts {
			_, ok := set[(*pop.Nexts)[i]]
			if ok {
				continue
			} else {
				stack = append(stack, (*pop.Nexts)[i])
				set[(*pop.Nexts)[i]] = struct{}{}

			}
		}
	}
	return ans
}
func sortedtopology(graph Graph) []int {
	fifo := make([]*Node, 0)
	inSet := make(map[*Node]int)
	ans := make([]int, 0)
	for _, v := range graph.Nodes {
		inSet[v] = v.In
		if v.In == 0 {
			fifo = append(fifo, v)
		}
	}
	fmt.Println(inSet)
	for len(fifo) != 0 {
		pop := fifo[0]
		if len(fifo) == 0 {
			fifo = []*Node{}
		} else {
			fifo = fifo[1:]
		}
		ans = append(ans, pop.Val)
		for j := 0; j < len(*pop.Nexts); j++ {
			temp := (*pop.Nexts)[j]
			inSet[temp] = inSet[temp] - 1
			if inSet[temp] == 0 {
				fifo = append(fifo, temp)
			}
		}
	}
	return ans
}

type UnionNode struct {
	v interface{}
}
type Unionset struct {
	Nodes   map[UnionNode]*UnionNode
	Parents map[*UnionNode]*UnionNode
	SizeMap map[*UnionNode]int
	Size    int
}

func Init(a []interface{}) *Unionset {
	if a == nil {
		return nil
	}
	if len(a) == 0 {
		return nil
	}
	var ans Unionset
	ans.Nodes = make(map[UnionNode]*UnionNode)
	ans.Parents = make(map[*UnionNode]*UnionNode)
	ans.SizeMap = make(map[*UnionNode]int)
	for i := range a {
		temp := UnionNode{a[i]}
		ans.Nodes[temp] = &temp
		ans.Parents[&temp] = &temp
		ans.SizeMap[&temp] = 1
		ans.Size++
	}
	return &ans
}

func CreateUnion(a [][]int) *Unionset {
	var ans Unionset
	ans.Nodes = make(map[UnionNode]*UnionNode)
	ans.Parents = make(map[*UnionNode]*UnionNode)
	ans.SizeMap = make(map[*UnionNode]int)

	for i := range a {
		for j := range a[i] {
			temp := UnionNode{a[i][j]}
			_, ok := ans.Nodes[temp]
			if !ok {
				ans.Nodes[temp] = &temp
				ans.Parents[&temp] = &temp
				ans.SizeMap[&temp] = 1
				ans.Size++
			}
		}
	}
	return &ans
}

func (u *Unionset) FindFather(a UnionNode) *UnionNode {
	ValueA := u.Nodes[a]
	stack := make([]*UnionNode, 0)
	stack = append(stack, ValueA)
	for ValueA != u.Parents[ValueA] {
		ValueA = u.Parents[ValueA]
		stack = append(stack, ValueA)
	}
	for i := range stack {
		u.Parents[stack[i]] = ValueA
	}

	return ValueA
}

func (u *Unionset) IsSameSet(a, b interface{}) bool {
	ParentA := u.FindFather(UnionNode{a})
	ParentB := u.FindFather(UnionNode{b})
	return ParentA == ParentB
}
func (u *Unionset) Sets() int {
	return u.Size
}

func (u *Unionset) Union(a, b interface{}) {
	if u.IsSameSet(a, b) {
		return
	}
	u.Size--
	ParentA := u.FindFather(UnionNode{a})
	ParentB := u.FindFather(UnionNode{b})
	SizeA := u.SizeMap[ParentA]
	SizeB := u.SizeMap[ParentB]
	if SizeA > SizeB {
		u.Parents[ParentB] = ParentA
		u.SizeMap[ParentA] = SizeA + SizeB
		delete(u.SizeMap, ParentB)
	} else {
		u.Parents[ParentA] = ParentB
		u.SizeMap[ParentB] = SizeA + SizeB
		delete(u.SizeMap, ParentA)
	}
}

type IntHeap []*Edge

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i].Weight < h[j].Weight }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(*Edge))
}
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func KruskalMST(graph *Graph) map[*Edge]struct{} {
	fifo := make([]interface{}, 0)
	for _, valueNode := range graph.Nodes {
		fifo = append(fifo, valueNode)
	}
	union := Init(fifo)
	var qp IntHeap
	heap.Init(&qp)
	for keyEdge := range graph.Edges {
		heap.Push(&qp, keyEdge.Weight)
	}
	ans := make(map[*Edge]struct{})
	for qp.Len() != 0 {
		temp := heap.Pop(&qp)
		if !union.IsSameSet(temp.(*Edge).Form, temp.(*Edge).To) {
			ans[temp.(*Edge)] = struct{}{}
			union.Union(temp.(*Edge).Form, temp.(*Edge).To)
		}

	}
	return ans
}

func findRedundantConnection(edges [][]int) []int {
	var unionSet Unionset
	unionSet.Nodes = make(map[UnionNode]*UnionNode)
	unionSet.Parents = make(map[*UnionNode]*UnionNode)
	unionSet.SizeMap = make(map[*UnionNode]int)
	ans := make([]int, 0)
	for i := range edges {
		for j := range edges[i] {
			temp := UnionNode{edges[i][j]}
			_, ok := unionSet.Nodes[temp]
			if !ok {
				unionSet.Nodes[temp] = &temp
				unionSet.Parents[&temp] = &temp
				unionSet.SizeMap[&temp] = 1
				unionSet.Size++
			}
		}
		if !unionSet.IsSameSet(edges[i][0], edges[i][1]) {
			unionSet.Union(edges[i][0], edges[i][1])
		} else {
			ans = append(ans, edges[i]...)
		}
	}
	return ans
}

func Krim(graph *Graph) map[*Edge]struct{} {
	if graph == nil {
		return nil
	}
	intHeap := make(IntHeap, 0)
	heap.Init(&intHeap)
	nodes := make(map[*Node]struct{})
	ans := make(map[*Edge]struct{})
	node := graph.Nodes[0]
	nodes[node] = struct{}{}
	for temp := range *node.Edges {
		heap.Push(&intHeap, temp)
	}

	for intHeap.Len() != 0 {
		edge := heap.Pop(&intHeap)
		_, ok := nodes[edge.(Edge).To]
		if !ok {
			nodes[edge.(Edge).To] = struct{}{}
			for temp := range *node.Edges {
				heap.Push(&intHeap, temp)
			}
		}
		ans[edge.(*Edge)] = struct{}{}
	}

	return ans
}

func min(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}
func dijkstra1(from *Node) map[*Node]int {
	ans := make(map[*Node]int)
	//打过对号的点
	selectedNodes := make(map[*Node]struct{})
	ans[from] = 0
	//ans中排除selectedNodes的最小值
	minNode := GetMinNodeAndUnselectedNode(ans, selectedNodes)
	for minNode != nil {
		for i := range *minNode.Edges {
			toNode := (*minNode.Edges)[i].To
			weight := (*minNode.Edges)[i].Weight
			if val, ok := ans[toNode]; ok {
				ans[toNode] = min(val, ans[minNode]+weight)
			} else {
				ans[toNode] = ans[minNode] + weight
			}
			selectedNodes[minNode] = struct{}{}
			minNode = GetMinNodeAndUnselectedNode(ans, selectedNodes)
		}
	}
	return ans
}

func GetMinNodeAndUnselectedNode(distance map[*Node]int, selected map[*Node]struct{}) *Node {
	min := math.MaxInt
	var ans *Node
	for k, v := range distance {
		_, ok := selected[k]
		if v < min && !ok {
			ans = k
		}
	}
	return ans
}

type NodeWeight struct {
	lable  *Node
	weight int
}

type zone struct {
	nodes    []NodeWeight
	indexMap map[Node]int
	heapsize int
}

func (z zone) Len() int {
	return z.heapsize
}

func (z zone) Less(i, j int) bool {
	return z.nodes[i].weight < z.nodes[j].weight
}

func (z zone) Swap(i, j int) {
	z.nodes[i], z.nodes[j] = z.nodes[j], z.nodes[i]
	z.indexMap[*z.nodes[i].lable] = i
	z.indexMap[*z.nodes[j].lable] = j
}

func (z *zone) Push(x interface{}) {
	if len(z.nodes) == z.heapsize {
		(*z).nodes = append((*z).nodes, x.(NodeWeight))
	} else {
		(*z).nodes[z.heapsize] = x.(NodeWeight)
	}
	z.indexMap[*z.nodes[z.heapsize].lable] = z.heapsize
	z.heapsize++
}

func (z *zone) Pop() interface{} {
	(*z).heapsize--
	x := (*z).nodes[z.heapsize]
	z.indexMap[*z.nodes[z.heapsize].lable] = -1
	return x
}

func (z *zone) contains(c NodeWeight) bool {
	if _, ok := z.indexMap[*c.lable]; ok {
		return true
	} else {
		return false
	}
}

func (z *zone) getIndex(c NodeWeight) int {
	return z.indexMap[*c.lable]
}

func (z *zone) isEmpty() bool {
	if z.heapsize == 0 {
		return true
	} else {
		return false
	}
}

func djikstra2(from *Node) map[*Node]int {
	if from == nil {
		return nil
	}
	ans := make(map[*Node]int)
	var z zone
	z.nodes = make([]NodeWeight, 0)
	z.nodes = append(z.nodes, NodeWeight{from, 0})
	heap.Init(&z)
	for !z.isEmpty() {
		temp := heap.Pop(&z)
		ans[temp.(NodeWeight).lable] = temp.(NodeWeight).weight
		for i := range *temp.(NodeWeight).lable.Edges {
			AddOrUpdateOrIgnore((*temp.(NodeWeight).lable.Edges)[i].To, (*temp.(NodeWeight).lable.Edges)[i].Weight, z, temp.(NodeWeight).weight)
		}
	}
	return ans
}

func AddOrUpdateOrIgnore(toNode *Node, toWeight int, z zone, baseWeight int) {

	if index, ok := z.indexMap[*toNode]; ok {
		//update
		if index != -1 {
			if baseWeight+toWeight < z.nodes[index].weight {
				z.nodes[index].weight = baseWeight + toWeight
				heap.Init(&z)
			}
		}
		//value == -1 ignore

	} else {
		//add
		heap.Push(&z, NodeWeight{toNode, toWeight})
	}
}

func main() {
	matrix := [][]int{{1, 0, 1}, {1, 1, 2}, {1, 2, 0}}
	ans := CreateGraph(matrix)
	fmt.Println(ans)
	fmt.Println(ans.Nodes[0])
	fmt.Println(ans.Nodes[1])
	fmt.Println(ans.Nodes[2])
	result := bfs(ans.Nodes[0])
	fmt.Println(result)
	matrix2 := [][]int{{1, 0, 1}, {1, 0, 2}, {1, 0, 3}, {1, 2, 4}, {1, 1, 5}, {1, 3, 6}, {1, 4, 7}, {1, 5, 7}, {1, 6, 7}, {1, 7, 0}}
	ans2 := CreateGraph(matrix2)
	result2 := dfs1(ans2.Nodes[0])
	fmt.Println(result2)
	result3 := dfs2(ans2.Nodes[0])
	fmt.Println(result3)
	matrix3 := [][]int{{1, 0, 1}, {1, 0, 2}, {1, 0, 3}, {1, 2, 4}, {1, 1, 5}, {1, 3, 6}, {1, 4, 7}, {1, 5, 7}, {1, 6, 7}}
	ans3 := CreateGraph(matrix3)
	result4 := sortedtopology(*ans3)
	fmt.Println(result4)
	fmt.Println("=========================")
	matrix5 := [][]int{{3, 7}, {1, 4}, {2, 8}, {1, 6}, {7, 9}, {6, 10}, {1, 7}, {2, 3}, {8, 9}, {5, 9}}
	result5 := findRedundantConnection(matrix5)
	fmt.Println(result5)
}
