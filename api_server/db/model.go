package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	Username string `bson:"_id"`
	Password string `bson:"password"`
}

// Create a new signed up user
func CreateNewUser(username string, password string) error {

	coll := GetClient().Database("users").Collection("users")
	user := User{Username: username, Password: password}
	result, err := coll.InsertOne(context.Background(), user)
	// Check for a duplicate key error
	if err != nil {
		if writeErr, ok := err.(mongo.WriteException); ok {
			for _, writeError := range writeErr.WriteErrors {
				if writeError.Code == 11000 { // MongoDB error code for duplicate key (username)
					return fmt.Errorf("username already exists")
				}
			}
		}
		return fmt.Errorf("registration failed")
	}

	fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
	return nil
}

// Validate if an username exists
func ValidateUser(username string) (string, error) {

	coll := GetClient().Database("users").Collection("users")
	filter := bson.M{"_id": username}
	var user User
	if err := coll.FindOne(context.Background(), filter).Decode(&user); err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Printf("Sucessfully validate user: %s\n", user.Username)

	return user.Password, nil

}
