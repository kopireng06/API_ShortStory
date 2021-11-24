package categories

import (
	"api_short_story/business/categories"
	"time"
)

type Category struct {
	Id        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
}

func (category Category) ToCategoryEntity() categories.CategoryEntity {
	return categories.CategoryEntity{
		Id:        category.Id,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
		Name:      category.Name,
	}
}

func FromAuthorEntity(category categories.CategoryEntity) Category {
	return Category{
		Id:        category.Id,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
		Name:      category.Name,
	}
}
