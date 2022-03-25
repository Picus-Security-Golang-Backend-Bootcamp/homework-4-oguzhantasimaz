package service

import "github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-oguzhantasimaz/pkg/models/authors"

func (a *AuthorService) GetAllAuthors() ([]*authors.Author, error) {
	return authors.GetAllAuthors(a.repository)
}
