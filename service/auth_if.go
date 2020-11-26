package service

import (
	"context"

	"gitlab.com/komentr/access/models"
)

//AuthService interface declaration
type AuthService interface {
	Register(ctx context.Context, req models.UserRequest) (string, error)
	Login(ctx context.Context, req models.LoginRequest) (interface{}, error)
	RegMail(ctx context.Context, req models.UserRequest) (interface{}, error)
}
