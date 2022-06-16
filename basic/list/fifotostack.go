package main

type MyStack struct {
	queueA *[]int
	help   *[]int
}

func queueInsert(queue *[]int, value int) {
	*queue = append(*queue, value)
}
func queueDelete(queue *[]int) int {
	res := (*queue)[0]
	*queue = (*queue)[1:]
	return res
}
func queueLen(queue []int) int {
	return len(queue)
}

func Constructor() MyStack {
	return MyStack{
		queueA: &([]int{}),
		help:   &([]int{}),
	}
}

func (t *MyStack) Push(x int) {
	queueInsert(t.queueA, x)
	// fmt.Println("push  lenA",queueLen(this.queueA),this.queueA,"lenB",queueLen(this.queueB),this.queueB)
}

func (t *MyStack) Pop() int {
	ans := 0
	if queueLen(*(t.queueA)) != 0 {
		for {
			if queueLen(*(t.queueA)) == 1 {
				ans = queueDelete(t.queueA)
				break
			}
			node := queueDelete(t.queueA)
			queueInsert(t.help, node)

		}
	}
	t.help, t.queueA = t.queueA, t.help
	//fmt.Println("pop  lenA",queueLen(this.queueA),this.queueA,"lenB",queueLen(this.queueB),this.queueB)
	return ans
}

func (t *MyStack) Top() int {
	ans := 0
	if queueLen(*(t.queueA)) != 0 {
		for {
			if queueLen(*(t.queueA)) == 1 {
				ans = queueDelete(t.queueA)
				queueInsert(t.help, ans)
				break
			}
			node := queueDelete(t.queueA)
			queueInsert(t.help, node)
		}
	}
	//fmt.Println("top lenA",queueLen(this.queueA),this.queueA,"lenB",queueLen(this.queueB),this.queueB)
	return ans
}

func (t *MyStack) Empty() bool {
	if queueLen(*(t.queueA)) == 0 {
		return true
	} else {
		return false
	}
}
