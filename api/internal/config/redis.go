package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

func SetupRedis(config *Config) *redis.Client {
	addr := fmt.Sprintf("%s:%s", config.redis_HOST, config.redis_PORT)
	rdb := redis.NewClient(&redis.Options{
		Addr:         addr,
		Password:     config.redis_PASS,
		DB:           config.redis_DB,
		PoolSize:     10,
		MinIdleConns: 5,
		DialTimeout:  3 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		MaxRetries:   3,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Printf("Redis unavailable: %v", err)
		log.Println("Application will continue without Redis.")
	} else {
		log.Println("Redis connected successfully")
	}
	return rdb
}
