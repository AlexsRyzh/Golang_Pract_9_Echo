package router

import (
	"github.com/labstack/echo/v4"
	"github.com/pract9/internal/db/api"
	"net/http"
)

func GetFileById(c echo.Context) error {

	fileId := c.Param("id")

	file, contentType, err := api.GetFileById(fileId)
	if err != nil {
		return c.String(err.StatusCode, err.Err.Error())
	}

	return c.Blob(http.StatusOK, *contentType, file.Bytes())
}
