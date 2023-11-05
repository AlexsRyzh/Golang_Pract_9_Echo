package router

import (
	"github.com/labstack/echo/v4"
	"github.com/pract9/internal/db/api"
	"net/http"
)

func GetFileInfoById(c echo.Context) error {

	fileId := c.Param("id")

	fileInfo, err := api.GetFileInfoById(fileId)
	if err != nil {
		return c.String(err.StatusCode, err.Err.Error())
	}

	if fileInfo == nil && err == nil {
		return c.String(http.StatusBadRequest, "Файл не найден")
	}

	return c.JSON(http.StatusOK, fileInfo)
}
