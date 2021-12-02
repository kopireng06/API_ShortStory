package response

import (
	shortstory "api_short_story/business/short_story"
	"time"
)

type Author struct {
	Id      uint   `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Profile string `json:"profile"`
}

type Category struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

type ShortStory struct {
	Id         uint      `json:"id"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	AuthorID   uint      `json:"authorId"`
	CategoryID uint      `json:"categoryId"`
	Author     Author    `json:"author"`
	Category   Category  `json:"category"`
}

func (shortStory ShortStory) ToShortStoryEntity() shortstory.ShortStoryEntity {
	return shortstory.ShortStoryEntity{
		Id:         shortStory.Id,
		CreatedAt:  shortStory.CreatedAt,
		UpdatedAt:  shortStory.UpdatedAt,
		Title:      shortStory.Title,
		Content:    shortStory.Content,
		AuthorID:   shortStory.AuthorID,
		CategoryID: shortStory.CategoryID,
		Author: shortstory.Author{
			Id:      shortStory.Author.Id,
			Name:    shortStory.Author.Name,
			Email:   shortStory.Author.Email,
			Profile: shortStory.Author.Profile,
		},
		Category: shortstory.Category{
			Id:   shortStory.Category.Id,
			Name: shortStory.Category.Name,
		},
	}
}

func FromShortStoryEntity(shortStory shortstory.ShortStoryEntity) ShortStory {
	return ShortStory{
		Id:         shortStory.Id,
		CreatedAt:  shortStory.CreatedAt,
		UpdatedAt:  shortStory.UpdatedAt,
		Title:      shortStory.Title,
		Content:    shortStory.Content,
		AuthorID:   shortStory.AuthorID,
		CategoryID: shortStory.CategoryID,
		Author: Author{
			Id:      shortStory.Author.Id,
			Name:    shortStory.Author.Name,
			Email:   shortStory.Author.Email,
			Profile: shortStory.Author.Profile,
		},
		Category: Category{
			Id:   shortStory.Category.Id,
			Name: shortStory.Category.Name,
		},
	}
}

func FromArrayShortStoryEntity(shortStory []shortstory.ShortStoryEntity) []ShortStory {
	shortStoryResponse := []ShortStory{}
	for i, _ := range shortStory {
		shortStoryResponse = append(shortStoryResponse, FromShortStoryEntity(shortStory[i]))
	}
	return shortStoryResponse
}

func ToArrayShortStoryEntity(shortStory []ShortStory) []shortstory.ShortStoryEntity {
	shortStoryResponse := []shortstory.ShortStoryEntity{}
	for i, _ := range shortStory {
		shortStoryResponse = append(shortStoryResponse, shortStory[i].ToShortStoryEntity())
	}
	return shortStoryResponse
}
