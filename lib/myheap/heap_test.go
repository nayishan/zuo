package heap

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {

	treeMap := new(heap)
	treeMap.HeapInit(100)
	treeMap.Add(1)
	treeMap.Add(2)
	treeMap.Add(3)
	t.Log(treeMap.Peek())
	for !treeMap.IsEmpty() {
		t.Log(treeMap.Poll())
	}
	for i := 0; i < 3; i++ {
		fmt.Println(treeMap.Data[i])
	}

}
