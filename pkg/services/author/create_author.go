package service

import (
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-oguzhantasimaz/pkg/models/authors"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-oguzhantasimaz/pkg/models/authors/validation"
)

type CreateAuthorRequest struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

func (r *CreateAuthorRequest) Validate() error {
	if r.Name == "" {
		return validation.ErrInvalidName
	}
	if r.Surname == "" {
		return validation.ErrInvalidSurname
	}
	return nil
}

func (a *AuthorService) CreateAuthor(req *CreateAuthorRequest) error {
	if err := req.Validate(); err != nil {
		return err
	}
	author := &authors.Author{
		Name:    req.Name,
		Surname: req.Surname,
	}
	return authors.CreateAuthor(a.repository, author)
}
