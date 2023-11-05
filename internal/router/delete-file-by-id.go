package router

import (
	"github.com/labstack/echo/v4"
	"github.com/pract9/internal/db/api"
	"net/http"
)

func DeleteFileById(c echo.Context) error {
	id := c.Param("id")

	err := api.DeleteFileById(id)
	if err != nil {
		return c.String(err.StatusCode, err.Err.Error())
	}

	return c.String(http.StatusOK, "Файл успешно удален")
}
