package service

import (
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-oguzhantasimaz/pkg/models/authors"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-oguzhantasimaz/pkg/models/authors/validation"
)

type GetAuthorByIdRequest struct {
	Id int `json:"id"`
}

func (r *GetAuthorByIdRequest) Validate() error {
	if r.Id == 0 {
		return validation.ErrInvalidID
	}
	return nil
}

func (a *AuthorService) GetAuthorByID(request *GetAuthorByIdRequest) (*authors.Author, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}
	return authors.GetAuthorByID(a.repository, request.Id)
}
