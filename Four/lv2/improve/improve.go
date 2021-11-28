package main

import (
	"fmt"
	"time"
)

func main()  {
	refresh := time.NewTicker(30 * time.Second)
	for{
		<- refresh.C
		fmt.Println("芜湖！起飞")
		t1 := time.Now()
		switch  {
		case t1.Hour()==2  && t1.Minute()==0 : fmt.Println("谁能比我卷")
		case t1.Hour()==6  && t1.Minute()==0 : fmt.Println("早八算什么，早六才是吾辈应起之时")
		}
	}
}
