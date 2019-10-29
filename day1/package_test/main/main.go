package main

import (
	"day1/package_test/calc"
	"fmt"
)

func main() {
	sum := calc.Add(7, 3)
	sub := calc.Sub(7, 4)
	fmt.Println("hello")
	fmt.Println("Add:",sum)
	fmt.Println("Sub:",sub)
}
