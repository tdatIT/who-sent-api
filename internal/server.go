//go:build wireinject
// +build wireinject

// /go:build wireinject
package server

import (
	"context"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/tdatIT/who-sent-api/config"
	"github.com/tdatIT/who-sent-api/internal/biz"
	"github.com/tdatIT/who-sent-api/internal/handle"
	"github.com/tdatIT/who-sent-api/internal/infrastructure/adapter"
	"github.com/tdatIT/who-sent-api/internal/infrastructure/repository"
	"github.com/tdatIT/who-sent-api/internal/middleware"
	"github.com/tdatIT/who-sent-api/internal/router"
	"github.com/tdatIT/who-sent-api/pkgs/database"
	"github.com/tdatIT/who-sent-api/pkgs/database/cacheDB"
	"github.com/tdatIT/who-sent-api/pkgs/database/ormDB"
	"github.com/tdatIT/who-sent-api/pkgs/logger"
	"time"
)

type Server struct {
	cfg      *config.AppConfig
	rest     *echo.Echo
	_cacheDB cacheDB.CacheEngine
	_ormDB   ormDB.Gorm
}

func New() (*Server, error) {
	panic(wire.Build(wire.NewSet(
		NewServer,
		config.Set,
		database.Set,
		repository.Set,
		adapter.Set,
		biz.Set,
		handle.Set,
		middleware.Set,
		router.Set,
	)))
}

func NewServer(
	cfg *config.AppConfig,
	userRouter router.UserRouter,
	authRouter router.AuthRouter,
	_ormDB ormDB.Gorm,
	_cacheDB cacheDB.CacheEngine,
) *Server {
	//init default log
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

	//init router here
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
	// Shutdown orm db
	if err := serv._ormDB.Close(); err != nil {
		logger.Errorf("Failed to close orm db cause:%v", err)
	}

	// Shutdown cache db
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
