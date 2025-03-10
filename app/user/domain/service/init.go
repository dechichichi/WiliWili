package service

import (
	"sync/atomic"
	"wiliwili/app/user/domain/repository"
)

type UserService struct {
	db repository.UserDB
	cache repository.UserCache
}
var RedisAvailable atomic.Bool

func NewUserService(db repository.UserDB, cache repository.UserCache) *UserService {
	if cache == nil {
		panic("userService`s cache should not be nil")
	}
	if db == nil {
		panic("userService`s db should not be nil")
	}
	svc:= &UserService{
		db: db,
		cache: cache,
	}
	return svc

}
