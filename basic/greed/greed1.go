package main

import (
	"container/heap"
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func GenerateRandomByte() byte {
	rand.Seed(time.Now().UnixNano())
	ans := rand.Intn('z'-'a'+1) + 'a'
	return byte(ans)
}

//给定一组字符串，求该组字符串拼接后字典序最小的组合。
func GenerateRandomString(arrayLen, stringLen int) []string {
	rand.Seed(time.Now().UnixNano())
	randArrayLen := rand.Intn(arrayLen)
	ans := make([]string, randArrayLen)
	for i := 0; i < randArrayLen; i++ {
		randStringLen := 0
		for randStringLen == 0 {
			randStringLen = rand.Intn(stringLen)
		}
		temp := make([]byte, randStringLen)
		for j := 0; j < randStringLen; j++ {
			temp[j] = GenerateRandomByte()
		}
		ans[i] = string(temp)
	}
	return ans
}

func lowestDictionary1(a []string) string {
	if a == nil {
		return ""
	}
	if len(a) == 0 {
		return ""
	}
	ans := processLD1(a)
	sort.Strings(ans)
	return ans[0]

}

//["av","bv","ac"] ,1 -> ["av","ac"]
func removeIndex(array []string, index int) []string {
	N := len(array)
	ans := make([]string, N-1)
	addOne := false
	for i := 0; i < len(ans); i++ {
		if i == index {
			addOne = true
		}
		if addOne {
<<<<<<< HEAD
			ans[i] = array[i+1]
=======
			if i+1 < len(array) {
				ans[i] = array[i+1]
			}
>>>>>>> b1b5096 (add generateRandomString)
		} else {
			ans[i] = array[i]
		}
	}
	return ans

}
func processLD1(array []string) []string {
	//只有最外层的ans才是最终的全局变量
	ans := make([]string, 0)
	if len(array) == 0 {
		//如果不是空字符的话，那么最底层的ans将无法做append，导致所有的
<<<<<<< HEAD
		//上层都无法append
		ans = append(ans, "")
=======
		//上层都无法append,加入“”之后len是1
		ans = append(ans, "")
		// fmt.Println(len(ans))
>>>>>>> b1b5096 (add generateRandomString)
		return ans
	}
	for i := range array {
		//为递归的输入参数做准备
		first := array[i]
		nexts := removeIndex(array, i)
		next := processLD1(nexts)
		for j := range next {
			temp := first + next[j]
			ans = append(ans, temp)
		}
	}
	return ans
}

type strings []string

func (s strings) Len() int {
	return len(s)
}

func (s strings) Less(i, j int) bool {
	ans := s[i]+s[j] < s[j]+s[i]
	return ans
}

func (s *strings) Swap(i, j int) {
	(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
}

func (s *strings) Push(x interface{}) {
	*s = append(*s, x.(string))
}

func (s *strings) Pop() interface{} {
	ans := (*s)[len(*s)-1]
	if len(*s) == 1 {
		*s = strings{}
	} else {
		(*s) = (*s)[:len(*s)-1]
	}
	return ans
}

func main() {
	for i := 0; i < 10; i++ {
		val := GenerateRandomString(10, 4)
		ans1 := lowestDictionary1(val)
		temp := strings(val)
		heap.Init(&temp)
		var ans2 string
		for temp.Len() > 0 {
			tempString := heap.Pop(&temp)
			ans2 = ans2 + tempString.(string)
		}
		if ans1 != ans2 {
			fmt.Println("Opps!", "heap", ans2, "recurse", ans1)
		}
	}
	fmt.Println("finish!")

}
