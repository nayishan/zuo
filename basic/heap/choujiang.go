package main

import (
	"container/heap"
	"fmt"
	"sort"
)

/* 1. 如果某个用户购买的商品数为0，但是又发生退货事件，则认为该事件无效，得奖名
单和上一个事件发生后一致，例子中的5用户
2.某用户发生购买商品事件，购买商品数+1,发生退货事件，购买商品数-1，
3.每次都是最多k个用户得奖，k也为传入的参数，如果根据全部规则，得奖人数确实不够
k个，那就以不够的情况输出结果
4.得奖系统分为得奖区和候选区，任何用户只要购买数>0，一定在这两个区域中的一个
5.购买最大的前k名用户进入得奖区，在最初的时候，如果得奖区没有达到k个用户，那么
新来的用户直接进入得奖区
6.如果购买数不足以进入得奖区，进入候选区
7.如果候选区购买数最多的用户，足以进入得奖区，该用户就会替换得奖区中购买数最少
的用户（大于才能替换），如果得奖区中购买数最少的用户有多个，就替换最早进入得奖
区的用户，如果候选区中购买数最多的用户有多个，机会给最早进入候选区的用户
8.候选区和得奖区是两套时间，因用户只会在其中的一个区域，所以只会有一个区域的时
间，从得奖区出来进入候选区的用户，得奖区时间删除，进入候选区的时间就
是当前事件的时间。（可以理解为arr[i]和op[i]中的i），从候选区出来进入得奖区的用
户，候选区时间删除，进入得奖区的时间就是当前事件的时间（可以理解为arr[i]和op[i]
中的i)
9.如果某用户的购买数==0,不管在哪个区域都离开，区域时间删除，离开是指彻底离开，
哪个区域也不会找到该用户，如果下次该用户又发生购买行为，产生>0的购买数。会再次
根据之前的规则回到某个区域，进入区域的时间重记。
arr = [3,3,1,2,1,2,5 ...
op = [T,T,T,T,F,T,F ...
*/

type Custom struct {
	Id        int
	Buy       int
	EnterTime int
}

func compare(arr []int, op []bool, k int) [][]int {
	customMap := make(map[int]Custom)
	cands := make([]Custom, 0)
	daddy := make([]Custom, 0)
	ans := make([][]int, 0)
	for i := 0; i < len(arr); i++ {
		id := arr[i]
		buyOrRefund := op[i]
		_, ok := customMap[id]
		//noid refund
		if !buyOrRefund && !ok {
			temp := make([]int, 0)
			for i := 0; i < len(daddy); i++ {
				temp = append(temp, daddy[i].Id)
			}
			ans = append(ans, temp)
			continue
		}
		//noid buy
		//id buy
		//id refund
		if !ok {
			customMap[id] = Custom{id, 0, 0}
		}
		c := customMap[id]
		if buyOrRefund {
			c.Buy++
		} else {
			c.Buy--
		}
		needClean := false
		if c.Buy == 0 {
			delete(customMap, id)
			needClean = true
		}
		inCands := false
		inDaddy := false
		for i := 0; i < len(cands); i++ {
			if id == cands[i].Id {
				inCands = true
			}
		}
		for i := 0; i < len(daddy); i++ {
			if id == daddy[i].Id {
				inDaddy = true
			}
		}
		c.EnterTime = i
		customMap[id] = Custom{id, c.Buy, c.EnterTime}
		if !inCands && !inDaddy {
			if len(daddy) <= k {
				daddy = append(daddy, Custom{id, c.Buy, c.EnterTime})
			} else {
				cands = append(cands, Custom{id, c.Buy, c.EnterTime})
			}
		}
		if needClean {
			cleanBuyZero(&daddy, &cands, id)
		}
		sortZone(&cands)
		sortZone(&daddy)
		move(&daddy, &cands, k, i)
		temp := make([]int, 0)
		for i := 0; i < len(daddy); i++ {
			temp = append(temp, daddy[i].Id)
		}
		ans = append(ans, temp)

	}
	return ans

}
func cleanBuyZero(daddy, cands *[]Custom, id int) {
	needClean := false
	d := *daddy
	index := 0
	for i := 0; i < len(d); i++ {
		if d[i].Id == id {
			needClean = true
		}
	}
	if needClean {
		d[index], d[len(d)-1] = d[len(d)-1], d[index]
		d = d[:len(d)-1]
	} else {
		c := *cands
		for i := 0; i < len(c); i++ {
			if c[i].Id == id {
				needClean = true
			}
		}
		if needClean {
			c[index], c[len(c)-1] = c[len(c)-1], c[index]
			c = c[:len(c)-1]
		}

	}
}

type IntSlice []Custom

func (s IntSlice) Len() int { return len(s) }
func (s IntSlice) Less(i, j int) bool {
	if s[i].Buy < s[j].Buy {
		return true
	} else if s[i].Buy > s[j].Buy {
		return false
	} else {
		if s[i].EnterTime < s[j].EnterTime {
			return true
		} else {
			return false
		}
	}
}
func (s IntSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func sortZone(zone *[]Custom) {
	z := *zone
	sort.Sort(IntSlice(z))
}

func move(daddy, cands *[]Custom, k int, enterTime int) {
	d := *daddy
	c := *cands
	if len(d) < k {
		if len(c) > 0 {
			d = append(d, c[len(c)-1])
			c = c[:len(c)-1]
			d[len(d)-1].EnterTime = enterTime
		}
	} else if len(*daddy) == k {
		if len(c) > 0 {
			if c[len(c)-1].Buy > d[0].Buy {
				c[len(c)-1], d[0] = d[0], c[len(c)-1]
			}
			d[0].EnterTime = enterTime
			c[len(c)-1].EnterTime = enterTime
		}
	} else {

	}
}

type zone struct {
	custom   []Custom
	indexMap map[Custom]int
	heapsize int
}

func (z zone) Len() int {
	return z.heapsize
}

func (z zone) Less(i, j int) bool {
	if z.custom[i].Buy != z.custom[j].Buy {
		return z.custom[i].Buy < z.custom[j].Buy
	} else {
		return z.custom[i].EnterTime < z.custom[j].EnterTime
	}
}

func (z zone) Swap(i, j int) {
	z.custom[i], z.custom[j] = z.custom[j], z.custom[i]
	z.indexMap[z.custom[i]] = i
	z.indexMap[z.custom[j]] = j
}

func (z *zone) Push(x interface{}) {
	if len(z.custom) == z.heapsize {
		(*z).custom = append((*z).custom, x.(Custom))
	} else {
		(*z).custom[z.heapsize] = x.(Custom)
	}
	z.indexMap[z.custom[z.heapsize]] = z.heapsize
	z.heapsize++
}

func (z *zone) Pop() interface{} {
	(*z).heapsize--
	x := (*z).custom[z.heapsize]
	delete(z.indexMap, z.custom[z.heapsize])
	return x
}
func (z *zone) contains(c Custom) bool {
	if _, ok := z.indexMap[c]; ok {
		return true
	} else {
		return false
	}
}
func (z *zone) getIndex(c Custom) int {
	return z.indexMap[c]
}
func (z *zone) isEmpty() bool {
	if z.heapsize == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	z := zone{
		custom:   []Custom{},
		indexMap: map[Custom]int{},
		heapsize: 0,
	}
	heap.Init(&z)
	heap.Push(&z, Custom{0, 0, 0})
	heap.Push(&z, Custom{0, 0, 1})
	for !z.isEmpty() {
		fmt.Println(heap.Pop(&z))

	}
}
