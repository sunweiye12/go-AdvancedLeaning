package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	// -----------------内置函数----------------
	// 1. close：主要用来关闭channe
	// 2. len：用来求长度，比如string、array、slice、map、channel
	// 3. new：用来分配内存，主要用来分配值类型，比如int、struct。返回的是指针
	// 4. make：用来分配内存，主要用来分配引用类型，比如chan、map、slice
	// 5. append：用来追加元素到 数组、slice中
	// 6. panic和recover：用来做错误处理   panic用于抛出异常,recover用来捕获异常

	test()

	var i int
	fmt.Println(i)	// 得到是的值

	j := new(int)	// 动过new来创建基本数据类型 返回的是地址值
	fmt.Println(j)
	fmt.Println(*j)	// 得到地址随对应的值

	var a []int		// 定义一个切片
	a = append(a,10,12,21)  // append用于在切片后面追加元素
	a = append(a,a...)  // a...的依次就是将a切片展开
	fmt.Println(a)

	test1()		// make和new区别

	// recusive() 	//递归函数,大问题可以由子问题解决

	for i:=0; i<10; i++ {
		fmt.Print(fab(i)," ")
	}
	fmt.Println()

	f := test2() // 闭包:一个函数和与其相关的饮用环境组合而成的实体
	fmt.Println(f(1))
	fmt.Println(f(10))
	fmt.Println(f(100))

	// 利用闭包来用一个函数实现两个功能
	f1 := makeSuffix(".jpg")
	fmt.Println(f1("test"))
	fmt.Println(f1("pic.jpg"))

	f2 := makeSuffix(".png")
	fmt.Println(f2("test"))
	fmt.Println(f2("121.png"))
}

func test() {
	defer func() {	// 在函数结束前判断函数结束前是否法伤错误导致停止,如果发生错误,则只想recover() 恢复一下
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	b := 0
	a := 100/b
	fmt.Println(a)
	return
}

// make和new的区别
func test1() {
	s1 := new([]int) // 用new 创建一个切片,他返回的是指向切片的一个地址.只是开辟可这样的一个类型,还没有初始化
	fmt.Println(s1)
	*s1 = make([]int,5)		// 还要通过make来将new创建出来的对象进行初始化
	(*s1)[0] = 10
	fmt.Println(*s1)

	s2 := make([]int,3) // 创建一个空间为3的切片,返回的是切片
	fmt.Println(s2)
}

// 递归函数
func recusive(){
	fmt.Println("递归函数")
	time.Sleep(time.Second)
	recusive()
}

// 菲波那切数列
func fab(n int) int {
	if n <= 1{
		return 1
	}
	return fab(n-1) + fab(n-2)
}

// 闭包:一个函数和与其相关的引用环境组合而成的实体
func test2() func(int) int {
	var x int		// 每次当调用函数的时候x会保留记录下来,在下一次调用的时候使用上一次的值
	return func(d int) int {
		x += d
		return x
	}
}

// 用闭包实现添加图片名的格式
func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name,suffix) {
			return name+suffix
		}
		return name
	}
}
