package service

import (
	"gitlab.com/komentr/access/repo"
)

type userService struct {
	userRepo repo.UserRepo
}

//NewUserService function declaration
func NewUserService(
	a repo.UserRepo,
) UserService {
	return &userService{
		userRepo: a,
	}
}
