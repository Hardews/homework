package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
)
var db,err = sql.Open("mysql","root:lmh123@tcp(127.0.0.1:3306)/userdata")

func main()  {
	if err!=nil{
		log.Fatal(err)
	}
	menu()
}

type users struct {
	id int
	usersName string
	password string
	question string
	answer string
}

func menu()  {
	var engine = gin.Default()
	var userOperation = engine.Group("/user")
	userOperation.POST("/login",Login)
	userOperation.POST("/register", register)
	userOperation.POST("/FindPassword",FindPassword)
	engine.Run()
}

func FindPassword(c *gin.Context)  {
	var user users
	var question,answer string

	var usersName = c.PostForm("username")
	var newPassword = c.PostForm("newPassword")
	question = c.PostForm("question")
	answer = c.PostForm("answer")

	sqlStr1 := "select username from userInfo where username = ?"
	err := db.QueryRow(sqlStr1,usersName).Scan(&user.usersName)
	if err!=nil{
		if err == sql.ErrNoRows{
			c.Writer.Write([]byte("用户名不存在"))
			return
		}else {
			fmt.Println("select failed, err :",err)
			return
		}
	}

	sqlStr2 := "select question,answer from question where username = ? and question = ?"
	err = db.QueryRow(sqlStr2,usersName,question).Scan(&question,&user.answer)
	c.Writer.Write([]byte("问题：" + question + "\n"))
	if answer==user.answer{
		goto update
	}else {
		c.Writer.Write([]byte("答案错误"))
		return
	}

	update:
		sqlStr3 := "update userInfo set password=? where username=?"
		_,err = db.Exec(sqlStr3,newPassword,user.usersName)
		if err!=nil{
			fmt.Println("update failed, err: ",err)
			return
		}
		c.Writer.Write([]byte("修改成功！"))

}


func Login(c *gin.Context)  {
	var user users
	c.Writer.Write([]byte("-----登陆页面----- \n" ))
	var usersName = c.PostForm("username")
	var password  = c.PostForm("password")
    sqlStr := "select username,password from userInfo where username = ?"
	err := db.QueryRow(sqlStr,usersName).Scan(&user.usersName,&user.password)
	if err!=nil{
		if err == sql.ErrNoRows{
			c.Writer.Write([]byte("用户名不存在"))
			return
		}else {
			fmt.Println("select failed, err :",err)
			return
		}
	}
	if user.usersName==usersName && user.password==password{
		c.Writer.Write([]byte("登陆成功！"))
	}else {
		c.Writer.Write([]byte("密码错误!"))
	}
}

func register(c *gin.Context)  {
	var u users
	c.Writer.Write([]byte("-----注册页面-----\n"))
	var username = c.PostForm("username")
	var password = c.PostForm("password")
	sqlStr := "select username from userInfo where username = ?"
	err := db.QueryRow(sqlStr,username).Scan(&u.usersName)
	if err!=nil{
		if err == sql.ErrNoRows{
			    goto register
		}else {
				fmt.Println("select failed, err :",err)
			    return
			}
		}
		register:
		if username == u.usersName{
			c.Writer.Write([]byte("用户名已存在"))
			return
		}
        if len(password) < 6 {
	    c.Writer.Write([]byte("密码长度不足"))
			return
		}
	    writeIn(username,password)
	    c.Writer.Write([]byte("注册成功！"))

}

func writeIn(username,password string)  {
	sqlStr := "insert userInfo (username,password) values (?,?)"
	stmt,err := db.Prepare(sqlStr)
	if err != nil{
		fmt.Println("prepare failed , err :",err)
		return
	}
	defer stmt.Close()
	_,err = stmt.Exec(username,password)
	if err != nil{
		fmt.Println("insert failed , err :",err)
		return
	}
	return
}

func auth(c *gin.Context){
	 _,err := c.Cookie("login_cookie")
	 if err!=nil{
		 c.JSON(403,gin.H{
			" message":"未登录",
		 })
		 return
	 }else{
		 return
	 }
}