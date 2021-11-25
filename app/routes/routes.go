package routes

import (
	"api_short_story/controllers/authors"
	"api_short_story/controllers/categories"
	shortstory "api_short_story/controllers/short_story"
	"api_short_story/middlewares/token"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type RouteControllerList struct {
	AuthorController     authors.AuthorController
	CategoryController   categories.CategoryController
	ShortStoryController shortstory.ShortStoryController
}

func (controller RouteControllerList) RouteRegister(c *echo.Echo) {
	authors := c.Group("/author")
	authors.POST("/login", controller.AuthorController.Login)
	authors.POST("", controller.AuthorController.AddAuthor)
	authors.PUT("/:id", controller.AuthorController.EditAuthor, token.MiddlewareAuthor)
	authors.GET("", controller.AuthorController.GetAllAuthors)
	authors.GET("/:id", controller.AuthorController.GetAuthorById)
	authors.DELETE("/:id", controller.AuthorController.DeleteAuthor, token.MiddlewareAdmin)

	categories := c.Group("/category")
	categories.POST("", controller.CategoryController.AddCategory, token.MiddlewareAdmin)
	categories.PUT("/:id", controller.CategoryController.EditCategory, token.MiddlewareAdmin)
	categories.GET("", controller.CategoryController.GetAllCategories)
	categories.GET("/:id", controller.CategoryController.GetCategoryById)
	categories.DELETE("/:id", controller.CategoryController.DeleteCategory, token.MiddlewareAdmin)

	stories := c.Group("/story")
	stories.POST("", controller.ShortStoryController.AddShortStory, token.MiddlewareAuthor)
	stories.PUT("/:id", controller.ShortStoryController.EditShortStory, token.MiddlewareAuthor)
	stories.GET("/:offset/:limit", controller.ShortStoryController.GetShortStories)
	stories.GET("/:id", controller.ShortStoryController.GetShortStoryById)
	stories.DELETE("/:id", controller.ShortStoryController.DeleteShortStory, token.MiddlewareAuthor)

	c.Pre(middleware.RemoveTrailingSlash())
}
