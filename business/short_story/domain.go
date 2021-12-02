package shortstory

import (
	"time"
)

type Author struct {
	Id        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Email     string
	Profile   string
}

type Category struct {
	Id   uint
	Name string
}

type ShortStoryEntity struct {
	Id         uint
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Title      string
	Content    string
	AuthorID   uint
	CategoryID uint
	Author     Author
	Category   Category
}

//milik usecase
type ShortStoryUseCaseInterface interface {
	GetShortStories(offset int, limit int) ([]ShortStoryEntity, error)
	GetShortStoryById(id int) (ShortStoryEntity, error)
	GetShortStoriesByTitle(title string) ([]ShortStoryEntity, error)
	GetShortStoriesByIdCategory(id int) ([]ShortStoryEntity, error)
	GetShortStoriesByIdAuthor(id int) ([]ShortStoryEntity, error)
	AddShortStory(story ShortStoryEntity) (ShortStoryEntity, error)
	EditShortStory(id int, story ShortStoryEntity) (ShortStoryEntity, error)
	DeleteShortStory(id int, story ShortStoryEntity) (ShortStoryEntity, error)
}

//milik repo
type ShortStoryRepoInterface interface {
	GetShortStories(offset int, limit int) ([]ShortStoryEntity, error)
	GetShortStoryById(id int) (ShortStoryEntity, error)
	GetShortStoriesByTitle(title string) ([]ShortStoryEntity, error)
	GetShortStoriesByIdCategory(id int) ([]ShortStoryEntity, error)
	GetShortStoriesByIdAuthor(id int) ([]ShortStoryEntity, error)
	AddShortStory(story ShortStoryEntity) (ShortStoryEntity, error)
	EditShortStory(id int, story ShortStoryEntity) (ShortStoryEntity, error)
	DeleteShortStory(id int, story ShortStoryEntity) (ShortStoryEntity, error)
}
