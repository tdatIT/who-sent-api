package router

import (
	"github.com/labstack/echo/v4"
	"github.com/tdatIT/who-sent-api/internal/handle/authHandle"
)

type AuthRouter struct {
	authHandle authHandle.AuthHandle
}

func NewAuthRoute(authHandle authHandle.AuthHandle) AuthRouter {
	return AuthRouter{authHandle: authHandle}
}

func (ar AuthRouter) Init(root *echo.Group) {
	authGroup := root.Group("/auth")
	authGroup.POST("/login", ar.authHandle.LoginByUsernameAndPassword)
	authGroup.POST("/register/email", ar.authHandle.RegisterByEmail)
}
