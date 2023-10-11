package mongostore

import (
	"context"
	"gosocial/internal/entity"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewImageStore(db *mongo.Database, collName string) *imageStore {
	collection := db.Collection(collName)

	return &imageStore{
		client:  collection,
		timeout: 3 * time.Second,
	}
}

type imageStore struct {
	client  *mongo.Collection
	timeout time.Duration
}

func (c *imageStore) Save(info entity.Image) error {

	doc := NewImageDocument(info)

	ctx, cancelFn := context.WithTimeout(context.Background(), c.timeout)
	defer cancelFn()

	_, err := c.client.InsertOne(ctx, doc)
	if err != nil {
		return err
	}

	return nil
}

func (c *imageStore) Get(username string) ([]entity.Image, error) {

	filter := bson.D{{Key: "user", Value: username}}

	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	cursor, err := c.client.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var images []entity.Image

	for cursor.Next(ctx) {
		var imgDoc ImageDoc
		if err := cursor.Decode(&imgDoc); err != nil {
			return nil, err
		}
		images = append(images, entity.Image{
			Path: imgDoc.Path,
			URL:  imgDoc.URL,
		})
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return images, nil
}

type ImageDoc struct {
	Doc  `bson:",inline"`
	Path string `json:"path" bson:"path"`
	URL  string `json:"url" bson:"url"`
}

func NewImageDocument(info entity.Image) *ImageDoc {
	return &ImageDoc{
		Doc:  NewDoc(),
		Path: info.Path,
		URL:  info.URL,
	}
}
