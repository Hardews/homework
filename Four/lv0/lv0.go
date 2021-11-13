package main

import (
	"fmt"
	"math/rand"
	"time"
)
//  这串代码主要是随机生成25个随机数，然后根据他们的区间去给他们一个“头衔”，让执行的人输入100以内的数字
//  看这些数字是不是这25个玩家中某一个玩家的等级
//  没啥用，写着玩，复习一下map和time 字符串转类型那个没想到咋用呢（或许可以生成验证码）
func main()  {
	for true{
		var num int  //声明一个变量，后面将随机数赋予它
		var player map[int]string  //声明，初始化一个map
		player = make(map[int]string)
		rand.Seed(time.Now().UnixNano()) //设置一个种子
		for i := 0; i < 25; i++ {
			num = rand.Intn(100) + 1  //生成1到100的随机数
			if num > 99{       //key-value
				player[num] = "该死的领头人"
			}
			if num < 10{
				player[num] = "一个准备出新手村的萌新"
			}
			if num <50{
				if num>10 {
					player[num] = "游戏里的还行玩家"
				}
			}
			if num>50{
				if num <98{
					player[num] = "游戏里的中流砥柱"
				}
			}
		}
		var a int
		fmt.Println("输入你要查找的玩家等级")
		fmt.Scan(&a)
		_,ok := player[a]
		if ok {
			if a == 99{
				fmt.Println("他可强了")
				fmt.Println("他是",player[a])
			}else {
				fmt.Println("该玩家是",player[a])
		}
		    }else {
			fmt.Println("该玩家不存在")
		}
	}
}
