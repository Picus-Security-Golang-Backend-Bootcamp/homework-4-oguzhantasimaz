package authors

type AuthorRepository interface {
	Migration()
	InsertSampleData()
	GetAllAuthors() ([]*Author, error)
	GetAuthorByID(id int) (*Author, error)
	CreateAuthor(author *Author) (*Author, error)
	UpdateAuthor(author *Author) (*Author, error)
	DeleteAuthor(id int) error
}
