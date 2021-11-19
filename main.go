package main

import (
	"api_short_story/app/routes"
	authorUsecase "api_short_story/business/authors"
	authorController "api_short_story/controllers/authors"
	authorRepo "api_short_story/drivers/databases/authors"
	"api_short_story/drivers/mysql"

	"github.com/labstack/echo/v4"
)

func main() {
	db := mysql.InitialDb()
	e := echo.New()
	authorRepoInterface := authorRepo.NewAuthorRepository(db)
	authorUseCaseInterface := authorUsecase.NewUseCase(authorRepoInterface)
	authorControllerInterface := authorController.NewAuthorController(authorUseCaseInterface)

	routesInit := routes.RouteControllerList{
		AuthorController: *authorControllerInterface,
	}

	routesInit.RouteRegister(e)
	e.Start(":8000")
}
