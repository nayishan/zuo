package bit

import (
	"math/rand"
	"testing"
	"time"
)

func generateArray(k *int, m *int) []int {

	rand.Seed(time.Now().UnixNano())
	maxK := 10
	maxM := 10
	maxCount := 10
	maxValue := 100
	minValue := -100
	*k = rand.Intn(maxK) + 1
	*m = rand.Intn(maxM) + 1
	if *k < *m {
		*k, *m = *m, *k
	}
	if *k == *m {
		*k = *k + 1
	}
	kTimes := rand.Intn(maxCount)
	array := make([]int, (*k)*kTimes+(*m))
	kmap := make(map[int]int)
	for i := 0; i < kTimes; i++ {
		kValue := rand.Intn(maxValue) - minValue
		_, isOk := kmap[kValue]
		if isOk {
			i++
			continue
		} else {
			kmap[kValue] = 1
		}
		for j := 0; j < *k; j++ {
			array[i*(*k)+j] = kValue
		}
	}
	mValue := 0
	for {
		mValue = rand.Intn(maxValue) - minValue
		_, isOk := kmap[mValue]
		if isOk {
			continue
		} else {
			break
		}
	}
	for i := 0; i < *m; i++ {
		array[kTimes*(*k)+i] = mValue
	}
	return array
}

func mapKM(a []int, k int, m int) int {
	kmMap := make(map[int]int)
	for i := 0; i < len(a); i++ {
		_, ok := kmMap[a[i]]
		if ok {
			kmMap[a[i]]++
		} else {
			kmMap[a[i]] = 1
		}
	}
	ans := -1
	for key, value := range kmMap {
		if value == m {
			ans = key
			break
		}
	}
	return ans

}

func TestKM(t *testing.T) {
	k := 0
	m := 0
	maxtimes := 10000
	for i := 0; i < maxtimes; i++ {
		array := generateArray(&k, &m)
		ans1 := KM(array, k, m)
		ans2 := mapKM(array, k, m)
		if ans1 != ans2 {
			t.Errorf("wrong!!, array:%v,ans1:%v, ans2:%v\n", array, ans1, ans2)
			break
		} else {
			t.Log("right", array, ans1, ans2)
		}
	}
}
