package service

import (
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-oguzhantasimaz/pkg/models/authors"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-oguzhantasimaz/pkg/models/books/validation"
)

type DeleteAuthorRequest struct {
	ID int `json:"id"`
}

func (r *DeleteAuthorRequest) Validate() error {
	if r.ID == 0 {
		return validation.ErrInvalidID
	}
	return nil
}

func (a *AuthorService) DeleteAuthor(req *DeleteAuthorRequest) error {
	if err := req.Validate(); err != nil {
		return err
	}
	return authors.DeleteAuthor(a.repository, req.ID)
}
