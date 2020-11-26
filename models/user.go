package models

import "time"

const (
	UserStatusNotActive = 1
	UserStatusActive    = 2
	UserStatusSuspend   = 3
	UserStatusOnlyName  = 4
	UserColl            = "user"
)

type User struct {
	ID       string    `json:"id" bson:"_id"`
	Username string    `json:"username" bson:"username"`
	Password string    `json:"-" bson:"password"`
	Name     string    `json:"name" bson:"name"`
	Role     int       `json:"role" bson:"role"`
	Image    string    `json:"image" bson:"image"`
	Active   int       `json:"active" bson:"active"`
	Verified int       `json:"verified" bson:"verified"`
	RegBy    string    `json:"regby" bson:"regby"`
	TID      string    `json:"tid" bson:"tid"`
	Created  time.Time `json:"created" bson:"created"`
	Updated  time.Time `json:"updated" bson:"updated"`
}

type UserRequest struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	Name     string `json:"name" bson:"name"`
	Role     int    `json:"role" bson:"role"`
	Image    string `json:"image" bson:"image"`
	Active   int    `json:"active" bson:"active"`
	RegBy    string `json:"regby" bson:"regby"`
	TID      string `json:"tid" bson:"tid"`
}

type LoginResponse struct {
	ID       string `json:"id" bson:"_id"`
	Username string `json:"username" bson:"username"`
	Name     string `json:"name" bson:"name"`
	Role     int    `json:"role" bson:"role"`
	Image    string `json:"image" bson:"image"`
	Active   int    `json:"active" bson:"active"`
	TID      string `json:"tid" bson:"tid"`
	Token    string `json:"token" bson:"token"`
}

type UserSearchRequest struct {
	Limit    int
	Page     int
	Sortby   string `json:"sortby,omitempty"`
	Order    string `json:"order,omitempty"`
	TID      string `json:"tid,omitempty"`
	Username string `json:"username,omitempty"`
	Name     string `json:"name,omitempty"`
	Role     int    `json:"role,omitempty"`
	Active   int    `json:"active,omitempty"`
	Verified int    `json:"verified,omitempty"`
	RegBy    string `json:"regby,omitempty"`
}

type UserFindRequest struct {
	Username string
}

type RegMailResponse struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}
