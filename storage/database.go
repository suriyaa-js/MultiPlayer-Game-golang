// /storage/database.go
package storage

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

// InitializeDB connects to the MongoDB database
func InitializeDB() error {
	mongoURL := os.Getenv("MONGO_DB_URL")
	dbName := os.Getenv("MONGO_DB_NAME")
	collection := os.Getenv("MONGO_DB_COLLECTION")
	log.Printf("Connecting to mongoDB at %s", mongoURL)

	var err error
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURL))
	if err != nil {
		return err
	}
	playerCollection = client.Database(dbName).Collection(collection) // Replace with your database name
	InitMongo(playerCollection)
	return nil
}

func GetClient() *mongo.Client {
	return client
}
