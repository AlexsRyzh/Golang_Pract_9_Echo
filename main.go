package main

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/pract9/internal/router"
	"os"
)

func main() {
	app := echo.New()

	app.GET("/files", router.GetFiles)                 // Получение списка файлов
	app.GET("/files/:id", router.GetFileById)          // Получение файла по id
	app.GET("/files/:id/info", router.GetFileInfoById) // Получение информации о файле по id
	app.POST("/files", router.PostFile)                // Загрузка файла
	app.PUT("files/:id", router.PutFileById)           // Обновление файла по id
	app.DELETE("/files/:id", router.DeleteFileById)    // Удаление файла по id

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Fatal(app.Start(":" + port))
}
