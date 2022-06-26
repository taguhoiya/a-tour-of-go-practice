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
			v1.GET("/todays_weather", controller.NeedUmbrella)
		}
	}
	engine.NoRoute(func(c *gin.Context) {
    c.JSON(404, gin.H{"status": 404, "error": "エンドポイントが不正です。"})
})
	engine.Run(":3000")
}

func ErrorHandle() {
	router := gin.Default()

	router.NoRoute(func (c *gin.Context) {

	})
}
