package radixsort

import (
	"testing"
	"zuo/lib/duishuqi"
)

func mergeSort(a []int) []int {
	if a == nil {
		return a
	}
	if len(a) < 2 {
		return a
	}

	process(&a, 0, len(a)-1)
	return a
}

func process(a *[]int, l, r int) {
	if r <= l {
		return
	}
	mid := l + (r-l)>>1
	process(a, l, mid)
	process(a, mid+1, r)
	merge(a, l, r, mid)
}

func merge(a *[]int, l, r, mid int) {
	tempL := l
	tempR := mid + 1
	helper := make([]int, r-l+1)
	index := 0

	for tempL <= mid && tempR <= r {
		if (*a)[tempL] <= (*a)[tempR] {
			helper[index] = (*a)[tempL]
			tempL++
		} else {
			helper[index] = (*a)[tempR]
			tempR++
		}
		index++
	}

	if tempL > mid {
		copy(helper[index:], (*a)[tempR:])
	}
	if tempR > r {
		copy(helper[index:], (*a)[tempL:])
	}
	copy((*a)[l:], helper)
}

func TestArraySort(t *testing.T) {
	flag := false
	for num := 0; num < 100; num++ {
		a := duishuqi.LenRandValueRand(100, 1000)
		b := make([]int, len(a))
		c := make([]int, len(a))
		copy(b, a)
		copy(c, a)
		mergeSort(a)
		arraySort(b)
		arraySort2(c)
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				t.Error("error", "merge sort", a, "radixsort", b)
				flag = true
				break
			}
			if a[i] != c[i] {
				t.Error("error", "merge sort", a, "RadixSort2", c)
			}
		}
		if flag {
			break
		}
	}
}

func TestGetBit(t *testing.T) {
	a := 100
	b := [3]int{0, 0, 1}
	m := maxBits(a)
	if m != 3 {
		t.Error("max bit", m)

	}
	for i := 1; i <= maxBits(a); i++ {
		bit := getBit(a, i)
		if bit != b[i-1] {
			t.Error("bit", bit, "should", b[i-1])
		}
	}
	c := 1
	d := [1]int{1}
	for i := 1; i <= maxBits(c); i++ {
		bit := getBit(c, i)
		if bit != d[i-1] {
			t.Error("bit", bit, "should", d[i-1])
		}
	}
}

func TestRadixSort(t *testing.T) {
	a := []int{3, 2, 1}
	b := []int{1, 2, 3}
	arraySort(a)
	for i := 0; i < 3; i++ {
		if a[i] != b[i] {
			t.Error("error", "a[i]", a[i], "b[i]", b[i])
		}

	}

}

func TestRadixSort2(t *testing.T) {
	a := []int{3, 2, 1}
	b := []int{1, 2, 3}
	arraySort2(a)
	for i := 0; i < 3; i++ {
		if a[i] != b[i] {
			t.Error("error", "a[i]", a[i], "b[i]", b[i])
		}
	}

}
