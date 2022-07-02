package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type program struct {
	start int
	end   int
}
type meeting []program

func (p meeting) Less(i, j int) bool {
	return p[i].end < p[j].end
}

func (p meeting) Len() int {
	return len(p)
}

func (p meeting) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p meeting) Sort() {
	sort.Sort(p)
}

func bestArrange1(a meeting) int {
	if a == nil {
		return 0
	}
	if a.Len() == 0 {
		return 0
	}
	a.Sort()
	ans := make(meeting, 0)
	endTime := 0
	for i := 0; i < a.Len(); i++ {
		if a[i].start >= endTime {
			ans = append(ans, a[i])
			endTime = a[i].end
		}
		//fmt.Println("a[i]", a[i], "endTime", endTime)
	}
	// fmt.Println("arrange1", ans)
	return len(ans)

}

func bestArrange2(a meeting) int {
	if a == nil {
		return 0
	}
	if len(a) == 0 {
		return 0
	}
	temp := processArrange(a)
	// fmt.Println("a", a, "temp", temp)
	max := 0
	for i := 0; i < len(temp); i++ {
		temp := len(temp[i])
		if temp > max {
			max = temp
		}
	}
	return max
}
func processArrange(a meeting) []meeting {
	ans := make([]meeting, 0)
	if a == nil {
		return nil
	}
	if len(a) == 0 {
		return nil
	}

	for i := 0; i < len(a); i++ {
		first := a[i]
		nexts := removeMeeting(a, i)
		next := processArrange(nexts)
		if next != nil {
			for j := 0; j < len(next); j++ {
				temp := make([]program, 0)
				temp = append(temp, first)
				for k := 0; k < len(next[j]); k++ {
					temp = append(temp, next[j][k])
				}
				ans = append(ans, temp)
			}
		} else {
			temp := make([]program, 0)
			temp = append(temp, first)
			ans = append(ans, temp)
		}
	}
	return ans
}
func removeMeeting(a meeting, index int) meeting {
	temp := make(meeting, len(a)-1)
	addOne := false
	for i := 0; i < len(a); i++ {
		if i == index {
			addOne = true
		}
		if addOne {
			if i+1 < len(a) {
				temp[i] = a[i+1]
			}
		} else {
			temp[i] = a[i]
		}
	}
	ans := make(meeting, 0)
	for i := 0; i < len(temp); i++ {
		if a[index].end <= temp[i].start {
			ans = append(ans, temp[i])
		}
	}
	return ans
}
func GenerateRandomSortedMember(members, memberRoot, memberCeil int) []int {
	ans := make([]int, members)
	rest := members
	max := memberRoot
	temp := memberRoot
	for i := 0; i < members; i++ {
		sorted := true
		for {
			rand.Seed(time.Now().UnixNano())
			temp = rand.Intn(memberCeil-rest+1-max) + max
			for j := 0; j < i; j++ {
				if temp < ans[j] {
					sorted = false
				} else if temp == ans[j] {
					temp = temp + 1
				}
			}
			if sorted {
				ans[i] = temp
				break
			}
		}
		if max < temp {
			max = temp
		}
		rest--
	}
	return ans
}

func GenerateRandomArrayMemberSorted(members int, arrayLen int, memberRoot int, memberCeil int) [][]int {
	rand.Seed(time.Now().UnixNano())
	randArrayLen := rand.Intn(arrayLen)
	ans := make([][]int, randArrayLen)
	for i := 0; i < randArrayLen; i++ {
		temp := GenerateRandomSortedMember(members, memberRoot, memberCeil)
		ans[i] = temp
	}
	return ans
}

func main() {
	// a := meeting{{1, 2}, {1, 4}, {2, 3}}
	// ans1 := bestArrange1(a)
	// fmt.Println(ans1)
	// b := meeting{{1, 2}, {1, 4}, {2, 3}}
	// ans2 := processArrange(b)
	// fmt.Println(ans2)
	// val := GenerateRandomArrayMemberSorted(2, 10, 0, 23)
	// fmt.Println(val)
	for i := 0; i < 1000; i++ {
		val := GenerateRandomArrayMemberSorted(2, 20, 0, 23)
		val1 := make(meeting, len(val))
		for i := 0; i < len(val); i++ {
			val1[i].start = val[i][0]
			val1[i].end = val[i][1]
		}
		val2 := make(meeting, len(val))
		copy(val2, val1)
		ans1 := bestArrange1(val1)
		ans2 := bestArrange2(val2)
		if ans1 != ans2 {
			fmt.Println("Oops!")
		}

	}
	fmt.Println("finish!")

}
