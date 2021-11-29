package categories

import (
	"api_short_story/business/categories"
	"api_short_story/controllers/categories/response"
	"api_short_story/models"
	"errors"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(gormDB *gorm.DB) categories.CategoryRepoInterface {
	return &CategoryRepository{
		db: gormDB,
	}
}

func (repo *CategoryRepository) GetAllCategories() ([]categories.CategoryEntity, error) {
	categoryDB := []response.Category{}
	result := repo.db.Find(&categoryDB)
	if result.Error != nil {
		return []categories.CategoryEntity{}, result.Error
	}
	categoryResponse := response.ToArrayCategoryEntity(categoryDB)
	return categoryResponse, result.Error
}

func (repo *CategoryRepository) AddCategory(category categories.CategoryEntity) (categories.CategoryEntity, error) {
	categoryDB := FromAuthorEntity(category)
	result := repo.db.Create(&categoryDB)
	if result.Error != nil {
		return categories.CategoryEntity{}, result.Error
	}
	return categoryDB.ToCategoryEntity(), nil
}

func (repo *CategoryRepository) EditCategory(id int, category categories.CategoryEntity) (categories.CategoryEntity, error) {
	categoryDB := response.Category{}
	result := repo.db.Find(&categoryDB, id)
	if result.Error != nil {
		return categories.CategoryEntity{}, result.Error
	}
	if result.RowsAffected == 0 {
		return categoryDB.ToCategoryEntity(), errors.New("record not found")
	}
	categoryDB.Name = category.Name
	result.Save(&categoryDB)
	return categoryDB.ToCategoryEntity(), nil
}

func (repo *CategoryRepository) DeleteCategory(id int) (categories.CategoryEntity, error) {
	storyDB := models.ShortStory{CategoryID: uint(id)}
	result := repo.db.Where(&storyDB).Delete(&storyDB)
	if result.Error != nil {
		return categories.CategoryEntity{}, result.Error
	}
	categoryDB := response.Category{Id: uint(id)}
	result = repo.db.Where(&categoryDB).Delete(&categoryDB)
	if result.Error != nil {
		return categories.CategoryEntity{}, result.Error
	}
	if result.RowsAffected == 0 {
		return categoryDB.ToCategoryEntity(), errors.New("record not found")
	}
	return categoryDB.ToCategoryEntity(), nil
}
