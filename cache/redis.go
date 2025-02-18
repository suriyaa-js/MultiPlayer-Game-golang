package cache

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client

func InitializeCache() {
	redisUrl := os.Getenv("REDIS_URL")
	log.Printf("Connecting to Redis at %s", redisUrl)
	rdb = redis.NewClient(&redis.Options{
		Addr: redisUrl, // Redis server address
		// Password: "",
		// DB:       0,
	})

	isConnected, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("redis cache connection is successful %v", isConnected)
	fmt.Println()
	if rdb == nil {
		fmt.Println("redis cache connection is empty() ")
	}
	flushAll(context.Background())
}
