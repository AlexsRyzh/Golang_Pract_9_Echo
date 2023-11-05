package api

import (
	"context"
	"github.com/pract9/internal/db"
	"github.com/pract9/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"os"
)

func GetFiles() ([]model.FileModel, error) {
	client, ctx, err := db.Connect()
	if err != nil {
		return nil, err
	}
	defer db.Disconnect(client, ctx)
	fileCollection := client.Database(os.Getenv("FILE_DATABASE")).Collection("fs.files")

	filter := bson.D{{}}
	var fileInfoList []model.FileModel

	cursor, err := fileCollection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	err = cursor.All(context.TODO(), &fileInfoList)
	if err != nil {
		return nil, err
	}

	return fileInfoList, nil
}
