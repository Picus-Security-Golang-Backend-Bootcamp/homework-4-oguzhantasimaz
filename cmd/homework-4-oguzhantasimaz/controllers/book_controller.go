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
	writer.Header().Set("Content-Type", "application/json")
	books, err := c.service.GetAllBooks()

	resp, _ := json.Marshal(books)

	if err != nil {
		log.Print(err)
	}

	_, err = writer.Write(resp)
}

func (c *BookController) GetBookByID(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
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

func (c *BookController) GetBookByTitle(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(request)
	title := vars["title"]

	book, err := c.service.GetBookByTitle(title)

	if err != nil {
		log.Print(err)
	}
	resp, _ := json.Marshal(book)

	_, err = writer.Write(resp)
}

func (c *BookController) CreateBook(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var book books.Book
	err := json.NewDecoder(request.Body).Decode(&book)

	if err != nil {
		log.Print(err)
	}

	newBook, err := c.service.CreateBook(&book)

	if err != nil {
		log.Print(err)
	}

	resp, _ := json.Marshal(newBook)

	_, err = writer.Write(resp)
}

func (c *BookController) UpdateBook(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var book books.Book
	err := json.NewDecoder(request.Body).Decode(&book)

	if err != nil {
		log.Print(err)
	}

	updatedBook, err := c.service.UpdateBook(&book)

	if err != nil {
		log.Print(err)
	}

	resp, _ := json.Marshal(updatedBook)

	_, err = writer.Write(resp)
}

func (c *BookController) DeleteBook(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		log.Print(err)
	}

	err = c.service.DeleteBook(id)

	if err != nil {
		log.Print(err)
	}
}

func (c *BookController) BuyBook(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	//get id and count from body
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	count, err := strconv.Atoi(vars["count"])

	if err != nil {
		log.Print(err)
	}

	boughtBook, err := c.service.BuyBook(id, count)

	if err != nil {
		log.Print(err)
	}

	resp, _ := json.Marshal(boughtBook)

	_, err = writer.Write(resp)
}
