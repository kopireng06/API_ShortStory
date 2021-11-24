package response

import (
	"api_short_story/business/categories"
	"api_short_story/controllers/short_story/response"
	"time"
)

type Category struct {
	Id         uint                  `json:"id"`
	CreatedAt  time.Time             `json:"createdAt"`
	UpdatedAt  time.Time             `json:"updatedAt"`
	Name       string                `json:"name"`
	ShortStory []response.ShortStory `json:"shortStory"`
}

func FromCategoryEntity(category categories.CategoryEntity) Category {
	return Category{
		Id:         category.Id,
		CreatedAt:  category.CreatedAt,
		UpdatedAt:  category.CreatedAt,
		Name:       category.Name,
		ShortStory: response.FromArrayShortStoryEntity(category.ShortStory),
	}
}

func (category Category) ToCategoryEntity() categories.CategoryEntity {
	return categories.CategoryEntity{
		Id:         category.Id,
		CreatedAt:  category.CreatedAt,
		UpdatedAt:  category.CreatedAt,
		Name:       category.Name,
		ShortStory: response.ToArrayShortStoryEntity(category.ShortStory),
	}
}

func ToArrayCategoryEntity(category []Category) []categories.CategoryEntity {
	categoryResponse := []categories.CategoryEntity{}
	for i, _ := range category {
		categoryResponse = append(categoryResponse, category[i].ToCategoryEntity())
	}
	return categoryResponse
}

func FromArrayCategoryEntity(category []categories.CategoryEntity) []Category {
	categoryResponse := []Category{}
	for i, _ := range category {
		categoryResponse = append(categoryResponse, FromCategoryEntity(category[i]))
	}
	return categoryResponse
}
