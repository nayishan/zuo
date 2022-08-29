package main

import "fmt"

func LeftToRight(n int) {
	if n == 1 {
		fmt.Println("Move 1 from left to right")
		return
	}
	LeftToMid(n - 1)
	fmt.Println("Move", n, "from left to right")
	MidToRight(n - 1)

}

func LeftToMid(n int) {
	if n == 1 {
		fmt.Println("Move 1 from Left to Mid")
		return
	}
	LeftToRight(n - 1)
	fmt.Println("Move", n, "from Left to Mid")
	RightToMid(n - 1)
}

func RightToMid(n int) {
	if n == 1 {
		fmt.Println("Move 1 from Right to Mid")
		return
	}
	RightToLeft(n - 1)
	fmt.Println("Move", n, "from Right to Mid")
	LeftToMid(n - 1)

}

func RightToLeft(n int) {
	if n == 1 {
		fmt.Println("Move 1 from Right to Left")
		return
	}
	RightToMid(n - 1)
	fmt.Println("Move", n, "from Right to Left")
	MidToLeft(n - 1)

}
func MidToLeft(n int) {
	if n == 1 {
		fmt.Println("Move 1 from Mid to Left")
		return
	}
	MidToRight(n - 1)
	fmt.Println("Move", n, "from Mid to Left")
	RightToLeft(n - 1)
}
func MidToRight(n int) {
	if n == 1 {
		fmt.Println("Move 1 from Mid to Right")
		return
	}
	MidToLeft(n - 1)
	fmt.Println("Move", n, "from Mid to Right")
	LeftToRight(n - 1)

}

func main() {
	LeftToRight(3)
}
