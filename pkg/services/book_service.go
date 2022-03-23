package services

import books "github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-oguzhantasimaz/pkg/models/books"

type BookService struct {
	repository books.BookRepository
}

func NewBookService(repository books.BookRepository) *BookService {
	return &BookService{
		repository: repository,
	}
}
func (b *BookService) GetAllBooks() ([]*books.Book, error) {
	return b.repository.GetAllBooks()
}

func (b *BookService) CreateBook(book *books.Book) (*books.Book, error) {
	return b.repository.CreateBook(book)
}

func (b *BookService) GetBookByID(id int) (*books.Book, error) {
	return b.repository.GetBookByID(id)
}

func (b *BookService) GetBookByTitle(title string) (*books.Book, error) {
	return b.repository.GetBookByTitle(title)
}

func (b *BookService) BuyBook(id int, count int) (*books.Book, error) {
	return b.repository.BuyBook(id, count)
}

func (b *BookService) UpdateBook(book *books.Book) (*books.Book, error) {
	return b.repository.UpdateBook(book)
}

func (b *BookService) DeleteBook(id int) error {
	return b.repository.DeleteBook(id)
}
