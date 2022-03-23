package book

import (
	"errors"

	authors "github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-oguzhantasimaz/pkg/models/authors"
	books "github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-oguzhantasimaz/pkg/models/books"
	"gorm.io/gorm"

	log "github.com/sirupsen/logrus"
)

type bookRepository struct {
	db *gorm.DB
}

// Create a new book repository
func NewBookRepository(db *gorm.DB) books.BookRepository {
	return &bookRepository{
		db: db,
	}
}

// Migrates the database to the latest schema
func (b *bookRepository) Migration() {
	b.db.Migrator().DropTable(&books.Book{})
	b.db.AutoMigrate(&books.Book{})
}

// GetAllBooks returns all books
func (b *bookRepository) GetAllBooks() ([]*books.Book, error) {
	books := make([]*books.Book, 0)
	err := b.db.Find(&books).Error
	return books, err
}

// GetBookByID returns a book by id
func (b *bookRepository) GetBookByID(id int) (*books.Book, error) {
	book := new(books.Book)
	err := b.db.Where("id = ?", id).First(book).Error
	return book, err
}

// GetBookByTitle returns a book by title
func (b *bookRepository) GetBookByTitle(title string) (*books.Book, error) {
	book := new(books.Book)
	//parameter title includes wildcard character '%'
	err := b.db.Where("title LIKE ?", "%"+title+"%").First(book).Error
	return book, err
}

// CreateBook creates a new book
func (b *bookRepository) CreateBook(book *books.Book) (*books.Book, error) {
	err := b.db.Create(book).Error
	return book, err
}

// UpdateBook updates a book
func (b *bookRepository) UpdateBook(book *books.Book) (*books.Book, error) {
	//control if book deleted
	if book.IsDeleted == true {
		return nil, errors.New("You cannot update deleted book")
	}
	err := b.db.Save(book).Error
	return book, err
}

// DeleteBook soft deletes a book
func (b *bookRepository) DeleteBook(id int) error {
	//update book isDeleted field to true
	book, err := b.GetBookByID(id)
	if err != nil {
		return err
	}
	//control if book deleted
	if book.IsDeleted == true {
		return errors.New("book already deleted")
	}

	book.IsDeleted = true
	book, err = b.UpdateBook(book)
	if err != nil {
		return err
	}
	return err
}

// BuyBook updates a book stock count
func (b *bookRepository) BuyBook(id int, count int) (*books.Book, error) {
	book, err := b.GetBookByID(id)
	if err != nil {
		return nil, err
	}

	if book.IsDeleted == true {
		return nil, errors.New("You cannot buy deleted book")
	}
	if book.StockCount < count {
		log.Fatal("Not enough stock")
	}
	book.StockCount -= count
	book, err = b.UpdateBook(book)
	if err != nil {
		return nil, err
	}
	return book, err
}

// InsertSampleData inserts sample data of books
func (r *bookRepository) InsertSampleData() {
	bookList := []books.Book{
		{Title: "Beyaz Ev", StockCode: "GPRG", StockCount: 10, Isbn: 123456789, PageCount: 200, Price: 10.0, IsDeleted: true, Author: authors.Author{Name: "Erkut", Surname: "Tackin"}},
		{Title: "Resimdeki Gozyasi", StockCode: "CPRG", StockCount: 10, Isbn: 123456788, PageCount: 200, Price: 10.0, IsDeleted: false, Author: authors.Author{Name: "Cem", Surname: "Karaca"}},
		{Title: "Hayat Bir Teselli", StockCode: "JPRG", StockCount: 10, Isbn: 123456787, PageCount: 200, Price: 10.0, IsDeleted: false, Author: authors.Author{Name: "Erkin", Surname: "Koray"}},
		{Title: "Unutamadim", StockCode: "PPRG", StockCount: 10, Isbn: 123456786, PageCount: 200, Price: 10.0, IsDeleted: false, Author: authors.Author{Name: "Baris", Surname: "Manco"}},
		{Title: "Amasra", StockCode: "RPRG", StockCount: 10, Isbn: 123456785, PageCount: 200, Price: 10.0, IsDeleted: false, Author: authors.Author{Name: "Baris", Surname: "Akarsu"}},
	}
	r.db.Create(&bookList)
}
