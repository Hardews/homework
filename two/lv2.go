package main

import "fmt"

type BasicAuthorInfo struct {
	TheName string
	TheFans int
}
type ExtraAuthorInfo struct {
	Gender       string
	WechatNumber string

}
type BasicVideoInfo struct {
	VideoName    string
	VideoAddress string
	VideoContent string
}
type OtherVideoInfo struct {
	TheNumberOfBarrage     int
	TheNumberOfLike        int
	TheNumberOfCollection  int
	TheNumberOfForwarding  int
	TheNumberOfCoins       int
	Comments               int
}
type VideoInfo struct {
	Basic1  BasicAuthorInfo
	Extra   ExtraAuthorInfo
	Basic2  BasicVideoInfo
	Other   OtherVideoInfo
}

func Add1(i *int){
	*i+=1
}
func Add2(l *int){
	*l+=2
}

func main()  {
	UpMaster   := VideoInfo{
		Basic1 :  BasicAuthorInfo{
			TheName : "好影课堂在熙",
			TheFans : 14000,
		},
		Extra: ExtraAuthorInfo{
			Gender       : "男",
			WechatNumber : "xy118030890",
		},
		Basic2: BasicVideoInfo{
			VideoName    : "【摄影后期】这次爽了，全是付费买的资源，请低调使用~",
			VideoAddress : "video/BV1mu411Z7jg",
			VideoContent : "摄影后期教学",
		},
		Other : OtherVideoInfo{
			TheNumberOfBarrage     : 369,
			TheNumberOfLike        : 6702,
			TheNumberOfCollection  : 280000,
			TheNumberOfForwarding  : 662,
			TheNumberOfCoins       : 6715,
			Comments               : 175,

		},
	}
	fmt.Println(UpMaster.Basic1.TheName)
	fmt.Println("欢迎来到视频详情页面")
	var a int
	fmt.Println("1.点   赞")
	fmt.Println("2.投   币")
	fmt.Println("3.转   发")
	fmt.Println("4.一键三连")
	fmt.Print("请输入您想做的操作（输入数字）：")
	fmt.Scan(&a)
	switch a {
	case 1: Add1(&UpMaster.Other.TheNumberOfLike)
	        fmt.Println("点赞成功！")
		    fmt.Println("现在的点赞数为",UpMaster.Other.TheNumberOfLike)
	case 2: Add1(&UpMaster.Other.TheNumberOfCoins)
	        fmt.Println("投币成功！")
	        fmt.Println("现在的硬币数为：",UpMaster.Other.TheNumberOfCoins)
	case 3: Add1(&UpMaster.Other.TheNumberOfForwarding)
		    fmt.Println("转发成功！")
		    fmt.Println("现在的转发量为：",UpMaster.Other.TheNumberOfForwarding)
	case 4: Add1(&UpMaster.Other.TheNumberOfLike)
		    Add2(&UpMaster.Other.TheNumberOfCoins)
		    Add1(&UpMaster.Other.TheNumberOfForwarding)
	        fmt.Println("一键三连成功！")
		    fmt.Println("现在的点赞数为",UpMaster.Other.TheNumberOfLike)
		    fmt.Println("现在的硬币数为：",UpMaster.Other.TheNumberOfCoins)
		    fmt.Println("现在的转发量为：",UpMaster.Other.TheNumberOfForwarding)
	}
}
