package main

import (
	"github.com/gin-gonic/gin"

	"github.com/macmaczhl/golightweb/handlers"
)

func InitServer() *gin.Engine {
	booksHandler := handlers.NewBooks()

	router := gin.Default()

	router.GET("/books/:id", booksHandler.Get)

	return router
}
