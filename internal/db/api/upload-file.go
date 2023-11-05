package api

import (
	"github.com/pract9/internal/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io"
	"mime/multipart"
	"os"
)

func UploadFile(file *multipart.FileHeader) (*primitive.ObjectID, error) {

	client, ctx, err := db.Connect()
	if err != nil {
		return nil, err
	}
	defer db.Disconnect(client, ctx)

	filesDb := client.Database(os.Getenv("FILE_DATABASE"))
	bucket, err := gridfs.NewBucket(filesDb)
	if err != nil {
		return nil, err
	}

	fileOpen, err := file.Open()
	if err != nil {
		return nil, err
	}

	uploadOpts := options.GridFSUpload().SetMetadata(bson.D{{"content-type", file.Header.Get("Content-Type")}})
	objectID, err := bucket.UploadFromStream(file.Filename, io.Reader(fileOpen), uploadOpts)

	return &objectID, err
}
