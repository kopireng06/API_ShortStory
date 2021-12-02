package shortstory

import (
	shortstory "api_short_story/business/short_story"
	"api_short_story/controllers"
	"api_short_story/controllers/short_story/request"
	"api_short_story/controllers/short_story/response"
	"api_short_story/middlewares/token"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ShortStoryController struct {
	usecase shortstory.ShortStoryUseCaseInterface
}

func NewController(shortStory shortstory.ShortStoryUseCaseInterface) *ShortStoryController {
	return &ShortStoryController{
		usecase: shortStory,
	}
}

func (controller *ShortStoryController) GetShortStories(c echo.Context) error {
	offset, _ := strconv.Atoi(c.Param("offset"))
	limit, _ := strconv.Atoi(c.Param("limit"))
	shortStories, err := controller.usecase.GetShortStories(offset, limit)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	return controllers.SuccessResponse(c, response.FromArrayShortStoryEntity(shortStories))
}

func (controller *ShortStoryController) GetShortStoryById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	shortStory, err := controller.usecase.GetShortStoryById(id)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusNotFound, err.Error())
	}
	return controllers.SuccessResponse(c, response.FromShortStoryEntity(shortStory))
}

func (controller *ShortStoryController) GetShortStoriesByTitle(c echo.Context) error {
	title := c.Param("title")
	shortStories, err := controller.usecase.GetShortStoriesByTitle(title)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	return controllers.SuccessResponse(c, response.FromArrayShortStoryEntity(shortStories))
}

func (controller *ShortStoryController) GetShortStoriesByIdCategory(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	shortStories, err := controller.usecase.GetShortStoriesByIdCategory(id)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	return controllers.SuccessResponse(c, response.FromArrayShortStoryEntity(shortStories))
}

func (controller *ShortStoryController) GetShortStoriesByIdAuthor(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	shortStories, err := controller.usecase.GetShortStoriesByIdAuthor(id)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	return controllers.SuccessResponse(c, response.FromArrayShortStoryEntity(shortStories))
}

func (controller *ShortStoryController) AddShortStory(c echo.Context) error {
	idAuthor := token.GetAuthorIdFromJWT(c)
	var storyAdd request.ShortStory
	err := c.Bind(&storyAdd)
	storyAdd.AuthorID = uint(idAuthor)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "error binding")
	}
	story, err := controller.usecase.AddShortStory(storyAdd.ToShortStoryEntity())
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	return controllers.SuccessResponse(c, response.FromShortStoryEntity(story))
}

func (controller *ShortStoryController) EditShortStory(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var storyEdit request.ShortStory
	idAuthor := token.GetAuthorIdFromJWT(c)
	storyEdit.AuthorID = uint(idAuthor)
	err := c.Bind(&storyEdit)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "error binding")
	}
	story, err := controller.usecase.EditShortStory(id, storyEdit.ToShortStoryEntity())
	if err != nil {
		if err.Error() == "unauthorized" {
			return controllers.ErrorResponse(c, http.StatusUnauthorized, err.Error())
		}
		return controllers.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	return controllers.SuccessResponse(c, response.FromShortStoryEntity(story))
}

func (controller *ShortStoryController) DeleteShortStory(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var storyDelete request.DeleteShortStory
	idAuthor := token.GetAuthorIdFromJWT(c)
	storyDelete.AuthorID = idAuthor
	story, err := controller.usecase.DeleteShortStory(id, storyDelete.ToShortStoryEntity())
	if err != nil {
		if err.Error() == "unauthorized" {
			return controllers.ErrorResponse(c, http.StatusUnauthorized, err.Error())
		}
		return controllers.ErrorResponse(c, http.StatusNotFound, err.Error())
	}
	deleteStory := response.DeleteShortStory{Id: story.Id}
	return controllers.SuccessResponse(c, deleteStory)
}
