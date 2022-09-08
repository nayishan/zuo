package main

import (
	"fmt"
	"math"
)

func min(a, b int) int {
	if a >= b {
		return b
	} else {
		return a
	}
}
func minStickers(stickers []string, target string) int {
	ans := process(stickers, target)
	if ans == math.MaxInt {
		ans = -1
	}
	return ans
}

func process(stickers []string, target string) int {
	if len(target) == 0 {
		return 0
	}
	ans := math.MaxInt
	for i := range stickers {
		//取一个
		rest := minus(stickers[i], target)
		if len(rest) != len(target) {
			ans = min(ans, process(stickers, rest))
		}
	}
	//加上上一个
	if ans != math.MaxInt {
		ans += 1
	}
	return ans
}

func minus(sticker, target string) string {
	bStricker := []byte(sticker)
	bTarget := []byte(target)
	var builder [26]int
	for i := range bTarget {
		index := bTarget[i] - 'a'
		builder[index] += 1
	}
	for i := range bStricker {
		index := bStricker[i] - 'a'
		builder[index] -= 1
	}
	ans := make([]byte, 0)
	for i := range builder {
		cha := byte(builder[i] + 'a')
		for j := builder[i]; j > 0; j-- {
			ans = append(ans, cha)
		}
	}
	return string(ans)
}

func main() {
	stickers := []string{"with", "example", "science"}
	target := "thehat"
	ans1 := minStickers(stickers, target)
	fmt.Println(ans1)
	ans2 := minStickers2(stickers, target)
	fmt.Println(ans2)
	ans3 := minStickers3(stickers, target)
	fmt.Println(ans3)
}

func minStickers2(stickers []string, target string) int {
	//辅助优化 用词频表代替贴纸数组
	var count [][26]int
	for i := 0; i < len(stickers); i++ {
		row := statistics(stickers[i])
		count = append(count, row)
	}
	ans := process2(count, target)
	if ans == math.MaxInt {
		ans = -1
	}
	return ans
}

func statistics(a string) [26]int {
	var ans [26]int
	bA := []byte(a)
	for i := 0; i < len(bA); i++ {
		index := bA[i] - 'a'
		ans[index] += 1
	}
	return ans
}

//sticker当初贴纸的字符统计
func process2(stickers [][26]int, target string) int {
	if len(target) == 0 {
		return 0
	}
	//统计target词频
	tCount := statistics(target)
	ans := math.MaxInt
	for i := range stickers {
		index := target[0] - 'a'
		sticker := stickers[i]
		//最关键优化剪枝，取target的第一个元素，如果sticker没有下一个sticker
		//优化的原因是路径xyz yxz zxy..都会在最优解中出现，而实际只需要一个就够了
		if sticker[index] > 0 {
			var rest [26]int
			restTarget := make([]byte, 0)
			for j := range tCount {
				rest[j] = tCount[j] - sticker[j]
				for k := rest[j]; k > 0; k-- {
					restTarget = append(restTarget, byte(j+'a'))
				}
			}
			ans = min(ans, process2(stickers, string(restTarget)))
		}
	}
	if ans != math.MaxInt {
		ans += 1
	}
	return ans
}

func minStickers3(stickers []string, target string) int {
	//辅助优化 用词频表代替贴纸数组
	var count [][26]int
	for i := 0; i < len(stickers); i++ {
		row := statistics(stickers[i])
		count = append(count, row)
	}
	dp := make(map[string]int)
	ans := process3(count, target, dp)
	if ans == math.MaxInt {
		ans = -1
	}
	return ans
}

func process3(stickers [][26]int, target string, dp map[string]int) int {
	if _, ok := dp[target]; ok {
		return dp[target]
	}

	if len(target) == 0 {
		return 0
	}
	//统计target词频
	tCount := statistics(target)
	ans := math.MaxInt
	for i := range stickers {
		index := target[0] - 'a'
		sticker := stickers[i]
		//最关键优化剪枝，取target的第一个元素，如果sticker没有下一个sticker
		//优化的原因是路径xyz yxz zxy..都会在最优解中出现，而实际只需要一个就够了
		if sticker[index] > 0 {
			var rest [26]int
			restTarget := make([]byte, 0)
			for j := range tCount {
				rest[j] = tCount[j] - sticker[j]
				for k := rest[j]; k > 0; k-- {
					restTarget = append(restTarget, byte(j+'a'))
				}
			}
			ans = min(ans, process2(stickers, string(restTarget)))
		}
	}
	if ans != math.MaxInt {
		ans += 1
	}
	dp[target] = ans
	return ans
}
