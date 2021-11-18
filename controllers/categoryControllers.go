package controllers

import (
	"api_short_story/configs"
	"api_short_story/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetCategory(c echo.Context) error {
	var category []models.Category
	result := configs.DB.Find(&category)

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
		category,
	}
	return c.JSON(http.StatusOK, response)
}

func InsertCategory(c echo.Context) error {
	var category models.Category
	c.Bind(&category)
	result := configs.DB.Create(&category)
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
		category,
	}
	return c.JSON(http.StatusOK, response)
}
