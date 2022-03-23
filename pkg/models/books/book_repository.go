package books

type BookRepository interface {
	Migration()
	InsertSampleData()
	GetAllBooks() ([]*Book, error)
	GetBookByID(id int) (*Book, error)
	GetBookByTitle(title string) (*Book, error)
	BuyBook(id int, count int) (*Book, error)
	CreateBook(book *Book) (*Book, error)
	UpdateBook(book *Book) (*Book, error)
	DeleteBook(id int) error
}
