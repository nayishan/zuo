package main

type MyQueue struct {
	pushStack []int
	popStack  []int
}

func Constructor() MyQueue {
	return MyQueue{
		pushStack: []int{},
		popStack:  []int{},
	}

}

func (t *MyQueue) Push(x int) {
	t.pushStack = append(t.pushStack, x)
}

func (t *MyQueue) Pop() int {
	node := 0
	if len(t.popStack) == 0 {
		for len(t.pushStack) != 0 {
			N := len(t.pushStack)
			node1 := t.pushStack[N-1]
			t.pushStack = t.pushStack[:N-1]
			t.popStack = append(t.popStack, node1)
		}
	}
	if len(t.popStack) != 0 {
		N := len(t.popStack)
		node = t.popStack[N-1]
		t.popStack = t.popStack[:N-1]
	}
	return node
}

func (t *MyQueue) Peek() int {
	node := 0
	if len(t.popStack) == 0 {
		for len(t.pushStack) != 0 {
			N := len(t.pushStack)
			node1 := t.pushStack[N-1]
			t.pushStack = t.pushStack[:N-1]
			t.popStack = append(t.popStack, node1)
		}
	}
	if len(t.popStack) != 0 {
		N := len(t.popStack)
		node = t.popStack[N-1]
	}
	return node

}

func (t *MyQueue) Empty() bool {
	if len(t.popStack) == 0 && len(t.pushStack) == 0 {
		return true
	} else {
		return false
	}

}
