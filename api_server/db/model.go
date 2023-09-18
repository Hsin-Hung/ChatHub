package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

type User struct {
	Username string `bson:"_id"`
	Password string `bson:"password"`
}

func CreateNewUser(username string, password string) error {

	// Send a ping to confirm a successful connection
	coll := GetClient().Database("admin").Collection("users")
	doc := User{Username: username, Password: password}
	result, err := coll.InsertOne(context.TODO(), doc)
	if err != nil {
		return err
	}

	fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
	return nil
}

func ValidateUser(username string) (string, error) {

	// Send a ping to confirm a successful connection
	coll := GetClient().Database("admin").Collection("users")
	filter := bson.M{"_id": username}
	var user User
	if err := coll.FindOne(context.TODO(), filter).Decode(&user); err != nil {
		fmt.Println(err)
		return "", err
	}else{
		fmt.Printf("Sucessfully validate user: %s\n", user.Username)
	}

	return user.Password, nil

}
