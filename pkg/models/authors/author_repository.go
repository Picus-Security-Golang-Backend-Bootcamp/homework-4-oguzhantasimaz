package authors

// AuthorRepository interface to define methods for author repository
type AuthorRepository interface {
	Migration()
	InsertSampleData()
	GetAllAuthors() ([]*Author, error)
	GetAuthorByID(id int) (*Author, error)
	CreateAuthor(author *Author) error
	UpdateAuthor(author *Author) error
	DeleteAuthor(id int) error
}
