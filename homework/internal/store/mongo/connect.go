package mongostore

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func Connect(mongoURI string, dbName string) (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(mongoURI)
	// Set the read preference
	clientOptions.SetReadPreference(readpref.Primary())
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	}

	db := client.Database("gocourse_db")

	return db, nil
}
