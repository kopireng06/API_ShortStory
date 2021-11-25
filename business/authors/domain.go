package authors

import (
	shortstory "api_short_story/business/short_story"
	"time"
)

//ditambahkan token
//struct milik usecase dan global
type AuthorEntity struct {
	Id              uint
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Name            string
	Email           string
	Password        string
	ConfirmPassword string
	Profile         string
	ShortStory      []shortstory.ShortStoryEntity
	Token           string
}

//milik usecase
type AuthorUseCaseInterface interface {
	Login(author AuthorEntity) (AuthorEntity, error)
	GetAllAuthors() ([]AuthorEntity, error)
	GetAuthorById(id int) (AuthorEntity, error)
	AddAuthor(author AuthorEntity) (AuthorEntity, error)
	EditAuthor(id int, author AuthorEntity) (AuthorEntity, error)
	DeleteAuthor(id int) (AuthorEntity, error)
}

//milik repo
type AuthorRepoInterface interface {
	Login(author AuthorEntity) (AuthorEntity, error)
	GetAllAuthors() ([]AuthorEntity, error)
	GetAuthorById(id int) (AuthorEntity, error)
	AddAuthor(author AuthorEntity) (AuthorEntity, error)
	EditAuthor(id int, author AuthorEntity) (AuthorEntity, error)
	DeleteAuthor(id int) (AuthorEntity, error)
}
