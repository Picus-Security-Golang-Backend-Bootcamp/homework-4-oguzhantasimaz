package books

// BookRepository interface to define methods for book repository
type BookRepository interface {
	Migration()
	InsertSampleData()
	GetAllBooks() ([]*Book, error)
	GetBookByID(id int) (*Book, error)
	GetBookByTitle(title string) (*Book, error)
	BuyBook(book *Book) error
	CreateBook(book *Book) error
	UpdateBook(book *Book) error
	DeleteBook(id int) error
}
