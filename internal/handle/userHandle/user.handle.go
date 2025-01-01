package userHandle

import "github.com/labstack/echo/v4"

type UserHandle interface {
	GetUserById(ctx echo.Context) error
}
