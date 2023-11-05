package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type FileModel struct {
	ID         primitive.ObjectID `bson:"_id"`
	Length     int64              `bson:"length"`
	ChunkSize  int32              `bson:"chunkSize"`
	UploadData time.Time          `bson:"uploadData"`
	Filename   string             `bson:"filename"`
	Metadata   MetaData           `bson:"metadata"`
}

type MetaData struct {
	ContentType string `bson:"content-type"`
}
