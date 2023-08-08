package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jacobintern/GoPractice/service"
)

func GetUserOrder(c *gin.Context) {

}

func GetOrder(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": service.OrderList(),
	})
}

func GetOne(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	c.JSON(200, gin.H{
		"data": service.OrderOne(id),
	})
}

func CreateOrder(c *gin.Context) {

	c.JSON(200, gin.H{
		"data":    nil,
		"message": "OK",
	})
}
