package fifo

//Push Pop IsEmpty Size
var fifo []int

func init() {
	fifo = make([]int, 0)
}

func Push(val int) {
	fifo = append(fifo, val)
}

func Pop() (int, bool) {
	N := len(fifo)
	if N == 0 {
		return 0, false
	}
	ans := fifo[0]
	fifo = fifo[1:]
	return ans, true
}
func IsEmpty() bool {
	N := len(fifo)
	if N == 0 {
		return true
	} else {
		return false
	}
}

func Size() int {
	return len(fifo)
}
