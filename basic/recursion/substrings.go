package main

import "fmt"

//固定参数str
//来到str[index]字符
//path已经确定的字符，不能改变了
//生成的子序列
func process1(str []byte, index int, ans []string, path string) {
	if len(str) == index {
		ans = append(ans, path)
		fmt.Println(path)
		return
	}
	//没要index位置的值
	process1(str, index+1, ans, path)
	//要了index位置的值
	path1 := path + string(str[index])
	process1(str, index+1, ans, path1)
}

func main() {
	str := "abc"
	ans := make([]string, 0)
	path := ""

	process1([]byte(str), 0, ans, path)
	fmt.Println("=================================")
	ans2 := make(map[string]struct{}, 0)
	str2 := "accc"
	process2([]byte(str2), 0, ans2, path)
	fmt.Println("=================================")
	fmt.Println(ans2)
}

func process2(str []byte, index int, ans map[string]struct{}, path string) {
	if len(str) == index {
		ans[path] = struct{}{}
		fmt.Println(path)
		return
	}
	//没要index位置的值
	process2(str, index+1, ans, path)
	//要了
	process2(str, index+1, ans, path+string(str[index]))
}
