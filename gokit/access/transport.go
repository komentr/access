package access

import (
	"context"

	"net/http"
	"gitlab.com/komentr/access/service"

	kitlog "github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/julienschmidt/httprouter"
)

//MakeHandler declaration
func MakeHandler(ctx context.Context, us service.UserService, logger kitlog.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(logger),
		kithttp.ServerErrorEncoder(encodeError),
	}

	NewUserHandler := kithttp.NewServer(
		NewUserEndpoint(us),
		decodeNewUserRequest,
		encodeResponse,
		opts...,
	)

	UpdateUserHandler := kithttp.NewServer(
		UpdateUserEndpoint(us),
		decodeUpdateUserRequest,
		encodeResponse,
		opts...,
	)

	SearchUserHandler := kithttp.NewServer(
		SearchUserEndpoint(us),
		decodeSearchRequest,
		encodeResponse,
		opts...,
	)

	GetUserHandler := kithttp.NewServer(
		GetUserEndpoint(us),
		decodeIDRequest,
		encodeResponse,
		opts...,
	)

	DeleteUserHandler := kithttp.NewServer(
		DeleteUserEndpoint(us),
		decodeIDRequest,
		encodeResponse,
		opts...,
	)

	r := httprouter.New()

	r.Handle("POST", "/api/v1/user/new", wrapHandler(NewUserHandler))
	r.Handle("PUT", "/api/v1/user/update", wrapHandler(UpdateUserHandler))
	r.Handle("GET", "/api/v1/user/search", wrapHandler(SearchUserHandler))
	r.Handle("GET", "/api/v1/user/get/:id", wrapHandler(GetUserHandler))
	r.Handle("DELETE", "/api/v1/user/delete/:id", wrapHandler(DeleteUserHandler))

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
