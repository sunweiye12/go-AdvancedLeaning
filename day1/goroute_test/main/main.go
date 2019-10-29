package main

import (
	"day1/goroute_test/add"
	"fmt"
)

func main() {
	fmt.Println("hello")
	c := make(chan int,1)
	go add.Add(4,9,c)
	sum := <- c
	fmt.Println("sum = ",sum)
}
