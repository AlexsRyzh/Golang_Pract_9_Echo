package router

import (
	"github.com/labstack/echo/v4"
	"github.com/pract9/internal/db/api"
	"log"
	"net/http"
)

func GetFiles(c echo.Context) error {

	fileListInfo, err := api.GetFiles()
	if err != nil {
		log.Fatal(err)
	}
	return c.JSON(http.StatusOK, fileListInfo)
}
