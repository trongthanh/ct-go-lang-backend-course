package mongostore

import (
	"context"
	"errors"
	"fmt"
	"log"
	"thanhtran/internal/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewUserStore(db *mongo.Database, collName string) *UserStore {
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

	return &UserStore{collection: collection}
}

type UserStore struct {
	collection *mongo.Collection
}

type UserDoc struct {
	Id       primitive.ObjectID `bson:"_id"`
	Username string             `bson:"username" unique:"true"`
	FullName string             `bson:"full_name"`
	Address  string             `bson:"address"`
	Password string             `bson:"password"`
}

func (u *UserStore) Save(info entity.UserInfo) error {

	userDoc := UserDoc{
		Id:       primitive.NewObjectID(),
		Username: info.Username,
		FullName: info.FullName,
		Address:  info.Address,
		Password: info.Password,
	}

	fmt.Println("UserStore.Save", userDoc)

	_, err := u.collection.InsertOne(context.Background(), userDoc)
	if err != nil {
		return err
	}

	fmt.Println("UserStore.Save", info)
	return nil
}

func (u *UserStore) Get(username string) (entity.UserInfo, error) {

	filter := bson.D{{Key: "username", Value: username}}

	var user UserDoc
	err := u.collection.FindOne(context.Background(), filter).Decode(&user)
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

var ErrUserNotFound = errors.New("user not found")
