package api

import (
	"errors"
	"github.com/pract9/internal/db"
	error2 "github.com/pract9/internal/error"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func UpdateFileById(file *multipart.FileHeader, hexId string) *error2.RequestError {
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

	fileOpen, err := file.Open()
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

	uploadOpts := options.GridFSUpload().SetMetadata(bson.D{{"content-type", file.Header.Get("Content-Type")}})
	err = bucket.UploadFromStreamWithID(objectID, file.Filename, io.Reader(fileOpen), uploadOpts)
	if err != nil {
		return &error2.RequestError{
			StatusCode: http.StatusBadRequest,
			Err:        errors.New("Неврзможно обновить файл"),
		}
	}
	return nil
}
