package main

import "zuo/lib/linklist"

func right(gas []int, cost []int) []bool {
	if gas == nil || cost == nil || len(gas) != len(cost) {
		return nil
	}
	N := len(gas)
	arr := make([]int, N)
	ans := make([]bool, N)
	for i := 0; i < N; i++ {
		arr[i] = gas[i] - cost[i]
		ans[i] = true
	}
	for i := 0; i < N; i++ {
		temp := 0
		for j := i; j < N; j++ {
			temp += arr[i]
			if temp < 0 {
				ans[i] = false
				break
			}
		}
		for j := 0; j < i-1; j++ {
			temp += arr[i]
			if temp < 0 {
				ans[i] = false
				break
			}
		}
	}
	return ans
}

//1.计算arr的累加和。
//2.制作一个两倍长度的累加和数组，该数组可以推出从任意点出发的累加和。
//3.用窗口解

func num(gas []int, cost []int) []bool {
	if gas == nil || cost == nil || len(gas) != len(cost) {
		return nil
	}
	N := len(gas)
	arr := make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = gas[i] - cost[i]
	}
	prefix := make([]int, 2*N)
	prefix[0] = arr[0]
	for i := 0; i < N; i++ {
		prefix[i] = prefix[i-1] + arr[i]
	}
	for i := N; i < 2*N; i++ {
		prefix[i] = prefix[i-1] + arr[i-N]
	}
	L := 0
	R := 0
	qMin := linklist.Acl_fifo_new()
	temp := make([]int, N)
	for R < N {
		for qMin.Acl_size() != 0 && prefix[qMin.Acl_tail().(int)] >= prefix[R] {
			qMin.Acl_pop_back()
		}
		qMin.Acl_push_back(R)
		R++
	}

	for L < N {
		for R < 2*N {
			for qMin.Acl_size() != 0 && prefix[qMin.Acl_tail().(int)] >= prefix[R] {
				qMin.Acl_pop_back()
			}
			qMin.Acl_push_back(R)
			R++
			temp[L] = qMin.Acl_head().(int)
		}
		if qMin.Acl_head().(int) == L {
			qMin.Acl_pop_front()
		}
		L++
	}
	ans := make([]bool, N)
	for i := 0; i < N; i++ {
		temp[i] += prefix[i]
		if temp[i] < 0 {
			ans[i] = false
		} else {
			ans[i] = true
		}
	}
	return ans
}
