package books

import (
	"errors"
	"time"

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
	AuthorID   int            `gorm:"foreignkey:AuthorID"`
	Author     authors.Author `gorm:"references:ID"`
	CreatedAt  time.Time      `gorm:"<-:create"`
}

func (b *Book) Print() {
	log.Infof("\nBook:\n %d | %s | %s | %d | %d | %d | %f | %v ", b.ID, b.Title, b.StockCode, b.StockCount, b.Isbn, b.PageCount, b.Price, b.IsDeleted)
}

func GetAllBooks(r BookRepository) ([]*Book, error) {
	books, err := r.GetAllBooks()
	return books, err
}

//BuyBook get book then look if its deleted and then update stock count
func BuyBook(r BookRepository, id int, count int) error {
	book, err := r.GetBookByID(id)
	if err != nil {
		return err
	}
	if book.IsDeleted {
		return errors.New("Book is already deleted")
	}
	if book.StockCount < count {
		return errors.New("Not enough stock")
	}
	book.StockCount -= count

	err = r.BuyBook(book)
	return err
}

func CreateBook(r BookRepository, book *Book) error {
	return r.CreateBook(book)
}

func GetBookByID(r BookRepository, id int) (*Book, error) {
	book, err := r.GetBookByID(id)
	if book.IsDeleted {
		return nil, errors.New("Book is already deleted")
	}
	return book, err
}

func GetBookByTitle(r BookRepository, title string) (*Book, error) {
	book, err := r.GetBookByTitle(title)
	if book.IsDeleted {
		return nil, errors.New("Book is already deleted")
	}
	return book, err
}

func UpdateBook(r BookRepository, book *Book) error {
	if book.IsDeleted {
		return errors.New("Book is already deleted")
	}
	return r.UpdateBook(book)
}

func DeleteBook(r BookRepository, id int) error {
	book, err := r.GetBookByID(id)
	if err != nil {
		return err
	}
	if book.IsDeleted {
		return errors.New("Book is already deleted")
	}
	return r.DeleteBook(id)
}
