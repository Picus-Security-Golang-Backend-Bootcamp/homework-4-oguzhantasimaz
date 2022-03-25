package service

import "github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-oguzhantasimaz/pkg/models/authors"

// AuthorService struct
type AuthorService struct {
	repository authors.AuthorRepository
}

// CreateAuthorService function to create author service
func CreateAuthorService(repository authors.AuthorRepository) *AuthorService {
	return &AuthorService{
		repository: repository,
	}
}
