package db

import (
	"chat-server/utils"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

type Message struct {
	Id        string `bson:"_id"`
	Content   string `bson:"content"`
	Sender    string `bson:"sender"`
	Upvotes   uint   `bson:"upvotes"`
	Downvotes uint   `bson:"downvotes"`
	Timestamp int64  `bson:"timestamp"`
}

func StoreMessage(message *Message) error {

	id, err := utils.Generate_id()
	if err != nil {
		return err
	}
	message.Id = id
	message.Timestamp = utils.Generate_timestamp()

	// Send a ping to confirm a successful connection
	coll := GetDBClient().Database("chat").Collection("messages")
	result, err := coll.InsertOne(context.TODO(), message)
	if err != nil {
		return err
	}

	fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
	return nil
}

func UpdateVotes(message Message) error {

	coll := GetDBClient().Database("chat").Collection("messages")

	var update bson.D
	filter := bson.D{{"_id", message.Id}}
	if message.Upvotes != 0 {
		update = bson.D{{"$inc", bson.D{{"upvotes", message.Upvotes}}}}
	} else if message.Downvotes != 0 {
		update = bson.D{{"$inc", bson.D{{"downvotes", message.Downvotes}}}}
	}

	result, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	fmt.Printf("Update document vote with _id: %v\n", result)
	return nil

}
