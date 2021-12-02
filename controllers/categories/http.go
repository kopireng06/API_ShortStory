package categories

import (
	"api_short_story/business/categories"
	"api_short_story/controllers"
	"api_short_story/controllers/categories/response"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CategoryController struct {
	usecase categories.CategoryUseCaseInterface
}

func NewAuthorController(category categories.CategoryUseCaseInterface) *CategoryController {
	return &CategoryController{
		usecase: category,
	}
}

func (controller *CategoryController) GetAllCategories(c echo.Context) error {
	categories, err := controller.usecase.GetAllCategories()
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	return controllers.SuccessResponse(c, response.FromArrayCategoryEntity(categories))
}

func (controller *CategoryController) AddCategory(c echo.Context) error {
	var categoryAdd response.Category
	err := c.Bind(&categoryAdd)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "error binding")
	}
	category, err := controller.usecase.AddCategory(categoryAdd.ToCategoryEntity())
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	return controllers.SuccessResponse(c, response.FromCategoryEntity(category))
}

func (controller *CategoryController) EditCategory(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var categoryEdit response.Category
	err := c.Bind(&categoryEdit)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "error binding")
	}
	category, err := controller.usecase.EditCategory(id, categoryEdit.ToCategoryEntity())
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	return controllers.SuccessResponse(c, response.FromCategoryEntity(category))
}

func (controller *CategoryController) DeleteCategory(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	category, err := controller.usecase.DeleteCategory(id)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusNotFound, err.Error())
	}
	deleteCategory := response.DeleteCategory{Id: category.Id}
	return controllers.SuccessResponse(c, deleteCategory)
}
