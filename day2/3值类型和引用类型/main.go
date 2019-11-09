package main

/*1.值类型：基本数据类型int、float、bool、string以及数组和struct。  内存在栈中占用
  2.引用类型：指针、slice、map、chan等都是引用类型。				内存在堆中分配
*/
import (
	"fmt"
)

func main() {
	var a = 100
	var b chan int = make(chan int, 1)
	fmt.Println("a=", a) // 值类型传递过来的就是值
	fmt.Println("b=", b) // 引用类型传递过来的就是地址值
	// 直接传递值,不会影响原来的值
	modify1(a)
	fmt.Println(a)
	// 传递的是地址值则会影响到原来的值
	modify2(&a) // 利用 & 可以获取值对象的地址值
	fmt.Println(a)
}

func modify1(a int) {
	a = 10
	return
}

func modify2(a *int) { // *int 代表int类型的地址值  *a得到a地址值上所存储的对象
	*a = 10
	return
}
