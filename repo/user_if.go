package repo

import (
	"context"

	"gitlab.com/komentr/access/models"
)

//UserRepo interface declaration
type UserRepo interface {
	Create(ctx context.Context, req models.User) (string, error)
	Update(ctx context.Context, req models.User) (string, error)
	Delete(ctx context.Context, req models.IDRequest) error
	Search(ctx context.Context, req models.UserSearchRequest) (interface{}, error)
	Get(ctx context.Context, req models.IDRequest) (models.User, error)
	Find(ctx context.Context, req models.UserFindRequest) (models.User, error)
}
