package repositories

import (
	"gorm.io/gorm"

	"github.com/macmaczhl/golightweb/config/db"
	"github.com/macmaczhl/golightweb/models"
)

type BooksRepository struct {
	db *gorm.DB
}

func NewBooksRepository() *BooksRepository {
	return &BooksRepository{db.Connection()}
}

func (r *BooksRepository) Get(id int) (*models.Book, error) {
	book := models.Book{Id: id}
	return &book, r.db.First(&book).Error
}
