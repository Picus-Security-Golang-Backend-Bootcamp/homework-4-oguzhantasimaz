package service

import (
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-oguzhantasimaz/pkg/models/books"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-oguzhantasimaz/pkg/models/books/validation"
)

type DeleteBookRequest struct {
	ID int `json:"id"`
}

func (r *DeleteBookRequest) Validate() error {
	if r.ID <= 0 {
		return validation.ErrInvalidID
	}
	return nil
}

func (b *BookService) DeleteBook(req *DeleteBookRequest) error {
	if err := req.Validate(); err != nil {
		return err
	}
	return books.DeleteBook(b.repository, req.ID)
}
