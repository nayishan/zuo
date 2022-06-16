package main

import "fmt"

type node struct {
	value int
}

func main() {
	var nodeMap map[node]string
	nodeMap = make(map[node]string)
	node1 := node{1}
	node2 := node{1}
	nodeMap[node1] = "hello there"
	nodeMap[node2] = "bye there"
	fmt.Println(nodeMap)
}
