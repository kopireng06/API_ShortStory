package request

import "api_short_story/business/authors"

type AuthorAdd struct {
	Name            string `json:"name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
	Profile         string `json:"profile"`
}

func (author AuthorAdd) ToAuthorEntity() authors.AuthorEntity {
	return authors.AuthorEntity{
		Name:            author.Name,
		Email:           author.Email,
		Password:        author.Password,
		Profile:         author.Profile,
		ConfirmPassword: author.ConfirmPassword,
	}
}
