package books

import (
	"gorm.io/gorm"

	authors "github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-oguzhantasimaz/pkg/models/authors"
	log "github.com/sirupsen/logrus"
)

// Book represents a book
type Book struct {
	gorm.Model
	ID         int
	Title      string
	StockCode  string
	StockCount int
	Isbn       int
	PageCount  int
	Price      float64
	IsDeleted  bool
	AuthorID   int
	Author     authors.Author `gorm:"references:ID"`
}

func (b *Book) Print() {
	log.Infof("\nBook:\n %d | %s | %s | %d | %d | %d | %f | %v ", b.ID, b.Title, b.StockCode, b.StockCount, b.Isbn, b.PageCount, b.Price, b.IsDeleted)
}
