package auth

import (
	"context"

	"gitlab.com/komentr/access/models"
	"gitlab.com/komentr/access/service"

	"github.com/go-kit/kit/endpoint"
)

//RegisterUserEndpoint func declaration
func RegisterUserEndpoint(s service.AuthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*models.UserRequest)
		id, err := s.Register(ctx, *req)
		if err != nil {
			return nil, err
		}
		return newResponse{ID: id, Err: err}, nil
	}
}

//LoginEndpoint func declaration
func LoginEndpoint(s service.AuthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*models.LoginRequest)
		user, err := s.Login(ctx, *req)
		if err != nil {
			return nil, err
		}
		return messageResponse{Message: user, Err: err}, nil
	}
}

//RegMailEndpoint func declaration
func RegMailEndpoint(s service.AuthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*models.UserRequest)
		res, err := s.RegMail(ctx, *req)
		if err != nil {
			return nil, err
		}
		return messageResponse{Message: res, Err: err}, nil
	}
}
