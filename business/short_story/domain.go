package shortstory

import "time"

type ShortStoryEntity struct {
	Id         uint
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Title      string
	Content    string
	AuthorID   uint
	CategoryID uint
}

//milik usecase
type ShortStoryUseCaseInterface interface {
	GetShortStories(offset int, limit int) ([]ShortStoryEntity, error)
	GetShortStoryById(id int) (ShortStoryEntity, error)
	AddShortStory(story ShortStoryEntity) (ShortStoryEntity, error)
	EditShortStory(id int, story ShortStoryEntity) (ShortStoryEntity, error)
	DeleteShortStory(id int) (ShortStoryEntity, error)
}

//milik repo
type ShortStoryRepoInterface interface {
	GetShortStories(offset int, limit int) ([]ShortStoryEntity, error)
	GetShortStoryById(id int) (ShortStoryEntity, error)
	AddShortStory(story ShortStoryEntity) (ShortStoryEntity, error)
	EditShortStory(id int, story ShortStoryEntity) (ShortStoryEntity, error)
	DeleteShortStory(id int) (ShortStoryEntity, error)
}
