package main

import (
	"github.com/HikaruEgashira/simple-server/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", controller.IndexHandler)
	router.GET("/user", controller.UserHandler)
	router.Run(":8080")
}
