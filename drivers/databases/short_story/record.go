package shortstory

import (
	shortstory "api_short_story/business/short_story"
	"time"
)

type ShortStory struct {
	Id         uint
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Title      string
	Content    string
	AuthorID   uint
	CategoryID uint
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

func FromAuthorEntity(shortStory shortstory.ShortStoryEntity) ShortStory {
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
