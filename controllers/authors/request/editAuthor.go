package request

import "api_short_story/business/authors"

type AuthorEdit struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Profile  string `json:"profile"`
}

func (author AuthorEdit) ToAuthorEntity() authors.AuthorEntity {
	return authors.AuthorEntity{
		Name:     author.Name,
		Email:    author.Email,
		Password: author.Password,
		Profile:  author.Profile,
	}
}
