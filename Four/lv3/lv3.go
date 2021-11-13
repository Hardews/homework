package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type Users struct {
	Name string
	Cipher string
}
func main()  {
	 var user Users
	 var usersData map[string]string
	 usersData = make(map[string]string)
	 fileName := "D:/GOProjects/src/homework_in_redrock/Four/lv3/usersData.txt"
	 file,err := os.OpenFile(fileName,os.O_RDWR,0)
	 if err!= nil{
		 fmt.Println("err:",err)
		 return
	 }
	 defer file.Close()
	 b1 := bufio.NewReader(file)
	for  {
		user.Name,_ = b1.ReadString(' ')
		user.Cipher,err = b1.ReadString('\n')
		if err!= nil {
			fmt.Println("err:",err)
			return
		}
		usersData[user.Cipher] = user.Name
        if err == io.EOF{
			break
		}
		break
	}
	for key, value := range usersData{
		fmt.Println(key,":",value)
	}
	var s int
	fmt.Println("请输入您要执行的操作")
    fmt.Printf("1.修改密码\n2.注册新账号\n3.退出程序\n")
	fmt.Scan(&s)
	if s == 1{
		change(usersData)
	}
	if s == 2{
		registered(usersData)
	}
	if s == 3{
		fmt.Println("程序已退出！")
	}else if s!=1 {
		if s != 2 {
			{
				fmt.Println("暂时还不明白您是什么操作")
				fmt.Println("再给您一次考虑的机会")
				fmt.Scan(&s)
				if s == 1 {
					change(usersData)
				} else if s == 2 {
					registered(usersData)
				} else if s != 1 {
					if s != 2 {
						fmt.Println("程序退出")
					}
				}
			}
		}
	}
}
func registered(usersData map[string]string)  {
     var name string
	 var cipher string
	 var t bool
	 for{
		 fmt.Println("请输入用户名")
		 fmt.Scan(&name)
		 for i,_ := range usersData{
			 a := bytes.EqualFold([]byte(i),[]byte(name))
			 if t == a {
				 fmt.Println("用户名已存在，请重新输入！")
				 continue
			 }else{
				 break
			 }
			 break
		 }
		 fmt.Println("请输入密码")
		 fmt.Scan(&cipher)
		 if len(cipher) <6 {
			 fmt.Println("密码太短，请重新输入！")
			 fmt.Scan(&cipher)
			 if len(cipher) < 6 {
				 fmt.Println("多次设置错误，请重新选择您要的服务")
				 fmt.Println("请输入您要执行的操作")
				 fmt.Printf("1.修改密码\n2.注册新账号\n3.退出程序\n")
			 }else {
			 usersData[cipher] = name
			 }
		 }else {
			 usersData[cipher] = name
		 }
		 break
	 }
	 writeIn(name,cipher,usersData)
	 fmt.Println("注册成功！")
}
func change(data map[string]string)  {
	 fileName := "D:/GOProjects/src/homework_in_redrock/Four/lv3/usersData.txt"
	 var userName string
	 var userCipher string
	 fmt.Println("请输入用户名")
	 fmt.Scan(&userName)

	 file,err := os.OpenFile(fileName,os.O_RDWR,0)
	 if err != nil{
		 fmt.Println("err: ",err)
		 return
	 }
	 reader := bufio.NewReader(file)
	 var str string
	 for {
		 str,err = reader.ReadString('\n')
		 if err == io.EOF{
			 break
		 }
	 }
	for  {
		if strings.Index(str,userName) == -1{
			fmt.Println("您输入的用户名不存在,请再次输入")
			fmt.Scan(&userCipher)
		}else {
			break
		}
	}
     fmt.Println("请输入密码")
	 for {
		 var i = 0
		 fmt.Scan(&userCipher)
		 value := data[userName]
		 if userName != value{
			 fmt.Println("密码错误")
			 i++
			 if i > 3 {
				 fmt.Println("密码错误过多，请十分钟后再尝试")
				 time.Sleep(10*time.Minute)
			 }else {
				 break
			 }
		 }else {
			 break
		 }
	 }
	 var newCipher string
	 fmt.Println("请输入新密码：")
	 fmt.Scan(&newCipher)
	 str = strings.Replace(str,userCipher,newCipher,1)
	 writer := bufio.NewWriter(file)
	 writer.WriteString(str)
	 writer.Flush()
	 defer file.Close()
	 writeIn(userName,newCipher,data)
}
func writeIn(User string,cipher string,data map[string]string)  {
	data[cipher] = User
	fileName := "D:/GOProjects/src/homework_in_redrock/Four/lv3/usersData.txt"
	file,err := os.OpenFile(fileName,os.O_WRONLY|os.O_APPEND,0)
	if err!= nil{
		fmt.Println("err：",err)
		return
	}
	file.Seek(0,2)
	file.WriteString("\n")
	file.WriteString(User)
	file.WriteString(" ")
	file.WriteString(cipher)
	file.WriteString("\n")
	file.Close()
}
