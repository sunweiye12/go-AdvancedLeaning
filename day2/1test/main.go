package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello")
	n := 5
	for i := 0; i <= n; i++ {
		//fmt.Println(i)
		//fmt.Println(i, "+", getB(n, i), "=", n)
		fmt.Printf("%d+%d=%d\n", i, getB(n, i), n)
		fmt.Printf("%d+%d=%d\n", i, getB(n, i), n)
	}
}

func getB(n int, a int) int {
	return n - a
}
