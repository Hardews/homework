package main

import (
	"bufio"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main()  {
	menu()
}

type users struct {
	usersName string
	password string
}

var dataMap = make(map[string]users)

func readUsersData() map[string]users{
	var usersData users
	fileName := "D:/GOProjects/src/homework_at_redrock/Five/lv2/Data.txt"
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
	var engine = gin.Default()
	var userOperation = engine.Group("/user")
	fmt.Println("-----主菜单-----")
	fmt.Println("请选择您的操作")
	fmt.Printf("1.登录\n2.注册\n3.退出\n")
	var chose int
	fmt.Scan(&chose)
	switch chose {
	case 1:
		userOperation.GET("/login",Login)
	case 2:
		userOperation.POST("/register", Register)
	case 3:
		return
	}
	engine.Run()

}

func Login(c *gin.Context)  {
	    var user users
	    readUsersData()
	    c.Writer.Write([]byte("-----登陆页面----- \n" ))
		user.usersName = c.Query("username")
		user.password  = c.Query("password")
		user.usersName = user.usersName + " "
		user.password = user.password + "|"
		if dataMap[user.usersName]==user{
			c.Writer.Write([]byte(user.usersName + "登录成功"))
		}else {
			c.Writer.Write([]byte(user.usersName + "登录失败"))
		}
}

func Register(c *gin.Context)  {
		readUsersData()
		c.Writer.Write([]byte("-----注册页面-----\n"))
		var username = c.PostForm("username")
		var password = c.PostForm("password")
		_,ok := dataMap[username + " "]
		if ok{
			c.Writer.Write([]byte(username + "注册失败,用户名已存在"))
		}else if len(password) < 6{
			c.Writer.Write([]byte(username + "注册失败，密码长度不足"))
		}else {
			writeIn(username,password)
			c.Writer.Write([]byte(username + "注册成功"))
		}
}

func writeIn(user,password string)  {
	var usersData users
	usersData.usersName = user
	usersData.password  = password
	fileName := "D:/GOProjects/src/homework_at_redrock/Five/lv2/Data.txt"
	file,err := os.OpenFile(fileName,os.O_RDWR|os.O_APPEND,os.ModePerm)
	if err!=nil{
		fmt.Println("err:",err)
	}
	file.Write([]byte(  user  + " "+ password + "|"))
	dataMap[user] = usersData
	return
}