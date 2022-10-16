package main

func c3(n int) int {
	if n < 1 {
		return 0
	}
	if n <= 3 {
		return n
	}
	base := [3][3]int{
		{1, 1, 0},
		{0, 0, 1},
		{1, 0, 0},
	}
	res := matrixPower(base, n-3)
	return 3*res[0][0] + 2*res[1][0] + res[2][0]
}

func matrixPower(m [3][3]int, p int) [3][3]int {
	//单位矩阵|1,0,0|
	//        |0,1,0|
	//        |0,0,1|
	N := len(m)
	res := [3][3]int{}
	for i := 0; i < N; i++ {
		res[i][i] = 1
	}
	t := m
	for ; p != 0; p >>= 1 {
		if (p & 1) != 0 {
			res = mutiMatrix(res, t)
		}
		t = mutiMatrix(t, t)
	}
	return res

}
func mutiMatrix(m1 [3][3]int, m2 [3][3]int) [3][3]int {
	res := [3][3]int{}
	for i := 0; i < len(m1); i++ {
		for j := 0; j < len(m2[0]); j++ {
			for k := 0; k < len(m2); k++ {
				res[i][j] += m1[j][k] * m2[k][j]
			}
		}
	}
	return res
}
