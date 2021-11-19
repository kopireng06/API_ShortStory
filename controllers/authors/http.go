package controllers

import (
	"api_short_story/business/authors"
	"api_short_story/controllers"
	"api_short_story/controllers/authors/request"
	"api_short_story/controllers/authors/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthorController struct {
	usecase authors.AuthorUseCaseInterface
}

func NewAuthorController(author authors.AuthorUseCaseInterface) *AuthorController {
	return &AuthorController{
		usecase: author,
	}
}

func (controller *AuthorController) Login(c echo.Context) error {
	//ctx := c.Request().Context()
	var authorLogin request.AuthorLogin
	err := c.Bind(&authorLogin)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "error binding")
	}
	author, _ := controller.usecase.Login(*authorLogin.ToAuthorEntity())
	return controllers.SuccessResponse(c, response.FromAuthorEntity(author))
}

func (controller *AuthorController) GetAllUsers(c echo.Context) error {
	return controllers.SuccessResponse(c, response.AuthorResponse{})
}

func (controller *AuthorController) AddAuthor(c echo.Context) error {
	var authorAdd request.AuthorAdd
	err := c.Bind(&authorAdd)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "error binding")
	}
	author, err := controller.usecase.AddAuthor(*authorAdd.ToAuthorEntity())
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	return controllers.SuccessResponse(c, response.FromAuthorEntity(author))
}
