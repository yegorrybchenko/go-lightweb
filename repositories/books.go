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

func (r *BooksRepository) Get(name string) (*models.Book, error) {
	book := models.Book{Name: name}
	return &book, r.db.First(&book).Error
}
