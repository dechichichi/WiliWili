package service

import "wiliwili/app/user/domain/repository"

type UserService struct {
	db repository.UserDB
}

