package router

import (
	"github.com/labstack/echo/v4"
	"github.com/pract9/internal/db/api"
	"log"
	"net/http"
)

func PostFile(c echo.Context) error {

	file, err := c.FormFile("file")
	if err != nil {
		log.Fatal(err)
	}

	objectID, err := api.UploadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	return c.String(http.StatusOK, "Создан файл с ID: "+objectID.Hex())
}
