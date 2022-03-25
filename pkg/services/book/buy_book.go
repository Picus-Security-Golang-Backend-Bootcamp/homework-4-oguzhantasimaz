package service

import (
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-oguzhantasimaz/pkg/models/books"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-oguzhantasimaz/pkg/models/books/validation"
)

type BuyBookRequest struct {
	ID    int `json:"id"`
	Count int `json:"count"`
}

func (r *BuyBookRequest) Validate() error {
	if r.ID <= 0 {
		return validation.ErrInvalidID
	}
	if r.Count <= 0 {
		return validation.ErrInvalidCount
	}
	return nil
}

func (b *BookService) BuyBook(req *BuyBookRequest) error {
	if err := req.Validate(); err != nil {
		return err
	}
	return books.BuyBook(b.repository, req.ID, req.Count)
}
