package main

import (
	"fmt"
	"time"
)
func main()  {
	t1 := time.Now()
	s1 := t1.Format("2006年1月2日 15时04分05秒")
	fmt.Println("您启动本程序的时间是",s1)
	for true {
		t2 := time.Now()  //定义一个现在的时间，输出了芜湖起飞后更新时间
		go clock1(t2)
		   clock2(t2)
		   clock3(t2)
		   remindTime(t2)
		time.Sleep(30*time.Second)
		fmt.Println("芜湖,起飞!")
	}


}
func clock1(t time.Time)  {
	 if t.Hour() == 6{  //早六宣言
		 if t.Minute()== 0{
			 fmt.Println("早八算什么，早六才是吾辈应起之时")
		 }
	 }
}
func clock2(t time.Time)  {
	if t.Hour() == 2{   //凌晨宣言
		if t.Minute()==00{
			fmt.Println("谁能比我卷！")
		}
	}
}
func clock3(t time.Time)  {
	if t.Hour() == 19{   //凌晨宣言
		if t.Minute()==36{
			fmt.Println("测试")
		}
	}
}
func remindTime(t time.Time)  {  //闹钟准备好时提醒现在的时间
	s2 := t.Format("2006年1月2日 15时04分05秒")
	if t.Hour() == 6 {
		if t.Minute()==0 {
			fmt.Println("现在的时间是",s2)
		}
	}else if t.Hour() == 2{
		if t.Minute() == 0{
			fmt.Println("现在的时间是",s2)
		}
	}
}
