package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jacobintern/GoPractice/controller"
)

func main() {
	app := gin.Default()

	app.GET("/api/orders", controller.GetOrder)
	app.GET("/api/userOrderDetail", controller.GetUserOrder)
	app.Run(":5001")
}
