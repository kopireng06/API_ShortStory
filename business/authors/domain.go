package authors

import "time"

type AuthorEntity struct {
	Id        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Email     string
	Password  string
	Profile   string
}

//milik usecase
type AuthorUseCaseInterface interface {
	Login(author AuthorEntity) (AuthorEntity, error)
	GetAllAuthors() ([]AuthorEntity, error)
	AddAuthor(authot AuthorEntity) (AuthorEntity, error)
}

//milik repo
type AuthorRepoInterface interface {
	Login(author AuthorEntity) (AuthorEntity, error)
	GetAllAuthors() ([]AuthorEntity, error)
	AddAuthor(authot AuthorEntity) (AuthorEntity, error)
}
