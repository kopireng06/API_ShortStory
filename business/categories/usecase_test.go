package categories_test

import (
	"api_short_story/business/categories"
	"api_short_story/business/categories/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepoInterfaceMock mocks.CategoryRepoInterface
var categoryUseCaseInterface categories.CategoryUseCaseInterface
var categoryArrayResponseFromRepo []categories.CategoryEntity
var categoryResponseFromRepo categories.CategoryEntity

func setupGetAllCategories() {
	categoryUseCaseInterface = categories.NewUseCase(&categoryRepoInterfaceMock)
	categoryArrayResponseFromRepo = []categories.CategoryEntity{
		{
			Id:   1,
			Name: "percintaan",
		},
	}
}

func setupAddCategory() {
	categoryUseCaseInterface = categories.NewUseCase(&categoryRepoInterfaceMock)
	categoryResponseFromRepo = categories.CategoryEntity{
		Id:   1,
		Name: "percintaan",
	}
}

func setupEditCategory() {
	categoryUseCaseInterface = categories.NewUseCase(&categoryRepoInterfaceMock)
	categoryResponseFromRepo = categories.CategoryEntity{
		Id:   1,
		Name: "percintaan",
	}
}

func setupDeleteCategory() {
	categoryUseCaseInterface = categories.NewUseCase(&categoryRepoInterfaceMock)
	categoryResponseFromRepo = categories.CategoryEntity{
		Id: 1,
	}
}

func TestGetAllCategories(t *testing.T) {
	setupGetAllCategories()
	t.Run("Success Get All Categories", func(t *testing.T) {
		categoryRepoInterfaceMock.On("GetAllCategories", mock.Anything).Return(categoryArrayResponseFromRepo, nil).Once()
		entity, err := categoryUseCaseInterface.GetAllCategories()
		assert.Equal(t, nil, err)
		assert.Equal(t, categoryArrayResponseFromRepo, entity)
	})

	t.Run("Fail Get All Categories", func(t *testing.T) {
		categoryRepoInterfaceMock.On("GetAllCategories", mock.Anything).Return([]categories.CategoryEntity{}, errors.New("error bro")).Once()
		entity, err := categoryUseCaseInterface.GetAllCategories()
		assert.Equal(t, errors.New("error bro"), err)
		assert.Equal(t, []categories.CategoryEntity{}, entity)
	})
}

func TestAddCategory(t *testing.T) {
	setupAddCategory()
	t.Run("Success Add Category", func(t *testing.T) {
		categoryRepoInterfaceMock.On("AddCategory", mock.Anything).Return(categoryResponseFromRepo, nil).Once()
		var requestAddCategory = categories.CategoryEntity{
			Name: "percintaan",
		}
		entity, err := categoryUseCaseInterface.AddCategory(requestAddCategory)
		assert.Equal(t, nil, err)
		assert.Equal(t, categoryResponseFromRepo, entity)
	})

	t.Run("Fail Add Category", func(t *testing.T) {
		categoryRepoInterfaceMock.On("AddCategory", mock.Anything).Return(categories.CategoryEntity{}, errors.New("error bro")).Once()
		var requestAddCategory = categories.CategoryEntity{
			Name: "percintaan",
		}
		entity, err := categoryUseCaseInterface.AddCategory(requestAddCategory)
		assert.Equal(t, "error bro", err.Error())
		assert.Equal(t, categories.CategoryEntity{}, entity)
	})

	t.Run("Add Category with Name Empty", func(t *testing.T) {
		var requestAddCategory = categories.CategoryEntity{
			Name: "",
		}
		entity, err := categoryUseCaseInterface.AddCategory(requestAddCategory)
		assert.Equal(t, "name empty", err.Error())
		assert.Equal(t, categories.CategoryEntity{}, entity)
	})
}

func TestEditCategory(t *testing.T) {
	setupEditCategory()
	t.Run("Success Edit Category", func(t *testing.T) {
		var requestAddCategory = categories.CategoryEntity{
			Id:   1,
			Name: "percintaan",
		}
		categoryRepoInterfaceMock.On("EditCategory", int(requestAddCategory.Id), mock.AnythingOfType("categories.CategoryEntity"), mock.Anything).Return(categoryResponseFromRepo, nil).Once()
		entity, err := categoryUseCaseInterface.EditCategory(int(requestAddCategory.Id), requestAddCategory)
		assert.Equal(t, nil, err)
		assert.Equal(t, categoryResponseFromRepo, entity)
	})

	t.Run("Fail Edit Category", func(t *testing.T) {
		var requestAddCategory = categories.CategoryEntity{
			Id:   1,
			Name: "percintaan",
		}
		categoryRepoInterfaceMock.On("EditCategory", int(requestAddCategory.Id), mock.AnythingOfType("categories.CategoryEntity"), mock.Anything).Return(categories.CategoryEntity{}, errors.New("error bro")).Once()
		entity, err := categoryUseCaseInterface.EditCategory(int(requestAddCategory.Id), requestAddCategory)
		assert.Equal(t, "error bro", err.Error())
		assert.Equal(t, categories.CategoryEntity{}, entity)
	})

	t.Run("Edit Category with Name Empty", func(t *testing.T) {
		var requestAddCategory = categories.CategoryEntity{
			Id:   1,
			Name: "",
		}
		entity, err := categoryUseCaseInterface.EditCategory(int(requestAddCategory.Id), requestAddCategory)
		assert.Equal(t, "name empty", err.Error())
		assert.Equal(t, categories.CategoryEntity{}, entity)
	})
}

func TestDeleteCategory(t *testing.T) {
	setupDeleteCategory()
	id := 1
	t.Run("Success Delete Category", func(t *testing.T) {
		categoryRepoInterfaceMock.On("DeleteCategory", id, mock.Anything).Return(categoryResponseFromRepo, nil).Once()
		entity, err := categoryUseCaseInterface.DeleteCategory(id)
		assert.Equal(t, nil, err)
		assert.Equal(t, categoryResponseFromRepo, entity)
	})

	t.Run("Fail Delete Category", func(t *testing.T) {
		categoryRepoInterfaceMock.On("DeleteCategory", id, mock.Anything).Return(categories.CategoryEntity{}, errors.New("error bro")).Once()
		entity, err := categoryUseCaseInterface.DeleteCategory(id)
		assert.Equal(t, "error bro", err.Error())
		assert.Equal(t, categories.CategoryEntity{}, entity)
	})
}
