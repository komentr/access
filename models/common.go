package models

import (
	jwt "github.com/dgrijalva/jwt-go"
)

const (
	TenantDefault  = "Default"
	TIDDefault     = "1"
	TenantCommon   = "Common"
	TIDCommon      = "99"
	Expiration     = 120
	ExpirationHour = 24
	ADMIN          = 1
	MANAGER        = 2
	SUPERVISOR     = 3
	OPERATOR       = 4
	AGENT          = 5
	STAFF          = 6
	USER           = 99
)

type IDRequest struct {
	ID string `json:"id" bson:"_id"`
}

type CustomClaims struct {
	UserID string `json:"userid"`
	TID    string `json:"tid"`
	Role   int    `json:"role"`
	jwt.StandardClaims
}

type JwtClaim struct {
	UserID string `json:"userid"`
	TID    string `json:"tid"`
	Role   int    `json:"role"`
}

type SearchRequest struct {
	Limit  int
	Page   int
	Sortby string
	Order  string
	ID     string
	TID    string
	Role   int
	Name   string
	Code   string
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
