package main

import (
	"api_short_story/app/routes"
	authorUsecase "api_short_story/business/authors"
	categoryUsecase "api_short_story/business/categories"
	shortStoryUseCase "api_short_story/business/short_story"
	authorController "api_short_story/controllers/authors"
	categoryController "api_short_story/controllers/categories"
	shortStoryController "api_short_story/controllers/short_story"
	authorRepo "api_short_story/drivers/databases/authors"
	categoryRepo "api_short_story/drivers/databases/categories"
	shortStoryRepo "api_short_story/drivers/databases/short_story"
	"api_short_story/drivers/mysql"

	"github.com/labstack/echo/v4"
)

func main() {
	db := mysql.InitialDb()
	e := echo.New()

	authorRepoInterface := authorRepo.NewAuthorRepository(db)
	authorUseCaseInterface := authorUsecase.NewUseCase(authorRepoInterface)
	authorControllerInterface := authorController.NewAuthorController(authorUseCaseInterface)

	categoryRepoInterface := categoryRepo.NewCategoryRepository(db)
	categoryUseCaseInterface := categoryUsecase.NewUseCase(categoryRepoInterface)
	categoryControllerInterface := categoryController.NewAuthorController(categoryUseCaseInterface)

	shortStoryInterface := shortStoryRepo.NewShortStoryRepository(db)
	shortStoryUseCaseInterface := shortStoryUseCase.NewUseCase(shortStoryInterface)
	shortStoryControllerInterface := shortStoryController.NewController(shortStoryUseCaseInterface)

	routesInit := routes.RouteControllerList{
		AuthorController:     *authorControllerInterface,
		CategoryController:   *categoryControllerInterface,
		ShortStoryController: *shortStoryControllerInterface,
	}

	routesInit.RouteRegister(e)
	e.Start(":8000")
}
