package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Message struct {
	Id        string `bson:"_id"`
	Content   string `bson:"content"`
	Sender    string `bson:"sender"`
	Upvotes   uint   `bson:"upvotes"`
	Downvotes uint   `bson:"downvotes"`
	Timestamp int64  `bson:"timestamp"`
}

func GetMessageHistory() (*mongo.Cursor, error) {

	coll := GetDBClient().Database("chat").Collection("messages")
	cur, err := coll.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	return cur, nil
}

func StoreMessage(message Message) error {

	// Send a ping to confirm a successful connection
	coll := GetDBClient().Database("chat").Collection("messages")
	result, err := coll.InsertOne(context.Background(), message)
	if err != nil {
		return err
	}

	fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
	return nil
}

func UpdateVotes(message Message) (Message, error) {

	coll := GetDBClient().Database("chat").Collection("messages")

	var update bson.D
	filter := bson.D{{"_id", message.Id}}
	if message.Upvotes != 0 {
		update = bson.D{{"$inc", bson.D{{"upvotes", message.Upvotes}}}}
	} else if message.Downvotes != 0 {
		update = bson.D{{"$inc", bson.D{{"downvotes", message.Downvotes}}}}
	}
	opts := options.FindOneAndUpdate().SetReturnDocument(1)
	var new_message Message
	err := coll.FindOneAndUpdate(context.Background(), filter, update, opts).Decode(&new_message)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Update document vote with _id: %v\n", new_message)
	return new_message, nil

}
