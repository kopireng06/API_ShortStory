package authors

import (
	"api_short_story/business/authors"
	"time"
)

// untuk database (repo)
//ditambahkan gorm
type Author struct {
	Id        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Email     string
	Password  string
	Profile   string
	Role      int
}

func (author Author) ToAuthorEntity() authors.AuthorEntity {
	return authors.AuthorEntity{
		Id:        author.Id,
		CreatedAt: author.CreatedAt,
		UpdatedAt: author.UpdatedAt,
		Name:      author.Name,
		Email:     author.Email,
		Password:  author.Password,
		Profile:   author.Profile,
		Role:      author.Role,
	}
}

func FromAuthorEntity(author authors.AuthorEntity) Author {
	return Author{
		Id:        author.Id,
		CreatedAt: author.CreatedAt,
		UpdatedAt: author.UpdatedAt,
		Name:      author.Name,
		Email:     author.Email,
		Password:  author.Password,
		Profile:   author.Profile,
		Role:      author.Role,
	}
}
