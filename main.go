package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jacobintern/GoPractice/controller"
)

func main() {
	app := gin.Default()

	app.GET("/api/orders", controller.GetOrder)
	app.GET("/api/order/:id", controller.GetOrderById)
	app.GET("/api/match", controller.MatchList)
	// app.GET("/api/order/:name", controller.GetOrderByUser)

	app.Run(":5001")
}
