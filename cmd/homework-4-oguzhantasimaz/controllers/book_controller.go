package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-oguzhantasimaz/pkg/models/books"
	service "github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-oguzhantasimaz/pkg/services/book"
	"github.com/gorilla/mux"
)

type BookController struct {
	service service.BookService
}

type ApiResponse struct {
	Data interface{} `json:"data"`
}

func CreateBookController(repository books.BookRepository) *BookController {
	return &BookController{
		service: *service.CreateBookService(repository),
	}
}

func (c *BookController) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := c.service.GetAllBooks()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp, _ := json.Marshal(books)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(resp)

	if err != nil {
		log.Print(err)
	}
}

func (c *BookController) GetBookByID(w http.ResponseWriter, r *http.Request) {
	request := new(service.GetBookByIDRequest)
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	request.Id = id
	book, err := c.service.GetBookByID(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	resp, err := json.Marshal(book)
	if err != nil {
		log.Print(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(resp)
	if err != nil {
		log.Print(err)
		return
	}
}

func (c *BookController) GetBookByTitle(w http.ResponseWriter, r *http.Request) {
	request := new(service.GetBookByTitleRequest)
	vars := mux.Vars(r)
	title := vars["title"]

	request.Title = title
	book, err := c.service.GetBookByTitle(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	resp, err := json.Marshal(book)
	if err != nil {
		log.Print(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(resp)
	if err != nil {
		log.Print(err)
		return
	}
}

func (c *BookController) CreateBook(w http.ResponseWriter, r *http.Request) {
	request := new(service.CreateBookRequest)
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = c.service.CreateBook(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Book created successfully"))
}

func (c *BookController) UpdateBook(w http.ResponseWriter, r *http.Request) {
	request := new(service.UpdateBookRequest)
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = c.service.UpdateBook(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Book updated successfully"))
}

func (c *BookController) DeleteBook(w http.ResponseWriter, r *http.Request) {
	request := new(service.DeleteBookRequest)
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	request.ID = id
	err = c.service.DeleteBook(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Book deleted successfully"))
}

func (c *BookController) BuyBook(w http.ResponseWriter, r *http.Request) {
	request := new(service.BuyBookRequest)
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = c.service.BuyBook(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Book bought successfully"))
}
