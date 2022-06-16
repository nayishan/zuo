package main

func countRangeSum(a []int, upper int, lower int) int {
	if a == nil {
		return 0
	}
	if len(a) == 0 {
		return 0
	}
	sum := presum(a)
	ans := process(sum, 0, len(a)-1, upper, lower)
	return ans
}

func presum(a []int) []int {
	sum := make([]int, len(a))
	sum[0] = a[0]
	for i := 1; i < len(a); i++ {
		sum[i] = sum[i-1] + a[i]
	}
	return sum
}

func process(sum []int, l int, r int, upper int, lower int) int {
	if l == r {
		if sum[l] >= lower && sum[l] <= upper {
			return 1
		} else {

			return 0

		}
	}
	mid := l + ((r - l) >> 1)
	num1 := process(sum, l, mid, upper, lower)
	num2 := process(sum, mid+1, r, upper, lower)
	num3 := merge(sum, l, mid, r, upper, lower)
	return num1 + num2 + num3
}

func merge(sum []int, l, mid, r, upper, lower int) int {
	ans := 0
	windowL := l
	windowR := l
	for i := mid + 1; i <= r; i++ {
		newupper, newlower := edge(sum[i], upper, lower)
		for x := windowL; x <= mid; x++ {
			if sum[x] >= newlower {
				break
			} else {
				windowL++
			}
		}
		for y := windowR; y <= mid; y++ {
			if sum[y] > newupper {
				break
			} else {
				windowR++
			}
		}
		if windowR > windowL {
			ans += (windowR - windowL)
		}
	}

	tempL := l
	tempR := mid + 1
	help := make([]int, r-l+1)
	tempH := 0
	for tempL <= mid && tempR <= r {
		if sum[tempL] < sum[tempR] {
			help[tempH] = sum[tempL]
			tempL++
		} else {
			help[tempH] = sum[tempR]
			tempR++
		}
		tempH++
	}
	if tempL > mid {
		copy(help[tempH:], sum[tempR:])
	}
	if tempR > r {
		copy(help[tempH:], sum[tempL:])
	}
	copy(sum[l:], help)
	return ans
}

func edge(num, upper, lower int) (int, int) {
	newupper := num - lower
	newlower := num - upper
	return newupper, newlower
}
