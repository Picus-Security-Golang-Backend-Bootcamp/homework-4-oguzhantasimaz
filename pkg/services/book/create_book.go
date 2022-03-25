package service

import (
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-oguzhantasimaz/pkg/models/books"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-oguzhantasimaz/pkg/models/books/validation"
)

type CreateBookRequest struct {
	Title      string  `json:"title"`
	StockCode  string  `json:"stock_code"`
	StockCount int     `json:"stock_count"`
	Isbn       int     `json:"isbn"`
	PageCount  int     `json:"page_count"`
	Price      float64 `json:"price"`
	IsDeleted  bool    `json:"is_deleted"`
	AuthorID   int     `json:"author_id"`
}

func (r *CreateBookRequest) Validate() error {
	if r.Title == "" {
		return validation.ErrInvalidTitle
	}
	if r.StockCode == "" {
		return validation.ErrInvalidStockCode
	}
	if r.StockCount < 0 {
		return validation.ErrInvalidStockCount
	}
	if r.Isbn < 0 {
		return validation.ErrInvalidIsbn
	}
	if r.PageCount < 0 {
		return validation.ErrInvalidPageCount
	}
	if r.Price < 0 {
		return validation.ErrInvalidPrice
	}
	if r.AuthorID < 0 {
		return validation.ErrInvalidAuthorID
	}
	return nil
}

func (b *BookService) CreateBook(req *CreateBookRequest) error {
	if err := req.Validate(); err != nil {
		return err
	}
	book := &books.Book{
		Title:      req.Title,
		StockCode:  req.StockCode,
		StockCount: req.StockCount,
		Isbn:       req.Isbn,
		PageCount:  req.PageCount,
		Price:      req.Price,
		IsDeleted:  req.IsDeleted,
		AuthorID:   req.AuthorID,
	}
	return books.CreateBook(b.repository, book)
}
