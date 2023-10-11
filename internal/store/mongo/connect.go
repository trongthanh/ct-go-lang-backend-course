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
	client, _ := mongo.Connect(ctx, clientOptions)

	// Use the Ping method to check the connection
	err := client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	db := client.Database("gosocial_db")

	return db, nil
}
