package service

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"gitlab.com/komentr/access/models"
	"gitlab.com/komentr/utils/errorcode"
)

func (s *userService) Create(ctx context.Context, req models.UserRequest) (string, error) {
	jwt := ctx.Value("jwtclaim").(models.JwtClaim)
	bss := strings.Split(jwt.TID, "/")
	bssReq := strings.Split(req.TID, "/")

	if jwt.Role != models.ADMIN {
		return "", errorcode.ErrUnAuthorized
	}

	if jwt.TID != models.TIDDefault {
		if req.TID == "" {
			req.TID = models.TIDCommon
		}

		if req.Role <= jwt.Role {
			fmt.Printf("req.Role <= jwt.Role: %d, %d \n", req.Role, jwt.Role)
			return "", errorcode.ErrInvalidArgument
		}
		if len(bssReq) < len(bss) {
			fmt.Printf("len(bssReq) < len(bss): %s, %s\n", bssReq, bss)
			return "", errorcode.ErrInvalidArgument
		}

		req.RegBy = jwt.UserID
		req.TID = jwt.TID

		for i, b := range bss {
			if bssReq[i] != b {
				return "", errorcode.ErrUnAuthorized
			}
		}
	}

	if req.Password == "" {
		req.Password = getDefaultPass()
	}

	if err := ValidateRegistration(req); err != nil {
		return "", err
	}
	// verify username not exist
	if user, err := s.userRepo.Find(ctx, models.UserFindRequest{Username: req.Username}); err == nil {
		return "", errors.New("username : " + user.Username + " exist")
	}

	newUser, err := New(req)
	if err != nil {
		return "", err
	}
	id, err := s.userRepo.Create(ctx, newUser)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (s *userService) Update(ctx context.Context, req models.User) (string, error) {
	jwt := ctx.Value("jwtclaim").(models.JwtClaim)
	bss := strings.Split(jwt.TID, "/")
	bssReq := strings.Split(req.TID, "/")

	if jwt.Role != models.ADMIN && jwt.Role != req.Role {
		return "", errorcode.ErrUnAuthorized
	}

	if jwt.TID != models.TIDDefault {
		if len(bssReq) <= len(bss) {
			req.TID = jwt.TID
		}
		if len(bssReq) > len(bss) {
			for i, b := range bss {
				if bssReq[i] != b {
					return "", errorcode.ErrUnAuthorized
				}
			}
		}
		if jwt.TID == "" && jwt.UserID != req.ID {
			return "", errorcode.ErrUnAuthorized
		}
	}

	// verify username
	if user, err := s.userRepo.Get(ctx, models.IDRequest{ID: req.ID}); err == nil {
		req.Username = user.Username
		req.Created = user.Created
		req.RegBy = user.RegBy
	} else {
		return "", err
	}

	status, err := s.userRepo.Update(ctx, req)
	if err != nil {
		return "", err
	}

	return status, err
}

func (s *userService) Search(ctx context.Context, req models.UserSearchRequest) (interface{}, error) {
	jwt := ctx.Value("jwtclaim").(models.JwtClaim)
	bss := strings.Split(jwt.TID, "/")
	bssReq := strings.Split(req.TID, "/")

	if jwt.TID == "" {
		return "", errorcode.ErrUnAuthorized
	}

	if jwt.TID != models.TIDDefault {
		if len(bssReq) <= len(bss) {
			req.TID = jwt.TID
		}
		if len(bssReq) > len(bss) {
			for i, b := range bss {
				if bssReq[i] != b {
					return "", errorcode.ErrUnAuthorized
				}
			}
		}

		if jwt.Role == models.USER {
			return "", errorcode.ErrUnAuthorized
		}

		if jwt.Role != models.ADMIN {
			req.Role = jwt.Role
		}
	}

	return s.userRepo.Search(ctx, req)
}

func (s *userService) Get(ctx context.Context, req models.IDRequest) (interface{}, error) {
	jwt := ctx.Value("jwtclaim").(models.JwtClaim)

	if jwt.TID == "" {
		return nil, errorcode.ErrUnAuthorized
	}

	return s.userRepo.Get(ctx, req)
}

func (s *userService) Delete(ctx context.Context, req models.IDRequest) (string, error) {
	jwt := ctx.Value("jwtclaim").(models.JwtClaim)

	if jwt.TID == "" {
		return "", errorcode.ErrUnAuthorized
	}

	if jwt.Role != models.ADMIN {
		return "", errorcode.ErrUnAuthorized
	}

	return "ok", s.userRepo.Delete(ctx, req)
}

func (s *userService) Find(ctx context.Context, req models.UserFindRequest) (models.User, error) {
	return s.userRepo.Find(ctx, req)
}
