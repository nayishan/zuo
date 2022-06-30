package heapgreater

import "testing"

type testData struct {
	value int
	index int
}

func (t *heapgreater) Less(index1, index2 int) bool {
	if t.Value[index1].(testData).value > t.Value[index2].(testData).value {
		return true
	} else {
		return false
	}
}

func TestHeapgreater(t *testing.T) {
	testheap := new(heapgreater)
	testheap.HeapInit()
	t.Log("hello")
	testheap.Add(testData{0, 0})
	testheap.Add(testData{2, 2})
	testheap.Add(testData{1, 1})
	testheap.Add(testData{1, 1})
	t.Log(testheap.IndexMap)
	t.Log(testheap.Poll())
	t.Log(testheap.Poll())
	t.Log(testheap.IndexMap)
	testheap.Add(testData{3, 3})
	t.Log(testheap.Value)
	for !testheap.IsEmpty() {
		t.Log(testheap.Poll())
	}

}
