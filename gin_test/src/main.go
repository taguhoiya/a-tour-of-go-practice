package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	engine := gin.Default()

	ua := ""
	// ミドルウェアを使用
	engine.Use(func(c *gin.Context) {
		ua = c.GetHeader("User-Agent")
		c.Next()
	})

	engine.GET("/ua", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":    "hello world",
			"User-Agent": ua,
		})
	})

	// htmlのディレクトリを指定
	engine.LoadHTMLGlob("templates/*")
	engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			// htmlに渡す変数を定義
			"message": "hello gin",
		})
	})

	engine.Static("/static", "./static")

	engine.POST("/upload", func(c *gin.Context) {
		file, header, err := c.Request.FormFile("image")
		if err != nil {
			c.String(http.StatusBadRequest, "Bad request", file)
			return
		}
		fileName := header.Filename
		dir, _ := os.Getwd()
		out, err := os.Create(dir + "\\images\\" + fileName)
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()
		_, err = io.Copy(out, file)
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"header": header,
			"file":   file,
			"dir":    dir,
			"out":    out,
		})
	})

	engine.Run(":3000")
}
