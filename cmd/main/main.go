package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	access "gitlab.com/komentr/access/gokit/access"
	userLogSvc "gitlab.com/komentr/access/gokit/access"
	auth "gitlab.com/komentr/access/gokit/auth"
	authLogSvc "gitlab.com/komentr/access/gokit/auth"

	"gitlab.com/komentr/access/repo"
	"gitlab.com/komentr/access/service"
	"gitlab.com/komentr/access/utils/middleware"
	"gitlab.com/komentr/access/utils/utilmongo"
)

//HTTPORT const
const (
	HTTPPORT = "HTTPPORT"
)

func main() {
	ctx := context.Background()
	httpPort := os.Getenv(HTTPPORT)
	if httpPort == "" {
		httpPort = ":5556"
	}

	// log
	logger := log.NewJSONLogger(os.Stderr)
	logger = level.Error(logger)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC, "caller", log.DefaultCaller)

	// mongoDB
	dbMongo, err := utilmongo.MongoDBLogin()
	if err != nil {
		logger.Log("err", err)
	}
	userRepo := repo.NewUserMgoRepo(dbMongo)

	//RedisDB
	// redisClient, err := utilsredis.RedisDBLogin(ctx)
	// if err != nil {
	// 	logger.Log("err", err)
	// }
	// userRepo := repo.NewUserRedis(redisClient)

	// service
	var auths service.AuthService
	var users service.UserService

	auths = service.NewAuthService(userRepo)
	auths = authLogSvc.NewLoggingService(log.With(logger, "component", "accesss"), auths)
	users = service.NewUserService(userRepo)
	users = userLogSvc.NewLoggingService(log.With(logger, "component", "accesss"), users)

	mux := http.NewServeMux()
	mux.Handle("/api/auth/", auth.MakeHandler(ctx, auths, logger))
	mux.Handle("/api/v1/", middleware.Auth(access.MakeHandler(ctx, users, logger)))

	http.Handle("/", middleware.CORS(mux))
	http.Handle("/metrics", promhttp.Handler())

	errs := make(chan error, 2)
	go func() {
		logger.Log(
			"service_name", "access",
			"transport", "http",
			"address", httpPort,
			"msg", "listening")
		errs <- http.ListenAndServe(httpPort, nil)
	}()
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	logger.Log("terminated", <-errs)

}
