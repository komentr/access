package access

import (
	"context"

	"gitlab.com/komentr/access/models"
	"gitlab.com/komentr/access/service"

	"github.com/go-kit/kit/endpoint"
)

//NewUserEndpoint endpoint
func NewUserEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		// r := request.(newRequest)
		// req := r.Req.(*models.UserRequest)
		// id, err := s.Create(r.Ctx, *req)
		req := request.(*models.UserRequest)
		id, err := s.Create(ctx, *req)
		if err != nil {
			return nil, err
		}
		return newResponse{ID: id, Err: err}, nil
	}
}

//UpdateUserEndpoint endpoint
func UpdateUserEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*models.User)
		stat, err := s.Update(ctx, *req)
		if err != nil {
			return nil, err
		}
		return messageResponse{Message: stat, Err: err}, nil
	}
}

//SearchUserEndpoint endpoint
func SearchUserEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*models.UserSearchRequest)
		message, err := s.Search(ctx, *req)
		if err != nil {
			return nil, err
		}
		return messageResponse{message, err}, nil
	}
}

//GetUserEndpoint endpoint
func GetUserEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*models.IDRequest)
		message, err := s.Get(ctx, *req)
		if err != nil {
			return nil, err
		}
		return messageResponse{Message: message, Err: err}, nil
	}
}

//DeleteUserEndpoint endpoint
func DeleteUserEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*models.IDRequest)
		stat, err := s.Delete(ctx, *req)
		if err != nil {
			return nil, err
		}
		return messageResponse{Message: stat, Err: err}, nil
	}
}
