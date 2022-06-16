package kmp

func GetNext(b string) []int {
	N := len(b)
	next := make([]int, N)

	if N == 1 {
		return []int{-1}
	}
	next[0] = -1
	next[1] = 0
	maxMatch := 0
	index := 2
	for index < N {
		if b[maxMatch] == b[index-1] {
			maxMatch++
			next[index] = maxMatch
			index++
		} else if maxMatch > 0 {
			maxMatch = next[maxMatch]
		} else {
			next[index] = 0
			index++
		}
	}
	return next
}

func strStr(haystack string, needle string) int {
	M := len(haystack)
	N := len(needle)
	if N == 0 {
		return 0
	}
	if M < 1 || M < N {
		return -1
	}
	next := GetNext(needle)
	haystackIndex := 0
	maxMatch := 0
	for haystackIndex < M && maxMatch < N {
		if haystack[haystackIndex] == needle[maxMatch] {
			haystackIndex++
			maxMatch++
		} else if next[maxMatch] == -1 {
			haystackIndex++
		} else {
			maxMatch = next[maxMatch]
		}
	}
	if maxMatch == N {
		return haystackIndex - maxMatch
	} else {
		return -1
	}
}
