package database

import (
	"context"
	"fmt"
	"log"
	"os"

	//"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/rabbice/ecommerce/logs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func DBInstance() *mongo.Client {
	//err := godotenv.Load(".env")
	//if err != nil {
	//log.Fatal("Error loading .env file")
	//}
	MongoDb := os.Getenv("MONGODB_URI")
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(MongoDb))
	if err != nil {
		log.Fatal(err)
	}
	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	prometheus.Register(logs.TotalRequests)
	prometheus.Register(logs.TotalHTTPMethods)
	prometheus.Register(logs.HTTPDuration)

	return client
}

var Client *mongo.Client = DBInstance()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database(os.Getenv("MONGODB_DATABASE")).Collection(collectionName)

	return collection
}
