package main

import (
	"github.com/gin-gonic/gin"
	"poker/controller"
	"poker/middleware"
)

func main() {
	engine := gin.Default()
	// ミドルウェア
	engine.Use(middleware.RecordUaAndTime)
	// CRUD 書籍
	bookEngine := engine.Group("/api")
	{
		v1 := bookEngine.Group("/v1")
		{
			v1.POST("/judge", controller.JudgeInput)
		}
	}
	engine.Run(":3000")
}
