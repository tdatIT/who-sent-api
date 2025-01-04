package router

import (
	"github.com/labstack/echo/v4"
	"github.com/tdatIT/who-sent-api/internal/handle/userHandle"
)

type UserRouter struct {
	userHandle userHandle.UserHandle
}

func NewUserRoute(userHandle userHandle.UserHandle) UserRouter {
	return UserRouter{userHandle: userHandle}
}

func (ur UserRouter) Init(root *echo.Group) {
	userGroup := root.Group("/users")
	userGroup.GET("/:id", ur.userHandle.GetUserById)
}
