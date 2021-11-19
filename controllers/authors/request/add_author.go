package request

import (
	"api_short_story/business/authors"
	"time"
)

type AuthorAdd struct {
	Id        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Profile   string    `json:"profile"`
}

func (author *AuthorAdd) ToAuthorEntity() *authors.AuthorEntity {
	return &authors.AuthorEntity{
		Id:        author.Id,
		CreatedAt: author.CreatedAt,
		UpdatedAt: author.UpdatedAt,
		Name:      author.Name,
		Email:     author.Email,
		Password:  author.Password,
		Profile:   author.Profile,
	}
}
