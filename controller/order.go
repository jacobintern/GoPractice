package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jacobintern/GoPractice/service"
)

// func GetOrderByUser(c *gin.Context) {
// 	name, _ := strconv.Atoi(c.Param("name"))
// 	c.JSON(200, gin.H{
// 		"data": service.GetOrderList(id),
// 	})
// }

func GetOrder(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": service.GetOrderList(),
	})
}

func GetOrderById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	c.JSON(200, gin.H{
		"data": service.GetOrder(id),
	})
}

func CreateOrder(c *gin.Context) {

	c.JSON(200, gin.H{
		"data":    nil,
		"message": "OK",
	})
}
