package main

import (
	"fmt"
	"os"
)


func main()  {
	 file,err := os.Create("D:/GOProjects/src/Three/three_lv2/plan.exe")//创建一个新文件
	 if err!= nil{                        //错误类型判断
		 fmt.Println("err:",err)
		 return
	 }

	 str := "I’m not afraid of difficulties and insist on learning programming" //定义写入字符串
	 n ,_ := file.Write([]byte(str))    //字符串转字节
	 if err != nil{
		 fmt.Println("err:",err)
		 return
	 }
	 fmt.Println("写入字节数",n)
	 p := make([]byte,n)            //定义一个字节切片p，使它等于n
	 m,_:= file.ReadAt(p,0)     //从头开始读取这个切片
	 if err!=nil {
		fmt.Println("err:",err)
		return
	 }
	 fmt.Println("读出字节数",m)
	 fmt.Println(string(p))        //使其转回为字符串
	 file.Close()                  //关闭文件

}
