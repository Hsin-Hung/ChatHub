package db

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db_client *mongo.Client

func ConnectDB() {
	db_uri, ok := os.LookupEnv("DB_URI")
	if !ok {
		db_uri = "mongodb://chat:chat@localhost:27017/messages"
	}
	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(db_uri).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	var err error
	db_client, err = mongo.Connect(context.Background(), opts)

	if err != nil {
		panic(err)
	}

	// Send a ping to confirm a successful connection
	var result bson.M
	if err := db_client.Database("messages").RunCommand(context.Background(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
}

// instead do it at database init script
func CreateIndexes() {

	chatIndexModel := mongo.IndexModel{
		Keys: bson.D{
			{"timestamp", 1},
		},
	}

	msg_coll := GetDBClient().Database("messages").Collection("messages")
	name, err := msg_coll.Indexes().CreateOne(context.Background(), chatIndexModel)
	if err != nil {
		panic(err)
	}
	fmt.Println("Messages: Name of Index Created: " + name)

}

func GetDBClient() *mongo.Client {
	return db_client
}
