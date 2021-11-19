package routes

import (
	controllers "api_short_story/controllers/authors"

	"github.com/labstack/echo/v4"
)

type RouteControllerList struct {
	AuthorController controllers.AuthorController
}

func (controller RouteControllerList) RouteRegister(c *echo.Echo) {
	users := c.Group("/author")
	users.POST("/login", controller.AuthorController.Login)
	users.POST("/", controller.AuthorController.AddAuthor)
}
