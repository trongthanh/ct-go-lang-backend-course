package mongostore

import (
	"context"
	"fmt"
	"gosocial/internal/entity"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewProfileStore(db *mongo.Database, collName string) *profileStore {
	collection := db.Collection(collName)
	// Create a unique index on the field
	useridIndex := mongo.IndexModel{
		Keys:    bson.D{{Key: "userid", Value: 1}},
		Options: options.Index().SetUnique(true),
	}

	usernameIndex := mongo.IndexModel{
		Keys:    bson.D{{Key: "username", Value: 1}},
		Options: options.Index().SetUnique(true),
	}

	_, err := collection.Indexes().CreateMany(context.Background(), []mongo.IndexModel{usernameIndex, useridIndex})
	if err != nil {
		log.Fatal(err)
	}

	return &profileStore{
		client:  collection,
		timeout: 3 * time.Second,
	}
}

type profileStore struct {
	client  *mongo.Collection
	timeout time.Duration
}

func (u *profileStore) Save(info entity.Profile) (ProfileDoc, error) {

	profileDoc := NewProfileDoc(info)

	ctx, cancelFn := context.WithTimeout(context.Background(), u.timeout)
	defer cancelFn()

	// generate default username to ensure uniqueness
	if profileDoc.Username == "" {
		count, _ := u.client.CountDocuments(ctx, bson.D{})
		if count == 0 {
			profileDoc.Username = "user_00001"
		} else {
			profileDoc.Username = fmt.Sprintf("user_%04d", count+1)
		}
	}

	result, err := u.client.InsertOne(ctx, profileDoc)
	if err != nil {
		return *profileDoc, err
	}
	// return profileDoc with new assigned ObjectID
	profileDoc.Id = result.InsertedID.(primitive.ObjectID)

	return *profileDoc, nil
}

func (u *profileStore) Get(userid string) (ProfileDoc, error) {

	filter := bson.D{{Key: "userid", Value: userid}}

	var profileDoc ProfileDoc
	err := u.client.FindOne(context.Background(), filter).Decode(&profileDoc)
	if err != nil {
		// if err == mongo.ErrNoDocuments {
		// 	return profileDoc, nil
		// }
		// no document found consider error
		return profileDoc, err
	}

	return profileDoc, nil
}

func (u *profileStore) Update(userid string, profile entity.Profile) error {
	filter := bson.D{{Key: "userid", Value: userid}}
	update := bson.M{
		"$set": bson.M{
			"bio":          profile.Bio,
			"account_type": profile.AccountType,
			"website":      profile.Website,
			"name":         profile.Name,
			"gender":       profile.Gender,
			"birthday":     profile.Birthday,
		}}

	ctx, cancel := context.WithTimeout(context.Background(), u.timeout)
	defer cancel()

	_, err := u.client.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}
