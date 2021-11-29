package categories

import (
	"errors"
)

type CategoryUseCase struct {
	repo CategoryRepoInterface
}

func NewUseCase(categoryRepo CategoryRepoInterface) CategoryUseCaseInterface {
	return &CategoryUseCase{
		repo: categoryRepo,
	}
}

func (usecase *CategoryUseCase) GetAllCategories() ([]CategoryEntity, error) {
	categories, err := usecase.repo.GetAllCategories()
	if err != nil {
		return []CategoryEntity{}, err
	}
	return categories, nil
}

func (usecase *CategoryUseCase) AddCategory(category CategoryEntity) (CategoryEntity, error) {
	if category.Name == "" {
		return CategoryEntity{}, errors.New("name empty")
	}
	category, err := usecase.repo.AddCategory(category)
	if err != nil {
		return CategoryEntity{}, err
	}
	return category, nil
}

func (usecase *CategoryUseCase) EditCategory(id int, category CategoryEntity) (CategoryEntity, error) {
	if category.Name == "" {
		return CategoryEntity{}, errors.New("name empty")
	}
	category, err := usecase.repo.EditCategory(id, category)
	if err != nil {
		return CategoryEntity{}, err
	}
	return category, nil
}

func (usecase *CategoryUseCase) DeleteCategory(id int) (CategoryEntity, error) {
	category, err := usecase.repo.DeleteCategory(id)
	if err != nil {
		return CategoryEntity{}, err
	}
	return category, nil
}
