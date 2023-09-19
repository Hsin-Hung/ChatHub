package db

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var redis_client *redis.Client

func ConnectRedis() {
	redis_client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func PublishMessage(message Message) error {
	fmt.Println("publish message")
	payload, err := json.Marshal(message)
	if err != nil {
		panic(err)
	}
	if err := redis_client.Publish(context.Background(), "new_message", payload).Err(); err != nil {
		fmt.Println(err)
		return err
	}
	return nil

}

func SubscribeMessage(subcribe chan Message){

	subscriber := redis_client.Subscribe(context.Background(), "new_message")
	defer subscriber.Close()
	// var message Message
	var message Message
	for {
		msg, err := subscriber.ReceiveMessage(context.Background())
		if err != nil {
			panic(err)
		}
		if err := json.Unmarshal([]byte(msg.Payload), &message); err != nil {
			panic(err)
		}
		fmt.Println(message)
		subcribe <- message
	}

}
