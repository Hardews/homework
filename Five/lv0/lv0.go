package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main()  {
	engine := gin.Default()

	engine.GET("/hello", func(c *gin.Context) {
		path := c.FullPath()
		fmt.Println("请求路径" + path)
		c.Writer.Write([]byte("hello"))
	})

	engine.Run()
}
