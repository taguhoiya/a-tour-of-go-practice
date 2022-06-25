package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"qiita/model"
	"qiita/service"
	"strconv"
)

func BookAdd(c *gin.Context) {
	book := model.Book{}
	err := c.Bind(&book)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	bookService := service.BookService{}
	bookService.SetBook(&book)
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}

func BookList(c *gin.Context) {
	bookService := service.BookService{}
	BookLists := bookService.GetBookList()
	c.JSONP(http.StatusOK, gin.H{
		"message": "ok",
		"data":    BookLists,
	})
}

func BookUpdate(c *gin.Context) {
	book := model.Book{}
	err := c.Bind(&book)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	bookService := service.BookService{}
	bookService.UpdateBook(&book)
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}

func BookDelete(c *gin.Context) {
	id := c.PostForm("id")
	intId, err := strconv.ParseInt(id, 10, 0)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("Bad request: %v", err))
		return
	}
	bookService := service.BookService{}
	bookService.DeleteBook(int(intId))
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}
