package auth

import (
	"context"

	"gitlab.com/komentr/access/service"

	"net/http"

	kitlog "github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/julienschmidt/httprouter"
)

//MakeHandler of auth transport
func MakeHandler(ctx context.Context, us service.AuthService, logger kitlog.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(logger),
		kithttp.ServerErrorEncoder(encodeError),
	}

	RegisterUserHandler := kithttp.NewServer(
		RegisterUserEndpoint(us),
		decodeNewUserRequest,
		encodeResponse,
		opts...,
	)

	LoginHandler := kithttp.NewServer(
		LoginEndpoint(us),
		decodeLoginRequest,
		encodeResponse,
		opts...,
	)

	RegMailHandler := kithttp.NewServer(
		RegMailEndpoint(us),
		decodeNewUserRequest,
		encodeResponse,
		opts...,
	)

	r := httprouter.New()

	r.Handle("POST", "/api/auth/register", wrapHandler(RegisterUserHandler))
	r.Handle("POST", "/api/auth/regmail", wrapHandler(RegMailHandler))
	r.Handle("POST", "/api/auth/login", wrapHandler(LoginHandler))

	return r
}

func wrapHandler(h http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// Take the context out from the request
		ctx := r.Context()
		ctx = context.WithValue(ctx, "params", ps)
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	}
}
