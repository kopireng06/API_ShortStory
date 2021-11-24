package response

import (
	shortstory "api_short_story/business/short_story"
	"time"
)

type ShortStory struct {
	Id         uint      `json:"id"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	AuthorID   uint      `json:"authorId"`
	CategoryID uint      `json:"categoryId"`
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
