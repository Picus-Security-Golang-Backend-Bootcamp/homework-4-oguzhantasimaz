package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-oguzhantasimaz/pkg/models/authors"
	service "github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-oguzhantasimaz/pkg/services/author"
	"github.com/gorilla/mux"
)

// AuthorController struct
type AuthorController struct {
	service service.AuthorService
}

// CreateAuthorController function to create author controller
func CreateAuthorController(repository authors.AuthorRepository) *AuthorController {
	return &AuthorController{
		service: *service.CreateAuthorService(repository),
	}
}

// GetAllAuthors function to get all authors from application layer
func (c *AuthorController) GetAllAuthors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	authors, err := c.service.GetAllAuthors()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp, _ := json.Marshal(authors)

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(resp)

	if err != nil {
		log.Print(err)
		return
	}
}

// CreateAuthor function to create author in application layer
func (c *AuthorController) CreateAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	request := new(service.CreateAuthorRequest)

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.service.CreateAuthor(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Author created successfully"))
}

// GetAuthorByID function to get author by id from application layer
func (c *AuthorController) GetAuthorByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	request := new(service.GetAuthorByIdRequest)
	request.Id = id
	author, err := c.service.GetAuthorByID(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	resp, err := json.Marshal(author)
	if err != nil {
		log.Print(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(resp)
	if err != nil {
		log.Print(err)
		return
	}
}

// UpdateAuthor function to update author in application layer
func (c *AuthorController) UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	request := new(service.UpdateAuthorRequest)
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.service.UpdateAuthor(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Author updated successfully"))
}

// DeleteAuthor function to delete author in application layer
func (c *AuthorController) DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	request := new(service.DeleteAuthorRequest)
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	request.ID = id
	err = c.service.DeleteAuthor(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Author deleted successfully"))
}
