package service

import (
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-oguzhantasimaz/pkg/models/authors"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-oguzhantasimaz/pkg/models/authors/validation"
)

type UpdateAuthorRequest struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

func (r *UpdateAuthorRequest) Validate() error {
	if r.ID == 0 {
		return validation.ErrInvalidID
	}
	if r.Name == "" {
		return validation.ErrInvalidName
	}
	if r.Surname == "" {
		return validation.ErrInvalidSurname
	}
	return nil
}

func (a *AuthorService) UpdateAuthor(req *UpdateAuthorRequest) error {
	if err := req.Validate(); err != nil {
		return err
	}
	author := &authors.Author{
		ID:      req.ID,
		Name:    req.Name,
		Surname: req.Surname,
	}
	return authors.UpdateAuthor(a.repository, author)
}
