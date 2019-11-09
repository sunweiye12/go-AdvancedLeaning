package main

import (
	"fmt"
)

func main() {
	var str1 string = "hello" // 反斜杠代表转译
	var str2 string = "world"
	// 拼接字符串
	str3 := str1 + " " + str2
	fmt.Println("str3=", str3)
	fmt.Println("len(str3)=", len(str3))
	//  切片
	substr := str3[0:5]
	fmt.Println(substr)

	substr = str3[6:]
	fmt.Println(substr)

}
