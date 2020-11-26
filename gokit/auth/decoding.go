package auth

import (
	"gitlab.com/komentr/access/models"
	"context"
	"encoding/json"
	"net/http"
)

func decodeNewUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var body = new(models.UserRequest)
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}
	return body, nil
}

func decodeLoginRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var body = new(models.LoginRequest)
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}
	return body, nil
}
