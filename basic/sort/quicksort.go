package main

func sortArray(a []int) []int {
	if a == nil {
		return a
	}
	if len(a) < 2 {
		return a
	}
	process(&a, 0, len(a)-1)
	return a
}

func process(a *[]int, l int, r int) {
	if r <= l {
		return
	}

	l1, r1 := netherland(a, l, r)
	process(a, l, l1-1)
	process(a, r1+1, r)

}

func netherland(a *[]int, l int, r int) (int, int) {
	b := *a
	left, right, index := l-1, r, l
	for index < right {
		if b[index] > b[r] {
			right--
			b[index], b[right] = b[right], b[index]
		} else if b[index] < b[r] {
			left++
			b[index], b[left] = b[left], b[index]
			index++
		} else {
			index++
		}
	}

	b[r], b[right] = b[right], b[r]
	return left + 1, right
}
