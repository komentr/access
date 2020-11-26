package access

import (
	"context"
	"time"

	"gitlab.com/komentr/access/models"
	"gitlab.com/komentr/access/service"

	"github.com/go-kit/kit/log"
)

type loggingService struct {
	logger log.Logger
	service.UserService
}

//NewLoggingService function
func NewLoggingService(logger log.Logger, s service.UserService) service.UserService {
	return &loggingService{logger, s}
}

func (s *loggingService) Create(ctx context.Context, p models.UserRequest) (id string, err error) {
	defer func(begin time.Time) {
		if err != nil {
			s.logger.Log(
				"took", time.Since(begin),
				"err", err,
				"method", "new user",
				"service_name", "access",
				"service_type", "command",
				"body", p,
			)
		}
	}(time.Now())
	return s.UserService.Create(ctx, p)
}

func (s *loggingService) Update(ctx context.Context, p models.User) (id string, err error) {
	defer func(begin time.Time) {
		if err != nil {
			s.logger.Log(
				"took", time.Since(begin),
				"err", err,
				"method", "update user",
				"service_name", "access",
				"service_type", "command",
				"body", p,
			)
		}
	}(time.Now())
	return s.UserService.Update(ctx, p)
}

func (s *loggingService) Search(ctx context.Context, p models.UserSearchRequest) (data interface{}, err error) {
	defer func(begin time.Time) {
		if err != nil {
			s.logger.Log(
				"took", time.Since(begin),
				"err", err,
				"method", "search user",
				"service_name", "access",
				"service_type", "command",
				"body", p,
			)
		}
	}(time.Now())
	return s.UserService.Search(ctx, p)
}

func (s *loggingService) Get(ctx context.Context, p models.IDRequest) (data interface{}, err error) {
	defer func(begin time.Time) {
		if err != nil {
			s.logger.Log(
				"took", time.Since(begin),
				"err", err,
				"method", "get user",
				"service_name", "access",
				"service_type", "command",
				"body", p,
			)
		}
	}(time.Now())
	return s.UserService.Get(ctx, p)
}

func (s *loggingService) Delete(ctx context.Context, p models.IDRequest) (id string, err error) {
	defer func(begin time.Time) {
		if err != nil {
			s.logger.Log(
				"took", time.Since(begin),
				"err", err,
				"method", "delete user",
				"service_name", "access",
				"service_type", "command",
				"body", p,
			)
		}
	}(time.Now())
	return s.UserService.Delete(ctx, p)
}

func (s *loggingService) Find(ctx context.Context, p models.UserFindRequest) (data models.User, err error) {
	defer func(begin time.Time) {
		if err != nil {
			s.logger.Log(
				"took", time.Since(begin),
				"err", err,
				"method", "find user",
				"service_name", "access",
				"service_type", "command",
				"body", p,
			)
		}
	}(time.Now())
	return s.UserService.Find(ctx, p)
}
