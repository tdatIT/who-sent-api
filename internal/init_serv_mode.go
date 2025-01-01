package server

import (
	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/tdatIT/who-sent-api/config"
	healthcheck "github.com/tdatIT/who-sent-api/pkgs/health"
	errors "github.com/tdatIT/who-sent-api/pkgs/utils/common/custom_error"
	"github.com/tdatIT/who-sent-api/pkgs/utils/valid"
)

func InitRestServ(cfg *config.AppConfig) *echo.Echo {
	echoApp := echo.New()
	echoApp.Use(echoMiddleware.TimeoutWithConfig(echoMiddleware.TimeoutConfig{
		Timeout:      cfg.Server.RequestTimeout,
		ErrorMessage: "Request Timeout",
	}))
	echoApp.Server.ReadTimeout = cfg.Server.ReadTimeout
	echoApp.Server.WriteTimeout = cfg.Server.WriteTimeout
	echoApp.Validator = &valid.Validator{Valid: valid.InitValidatorInstance()}
	echoApp.HTTPErrorHandler = errors.CustomErrorHandler
	echoApp.Use(echoMiddleware.Logger())

	echoApp.Use(echoMiddleware.RecoverWithConfig(echoMiddleware.RecoverConfig{
		StackSize: 1 << 10, // 1 KB
		LogLevel:  log.ERROR,
	}))

	metricsMiddleware := echoprometheus.NewMiddleware("echo")
	echoApp.Use(metricsMiddleware)
	echoApp.GET("/metrics", echoprometheus.NewHandler())

	h, _ := healthcheck.NewHealthCheckService(cfg)
	echoApp.GET("/readiness", echo.WrapHandler(h.Handler()))
	echoApp.GET("/liveness", func(ctx echo.Context) error {
		return ctx.String(200, "OK")
	})

	return echoApp
}
