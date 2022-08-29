package main

import (
	"fmt"
)

func remove(s []byte, i int) []byte {
	N := len(s) - 1
	ans := make([]byte, N)
	tempIndex := 0
	for index := 0; index < N; index++ {
		if tempIndex == i {
			tempIndex++
		}
		ans[index] = s[tempIndex]
		tempIndex++
	}
	return ans
}

func process(s []byte, path string, ans *[]string) {
	if len(s) == 0 {
		*ans = append(*ans, path)
		fmt.Println(path)
		return
	}
	for i := 0; i < len(s); i++ {
		temp := remove(s, i)
		process(temp, path+string(s[i]), ans)
	}
}

func main() {
	str := "abcca"
	ans := make([]string, 0)
	ans2 := make([]string, 0)
	ans3 := make([]string, 0)
	path := ""
	process([]byte(str), path, &ans)
	fmt.Println("=========================")
	process2([]byte(str), 0, &ans2)
	fmt.Println("=========================")
	process3([]byte(str), 0, &ans3)

	fmt.Println(len(ans), len(ans2), len(ans3))
}
func swap(s []byte, i, j int) {
	s[i], s[j] = s[j], s[i]
}

func process2(s []byte, index int, ans *[]string) {
	if len(s) == index {
		*ans = append(*ans, string(s))
		fmt.Println(string(s))
		return
	}
	for i := index; i < len(s); i++ {
		swap(s, index, i)
		process2(s, index+1, ans)
		swap(s, index, i)
	}
}

func process3(s []byte, index int, ans *[]string) {
	if len(s) == index {
		*ans = append(*ans, string(s))
		fmt.Println(string(s))
		return
	}
	var visited = [256]bool{}
	for i := index; i < len(s); i++ {
		if !visited[s[i]] {
			visited[s[i]] = true
			swap(s, index, i)
			process3(s, index+1, ans)
			swap(s, index, i)

		}
	}

}
