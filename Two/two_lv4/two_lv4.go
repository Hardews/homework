package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Player struct {
	PlayerName string
	Level  int
}
type PlayerSlice []Player

func (s PlayerSlice) Len() int {
	return len(s)
}//实现接口
func (s PlayerSlice) Less(i,j int)  bool{
	return s[i].Level > s[j].Level
}//按等级从高到低排序
func (s PlayerSlice) Swap(i,j int)  {
	s[i],s[j] = s[j],s[i]
}
func main()  {
	var play PlayerSlice
	var n int
	fmt.Println("输入排序个数")
	fmt.Scanln(&n)
	rand.Seed(time.Now().UnixNano())
	for i := 0 ; i < n ; i++{
		player := Player{
			PlayerName : fmt.Sprintf("No.%d player's level is",rand.Intn(1000)),  //生成玩家
			Level:      rand.Intn(100),                                      //生成等级
		}
		play = append(play,player)  //加入到切片中
	}
	sort.Sort(play)          //排序
	for _,v := range play{  //遍历，输出
		fmt.Println(v)
	}

}
