package main

import "fmt"

var AName string
var VName string
type NewVideo struct {
	AuthorName string
	VideoName  string
}

func (N *NewVideo)create() {
	N.AuthorName = AName
	N.VideoName  = VName
}
func main()  {
	fmt.Println("您的昵称")
	fmt.Scan(&AName)
	fmt.Println("输入您要发布视频的名称")
	fmt.Scan(&VName)
	daryl := NewVideo{AName,VName}
	fmt.Println("亲爱的“",daryl.AuthorName,"”")
	fmt.Println("您名为“",daryl.VideoName,"”的视频发布成功")
}