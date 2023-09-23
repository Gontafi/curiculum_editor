package db

import (
	"educational_program_creator/internal/config"
	"github.com/redis/go-redis/v9"
)

func NewRedisClient(cfg *config.AppConfig) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return rdb
}
