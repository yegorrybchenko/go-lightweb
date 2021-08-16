package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/macmaczhl/golightweb/repositories"
)

type Books struct {
	repo *repositories.BooksRepository
}

func NewBooks() *Books {
	return &Books{repositories.NewBooksRepository()}
}

func (h *Books) Get(c *gin.Context) {
	id := c.Param("id")
	intId, _ := strconv.Atoi(id)
	book, err := h.repo.Get(intId)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, book)
}
