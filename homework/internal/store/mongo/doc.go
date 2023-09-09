package mongostore

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Doc struct {
	Id        primitive.ObjectID `bson:"_id"`
	Version   int64              `bson:"version"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

func NewDoc() Doc {
	docId := primitive.NewObjectID()
	return Doc{
		Id:        docId,
		Version:   1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
