package controller

import (
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
