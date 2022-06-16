package bit

//假设这两个数是a,b
//得出a^b
func aXorB(a []int) int {
	ans := a[0]
	for i := 1; i < len(a); i++ {
		ans = ans ^ a[i]
	}
	return ans
}

//取ans的最右侧为1的数
func rightOne(a int) int {
	return a & (-a)
}

//异或为1一定是a 该位为0  b该位为1
//根据上面的数将该数组分成两组，只取其中的一组做异或，结果一定是a或者b
func aOrB(a []int, num int) int {
	ans := 0
	for i := 0; i < len(a); i++ {
		if num&a[i] == 0 {
			ans = ans ^ a[i]
		}
	}
	return ans
}

// a ^ a ^ b 或者 b ^ a ^ b
func another(aOrb int, aXorB int) int {
	return aOrb ^ aXorB
}

func twoNum(a []int, num1 *int, num2 *int) {
	temp := aXorB(a)
	devideNum := rightOne(temp)
	*num1 = aOrB(a, devideNum)
	*num2 = *num1 ^ temp
}

func KM(a []int, k int, m int) int {
	var t [32]int
	for i := 0; i < len(a); i++ {
		for j := 0; j < 32; j++ {
			if (a[i]>>j)&1 != 0 {
				t[j]++
			}
		}
	}
	ans := 0
	for i := 0; i < 32; i++ {
		if t[i]%k != 0 {
			if t[i]%k == m {
				ans |= (1 << i)
			} else {
				return -1
			}

		}
	}
	return ans
}
