package main

import "fmt"

func main() {
	var a [10]int   //数组的定义,并默认初始化,长度信息是数组类型的一部分
	a[0] = 12
	a[9] = 3

	fmt.Println(a)

	//遍历数组的方法
	for i := 0;i<len(a) ;i++  {
		fmt.Print(a[i]," ")
	}
	fmt.Println()

	for index,value := range a {
		fmt.Printf("a[%d] = %d\n",index,value,)
	}

	创建数组()
}

func 创建数组(){
	var a1 [5]int = [5]int{1,3,4}
	var a2 = [5]int{1,2,3,4,5}
	var a3  = [...]int{1,4,3,2,5}
	var a4  = [...]int{1:5, 6:7}
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a4)

	// 二维数组
	var arr [2][3]int  = [...][3]int{{3,2,1},{1,2,3}} // 第一个可以推导第二个不可以推导
	fmt.Println(arr)
}
