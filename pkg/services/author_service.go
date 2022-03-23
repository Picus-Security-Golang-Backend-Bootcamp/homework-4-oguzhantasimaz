package services

import "github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-oguzhantasimaz/pkg/models/authors"

type AuthorService struct {
	repository authors.AuthorRepository
}

func NewAuthorService(repository authors.AuthorRepository) *AuthorService {
	return &AuthorService{
		repository: repository,
	}
}

func (a *AuthorService) GetAllAuthors() ([]*authors.Author, error) {
	return a.repository.GetAllAuthors()
}

func (a *AuthorService) CreateAuthor(author *authors.Author) (*authors.Author, error) {
	return a.repository.CreateAuthor(author)
}

func (a *AuthorService) GetAuthorByID(id int) (*authors.Author, error) {
	return a.repository.GetAuthorByID(id)
}

func (a *AuthorService) UpdateAuthor(author *authors.Author) (*authors.Author, error) {
	return a.repository.UpdateAuthor(author)
}

func (a *AuthorService) DeleteAuthor(id int) error {
	return a.repository.DeleteAuthor(id)
}
