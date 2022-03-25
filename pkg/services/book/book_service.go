package service

import "github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-oguzhantasimaz/pkg/models/books"

// BookService struct
type BookService struct {
	repository books.BookRepository
}

// CreateBookService function to create book service
func CreateBookService(repository books.BookRepository) *BookService {
	return &BookService{
		repository: repository,
	}
}
