package db

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Message struct {
	Id             string   `bson:"_id"`
	Content        string   `bson:"content"`
	Sender         string   `bson:"sender"`
	UpvotesCount   uint     `bson:"upvotesCount"`
	DownvotesCount uint     `bson:"downvotesCount"`
	Upvotes        []string `bson:"upvotes"`
	Downvotes      []string `bson:"downvotes"`
	Timestamp      int64    `bson:"timestamp"`
}

// get all message history
func GetMessageHistory() (*mongo.Cursor, error) {

	coll := GetDBClient().Database("chat").Collection("messages")
	filter := bson.D{}
	opts := options.Find().SetSort(bson.D{{"timestamp", 1}}) // sort the message based on timestamp
	cur, err := coll.Find(context.Background(), filter, opts)
	if err != nil {
		return nil, err
	}
	return cur, nil
}

// store new message into database
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

// update the votes of given message in the database
func UpdateVotes(message Message) (Message, error) {

	var votesCount string
	var votesArray string
	if message.UpvotesCount != 0 {
		votesCount = "upvotesCount"
		votesArray = "upvotes"
	} else if message.DownvotesCount != 0 {
		votesCount = "downvotesCount"
		votesArray = "downvotes"
	} else {
		return Message{}, errors.New("invalid votes")
	}

	coll := GetDBClient().Database("chat").Collection("messages")
	var update bson.D
	filter := bson.D{{"_id", message.Id}, {votesArray, message.Sender}}
	// filter["_id"] = message.Id
	// filter[votesArray] = message.Sender
	update = bson.D{{"$inc", bson.D{{votesCount, -1}}}, {"$pull", bson.D{{votesArray, message.Sender}}}}
	opts := options.FindOneAndUpdate().SetReturnDocument(1) // return the new updated message
	var new_message Message
	err := coll.FindOneAndUpdate(context.Background(), filter, update, opts).Decode(&new_message)
	if err != nil {
		fmt.Println(err)
		filter := bson.D{{"_id", message.Id}}
		update = bson.D{{"$inc", bson.D{{votesCount, 1}}}, {"$push", bson.D{{votesArray, message.Sender}}}}
		opts := options.FindOneAndUpdate().SetReturnDocument(1) // return the new updated message
		err := coll.FindOneAndUpdate(context.Background(), filter, update, opts).Decode(&new_message)
		if err != nil {
			fmt.Println(err)
			return Message{}, err
		}
		fmt.Printf("Update document vote with _id: %v\n", new_message)
		return new_message, nil
	}

	return new_message, nil

}
