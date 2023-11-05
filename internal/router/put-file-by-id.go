package router

import (
	"github.com/labstack/echo/v4"
	"github.com/pract9/internal/db/api"
	"net/http"
)

func PutFileById(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.String(http.StatusBadRequest, "Файл не указан")
	}

	id := c.Param("id")

	err1 := api.UpdateFileById(file, id)
	if err1 != nil {
		return c.String(err1.StatusCode, err1.Err.Error())
	}

	return c.String(http.StatusOK, "Файл изменен")
}
