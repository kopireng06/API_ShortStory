package controllers

import (
	"api_short_story/configs"
	"api_short_story/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAuthor(c echo.Context) error {
	var author []models.Author
	result := configs.DB.Find(&author)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			http.StatusInternalServerError,
			result.Error.Error(),
			nil,
		})
	}

	var response = models.BaseResponse{
		http.StatusOK,
		"success",
		author,
	}
	return c.JSON(http.StatusOK, response)
}

func GetAuthorById(c echo.Context) error {
	id := c.Param("id")
	var author models.Author
	result := configs.DB.Find(&author, id)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			http.StatusInternalServerError,
			result.Error.Error(),
			nil,
		})
	}

	var response = models.BaseResponse{
		http.StatusOK,
		"success",
		author,
	}
	return c.JSON(http.StatusOK, response)
}

func GetAuthorByName(c echo.Context) error {
	name := c.Param("name")
	var author []models.Author
	result := configs.DB.Where("name LIKE ?", "%"+name+"%").Find(&author)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			http.StatusInternalServerError,
			result.Error.Error(),
			nil,
		})
	}

	var response = models.BaseResponse{
		http.StatusOK,
		"success",
		author,
	}
	return c.JSON(http.StatusOK, response)
}

func InsertAuthor(c echo.Context) error {
	var author models.Author
	c.Bind(&author)
	result := configs.DB.Create(&author)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			http.StatusInternalServerError,
			result.Error.Error(),
			nil,
		})
	}

	var response = models.BaseResponse{
		http.StatusOK,
		"success",
		author,
	}
	return c.JSON(http.StatusOK, response)
}
