package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Books(c *gin.Context) {
	c.String(http.StatusOK, "Hello world")
}
