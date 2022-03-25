package service

import (
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-oguzhantasimaz/pkg/models/authors/validation"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-oguzhantasimaz/pkg/models/books"
)

type GetBookByIDRequest struct {
	Id int `json:"id"`
}

func (r GetBookByIDRequest) Validate() error {
	if r.Id <= 0 {
		return validation.ErrInvalidID
	}
	return nil
}

func (b *BookService) GetBookByID(request *GetBookByIDRequest) (*books.Book, error) {
	err := request.Validate()
	if err != nil {
		return nil, err
	}
	return books.GetBookByID(b.repository, request.Id)
}
