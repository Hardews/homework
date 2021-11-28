package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("充钱榜！！！")
	for true {
		var num int               //声明一个变量，后面将随机数赋予它
		var player map[int]string //声明，初始化一个map
		player = make(map[int]string)
		rand.Seed(time.Now().UnixNano()) //设置一个种子
		for i := 0; i < 100; i++ {
			num = rand.Intn(100) + 1 //生成1到100的随机数
			switch  {
			case num == 99:
				player[i] = "该死的领头人~~~"
			case num > 50&& num < 99:
				player[i] = "中流砥柱董事会成员~"
			case num < 10:
				player[i] = "小菜鸟~~~~~~~"
			case num > 10&& num < 50:
				player[i] = "勉强能上去打个架"
			default:
				fmt.Println("他可能已经是神了！！！")
			}
			var a int
			fmt.Println("输入你要查找的玩家等级")
			fmt.Scan(&a)
			_, ok := player[a]
			if ok {
				switch  {
				case a == 99:
					fmt.Println("何人敢搜我大哥的名号，他可是",player[a])
				case a > 50&&a < 99:
					fmt.Println("????，他可是",player[a])
				case a < 10:
					fmt.Println("你说的这个人不太行噻，它是",player[a])
				case a > 10&&a < 50:
					fmt.Println("似乎他有点厉害，它是",player[a])
				default:
					fmt.Println("该玩家未上榜！")
				     }
				  }else {
					  fmt.Println("该玩家不存在")
			      }
			   }
			}
		}