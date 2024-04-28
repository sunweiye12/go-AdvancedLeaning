package main

import (
	"fmt"
	"time"
)


func main() {
	const Layout = "2006-01-02 15:00:00"
	//
	//now := time.Now().Unix()+3600*4
	//tm := time.Unix(now, 0)
	//hourEndTime, _ := time.Parse("2006-01-02 15:00:00", tm.Format("2006-01-02 15:00:00"))
	//hourEndTime, _ := time.Parse(Layout, tm.Format(Layout))

	//fmt.Println(now)
	//fmt.Println(tm)
	//fmt.Println(hourEndTime)
	//fmt.Println("test")
	fmt.Println(time.Now().Unix())
	fmt.Println(time.Unix(1675681200, 0))

}
