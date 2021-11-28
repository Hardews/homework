package main

import (
	"github.com/gin-gonic/gin"
)

func main()  {
	engine := gin.Default()

    engine.GET("/homework", func(c *gin.Context) {

	})

	engine.Run()
}