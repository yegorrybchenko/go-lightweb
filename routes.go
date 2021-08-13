package main

import (
	"github.com/gin-gonic/gin"

	"github.com/macmaczhl/golightweb/handlers"
)

func InitServer() *gin.Engine {
	router := gin.Default()

	router.GET("/hello", handlers.HelloWorld)

	return router
}
