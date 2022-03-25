package validation

import "errors"

var (
	ErrInvalidID         = errors.New("ID must be greater than 0")
	ErrInvalidCount      = errors.New("Count must be greater than 0")
	ErrInvalidTitle      = errors.New("Title must not be empty")
	ErrInvalidStockCode  = errors.New("StockCode must not be empty")
	ErrInvalidStockCount = errors.New("StockCount must be greater than 0")
	ErrInvalidIsbn       = errors.New("Isbn must be greater than 0")
	ErrInvalidPageCount  = errors.New("PageCount must be greater than 0")
	ErrInvalidPrice      = errors.New("Price must be greater than 0")
	ErrInvalidAuthorID   = errors.New("AuthorID must be greater than 0")
)
