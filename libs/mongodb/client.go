package mongodb

import (
	"context"
	"graphql_api/keys"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var db *mongo.Database

/*
	Creates a MongoDB client and establishes a connection and
	assigns a pointer to the database to db.
*/
func InitiateDatabase() {
	if db != nil {
		return
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	client, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI(keys.GetKeys().MONGO_URI),
	)

	if err != nil {
		log.Fatal(err)
	}

	failedConnection := client.Ping(ctx, readpref.Primary())

	if failedConnection != nil {
		log.Fatal("❌", err)
	}

	log.Println("✅ Connection to MongoDB established")

	db = client.Database(keys.GetKeys().MONGO_DB_NAME)
}

/*
	Returns a pointer to the MongoDB database that can
	be used throughout the application.
*/
func GetClient() *mongo.Database {
	if db != nil {
		return db
	}

	InitiateDatabase()

	return db
}
