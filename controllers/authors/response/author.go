package response

import (
	"api_short_story/business/authors"
	"time"
)

type AuthorResponse struct {
	Id        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Profile   string    `json:"profile"`
}

func FromAuthorEntity(author authors.AuthorEntity) AuthorResponse {
	return AuthorResponse{
		Id:        author.Id,
		CreatedAt: author.CreatedAt,
		UpdatedAt: author.UpdatedAt,
		Name:      author.Name,
		Email:     author.Email,
		Password:  author.Password,
		Profile:   author.Profile,
	}
}
