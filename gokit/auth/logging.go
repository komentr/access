package auth

import (
	"context"
	"time"

	"gitlab.com/komentr/access/models"
	"gitlab.com/komentr/access/service"

	"github.com/go-kit/kit/log"
)

type loggingService struct {
	logger log.Logger
	service.AuthService
}

//NewLoggingService function
func NewLoggingService(logger log.Logger, s service.AuthService) service.AuthService {
	return &loggingService{logger, s}
}

func (s *loggingService) Register(ctx context.Context, p models.UserRequest) (id string, err error) {
	defer func(begin time.Time) {
		if err != nil {
			s.logger.Log(
				"took", time.Since(begin),
				"err", err,
				"method", "register",
				"service_name", "access",
				"service_type", "command",
				"body", p,
			)
		}
	}(time.Now())
	return s.AuthService.Register(ctx, p)
}

func (s *loggingService) Login(ctx context.Context, p models.LoginRequest) (data interface{}, err error) {
	defer func(begin time.Time) {
		if err != nil {
			s.logger.Log(
				"took", time.Since(begin),
				"err", err,
				"method", "login",
				"service_name", "access",
				"service_type", "command",
				"body", p,
			)
		}
	}(time.Now())
	return s.AuthService.Login(ctx, p)
}

func (s *loggingService) RegMail(ctx context.Context, p models.UserRequest) (data interface{}, err error) {
	defer func(begin time.Time) {
		if err != nil {
			s.logger.Log(
				"took", time.Since(begin),
				"err", err,
				"method", "regmail",
				"service_name", "access",
				"service_type", "command",
				"body", p,
			)
		}
	}(time.Now())
	return s.AuthService.RegMail(ctx, p)
}
