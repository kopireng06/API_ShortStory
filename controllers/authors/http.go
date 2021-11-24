package authors

import (
	"api_short_story/business/authors"
	"api_short_story/controllers"
	"api_short_story/controllers/authors/request"
	"api_short_story/controllers/authors/response"
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
	return controllers.SuccessResponse(c, response.FromArrayAuthorEntityToAuthorOnly(authors))
}

func (controller *AuthorController) GetAuthorById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	author, err := controller.usecase.GetAuthorById(id)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	return controllers.SuccessResponse(c, response.FromAuthorEntity(author))
}

func (controller *AuthorController) AddAuthor(c echo.Context) error {
	var authorAdd response.Author
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
	id, _ := strconv.Atoi(c.Param("id"))
	var authorEdit response.Author
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
		return controllers.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	deleteAuthor := response.DeleteAuthor{Id: author.Id}
	return controllers.SuccessResponse(c, deleteAuthor)
}
