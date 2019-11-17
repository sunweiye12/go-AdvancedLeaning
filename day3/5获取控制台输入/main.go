package main

import (
	"fmt"
)

func main() {
	add1()
	add2(1,2)
	add3(1,2)
	add4(1,2)
	add5(1,2)
}

func add1() {
	fmt.Println("第一个无参无返回值函数")
}

func add2(a int, b int) {
	fmt.Println("第一个有参无返回值函数")
}

func add3(a int, b int) int {
	fmt.Println("第一个有参有返回值函数")
	return 1
}

func add4(a int, b int) (int, int)  {
	fmt.Println("第一个有参多返回值函数")
	return 1,2
}

func add5(a, b int) (int, int)  {
	fmt.Println("第一个有参多返回值函数")
	return 1,2
}
//==================匿名函数====================
func add6(a, b int) int {
	// 声明一个匿名函数
	result := func(a1 int ,b1 int) int {
		return a1+b1
	}(a,b)
	return result
}

func main1() {
	c := add6
	fmt.Println(c)
	sum := c(10, 20)
	fmt.Println(sum)
	if ( c == add6 ) {
		fmt.Println("c equal add")
	}
}


