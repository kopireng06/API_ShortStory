package shortstory

import (
	shortstory "api_short_story/business/short_story"
	"api_short_story/controllers"
	"api_short_story/controllers/short_story/response"
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
		return controllers.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	return controllers.SuccessResponse(c, response.FromShortStoryEntity(shortStory))
}

func (controller *ShortStoryController) AddShortStory(c echo.Context) error {
	var storyAdd response.ShortStory
	err := c.Bind(&storyAdd)
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
	var storyEdit response.ShortStory
	err := c.Bind(&storyEdit)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "error binding")
	}
	story, err := controller.usecase.EditShortStory(id, storyEdit.ToShortStoryEntity())
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	return controllers.SuccessResponse(c, response.FromShortStoryEntity(story))
}

func (controller *ShortStoryController) DeleteShortStory(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	story, err := controller.usecase.DeleteShortStory(id)
	if err != nil {
		// c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		// c.Response().WriteHeader(http.StatusBadRequest)
		return controllers.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	deleteStory := response.DeleteShortStory{Id: story.Id}
	return controllers.SuccessResponse(c, deleteStory)
}
