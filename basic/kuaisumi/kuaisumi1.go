package main

// O(log(n))
func f3(n int) int {
	if n < 1 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	base := [2][2]int{{1, 1}, {1, 0}}
	res := matrixPower(base, n-2)
	//|1,1| * |x ,y|
	//        |z ,t|
	return res[0][0] + res[1][0]
}
func matrixPower(m [2][2]int, p int) [2][2]int {
	//单位矩阵|1,0|
	//        |0,1|
	res := [2][2]int{{1, 0}, {0, 1}}
	t := m
	for ; p != 0; p >>= 1 {
		if (p & 1) != 0 {
			res = mutiMatrix(res, t)
		}
		t = mutiMatrix(t, t)
	}
	return res

}
func mutiMatrix(m1 [2][2]int, m2 [2][2]int) [2][2]int {
	// m1|a,b| * m2|x,y|
	//   |c,d|     |z,t|
	res := [2][2]int{}
	for i := 0; i < len(m1); i++ {
		for j := 0; j < len(m2[0]); j++ {
			for k := 0; k < len(m2); k++ {
				res[i][j] += m1[j][k] * m2[k][j]
			}
		}
	}
	return res

}
