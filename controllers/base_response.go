package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type BaseReponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SuccessResponse(c echo.Context, data interface{}) error {
	response := BaseReponse{}
	response.Status = http.StatusOK
	response.Message = "success"
	response.Data = data
	return c.JSON(http.StatusOK, response)
}

func ErrorResponse(c echo.Context, status int, err string) error {
	response := BaseReponse{}
	response.Status = status
	response.Message = err
	return c.JSON(http.StatusOK, response)
}
