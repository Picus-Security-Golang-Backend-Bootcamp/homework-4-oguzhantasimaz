package author

import (
	authors "github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-oguzhantasimaz/pkg/models/authors"
	"gorm.io/gorm"
)

type authorRepository struct {
	db *gorm.DB
}

// NewAuthorRepository creates a new author repository
func NewAuthorRepository(db *gorm.DB) authors.AuthorRepository {
	return &authorRepository{
		db: db,
	}
}

// Migrates the database to the latest schema
func (a *authorRepository) Migration() {
	a.db.Migrator().DropTable(&authors.Author{})
	a.db.AutoMigrate(&authors.Author{})
}

// GetAllAuthors returns all authors
func (a *authorRepository) GetAllAuthors() ([]*authors.Author, error) {
	authors := make([]*authors.Author, 0)
	err := a.db.Find(&authors).Error
	return authors, err
}

// GetAuthorByID returns a author by id
func (a *authorRepository) GetAuthorByID(id int) (*authors.Author, error) {
	author := new(authors.Author)
	err := a.db.Where("id = ?", id).First(author).Error
	return author, err
}

// CreateAuthor creates a new author
func (a *authorRepository) CreateAuthor(author *authors.Author) (*authors.Author, error) {
	err := a.db.Create(author).Error
	return author, err
}

// UpdateAuthor updates a author
func (a *authorRepository) UpdateAuthor(author *authors.Author) (*authors.Author, error) {
	err := a.db.Save(author).Error
	return author, err
}

// DeleteAuthor deletes a author
func (a *authorRepository) DeleteAuthor(id int) (err error) {
	err = a.db.Where("id = ?", id).Delete(&authors.Author{}).Error
	return
}

// GetAuthorByName returns a author by name
func (a *authorRepository) GetAuthorByName(name string) (*authors.Author, error) {
	author := new(authors.Author)
	err := a.db.Where("name = ?", name).First(author).Error
	return author, err
}

// InsertSampleData inserts sample data of authors
func (a *authorRepository) InsertSampleData() {
	authorList := []*authors.Author{
		{Name: "Erkut", Surname: "Tackin"},
		{Name: "Cem", Surname: "Karaca"},
		{Name: "Erkin", Surname: "Koray"},
		{Name: "Baris", Surname: "Manco"},
		{Name: "Baris", Surname: "Akarsu"},
	}
	for _, auth := range authorList {
		a.db.Create(&authors.Author{Name: auth.Name, Surname: auth.Surname})
	}
}
