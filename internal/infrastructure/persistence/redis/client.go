package redis

import (
	redislib "github.com/redis/go-redis/v9"
)

func Connect(addr string, password string, db int) *redislib.Client {
	client := redislib.NewClient(&redislib.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	return client
}
