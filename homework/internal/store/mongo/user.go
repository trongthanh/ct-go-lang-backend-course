package mongostore

import (
	"context"
	"errors"
	"fmt"
	"log"
	"thanhtran/internal/entity"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewUserStore(db *mongo.Database, collName string) *userStore {
	collection := db.Collection(collName)
	// Create a unique index on the "username" field
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "username", Value: 1}},
		Options: options.Index().SetUnique(true),
	}

	_, err := collection.Indexes().CreateOne(context.Background(), indexModel)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Unique index at username created successfully")

	return &userStore{
		client:  collection,
		timeout: 3 * time.Second,
	}
}

type userStore struct {
	client  *mongo.Collection
	timeout time.Duration
}

type UserDoc struct {
	Id       primitive.ObjectID `bson:"_id"`
	Username string             `bson:"username" unique:"true"`
	FullName string             `bson:"full_name"`
	Address  string             `bson:"address"`
	Password string             `bson:"password"`
}

func (u *userStore) Save(info entity.UserInfo) error {

	userDoc := NewUserDocument(info)

	ctx, cancelFn := context.WithTimeout(context.Background(), u.timeout)
	defer cancelFn()

	_, err := u.client.InsertOne(ctx, userDoc)
	if err != nil {
		return err
	}

	fmt.Println("UserStore.Save", info)
	return nil
}

func (u *userStore) Get(username string) (entity.UserInfo, error) {

	filter := bson.D{{Key: "username", Value: username}}

	var user UserDoc
	err := u.client.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return entity.UserInfo{}, nil // Return empty UserInfo if no documents found
		}
		return entity.UserInfo{}, err
	}

	userInfo := entity.UserInfo{
		Username: user.Username,
		FullName: user.FullName,
		Address:  user.Address,
		Password: user.Password,
	}

	return userInfo, nil
}

func (u *userStore) Update(info entity.UserInfo) error {
	filter := bson.D{{Key: "username", Value: info.Username}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "full_name", Value: info.FullName},
			{Key: "address", Value: info.Address},
			{Key: "password", Value: info.Password},
		}},
	}

	ctx, cancel := context.WithTimeout(context.Background(), u.timeout)
	defer cancel()

	_, err := u.client.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func NewUserDocument(info entity.UserInfo) UserDoc {
	return UserDoc{
		Id:       primitive.NewObjectID(),
		Username: info.Username,
		FullName: info.FullName,
		Address:  info.Address,
		Password: info.Password,
	}
}

var ErrUserNotFound = errors.New("user not found")
