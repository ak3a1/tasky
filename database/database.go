package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client = CreateMongoClient()

func CreateMongoClient() *mongo.Client {
	err := godotenv.Overload()
	MongoDbURI := os.Getenv("MONGODB_URI")
	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(MongoDbURI))
	if err != nil {
		log.Fatal(err)
	}
	defer cancel()
	fmt.Println("Connected to MONGO -> ", MongoDbURI)
	return client
}

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database("go-mongodb").Collection(collectionName)
}
