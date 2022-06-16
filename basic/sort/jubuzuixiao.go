package main

import "fmt"
import "zuo/duishuqi"

func jubuzuixiao(a []int) int {
	INVAL := -1
	if a == nil {
		return INVAL
	}
	N := len(a)
	if N == 0 {
		return INVAL
	}
	if N == 1 {
		return 0
	}
	//非单递减
	if a[0] < a[1] {
		return 0
	}
	//非单递增
	if a[N-1] < a[N-2] {
		return N - 1
	}
	L := 0
	R := N - 1
	//此时只剩一种情况，该数组一定是递减转化为递增
	for L < R -1 {
		mid := (L + R) / 2
		// a[i-1]<a[i]   a[i]<a[i+1]
		if a[mid] < a[mid+1] && a[mid] < a[mid-1] {
			return mid
		} else if a[mid] > a[mid+1] { //a[i] > a[i+1] remove all left
			L = mid + 1
		} else { //remove all right
			R = mid - 1
		}
	}
	if a[L]< a[R]{
		return L
	}else{
		return R
	}
}
func equalNum(m int , n int)bool{
	if m == n {
		return true
	}else{
		return false
	}

}
func testLocalMin(a[]int,m int) bool{
	if a ==  nil {
		return equalNum(m,-1)
	}
	N := len(a)
	if(N == 0){
		return equalNum(m,-1)
	}
	if N ==1{
		return equalNum(m,0)
	}
	if m ==0 {
		if a[m] < a[m+1]{
			return true
		}else{
			return false
		}
	}
	if m == N-1{
		if a[m] <a[m-1]{
			return true
		}else{
			return false
		}
	}
	if a[m] < a[m-1] && a[m] < a[m+1]{
		return true
	}else{
		return false
	}

}
func main(){
	tryTimes := 100000
	maxLen := 30
	maxValue := 100
	for i:=0;i<tryTimes;i++{
		a := duishuqi.WuXuBuDeng(maxLen,maxValue)

		localMin := jubuzuixiao(a)
		flag := testLocalMin(a,localMin)
		if !flag{
			fmt.Println("jubuzuixiao is wrong !!!!")
			fmt.Println(a)
			fmt.Println("jubuzuixiao index",localMin," value ",a[localMin])
			break
		}
	}

}
