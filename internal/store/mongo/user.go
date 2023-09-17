package mongostore

import (
	"context"
	"errors"
	"fmt"
	"gosocial/internal/entity"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewUserStore(db *mongo.Database, collName string) *userStore {
	collection := db.Collection(collName)
	// Create a unique index on the "email" field
	indexModel := mongo.IndexModel{
		Keys: bson.D{
			{Key: "email", Value: 1},
		},
		Options: options.Index().SetUnique(true),
	}

	_, err := collection.Indexes().CreateOne(context.Background(), indexModel)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Unique index at email created successfully")

	return &userStore{
		client:  collection,
		timeout: 3 * time.Second,
	}
}

type userStore struct {
	client  *mongo.Collection
	timeout time.Duration
}

func (u *userStore) Save(info entity.User) (primitive.ObjectID, error) {

	userDoc := NewUserDoc(info)

	ctx, cancelFn := context.WithTimeout(context.Background(), u.timeout)
	defer cancelFn()

	result, err := u.client.InsertOne(ctx, userDoc)
	if err != nil {
		return primitive.ObjectID{}, err
	}

	fmt.Println("UserStore.Save", result)
	return result.InsertedID.(primitive.ObjectID), nil
}

func (u *userStore) Get(id string) (UserDoc, error) {

	filter := bson.D{{Key: "_id", Value: id}}

	var userDoc UserDoc
	err := u.client.FindOne(context.Background(), filter).Decode(&userDoc)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return userDoc, nil // Return empty User if no documents found
		}
		return userDoc, err
	}

	return userDoc, nil
}

func (u *userStore) GetByEmail(email string) (UserDoc, error) {

	filter := bson.D{{Key: "email", Value: email}}

	var userDoc UserDoc
	err := u.client.FindOne(context.Background(), filter).Decode(&userDoc)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return userDoc, nil // Return empty User if no documents found
		}
		return userDoc, err
	}

	return userDoc, nil
}

func (u *userStore) Update(id string, user entity.User) error {
	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.M{
		"$set": bson.M{
			"email":           user.Email,
			"hashed_password": user.HashedPassword,
			"active":          user.Active,
		}}

	ctx, cancel := context.WithTimeout(context.Background(), u.timeout)
	defer cancel()

	_, err := u.client.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

var ErrUserNotFound = errors.New("user not found")
