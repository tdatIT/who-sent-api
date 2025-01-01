// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package server

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/tdatIT/who-sent-api/config"
	"github.com/tdatIT/who-sent-api/internal/biz/userServ"
	"github.com/tdatIT/who-sent-api/internal/handle/authHandle"
	"github.com/tdatIT/who-sent-api/internal/handle/userHandle"
	"github.com/tdatIT/who-sent-api/internal/infrastructure/repository/userRepo"
	"github.com/tdatIT/who-sent-api/internal/router"
	"github.com/tdatIT/who-sent-api/pkgs/database/cacheDB"
	"github.com/tdatIT/who-sent-api/pkgs/database/ormDB"
	"github.com/tdatIT/who-sent-api/pkgs/logger"
	"time"
)

// Injectors from server.go:

func New() (*Server, error) {
	appConfig, err := config.NewConfig()
	if err != nil {
		return nil, err
	}
	gorm := ormDB.NewDBConnection(appConfig)
	userRepository := userRepo.NewUserRepositoryImpl(gorm)
	userService := userServ.NewUserService(appConfig, userRepository)
	userHandleUserHandle := userHandle.NewUserHandle(userService)
	userRouter := router.NewUserRoute(userHandleUserHandle)
	authHandleAuthHandle := authHandle.NewAuthHandle(userService)
	authRouter := router.NewAuthRoute(authHandleAuthHandle)
	cacheEngine, err := cacheDB.NewCacheEngine(appConfig)
	if err != nil {
		return nil, err
	}
	server := NewServer(appConfig, userRouter, authRouter, gorm, cacheEngine)
	return server, nil
}

// server.go:

type Server struct {
	cfg      *config.AppConfig
	rest     *echo.Echo
	_cacheDB cacheDB.CacheEngine
	_ormDB   ormDB.Gorm
}

func NewServer(
	cfg *config.AppConfig,
	userRouter router.UserRouter,
	authRouter router.AuthRouter,
	_ormDB ormDB.Gorm,
	_cacheDB cacheDB.CacheEngine,
) *Server {

	appLog := logger.NewLogger(&logger.LogConfig{
		ServiceName: cfg.Server.Name,
		Level:       cfg.LogConfig.Level,
		LogFormat:   cfg.LogConfig.Encoding,
		TimeFormat:  logger.ISO8601TimeEncoder,
	})
	logger.SetLogger(appLog)

	serv := &Server{
		cfg:      cfg,
		_ormDB:   _ormDB,
		_cacheDB: _cacheDB,
	}
	serv.rest = InitRestServ(cfg)

	v1 := serv.rest.Group("/v1")
	userRouter.Init(v1)
	authRouter.Init(v1)

	return serv
}

func (serv Server) REST() *echo.Echo {
	return serv.rest
}

func (serv Server) Config() *config.AppConfig {
	return serv.cfg
}

func (serv Server) Shutdown() {

	if err := serv._ormDB.Close(); err != nil {
		logger.Errorf("Failed to close orm db cause:%v", err)
	}

	if err := serv._cacheDB.Close(); err != nil {
		logger.Errorf("Failed to close cache db cause:%v", err)
	}

	if serv.rest != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := serv.rest.Shutdown(ctx); err != nil {
			logger.Fatalf("Failed to shutdown echo app cause:%v", err)
		}
	}

}