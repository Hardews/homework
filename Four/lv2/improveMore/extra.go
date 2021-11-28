package main

import (
	"fmt"
	"time"
)

func main()  {
	menu()
}
/*
还没写完，因为还是不大会，用了一些很笨的方法实现
准备去想一下怎么存储这些已经输入的切片
有些功能还没有完善，以后再写，先弄注册登录，跟上数据库啥的
 */
type timer struct {
	title string
	setTime string
	time time.Time
}

func checkContinue()  {
	var chose int
	fmt.Println("是否继续留在执行页面\n1.是\n2.否")
	switch chose {
	case 1:executionMenu(time.Now()," ","Now")
	case 2:menu()
	}
}
func executionMenu(Time time.Time, output,title string)  {
	fmt.Printf("\n\n\n")
	fmt.Println("----执行页面----")
	fmt.Printf("\n")
	var n = 0
	for  {
		refresh := time.NewTicker(2 * time.Second)
		<- refresh.C
		t := time.Now()
		s := time.Now().Format("2006-01-02 15:04:05")
		switch  {
		case t.Hour()==2  && t.Minute()==0 :
			fmt.Println()
			fmt.Println("都凌晨两点了，没想到吧，我还在学！！")
			time.Sleep(time.Minute)
			checkContinue()
		case t.Hour()==6  && t.Minute()==0 :
			fmt.Println()
			fmt.Println("早八算个屁，我早六卷四你")
			time.Sleep(time.Minute)
			checkContinue()
		case t.Hour()==8  && t.Minute()==0 :
			fmt.Println()
			fmt.Println("早八好好睡，学习跟不上，精神要跟上")
			time.Sleep(time.Minute)
			checkContinue()
		case t.Hour()==Time.Hour()  && t.Minute()==Time.Minute():
			fmt.Println()
			fmt.Printf("名为%s的定时器提醒您 : 现在的时间是 : %s\n%s",title,s,output)
			time.Sleep(time.Minute)
			checkContinue()
		default:
			for i := 0; i >= 0; {
				if i==1 {
					fmt.Print("程序正在运行")
					break
				}else if i==20*n {
					fmt.Print(".")
					break
				}else if i==100*n{
					fmt.Println("检测到您停留的时间过长")
					checkContinue()
				}else {
					i++
					break
				}
			}
		}
	}
}

var timerMap = make(map[int]timer) // ......

func storageNowClock()  { //存储的地方，还在想办法完善....
	timerMap[0] = timer{
		title: "早六人",
		setTime: "早上六点",
		time: time.Date(0,0,0,6,0,0,0,time.Local),
	}
	timerMap[1] = timer{
		title: "早八人",
		setTime: "早上八点",
		time: time.Date(0,0,0,8,0,0,0,time.Local),
	}
	timerMap[2] = timer{
		title: "卷王哥哥",
		setTime: "凌晨两点",
		time: time.Date(0,0,0,2,0,0,0,time.Local),
	}
	for j, i := range timerMap {
		fmt.Printf("定时器%d ,标题 : %s\n时间 : %s\n",j,i.title,i.setTime)
	}
	fmt.Printf("\n\n")
	time.Sleep(2*time.Second)
	fmt.Println("显示完毕，即将返回主菜单")
	time.Sleep(3*time.Second)
	menu()
}

func storageClock(title string,Time time.Time)  { //想办法中.....
	timerMap[len(timerMap) + 1] = timer{
		title: title,
		time: Time,
	}
}

func menu()  {
	var Input int
	 fmt.Println("------主菜单------")
	 fmt.Printf("1.设置新定时器\n2.执行页面\n3.退出\n4.查看定时器\n")
	 fmt.Scan(&Input)
	 switch Input {
	 case 1: setClock()
	 case 2: executionMenu(time.Now()," ","Now")
	 case 3: return
	 case 4: storageNowClock()

	 default:
		 fmt.Println("error!!!")
	}
}

func setClock()  {
	 var Input int
	 fmt.Println("-----请选择想要设置的类型------")
	 fmt.Printf("1.重复定时器\n2.一次性定时器\n")
	 fmt.Scan(&Input)
	switch Input {
	case 1:setRepeatClock()
	case 2:setDisposableClock()
	default:
		fmt.Println("error")
	}
	return
}
func setRepeatClock()  {
	var InputTimeHour int
	var InputTimeMinute int
	var outputWhat string
	var title string
	fmt.Println("请输入想要设置的定时器的标题")
	fmt.Scan(&title)
	fmt.Println("请输入想要设置的定时器为几点")
	fmt.Scan(&InputTimeHour)
	t := time.Now()
	switch {
	case t.Hour() <= InputTimeHour:fmt.Printf("------已默认%d年%d月%d日开始重复-----",time.Now().Year(),time.Now().Month(),time.Now().Day())
	case t.Hour() > InputTimeHour :fmt.Printf("------已默认%d年%d月%d日开始重复-----",time.Now().Year(),time.Now().Month(),time.Now().Day()+1)
	}
	fmt.Println("请输入想要设置的定时器为几分")
	fmt.Scan(&InputTimeMinute)
	fmt.Println("你想在这个时候输出什么？")
	fmt.Scan(&outputWhat)
	repeatClock(InputTimeHour,InputTimeMinute,outputWhat,title)
}

func repeatClock(Hour,Minute int,YouWantSay,title string)  {
	 ANewClock := time.Date(0,0,0,Hour,Minute,0,0,time.Local)
	 executionMenu(ANewClock,YouWantSay,title)
}

func setDisposableClock()  {
	var chose,InputTimeHour,InputTimeMinute int
	var InputTimeYear,InputTimeMonth,InputTimeDay int
	var outputWhat string
	fmt.Printf("------默认年份设置的日期为%d年%d月%d日-----",time.Now().Year(),time.Now().Month(),time.Now().Day())
	fmt.Printf("您是否想更换年月日？\n1.是\n2.否\n")
	fmt.Scan(&chose)
	switch chose {
	case 1:
		fmt.Println("请输入想要设置的定时器的年份")
		fmt.Scan(&InputTimeYear)
		fmt.Println("请输入想要设置的定时器的月份")
		fmt.Scan(&InputTimeMonth)
		fmt.Println("请输入想要设置的定时器为何日")
		fmt.Scan(&InputTimeDay)
		fmt.Println("请输入想要设置的定时器为何时")
		fmt.Scan(&InputTimeHour)
		fmt.Println("请输入想要设置的定时器为几分")
		fmt.Scan(&InputTimeMinute)
		fmt.Println("你想在这个时候输出什么？")
		fmt.Scan(&outputWhat)
		disposeClock(InputTimeYear,InputTimeDay,InputTimeHour,InputTimeMinute,outputWhat,inputMonth(InputTimeMonth))
	case 2:
		fmt.Println("请输入想要设置的定时器为几点")
		fmt.Scan(&InputTimeHour)
		fmt.Println("请输入想要设置的定时器为几分")
		fmt.Scan(&InputTimeMinute)
		fmt.Println("你想在这个时候输出什么？")
		fmt.Scan(&outputWhat)
		disposeClock(time.Now().Year(),time.Now().Day(),InputTimeHour,InputTimeMinute,outputWhat,time.Now().Month())
	}
}

func disposeClock(Year,Day,Hour,Minute int, output string,Month time.Month)  {
	var chose int
	var title string
	var Time time.Time
	Time = time.Date(Year,Month,Day,Hour,Minute,time.Now().Second(),time.Now().Nanosecond(),time.Local)
	for {
		refresh := time.NewTicker(time.Second)
		<- refresh.C
		t := time.Now()
		fmt.Println("现在的时间是：",t)
		fmt.Printf("选择接下来的操作\n1.等定时器输出\n2.退出\n")
		fmt.Scan(&chose)
		break
	}
	switch chose {
	case 1:
		for true {
			refresh := time.NewTicker(time.Second)
			<- refresh.C
			t := time.Now()
			if t.Year()==Year && t.Month()== Month&&t.Day()==Day &&t.Hour()==Hour && t.Minute()==Minute{
				fmt.Println("时间到啦时间到啦")
				fmt.Println(output)
			}else {
				fmt.Printf("检测到定时器设定时间还有点远，选择操作\n1.留在此页面五分钟\n2.退出\n3.返回空页面并保持运行\n")
				fmt.Scan(&chose)
				switch chose {
				case 1:
					time.Sleep(5 * time.Second)
					continue
				case 2:
					break
				case 3:
					fmt.Println("请为它设置标题")
				    fmt.Scan(&title)
					storageClock(title,Time)
					fmt.Print("正在前往执行页面")
					for i := 0; i < 6; i++ {
						time.Sleep(time.Second)
						fmt.Print(".")
					}
					fmt.Println()
					time.Sleep(3*time.Second)
					executionMenu(Time,output,title)
				}
			}
		}
	case 2:
		return
	}
	return

}
func inputMonth(m int)  time.Month{
	var t time.Month
	switch m {
	case 1:  t = time.January
	case 2:  t = time.February
	case 3:  t = time.March
	case 4:  t = time.April
	case 5:  t = time.May
	case 6:  t = time.June
	case 7:  t = time.July
	case 8:  t = time.August
	case 9:  t = time.September
	case 10: t = time.October
	case 11: t = time.November
	case 12: t = time.December
	default:
		fmt.Println("error")
	}
	return t
}


