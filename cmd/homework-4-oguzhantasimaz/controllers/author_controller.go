package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-oguzhantasimaz/pkg/models/authors"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-oguzhantasimaz/pkg/services"
	"github.com/gorilla/mux"
)

type AuthorController struct {
	service services.AuthorService
}

func NewAuthorController(repository authors.AuthorRepository) *AuthorController {
	return &AuthorController{
		service: *services.NewAuthorService(repository),
	}
}

func (c *AuthorController) GetAllAuthors(writer http.ResponseWriter, request *http.Request) {
	authors, err := c.service.GetAllAuthors()

	resp, _ := json.Marshal(authors)

	if err != nil {
		log.Print(err)
	}

	_, err = writer.Write(resp)
}

func (c *AuthorController) CreateAuthor(writer http.ResponseWriter, request *http.Request) {
	var author authors.Author

	err := json.NewDecoder(request.Body).Decode(&author)

	if err != nil {
		log.Print(err)
	}

	newAuthor, err := c.service.CreateAuthor(&author)

	if err != nil {
		log.Print(err)
	}

	resp, _ := json.Marshal(newAuthor)

	_, err = writer.Write(resp)
}

func (c *AuthorController) GetAuthorByID(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		log.Print(err)
	}

	author, err := c.service.GetAuthorByID(id)

	if err != nil {
		log.Print(err)
	}

	resp, _ := json.Marshal(author)

	_, err = writer.Write(resp)
}

func (c *AuthorController) UpdateAuthor(writer http.ResponseWriter, request *http.Request) {
	var author authors.Author

	err := json.NewDecoder(request.Body).Decode(&author)

	if err != nil {
		log.Print(err)
	}

	updatedAuthor, err := c.service.UpdateAuthor(&author)

	if err != nil {
		log.Print(err)
	}

	resp, _ := json.Marshal(updatedAuthor)

	_, err = writer.Write(resp)
}

func (c *AuthorController) DeleteAuthor(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		log.Print(err)
	}

	err = c.service.DeleteAuthor(id)

	if err != nil {
		log.Print(err)
	}
}
