package authors

import (
	"api_short_story/business/authors"
	"api_short_story/controllers"
	"api_short_story/controllers/authors/request"
	"api_short_story/controllers/authors/response"
	"api_short_story/middlewares/token"
	"net/http"
	"strconv"

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
	var authorLogin request.AuthorLogin
	err := c.Bind(&authorLogin)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "error binding")
	}
	author, err := controller.usecase.Login(*authorLogin.ToAuthorEntity())
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	responeLogin := response.AuthorLogin{
		Token: author.Token,
	}
	return controllers.SuccessResponse(c, responeLogin)
}

func (controller *AuthorController) GetAllAuthors(c echo.Context) error {
	authors, err := controller.usecase.GetAllAuthors()
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	return controllers.SuccessResponse(c, response.FromArrayAuthorEntity(authors))
}

func (controller *AuthorController) GetAuthorById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	author, err := controller.usecase.GetAuthorById(id)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusNotFound, err.Error())
	}
	return controllers.SuccessResponse(c, response.FromAuthorEntity(author))
}

func (controller *AuthorController) GetAuthorsByName(c echo.Context) error {
	name := c.Param("name")
	authors, err := controller.usecase.GetAuthorsByName(name)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusNotFound, err.Error())
	}
	return controllers.SuccessResponse(c, response.FromArrayAuthorEntity(authors))
}

func (controller *AuthorController) AddAuthor(c echo.Context) error {
	var authorAdd request.AuthorAdd
	err := c.Bind(&authorAdd)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "error binding")
	}
	author, err := controller.usecase.AddAuthor(authorAdd.ToAuthorEntity())
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	return controllers.SuccessResponse(c, response.FromAuthorEntity(author))
}

func (controller *AuthorController) EditAuthor(c echo.Context) error {
	id := token.GetAuthorIdFromJWT(c)
	var authorEdit request.AuthorEdit
	err := c.Bind(&authorEdit)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "error binding")
	}
	author, err := controller.usecase.EditAuthor(id, authorEdit.ToAuthorEntity())
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	return controllers.SuccessResponse(c, response.FromAuthorEntity(author))
}

func (controller *AuthorController) DeleteAuthor(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	author, err := controller.usecase.DeleteAuthor(id)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusNotFound, err.Error())
	}
	deleteAuthor := response.DeleteAuthor{Id: author.Id}
	return controllers.SuccessResponse(c, deleteAuthor)
}
