package categories

import (
	shortstory "api_short_story/business/short_story"
	"time"
)

type CategoryEntity struct {
	Id              uint
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Name            string
	NumberOfStories int
	ShortStory      []shortstory.ShortStoryEntity
}

//milik usecase
type CategoryUseCaseInterface interface {
	GetAllCategories() ([]CategoryEntity, error)
	AddCategory(category CategoryEntity) (CategoryEntity, error)
	EditCategory(id int, author CategoryEntity) (CategoryEntity, error)
	DeleteCategory(id int) (CategoryEntity, error)
}

//milik repo
type CategoryRepoInterface interface {
	GetAllCategories() ([]CategoryEntity, error)
	AddCategory(category CategoryEntity) (CategoryEntity, error)
	EditCategory(id int, author CategoryEntity) (CategoryEntity, error)
	DeleteCategory(id int) (CategoryEntity, error)
}
