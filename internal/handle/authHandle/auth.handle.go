package authHandle

import "github.com/labstack/echo/v4"

type AuthHandle interface {
	RegisterByEmail(ctx echo.Context) error
}
