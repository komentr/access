package service

import (
	"context"
	"errors"
	"os"

	"gitlab.com/komentr/access/models"
	"gitlab.com/komentr/utils/errorcode"
)

//TokenSecret const
const (
	TokenSecret = "TOKENSECRET"
)

func (s *authService) Register(ctx context.Context, req models.UserRequest) (string, error) {

	if err := ValidateRegistration(req); err != nil {
		return "", err
	}
	// verify username not exist
	if user, err := s.userRepo.Find(ctx, models.UserFindRequest{Username: req.Username}); err == nil {
		return "", errors.New("username : " + user.Username + " exist")
	}

	req.Active = models.UserStatusActive

	newUser, err := Register(req)
	if err != nil {
		return "", err
	}
	newUser.Role = models.USER

	id, err := s.userRepo.Create(ctx, newUser)
	newUser.ID = id
	if err != nil {
		return "", err
	}

	return id, nil
}

func (s *authService) Login(ctx context.Context, req models.LoginRequest) (interface{}, error) {
	var err error
	var user models.User
	var userSearchReq models.UserFindRequest

	secret := os.Getenv(TokenSecret)
	if secret == "" {
		secret = "supersecret"
	}

	if req.Username != "" {
		userSearchReq.Username = req.Username
		user, err = s.userRepo.Find(ctx, userSearchReq)
		if err != nil {
			return nil, err
		}
	}

	if !CheckPasswordHash(req.Password, user.Password) {
		return nil, errorcode.ErrPassInvalid
	}

	token, err := GenerateToken([]byte(secret), user.ID, user.TID, user.Role, 0)
	if err != nil {
		return nil, err
	}

	return models.LoginResponse{
		ID:       user.ID,
		Username: user.Username,
		Token:    token,
		Name:     user.Name,
		Role:     user.Role,
		Image:    user.Image,
		TID:      user.TID,
		Active:   user.Active,
	}, nil
}

func (s *authService) RegMail(ctx context.Context, req models.UserRequest) (interface{}, error) {
	var id, token string
	var err error

	if err := validateEmail(req.Username); err != nil {
		return nil, err
	}

	secret := os.Getenv(TokenSecret)
	if secret == "" {
		secret = "supersecret"
	}

	// verify username is exist
	if user, err := s.userRepo.Find(ctx, models.UserFindRequest{Username: req.Username}); err == nil {
		// return "", errors.New("username : " + user.Username + " exist")
		token, err = GenerateToken([]byte(secret), user.ID, user.TID, user.Role, 0)
		if err != nil {
			return nil, err
		}
		id = user.ID

		return models.RegMailResponse{ID: id, Token: token}, nil
	}

	req.Password = DefaultPass
	req.Active = models.UserStatusOnlyName

	newUser, err := Register(req)
	if err != nil {
		return nil, err
	}
	newUser.Role = models.USER

	id, err = s.userRepo.Create(ctx, newUser)
	newUser.ID = id
	if err != nil {
		return nil, err
	}

	token, err = GenerateToken([]byte(secret), newUser.ID, newUser.TID, newUser.Role, 0)
	if err != nil {
		return nil, err
	}

	return models.RegMailResponse{ID: id, Token: token}, nil
}
