package userservice

import (
	"PhoneCall/repository"
	"PhoneCall/service/redisservice"
)

type UserService struct {
	RedisService redisservice.RedisService
	UserRepo     repository.UserRepo
}

func NewUserService(UserRepo repository.UserRepo, RedisService redisservice.RedisService) *UserService {
	return &UserService{
		UserRepo:     UserRepo,
		RedisService: RedisService,
	}
}
