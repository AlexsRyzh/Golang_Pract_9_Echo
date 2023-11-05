package api

import (
	"errors"
	"github.com/pract9/internal/db"
	error2 "github.com/pract9/internal/error"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"net/http"
	"os"
)

func DeleteFileById(hexId string) *error2.RequestError {
	client, ctx, err := db.Connect()
	if err != nil {
		return &error2.RequestError{
			StatusCode: http.StatusInternalServerError,
			Err:        errors.New("Ошибка сервера"),
		}
	}
	defer db.Disconnect(client, ctx)

	filesDb := client.Database(os.Getenv("FILE_DATABASE"))
	bucket, err := gridfs.NewBucket(filesDb)
	if err != nil {
		return &error2.RequestError{
			StatusCode: http.StatusInternalServerError,
			Err:        errors.New("Ошибка сервера"),
		}
	}

	objectID, err := primitive.ObjectIDFromHex(hexId)
	if err != nil {
		return &error2.RequestError{
			StatusCode: http.StatusBadRequest,
			Err:        errors.New("Не верный ID"),
		}
	}

	err = bucket.Delete(objectID)
	if err != nil {
		return &error2.RequestError{
			StatusCode: http.StatusBadRequest,
			Err:        errors.New("Файл не найден"),
		}
	}

	return nil
}
