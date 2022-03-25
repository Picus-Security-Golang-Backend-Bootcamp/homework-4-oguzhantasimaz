package service

import "github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-oguzhantasimaz/pkg/models/authors"

type AuthorService struct {
	repository authors.AuthorRepository
}

func CreateAuthorService(repository authors.AuthorRepository) *AuthorService {
	return &AuthorService{
		repository: repository,
	}
}
