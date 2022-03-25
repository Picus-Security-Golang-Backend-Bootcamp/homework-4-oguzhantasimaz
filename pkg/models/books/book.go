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

// GetAllBooks returns all books from the database
func GetAllBooks(r BookRepository) ([]*Book, error) {
	books, err := r.GetAllBooks()
	return books, err
}

//BuyBook gets id of book then look if its deleted and then checks cupdates after stock count is enough
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

// CreateBook creates a book in the database
func CreateBook(r BookRepository, book *Book) error {
	return r.CreateBook(book)
}

// GetBookByID returns a book from the database by id if its not deleted
func GetBookByID(r BookRepository, id int) (*Book, error) {
	book, err := r.GetBookByID(id)
	if book.IsDeleted {
		return nil, errors.New("Book is already deleted")
	}
	return book, err
}

//GetBookByTitle returns a book from the database by title if its not deleted
func GetBookByTitle(r BookRepository, title string) (*Book, error) {
	book, err := r.GetBookByTitle(title)
	if book.IsDeleted {
		return nil, errors.New("Book is already deleted")
	}
	return book, err
}

// UpdateBook updates a book in the database if its not deleted
func UpdateBook(r BookRepository, book *Book) error {
	if book.IsDeleted {
		return errors.New("Book is already deleted")
	}
	return r.UpdateBook(book)
}

// DeleteBook deletes a book from the database if its not deleted
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
