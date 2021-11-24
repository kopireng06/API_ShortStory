package response

import (
	"api_short_story/business/authors"
	"time"
)

type AuthorOnly struct {
	Id        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Profile   string    `json:"profile"`
}

func FromAuthorEntityToAuthorOnly(author authors.AuthorEntity) AuthorOnly {
	return AuthorOnly{
		Id:        author.Id,
		CreatedAt: author.CreatedAt,
		UpdatedAt: author.UpdatedAt,
		Name:      author.Name,
		Email:     author.Email,
		Password:  author.Password,
		Profile:   author.Profile,
	}
}

func (author AuthorOnly) ToAuthorEntity() authors.AuthorEntity {
	return authors.AuthorEntity{
		Id:        author.Id,
		CreatedAt: author.CreatedAt,
		UpdatedAt: author.UpdatedAt,
		Name:      author.Name,
		Email:     author.Email,
		Password:  author.Password,
		Profile:   author.Profile,
	}
}

func FromArrayAuthorEntityToAuthorOnly(author []authors.AuthorEntity) []AuthorOnly {
	authorResponse := []AuthorOnly{}
	for i, _ := range author {
		authorResponse = append(authorResponse, FromAuthorEntityToAuthorOnly(author[i]))
	}
	return authorResponse
}

func ToArrayAuthorEntityToAuthorOnly(author []AuthorOnly) []authors.AuthorEntity {
	authorResponse := []authors.AuthorEntity{}
	for i, _ := range author {
		authorResponse = append(authorResponse, author[i].ToAuthorEntity())
	}
	return authorResponse
}
