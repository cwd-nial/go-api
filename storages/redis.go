package storages

import (
	"github.com/go-redis/redis"
	"log"
	"sync"
)

type RedisClient struct{ *redis.Client }

const key = "artists"

var once sync.Once
var redisClient *RedisClient

func GetRedisClient() *RedisClient {

	once.Do(func() {
		client := redis.NewClient(&redis.Options{
			Addr:     "redis:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		})

		redisClient = &RedisClient{client}
	})

	_, err := redisClient.Ping().Result()
	if err != nil {
		log.Fatalf("Could not connect to redis %v", err)
	}

	return redisClient
}
