package service

import "github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-oguzhantasimaz/pkg/models/books"

func (b *BookService) GetAllBooks() ([]*books.Book, error) {
	return books.GetAllBooks(b.repository)
}
