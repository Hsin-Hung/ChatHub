package db

import (
	"github.com/redis/go-redis/v9"
	"context"
)

var redis_client *redis.Client

func ConnectRedis() {
	redis_client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func GetRedisClient() *redis.Client {
	return redis_client
}

func RegisterClient(username string){

	ctx := context.Background()

	err := redis_client.Set(ctx, username, "online", 0).Err()
	if err != nil {
		panic(err)
	}


}

// func DoesClientExist(username string){

// 	ctx := context.Background()

// 	val, err := redis_client.Get(ctx, username).Result()
// 	if err != nil {
// 		panic(err)
// 	}



// }

