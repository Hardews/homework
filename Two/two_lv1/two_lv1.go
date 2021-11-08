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
}


