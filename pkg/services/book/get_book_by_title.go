package service

import (
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-oguzhantasimaz/pkg/models/books"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-oguzhantasimaz/pkg/models/books/validation"
)

type GetBookByTitleRequest struct {
	Title string `json:"title"`
}

func (r GetBookByTitleRequest) Validate() error {
	if r.Title == "" {
		return validation.ErrInvalidTitle
	}
	return nil
}

func (b *BookService) GetBookByTitle(request *GetBookByTitleRequest) (*books.Book, error) {
	return books.GetBookByTitle(b.repository, request.Title)
}
