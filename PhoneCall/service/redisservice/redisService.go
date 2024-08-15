package redisservice

import (
	"github.com/redis/go-redis/v9"
)

type RedisService struct {
	Client *redis.Client
}

func NewRedisService(Client *redis.Client) *RedisService {
	return &RedisService{Client: Client}
}
