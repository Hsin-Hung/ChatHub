package db

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

var redis_client *redis.Client

func ConnectRedis() {
	redis_uri, ok := os.LookupEnv("REDIS_URI")
	if !ok {
		redis_uri = "redis://localhost:6379"
	}

	opt, err := redis.ParseURL(redis_uri)
	if err != nil {
		panic(err)
	}

	redis_client = redis.NewClient(opt)
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

func SubscribeMessage(subcribe chan Message) {

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
