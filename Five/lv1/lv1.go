package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main()  {
	engine := gin.Default()
	auth := func(c *gin.Context){
		//获取cookie
		value, err := c.Cookie("gin_cookie")
		if err != nil {
			c.JSON(403,gin.H{
				"message":"认证失败,没有cookie",
			})
			//认证失败
			//终止的意思，也就是说执行该函数，会终止后面所有的该请求下的函数
			c.Abort()
		}else{
			//将获取到的cookie的值写入上下文
			c.Set("cookie",value)
			v,_:=c.Get("next")
			fmt.Println(v)
		}
	}
	engine.POST("/login",func (c *gin.Context){
		username:=c.PostForm("username")
		password:=c.PostForm("password")
		if username == "123" && password == "321"{
			c.SetCookie("gin_cookie", username, 3600, "/", "", false, true)
			c.JSON(200,gin.H{
				"msg": "login successfully",
			})
		}else{
			c.JSON(403,gin.H{
				"message":"认证失败,账号密码错误",
			})
		}
	})
	//在中间放入鉴权中间件
	engine.GET("/hello",auth, func(c *gin.Context) {
		cookie,_:=c.Get("cookie")
		str:=cookie.(string)
		c.String(200,"hello world"+str)
		//测试next函数
		c.Set("next","test next")
	})


	engine.Run()
}
