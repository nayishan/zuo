package bitmap

type bitmap struct {
	Val []int64
}

func (t *bitmap) init(max int) {
	t.Val = make([]int64, (max+64)>>6)
}

func (t *bitmap) add(num int) {
	t.Val[num>>6] |= (int64(1) << (num & 63))
}

func (t *bitmap) delete(num int) {
	t.Val[num>>6] &= ^(int64(1) << (num & 63))
}

func (t *bitmap) contains(num int) bool {
	if (t.Val[num>>6] & int64(1) << (num & 63)) != 0 {
		return true
	} else {
		return false
	}
}
