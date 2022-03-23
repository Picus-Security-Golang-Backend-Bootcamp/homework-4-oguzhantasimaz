package authors

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// Author represents an author
type Author struct {
	gorm.Model
	ID      int
	Name    string
	Surname string
}

func (b *Author) Print() {
	log.Infof("\nAuthor:\n %s | %s", b.Name, b.Surname)
}

func NewAuthor(r AuthorRepository, author *Author) (*Author, error) {
	return r.CreateAuthor(author)
}

func GetAllAuthors(r AuthorRepository) ([]*Author, error) {
	return r.GetAllAuthors()
}

func GetAuthorByID(r AuthorRepository, id int) (*Author, error) {
	return r.GetAuthorByID(id)
}

func UpdateAuthor(r AuthorRepository, author *Author) (*Author, error) {
	return r.UpdateAuthor(author)
}

func DeleteAuthor(r AuthorRepository, id int) error {
	return r.DeleteAuthor(id)
}
