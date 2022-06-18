package sort

import (
	"fmt"
)

func sortArray(a []int) []int {
	if a == nil {
		return nil
	}
	if len(a) == 0 {
		return nil
	}
	process(a, 0, len(a)-1)
	return a
}

func process(a []int, L int, R int) {
	if L == R {
		return
	}
	mid := L + ((R - L) >> 1)
	process(a, L, mid)
	process(a, mid+1, R)
	merge(a, L, mid, R)
}

func merge(a []int, L int, mid int, R int) {
	temp := make([]int, R-L+1)
	p1 := L
	p2 := mid + 1
	i := 0

	for p1 <= mid && p2 <= R {
		if a[p1] <= a[p2] {
			temp[i] = a[p1]
			p1++
			i++
		} else {
			temp[i] = a[p2]
			p2++
			i++
		}
	}
	if p1 > mid {
		for p2 <= R {
			temp[i] = a[p2]
			p2++
			i++
		}
	}
	if p2 > R {
		for p1 <= mid {
			temp[i] = a[p1]
			p1++
			i++
		}
	}
	for i := 0; i < R-L+1; i++ {
		a[L+i] = temp[i]
	}
}

func MergeSort2(a []int) []int {

	if a == nil {
		return nil
	}
	N := len(a)
	if N < 2 {
		return a
	}
	step := 1
	for step < N {
		start := 0
		stop := start + 2*step - 1
		for {
			mid := start + step - 1
			fmt.Println(start, stop, mid)
			if mid >= N {
				break
			} else if stop >= N {
				stop = N - 1
				merge2(start, mid, stop, &a)
				break
			} else {
				fmt.Println("else")
				merge2(start, mid, stop, &a)
			}
			start += 2 * step
			stop += 2 * step
		}
		fmt.Println("increase step")
		step *= 2
	}
	return a
}

func merge2(left int, mid int, right int, a *[]int) {
	temp := make([]int, right-left+1)
	indexL := left
	indexR := mid + 1
	index := 0
	endL := mid
	endR := right
	fmt.Println("merge", endL, endR, "len(temp)", len(temp))
	for indexL <= endL && indexR <= endR {
		if (*a)[indexL] <= (*a)[indexR] {
			temp[index] = (*a)[indexL]
			indexL++
		} else {
			temp[index] = (*a)[indexR]
			indexR++
		}
		index++
	}
	fmt.Println(temp)
	if indexL > endL {
		copy(temp[index:], (*a)[indexR:])
	}
	if indexR > endR {
		copy(temp[index:], (*a)[indexL:])
	}
	copy((*a)[left:], temp)
}

func swap(a *[]int, m int, n int) {
	temp := (*a)[m]
	(*a)[m] = (*a)[n]
	(*a)[n] = temp
}

func twoZone(a []int) {
	index := 0
	indexL := -1
	N := len(a)
	temp := a[N-1]
	for index < N {
		if a[index] <= temp {
			swap(&a, index, indexL+1)
			index++
			indexL++
		} else {
			index++
		}
	}
}

func partition(a []int) {
	N := len(a)
	if N < 2 {
		return
	}
	lessR := -1
	moreL := N - 1
	index := 0
	for index <= moreL {
		if a[index] < a[N-1] {
			swap(&a, index, lessR+1)
			lessR++
			index++
		} else if a[index] > a[N-1] {
			swap(&a, index, moreL-1)
			moreL--
		} else {
			index++
		}
	}
	swap(&a, moreL, N-1)
}

func quickPartition(a *[]int, l int, r int) []int {
	lessR := l - 1
	moreL := r
	index := l
	for index < moreL {
		if (*a)[index] < (*a)[r] {
			swap(a, index, lessR+1)
			lessR++
			index++
			fmt.Println("<", *a)

		} else if (*a)[index] > (*a)[r] {
			swap(a, index, moreL-1)
			moreL--
			fmt.Println(">", *a)
		} else {
			index++
		}
	}
	swap(a, moreL, r)
	res := []int{lessR + 1, moreL}
	return res
}

func quickProcess(a *[]int, l int, r int) {
	if l > r {
		return
	}
	if l == r {
		return
	}
	res := quickPartition(a, l, r)
	quickProcess(a, l, res[0]-1)
	quickProcess(a, res[1]+1, r)
}

func quickSort(a []int) {
	if a == nil {
		return
	}
	N := len(a)
	if N < 2 {
		return
	}
	quickProcess(&a, 0, N-1)
}

func push(stack *[][]int, val []int) {
	*stack = append(*stack, val)
}
func pop(stack *[][]int) []int {
	val := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]
	return val
}
func isEmpty(stack *[][]int) bool {
	fmt.Println(len(*stack))
	if len(*stack) == 0 {
		return true
	} else {
		return false
	}
}
func quickSort2(a []int) {
	if a == nil {
		return
	}
	N := len(a)
	if N < 2 {
		return
	}
	stack := make([][]int, 0)
	push(&stack, []int{0, N - 1})
	index := 0
	for !isEmpty(&stack) {
		fmt.Println(stack)
		index++
		val := pop(&stack)
		res := quickPartition(&a, val[0], val[1])
		if res[0] > val[0] {
			fmt.Println("push left")
			push(&stack, []int{val[0], res[0] - 1})
		}
		if res[1] < val[1] {
			fmt.Println("push right")
			push(&stack, []int{res[1] + 1, val[1]})
		}
	}
	fmt.Println("index", index)
}

func insertSort(a []int) []int {
	if a == nil {
		return a
	}
	if len(a) < 2 {
		return a
	}

	for i := 0; i < len(a)-1; i++ {
		newIndex := i + 1
		for j := i; j >= 0; j-- {
			if a[newIndex] < a[j] {
				a[newIndex], a[j] = a[j], a[newIndex]
				newIndex = j
			} else {
				break
			}
		}
	}
	return a

}

//复习merge递归
func mergeSort(a *[]int) {
	if *a == nil || len(*a) < 2 {
		return
	}
	mergeProcess(a, 0, len(*a)-1)
}

func mergeProcess(a *[]int, L int, R int) {
	if L == R {
		return
	}
	mid := L + ((R - L) >> 1)
	mergeProcess(a, L, mid)
	mergeProcess(a, mid+1, R)
	merge3(a, L, mid, R)
}

func merge3(a *[]int, L int, mid int, R int) {
	tempL := L
	tempR := mid + 1
	tempIndex := 0
	res := make([]int, R-L+1)
	for tempL <= mid && tempR <= R {
		if (*a)[tempL] < (*a)[tempR] {
			res[tempIndex] = (*a)[tempL]
			tempIndex++
			tempL++
		} else {
			res[tempIndex] = (*a)[tempR]
			tempIndex++
			tempR++
		}
	}
	if tempL > mid {
		copy(res[tempIndex:], (*a)[tempR:])
	}
	if tempR > R {
		copy(res[tempIndex:], (*a)[tempL:])
	}
	copy((*a)[L:], res)
}

//复习merge非递归
func mergeSort4(a *[]int) {
	if *a == nil || len(*a) < 2 {
		return
	}
	mergeprocess2(a, 0, len(*a)-1)
}
func mergeprocess2(a *[]int, L int, R int) {
	step := 1
	for {
		tempL := L
		for {
			lBegin := tempL
			lEnd := tempL + step - 1
			rEnd := tempL + 2*step - 1
			if rEnd > R {
				rEnd = R
			}
			if lEnd >= R {
				break
			} else {
				merge4(a, lBegin, lEnd, rEnd)
			}
			tempL += 2 * step
		}
		if step <= (len(*a) >> 1) {
			step = 2 * step
		} else {
			break
		}
	}
}

func merge4(a *[]int, L int, mid int, R int) {
	res := make([]int, R-L+1)
	index := 0
	lBegin := L
	rBegin := mid + 1
	for rBegin <= R && lBegin <= mid {
		if (*a)[lBegin] > (*a)[rBegin] {
			res[index] = (*a)[rBegin]
			rBegin++
		} else {
			res[index] = (*a)[lBegin]
			lBegin++
		}
		index++
	}
	if lBegin > mid {
		copy(res[index:], (*a)[rBegin:])
	}
	if rBegin > R {
		copy(res[index:], (*a)[lBegin:])
	}

	copy((*a)[L:], res)
}

func jubuzuixiao(a []int) int {
	if len(a) < 2 {
		return -1
	}
	ans := 0
	if a[0] < a[1] {
		return 0
	} else if a[len(a)-1] < a[len(a)-2] {
		return len(a) - 1
	} else {
		L := 0
		R := len(a) - 1
		for L <= R {
			mid := L + ((R - L) >> 1)
			if a[mid] < a[mid+1] && a[mid] < a[mid-1] {
				ans = mid
				break
			} else if a[mid] > a[mid+1] {
				L = mid + 1
			} else {
				R = mid - 1
			}
		}
	}
	return ans
}
func decode(encoded []int, first int) []int {
	ans := make([]int, len(encoded)+1)
	ans[0] = first
	for i := 0; i < len(encoded); i++ {
		ans[i+1] = encoded[i] ^ ans[i]
	}
	return ans

}
