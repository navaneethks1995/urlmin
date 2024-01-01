package store

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type StorageService struct {
	redisClient *redis.Client
}

var (
	storeService = &StorageService{}
	ctx          = context.Background()
)

const cacheDuration = 6 * time.Hour

func InitializeStorage() *StorageService {
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		// Password: "redis_password",
		DB: 0,
	})

	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Error initialising Redis: %v", err))
	}
	fmt.Printf("\nRedis started successfully: pong message = {%s}", pong)
	storeService.redisClient = redisClient
	return storeService
}

func SaveUrlMapping(longUrl string, minUrl string, userid string) {
	err := storeService.redisClient.Set(ctx, minUrl, longUrl, cacheDuration).Err()
	if err != nil {
		panic(fmt.Sprintf("Failed saving key url | Error: %v - shortUrl: %s - originalUrl: %s\n", err, minUrl, longUrl))
	}
}

func GetOriginalUrl(minUrl string) string {
	result, err := storeService.redisClient.Get(ctx, minUrl).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed RetrieveInitialUrl url | Error: %v - shortUrl: %s\n", err, minUrl))
	}
	return result
}
