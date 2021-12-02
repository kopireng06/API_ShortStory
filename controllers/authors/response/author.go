package response

import (
	"api_short_story/business/authors"
	"time"
)

//untuk controller

type Author struct {
	Id        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Profile   string    `json:"profile"`
}

func FromAuthorEntity(author authors.AuthorEntity) Author {
	return Author{
		Id:        author.Id,
		CreatedAt: author.CreatedAt,
		UpdatedAt: author.UpdatedAt,
		Name:      author.Name,
		Email:     author.Email,
		Profile:   author.Profile,
	}
}

func (author Author) ToAuthorEntity() authors.AuthorEntity {
	return authors.AuthorEntity{
		Id:        author.Id,
		CreatedAt: author.CreatedAt,
		UpdatedAt: author.UpdatedAt,
		Name:      author.Name,
		Email:     author.Email,
		Profile:   author.Profile,
	}
}

func FromArrayAuthorEntity(author []authors.AuthorEntity) []Author {
	authorResponse := []Author{}
	for i, _ := range author {
		authorResponse = append(authorResponse, FromAuthorEntity(author[i]))
	}
	return authorResponse
}

func ToArrayAuthorEntity(author []Author) []authors.AuthorEntity {
	authorResponse := []authors.AuthorEntity{}
	for i, _ := range author {
		authorResponse = append(authorResponse, author[i].ToAuthorEntity())
	}
	return authorResponse
}
