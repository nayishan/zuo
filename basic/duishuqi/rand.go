package main

import (
	"fmt"
	"math"
	"math/rand"
)

func testRand() {
	maxTimes := 1000000
	count := 0
	for i := 0; i < maxTimes; i++ {
		num := rand.Float64()
		if num < 0.3 {
			count = count + 1
		}
	}
	fmt.Println(float64(count) / float64(maxTimes))
}

func testRand2() {
	maxTimes := 1000000
	maxIndex := 9
	count := make([]int, maxIndex)
	num := 0
	for i := 0; i < maxTimes; i++ {
		num = rand.Intn(maxIndex)
		count[num]++
	}
	for i := 0; i < maxIndex; i++ {
		fmt.Println("num ", i, "accur", count[i], "times")
	}
}
func testRand3() {
	maxTimes := 1000000
	count := 0
	for i := 0; i < maxTimes; i++ {
		num := xToXPower2()
		if num < 0.3 {
			count++
		}
	}
	fmt.Println(float64(count) / float64(maxTimes))
}

//return [0,1)
//scope [0~x] return x2
func xToXPower2() float64 {
	return math.Max(rand.Float64(), rand.Float64())
}

func testRand4() {
	maxTimes := 1000000
	count := 0
	for i := 0; i < maxTimes; i++ {
		num := xToXPower3()
		if num < 0.3 {
			count++
		}
	}
	fmt.Println(float64(count) / float64(maxTimes))
}
func xToXPower3() float64 {
	return math.Max(rand.Float64(), math.Max(rand.Float64(), rand.Float64()))
}

func main() {
	testRand()
	testRand2()
	testRand3()
	testRand4()
}
