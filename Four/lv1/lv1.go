package main

import (
	"fmt"
	"time"
)

func main() {
	var skillName string
	fmt.Println("您的对手准备发起进攻")
	fmt.Println("请输入您要释放的技能")
	fmt.Scan(&skillName)
	switch skillName {
	case "龙卷风摧毁停车场": ReleaseSkill("龙卷风摧毁停车场", func(skillName string) {
		fmt.Println("尝尝我的厉害吧")
	})
	case "乌鸦坐灰机": ReleaseSkill("乌鸦坐灰机", func(skillName string) {
		fmt.Println(skillName)
	})
	case "组合技": ReleaseSkill("组合技", func(skillName string) {
		fmt.Println("一套组合拳！")
	})
	case "大招":  ReleaseSkill("大招", func(skillName string) {
		fmt.Println("开挂的人生不需要解释")
		fmt.Println(skillName,"就是")
	})
	}
}
func ReleaseSkill(skillNames string, releaseSkillFunc func(string)) {
	releaseSkillFunc(skillNames)
	switch skillNames {
	case "大招":
			fmt.Println("直接KO对面")
	case "龙卷风摧毁停车场":
		for i := 0; i < 10; i++ {
			time.Sleep(1*time.Second)
			fmt.Printf("(ﾉ｀⊿´)ﾉ！！！\n对手血量-1\n")
		}
			fmt.Println("K.O")
	case"乌鸦坐灰机":
			fmt.Printf("对手已经被你坐扁！！\n您获得胜利")
	case "组合技":
			fmt.Printf("亢龙有悔!!\n对手血量-2\n")
			time.Sleep(1*time.Second)
			fmt.Printf("神龙摆尾！！\n对手血量-3\n")
			time.Sleep(1*time.Second)
			fmt.Printf("左右勾拳\n对手血量-5\nK.O")
	}
}
