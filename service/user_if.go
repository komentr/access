package service

import (
	"context"
	"gitlab.com/komentr/access/models"
)

//UserService type interface declaration
type UserService interface {
	Create(ctx context.Context, req models.UserRequest) (string, error)
	Update(ctx context.Context, req models.User) (string, error)
	Search(ctx context.Context, req models.UserSearchRequest) (interface{}, error)
	Get(ctx context.Context, req models.IDRequest) (interface{}, error)
	Delete(ctx context.Context, req models.IDRequest) (string, error)
	Find(ctx context.Context, req models.UserFindRequest) (models.User, error)
}
