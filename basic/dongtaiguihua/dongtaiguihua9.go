package main

import "container/heap"

type Machine struct {
	TimePoint int
	WorkTime  int
}
type Machines []Machine

func (h Machines) Len() int { return len(h) }
func (h Machines) Less(i, j int) bool {
	return (h[i].TimePoint + h[i].WorkTime) > (h[j].TimePoint + h[j].WorkTime)
}
func (h Machines) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *Machines) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(Machine))
}
func (h *Machines) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func min(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}
func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func minTime(arr []int, n int, a, b int) int {
	m := make(Machines, 0)
	for i := 0; i < len(arr); i++ {
		mTemp := Machine{0, arr[i]}
		m = append(m, mTemp)
	}
	heap.Init(&m)
	//每个人喝完咖啡的时间
	drinks := make([]int, n)
	for i := 0; i < n; i++ {
		dTemp := Machine{
			TimePoint: 0,
			WorkTime:  0,
		}
		x := heap.Pop(&m)
		dTemp.TimePoint = x.(Machine).TimePoint + x.(Machine).WorkTime
		dTemp.WorkTime = x.(Machine).WorkTime
		heap.Push(&m, dTemp)
		drinks[i] = dTemp.TimePoint
	}
	return bestTime(drinks, a, b, 0, 0)
}

//每个人可以开始刷杯子的时间，洗杯子和挥发的时间，第几号，咖啡机什么时间能用
//drinks从index出发所有的杯子都变干净，最早的结束时间
func bestTime(drinks []int, wash, air int, index int, free int) int {
	if index == len(drinks) {
		return 0
	}
	//洗杯子
	selfClean1 := max(drinks[index], free) + wash
	restClean1 := bestTime(drinks, wash, air, index+1, selfClean1)
	p1 := max(selfClean1, restClean1)
	//挥发
	selfClean2 := drinks[index] + air
	restClean2 := bestTime(drinks, wash, air, index+1, free)
	p2 := max(selfClean2, restClean2)
	return min(p1, p2)
}
func minTime2(arr []int, n int, a, b int) int {
	m := make(Machines, 0)
	for i := 0; i < len(arr); i++ {
		mTemp := Machine{0, arr[i]}
		m = append(m, mTemp)
	}
	heap.Init(&m)
	//每个人喝完咖啡的时间
	drinks := make([]int, n)
	for i := 0; i < n; i++ {
		dTemp := Machine{
			TimePoint: 0,
			WorkTime:  0,
		}
		x := heap.Pop(&m)
		dTemp.TimePoint = x.(Machine).TimePoint + x.(Machine).WorkTime
		dTemp.WorkTime = x.(Machine).WorkTime
		heap.Push(&m, dTemp)
		drinks[i] = dTemp.TimePoint
	}
	return bestTime2(drinks, a, b)
}

//index 0~n+1
//free 0～所有人都用洗咖啡机器的值+1
func bestTime2(drinks []int, wash, air int) int {
	N := len(drinks)
	maxFree := 0
	for i := 0; i < N; i++ {
		maxFree = max(drinks[i], maxFree) + wash
	}
	//dp[N+1][maxFree+1]
	var dp [][]int
	for i := 0; i <= N; i++ {
		dp[i] = make([]int, maxFree+1)
	}
	// if index == len(drinks) {
	// 	return 0
	// }
	for j := 0; j <= maxFree; j++ {
		dp[N][j] = 0
	}
	// selfClean1 := max(drinks[index], free) + wash
	// restClean1 := bestTime(drinks, wash, air, index+1, selfClean1)
	// p1 := max(selfClean1, restClean1)
	// selfClean2 := drinks[index] + air
	// restClean2 := bestTime(drinks, wash, air, index+1, free)
	// p2 := max(selfClean2, restClean2)
	// return min(p1, p2)
	for i := N - 1; i >= 0; i-- {
		for j := 0; j <= maxFree; j++ {
			selfClean1 := max(drinks[i], j) + wash
			//本质上就是不会被递归掉到。和范围模型的L..R情况类似
			if selfClean1 > maxFree {
				continue
			}
			restClean1 := dp[i+1][selfClean1]
			p1 := max(selfClean1, restClean1)
			selfClean2 := drinks[i] + air
			restClean2 := dp[i+1][j]
			p2 := max(selfClean2, restClean2)
			dp[i][j] = min(p1, p2)
		}
	}
	return dp[0][0]
}

func main() {

}
