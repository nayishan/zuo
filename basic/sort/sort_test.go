package sort

import (
	"testing"
	"zuo/lib/duishuqi"
)

func TestMergeSort(t *testing.T) {
	maxLen := 100
	maxValue := 50
	flag := true
	for i := 0; i < 10000; i++ {
		tempArray := duishuqi.LenRandValueRand(maxLen, maxValue)
		array1 := make([]int, len(tempArray))
		array2 := make([]int, len(tempArray))
		copy(array1, tempArray)
		copy(array2, tempArray)
		mergeSort(&array1)
		duishuqi.InsertSort(&array2)
		for j := 0; j < len(tempArray); j++ {
			if array1[j] != array2[j] {
				t.Errorf(" tempArray:%v\n array1:%v\n array2:%v\n", tempArray, array1, array2)
				flag = false
				break
			}
		}
		if !flag {
			break
		}
	}
}

func TestMergeSort4(t *testing.T) {
	maxLen := 100
	maxValue := 50
	flag := true
	for i := 0; i < 10000; i++ {
		tempArray := duishuqi.LenRandValueRand(maxLen, maxValue)
		array1 := make([]int, len(tempArray))
		array2 := make([]int, len(tempArray))
		copy(array1, tempArray)
		copy(array2, tempArray)
		mergeSort4(&array1)
		duishuqi.InsertSort(&array2)
		for j := 0; j < len(tempArray); j++ {
			if array1[j] != array2[j] {
				t.Errorf("\n tempArray:%v\n array1:%v\n array2:%v\n", tempArray, array1, array2)
				flag = false
				break
			}
		}
		if !flag {
			break
		}
	}

}
func equalNum(m int, n int) bool {
	if m == n {
		return true
	} else {
		return false
	}

}
func testLocalMin(a []int, m int) bool {
	if a == nil {
		return equalNum(m, -1)
	}
	N := len(a)
	if N == 0 {
		return equalNum(m, -1)
	}
	if N == 1 {
		return equalNum(m, -1)
	}
	if m == 0 {
		if a[m] < a[m+1] {
			return true
		} else {
			return false
		}
	}
	if m == N-1 {
		if a[m] < a[m-1] {
			return true
		} else {
			return false
		}
	}
	if a[m] < a[m-1] && a[m] < a[m+1] {
		return true
	} else {
		return false
	}

}

func TestJubuzuixiao(t *testing.T) {
	maxLen := 100
	maxValue := 50
	for i := 0; i < 10000; i++ {
		tempArray := duishuqi.WuXuBuDeng(maxLen, maxValue)
		index1 := jubuzuixiao(tempArray)
		flag := testLocalMin(tempArray, index1)
		if !flag {
			t.Errorf("\n tempArray:%v index1:%v\n", tempArray, index1)
			break
		}
	}
}
