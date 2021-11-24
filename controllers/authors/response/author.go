package response

import (
	"api_short_story/business/authors"
	"api_short_story/controllers/short_story/response"
	"time"
)

//untuk controller

type Author struct {
	Id         uint                  `json:"id"`
	CreatedAt  time.Time             `json:"createdAt"`
	UpdatedAt  time.Time             `json:"updatedAt"`
	Name       string                `json:"name"`
	Email      string                `json:"email"`
	Password   string                `json:"password"`
	Profile    string                `json:"profile"`
	ShortStory []response.ShortStory `json:"shortStory"`
}

func FromAuthorEntity(author authors.AuthorEntity) Author {
	return Author{
		Id:         author.Id,
		CreatedAt:  author.CreatedAt,
		UpdatedAt:  author.UpdatedAt,
		Name:       author.Name,
		Email:      author.Email,
		Password:   author.Password,
		Profile:    author.Profile,
		ShortStory: response.FromArrayShortStoryEntity(author.ShortStory),
	}
}

func (author Author) ToAuthorEntity() authors.AuthorEntity {
	return authors.AuthorEntity{
		Id:         author.Id,
		CreatedAt:  author.CreatedAt,
		UpdatedAt:  author.UpdatedAt,
		Name:       author.Name,
		Email:      author.Email,
		Password:   author.Password,
		Profile:    author.Profile,
		ShortStory: response.ToArrayShortStoryEntity(author.ShortStory),
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
