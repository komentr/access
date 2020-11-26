package middleware

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"os"
	"time"

	"net/http"
	"strings"

	"gitlab.com/komentr/access/models"

	jwt "github.com/dgrijalva/jwt-go"
)

type KeyContext interface{}

//TokenSecret const
const (
	TokenSecret = "TOKENSECRET"
)

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "DELETE,PATCH,GET,POST,PUT,OPTIONS")
		w.Header().Add("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Content-Disposition, Content-Transfer-Encoding, Content-Description")
		w.Header().Add("Access-Control-Expose-Headers", "Link")
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		log.Println("url : path ", r.URL.Path, r.Method, time.Now())
		next.ServeHTTP(w, r)
	})
}

type customClaims struct {
	UserID string `json:"userid"`
	TID    string `json:"tid"`
	Role   int    `json:"role"`
	jwt.StandardClaims
}

type rsp struct {
	Error string `json:"error,omitempty"`
}

func Auth(next http.Handler) http.Handler {
	secret := os.Getenv(TokenSecret)
	if secret == "" {
		secret = "supersecret"
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var errresponse rsp
		tokAuth := r.Header.Get("Authorization")
		if len(tokAuth) == 0 {
			w.WriteHeader(http.StatusForbidden)
			errresponse.Error = "null token"
			json.NewEncoder(w).Encode(errresponse)
			return
		}
		authHeaderParts := strings.Split(tokAuth, " ")
		if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
			w.WriteHeader(http.StatusForbidden)
			errresponse.Error = "token not full"
			json.NewEncoder(w).Encode(errresponse)
			return
		}

		token, err := jwt.ParseWithClaims(
			authHeaderParts[1],
			&customClaims{},
			func(token *jwt.Token) (interface{}, error) {
				if token.Method != jwt.SigningMethodHS256 {
					return nil, errors.New("jwt: unexpected signing method")
				}
				return []byte(secret), nil
			})

		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			errresponse.Error = "err " + err.Error()
			json.NewEncoder(w).Encode(errresponse)
			return
		}

		if !token.Valid {
			w.WriteHeader(http.StatusForbidden)
			errresponse.Error = "JWT Token was invalid"
			json.NewEncoder(w).Encode(errresponse)
			return
		}
		c, ok := token.Claims.(*customClaims)
		if !ok {
			log.Println("fail parsing claim")
		}
		jwtclaim := models.JwtClaim{
			UserID: c.UserID,
			TID:    c.TID,
			Role:   c.Role,
		}

		k := KeyContext("jwtclaim")
		ctx := context.WithValue(r.Context(), k, jwtclaim)

		r = r.WithContext(ctx)
		// log.Println("userid : ", c.UserID)
		// log.Println("tid : ", c.TID)
		// log.Println("role : ", c.Role)

		next.ServeHTTP(w, r)
		return
	})
}
