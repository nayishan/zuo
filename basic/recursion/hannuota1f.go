package main

import "fmt"

func f1(n int, from, to, other string) {
	if n == 1 {
		fmt.Println("Move 1 from", from, "to", to)
	} else {
		f1(n-1, from, other, to)
		fmt.Println("Move", n, "from", from, "to", other)
		f1(n-1, other, to, from)
	}
}

func main() {
	f1(3, "Left", "Right", "Mid")
}
