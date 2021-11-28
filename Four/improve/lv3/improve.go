package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

/*
这个登陆注册这个，怎么写json我还是不太懂，还有读我也是有点懵
所以这个算是不好的代码，但我以后会改的
我再学学...
*/

func main()  {
	readUsersData()
	menu()
}

type users struct {
	 usersName string
	 password string
}

var dataMap = make(map[string]users)

func readUsersData() map[string]users{
	var usersData users
	fileName := "D:/GOProjects/src/homework_at_redrock/Four/improve/lv3/Data.txt"
	file1,err := os.OpenFile(fileName,os.O_RDONLY,os.ModePerm)
	if err!=nil{
		fmt.Println("err1:",err)
	}
	defer file1.Close()
    data := bufio.NewReader(file1)
	for {
		usersData.usersName,_= data.ReadString(' ')
		usersData.password, err = data.ReadString('|')
		dataMap[usersData.usersName] = usersData
		if err == io.EOF {
			break
		}
	}
	return dataMap
}

func menu()  {
	fmt.Println("-----主菜单-----")
	fmt.Println("请选择您的操作")
	fmt.Printf("1.登录\n2.注册\n3.退出\n")
	var chose int
	fmt.Scan(&chose)
	switch chose {
	case 1:
		login()
	case 2:
		register()
	case 3:
		return
	}
}

func login()  {
	 readUsersData()
	 var i = 0
	 var user users
     fmt.Println("-----登陆页面-----")
try:
	 fmt.Println("请输入账号")
	 fmt.Scan(&user.usersName)
	 user.usersName = user.usersName + " "
		 i++
		 if i > 4 {
			 fmt.Println("错误次数过多，锁定五分钟后返回主页面")
			 time.Sleep(5 * time.Minute)
		 }
	 fmt.Println("请输入密码")
	 fmt.Scan(&user.password)
	user.password = user.password + "|"
     if dataMap[user.usersName]==user{
		 fmt.Println("登陆成功")
		 time.Sleep(3 * time.Second)
		 fmt.Print("即将返回主菜单")
		 for i := 0; i < 6; i++ {
			 time.Sleep(time.Second)
			 fmt.Print(".")
		 }
		 fmt.Println()
		 time.Sleep(time.Second*2)
		 menu()
	 }else {
		 fmt.Println("登陆失败,请重新输入")
		 goto try
	 }
}

func register()  {
	var i = 0
	var NewUsername,NewPassword string
	readUsersData()
	fmt.Println("-----注册页面-----")
	try:
	fmt.Println("请输入您想要注册的账号")
	fmt.Scan(&NewUsername)
	NewUsername = NewUsername + " "
	_, ok := dataMap[NewUsername]
	if ok{
		fmt.Println("账号已存在，请重新输入")
		goto try
	}else {
		goto again
	}
	again:
	fmt.Println("输入您的密码 （大于等于六位）")
	fmt.Scan(&NewPassword)
	switch {
	case len(NewPassword) >= 6:
		break
	case len(NewPassword) <= 6:
		i++
		if i > 3{
			fmt.Println("错误多次，已锁定，五分钟后返回主菜单")
			time.Sleep(5 * time.Minute)
			menu()
		}
		fmt.Println("密码长度不够，请重新输入")
		goto again
	}
	writeIn(NewUsername,NewPassword)
	fmt.Println("注册成功，返回主菜单")
	menu()
}

func writeIn(user,password string)  {
	var usersData users
	usersData.usersName = user
	usersData.password  = password
	fileName := "D:/GOProjects/src/homework_at_redrock/Four/improve/lv3/Data.txt"
	file,err := os.OpenFile(fileName,os.O_RDWR|os.O_APPEND,os.ModePerm)
	if err!=nil{
		fmt.Println("err:",err)
	}
	file.Write([]byte(  user  + password + "|"))
    dataMap[user] = usersData
	return
}