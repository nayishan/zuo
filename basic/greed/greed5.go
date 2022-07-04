package main

func light(a []byte) int {

	if a == nil {
		return 0
	}
	if len(a) == 0 {
		return 0
	}
	i := 0
	ans := 0
	for i < len(a) {
		if a[i] == 'x' {
			i++
		} else {
			ans++
			if i+1 < len(a) {
				if a[i+1] == 'x' {
					i = i + 2
				} else {
					i = i + 3
				}
			} else {
				break
			}
		}
	}
	return ans
}
