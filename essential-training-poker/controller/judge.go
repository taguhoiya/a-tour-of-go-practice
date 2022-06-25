package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"poker/model"
	"poker/service"
)

func JudgeInput(c *gin.Context) {
	input := model.Input{}
	err := c.Bind(&input)
	if err != nil {
		c.String(http.StatusBadRequest, "Please input valid request body")
		return
	}
	result := service.Judge(input)
	c.JSON(http.StatusCreated, gin.H{
		"status": 200,
		"input":  input.Cards,
		"result": result,
	})
}
