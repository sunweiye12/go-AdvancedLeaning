package main

import (
	"fmt"
)

func main() {
	var str string = "hello world\n" // 反斜杠代表转译
	// 可以保留书写格式
	var str1 string = `	窗前明月光,
	疑是地上霜.
	举头望明月,
	我是郭德纲.
	`
	var a byte = 'c'
	fmt.Println(str)
	fmt.Println(str1)
	fmt.Println(a)
	fmt.Printf("%c\n", a) // 格式化输出才能够打印出字符
	fmt.Printf("%T\n", a) // %T打印类型
}
