package api

import (
	"context"
	"errors"
	"github.com/pract9/internal/db"
	error2 "github.com/pract9/internal/error"
	"github.com/pract9/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"os"
)

func GetFileInfoById(hexId string) (*model.FileModel, *error2.RequestError) {

	client, ctx, err := db.Connect()
	if err != nil {
		return nil, &error2.RequestError{
			StatusCode: http.StatusInternalServerError,
			Err:        errors.New("Ошибка сервера"),
		}
	}
	defer db.Disconnect(client, ctx)
	fileCollection := client.Database(os.Getenv("FILE_DATABASE")).Collection("fs.files")

	id, err := primitive.ObjectIDFromHex(hexId)
	if err != nil {
		return nil, &error2.RequestError{
			StatusCode: http.StatusBadRequest,
			Err:        errors.New("Не верный ID"),
		}
	}
	filter := bson.D{{
		"_id", id,
	}}

	var fileInfo model.FileModel
	err = fileCollection.FindOne(context.TODO(), filter).Decode(&fileInfo)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, &error2.RequestError{
			StatusCode: http.StatusBadRequest,
			Err:        errors.New("Информация не найдена"),
		}
	}

	return &fileInfo, nil
}
