package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// MongoDB is the type that has access to the
// Mongo database and collection instances.
type MongoDB struct {
	Database   *mongo.Database
	Collection *mongo.Collection
}

// NewMongoDB instantiates a new object of type MongoDB.
func NewMongoDB(mongoURI, collection string) *MongoDB {
	var (
		db   = ConnectMongoDB(mongoURI)
		coll = db.Collection(collection)
	)
	return &MongoDB{
		Database:   db,
		Collection: coll,
	}
}

// ConnectMongoDB starts the connection of this client to the MongoDB.
func ConnectMongoDB(uri string) *mongo.Database {
	ctx := context.Background()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("Database Connection Error ", err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("Database Connection Error ", err)
	}

	return client.Database("sigma")
}
