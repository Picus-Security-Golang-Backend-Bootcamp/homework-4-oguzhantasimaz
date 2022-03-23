package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-oguzhantasimaz/pkg/models/books"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-oguzhantasimaz/pkg/services"
	"github.com/gorilla/mux"
)

type BookController struct {
	service services.BookService
}

type ApiResponse struct {
	Data interface{} `json:"data"`
}

func NewBookController(repository books.BookRepository) *BookController {
	return &BookController{
		service: *services.NewBookService(repository),
	}
}

func (c *BookController) GetAllBooks(writer http.ResponseWriter, request *http.Request) {
	books, err := c.service.GetAllBooks()

	resp, _ := json.Marshal(books)

	if err != nil {
		log.Print(err)
	}

	_, err = writer.Write(resp)
}

func (c *BookController) GetBookByID(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		log.Print(err)
	}

	book, err := c.service.GetBookByID(id)

	if err != nil {
		log.Print(err)
	}
	resp, _ := json.Marshal(book)

	_, err = writer.Write(resp)

}
