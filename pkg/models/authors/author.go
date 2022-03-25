package authors

import (
	"time"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// Author represents an author
type Author struct {
	gorm.Model
	ID        int
	Name      string
	Surname   string
	CreatedAt time.Time `gorm:"<-:create"`
}

func (b *Author) Print() {
	log.Infof("\nAuthor:\n %s | %s", b.Name, b.Surname)
}

// CreateAuthor creates an author in the database
func CreateAuthor(r AuthorRepository, author *Author) error {
	return r.CreateAuthor(author)
}

// GetAllAuthors returns all authors from the database
func GetAllAuthors(r AuthorRepository) ([]*Author, error) {
	return r.GetAllAuthors()
}

// GetAuthorByID returns an author from the database by id
func GetAuthorByID(r AuthorRepository, id int) (*Author, error) {
	return r.GetAuthorByID(id)
}

// UpdateAuthor updates an author in the database
func UpdateAuthor(r AuthorRepository, author *Author) error {
	return r.UpdateAuthor(author)
}

// DeleteAuthor deletes an author from the database
func DeleteAuthor(r AuthorRepository, id int) error {
	return r.DeleteAuthor(id)
}
