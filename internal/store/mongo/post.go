package mongostore

import (
	"context"
	"gosocial/internal/entity"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewPostStore(db *mongo.Database, collName string) *postStore {
	collection := db.Collection(collName)

	return &postStore{
		client:  collection,
		timeout: 3 * time.Second,
	}
}

type postStore struct {
	client  *mongo.Collection
	timeout time.Duration
}

func (u *postStore) Save(info entity.Post) (PostDoc, error) {

	postDoc := NewPostDoc(info)

	ctx, cancelFn := context.WithTimeout(context.Background(), u.timeout)
	defer cancelFn()

	result, err := u.client.InsertOne(ctx, postDoc)
	if err != nil {
		return *postDoc, err
	}
	// return postDoc with new assigned ObjectID
	postDoc.DocId = result.InsertedID.(primitive.ObjectID)

	return *postDoc, nil
}

func (u *postStore) GetManyByUser(userid string) ([]PostDoc, error) {

	filter := bson.D{{Key: "userid", Value: userid}}

	// Pagination parameters
	pageNumber := 1 // Current page number (start from 1)
	pageSize := 10  // Number of documents per page
	skip := (pageNumber - 1) * pageSize

	//
	ctx, cancelFn := context.WithTimeout(context.Background(), u.timeout)
	defer cancelFn()
	findOptions := options.Find().SetSkip(int64(skip)).SetLimit(int64(pageSize))
	cursor, err := u.client.Find(ctx, filter, findOptions)
	if err != nil {
		return []PostDoc{}, err
	}

	var postDocs []PostDoc
	if err = cursor.All(context.Background(), &postDocs); err != nil {
		return []PostDoc{}, err
	}

	return postDocs, nil
}

func (u *postStore) GetMany() ([]PostDoc, error) {
	// Pagination parameters
	pageNumber := 1 // Current page number (start from 1)
	pageSize := 10  // Number of documents per page
	skip := (pageNumber - 1) * pageSize

	//
	ctx, cancelFn := context.WithTimeout(context.Background(), u.timeout)
	defer cancelFn()
	findOptions := options.Find().SetSkip(int64(skip)).SetLimit(int64(pageSize))
	cursor, err := u.client.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return []PostDoc{}, err
	}

	var postDocs []PostDoc
	if err = cursor.All(context.Background(), &postDocs); err != nil {
		return []PostDoc{}, err
	}

	return postDocs, nil
}

func (u *postStore) DeleteOne(postid string) error {

	ctx, cancelFn := context.WithTimeout(context.Background(), u.timeout)
	defer cancelFn()

	_, err := u.client.DeleteOne(ctx, bson.M{"_id": postid})
	if err != nil {
		return err
	}

	return nil
}

func (u *postStore) LikePost(postid string, userid string) (int, error) {

	ctx, cancelFn := context.WithTimeout(context.Background(), u.timeout)
	defer cancelFn()

	_, err := u.client.UpdateOne(ctx, bson.M{"_id": postid},
		bson.M{"$addToSet": bson.M{"likes": bson.M{"userid": userid}}})
	if err != nil {
		return -1, err
	}

	var postDoc PostDoc
	err = u.client.FindOne(ctx, bson.M{"_id": postid}).Decode(&postDoc)
	if err != nil {
		return -1, err
	}

	return len(postDoc.Post.Likes), nil
}
