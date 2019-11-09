package main

import "fmt"

// 设计一个函数实现交换两个数的值

func main() {
	a := 5
	b := 10
	//swap(&a,&b)		// 调用函数,来交换两个数
	//a,b = swap1(a,b)
	a, b = b, a
	fmt.Println(a)
	fmt.Println(b)
}

func swap(a *int, b *int) {
	tem := *a
	*a = *b
	*b = tem
	return
}

func swap1(a int, b int) (int, int) {
	return b, a
}
