package service

import (
	"gitlab.com/komentr/access/repo"
)

type authService struct {
	userRepo repo.UserRepo
}

//NewAuthService function declaration
func NewAuthService(
	a repo.UserRepo,
) AuthService {
	return &authService{
		userRepo: a,
	}
}
