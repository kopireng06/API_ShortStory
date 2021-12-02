package request

import "api_short_story/business/authors"

type AuthorLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (author *AuthorLogin) ToAuthorEntity() *authors.AuthorEntity {
	return &authors.AuthorEntity{
		Email:    author.Email,
		Password: author.Password,
	}
}
