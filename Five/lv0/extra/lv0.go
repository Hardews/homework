package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main()  {
	engine := gin.Default()

	//GET请求
	engine.GET("/ping", func(context *gin.Context) {
		context.JSON(200,gin.H{
			"message": "pong",
		})
	})

	engine.GET("/user/:name", func(context *gin.Context) {
		name := context.Param("name")
		context.String(http.StatusOK,"Hello %s",name)
	})

	engine.GET("/welcome", func(context *gin.Context) {
		firstName := context.DefaultQuery("firstName","Guest")
		lastName  := context.Query("lastName")

		context.String(http.StatusOK,"Hello %s %s",firstName,lastName)
	})

	//
	engine.POST("/form_post", func(context *gin.Context) {
		message := context.PostForm("message")
		nick := context.DefaultQuery("nick","anonymous")
		context.JSON(200,gin.H{
			"status":"posted",
			"message":message,
			"nick":nick,
		})
	})

	//POST请求
	engine.POST("/login", func(context *gin.Context) {
		var u login
		if err := context.ShouldBind(&u); err != nil{
			context.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
			return
		}
		context.JSON(200,gin.H{
			"status":"you are logged in",
			"user":u,
		})
	})
    //JSON渲染
    engine.GET("/someJSON", func(context *gin.Context) {
		context.JSON(200,gin.H{"message":"hey","status":http.StatusOK})
	})

	//string型
	engine.GET("/someString", func(c *gin.Context) {
		c.String(200,"hello world!")
	})

	//路由分组
	v1 := engine.Group("/v1")
	{
		//path：127.0.0.1:8080/v1/hello
		v1.GET("hello",func(c *gin.Context){
			c.String(200,"hello this is v1")
		})
	}

	// Simple group: v2
	v2 := engine.Group("/v2")
	{
		//path：127.0.0.1:8080/v2/hello
		v2.GET("hello",func(c *gin.Context){
			c.String(200,"hello this is v2")
		})
	}

	engine.GET("/test", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently,"http://www.baidu.com/")
	})

	engine.Run()
}

type login struct {
	User string `form:"user" json:"user" xml:"user" binding:"-"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}