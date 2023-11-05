package api

import (
	"bytes"
	"context"
	"errors"
	"github.com/pract9/internal/db"
	error2 "github.com/pract9/internal/error"
	"github.com/pract9/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"net/http"
	"os"
)

func GetFileById(hexId string) (*bytes.Buffer, *string, *error2.RequestError) {
	client, ctx, err := db.Connect()
	if err != nil {
		return nil, nil, &error2.RequestError{
			StatusCode: http.StatusInternalServerError,
			Err:        errors.New("Ошибка сервера"),
		}
	}
	defer db.Disconnect(client, ctx)

	filesDb := client.Database(os.Getenv("FILE_DATABASE"))
	bucket, err := gridfs.NewBucket(filesDb)
	if err != nil {
		return nil, nil, &error2.RequestError{
			StatusCode: http.StatusInternalServerError,
			Err:        errors.New("Ошибка сервера"),
		}
	}

	id, err := primitive.ObjectIDFromHex(hexId)
	if err != nil {
		return nil, nil, &error2.RequestError{
			StatusCode: http.StatusBadRequest,
			Err:        errors.New("Не верный id"),
		}
	}
	fileBuffer := bytes.NewBuffer(nil)
	_, err = bucket.DownloadToStream(id, fileBuffer)
	if err != nil {
		return nil, nil, &error2.RequestError{
			StatusCode: http.StatusBadRequest,
			Err:        errors.New("Файл не найден"),
		}
	}

	filter := bson.D{{
		"_id", id,
	}}
	cursor, err := bucket.Find(filter)
	if err != nil {
		return nil, nil, &error2.RequestError{
			StatusCode: http.StatusBadRequest,
			Err:        errors.New("Файл не найден"),
		}
	}

	var fileInfo []model.FileModel
	err = cursor.All(context.TODO(), &fileInfo)

	var contentType string
	for _, file := range fileInfo {
		contentType = file.Metadata.ContentType
	}

	return fileBuffer, &contentType, nil
}
