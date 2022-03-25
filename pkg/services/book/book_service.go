package service

import "github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-oguzhantasimaz/pkg/models/books"

type BookService struct {
	repository books.BookRepository
}

func CreateBookService(repository books.BookRepository) *BookService {
	return &BookService{
		repository: repository,
	}
}
