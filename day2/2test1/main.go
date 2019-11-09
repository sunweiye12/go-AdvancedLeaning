package main

/*定义两个常量Man=1和Female=2，获取当前时间的秒数，如果能被Female整除，则
  在终端打印female，否则打印man*/

import (
	"fmt"
	"time"
)

const (
	Male   = 1
	Female = 2
)

func main() {
	for {
		second := time.Now().Unix() //获取当前的秒数
		//fmt.Println(second)
		if second%Female == 0 {
			fmt.Println("female")
		} else {
			fmt.Println("male")
		}
		time.Sleep(1 * time.Second)
	}
}
