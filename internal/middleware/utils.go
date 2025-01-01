package middleware

import (
	"github.com/labstack/echo/v4"
	errors "github.com/tdatIT/who-sent-api/pkgs/utils/common/servErr"
)

func GetUserIdFromContext(ctx echo.Context) (int, error) {
	// Retrieve user id from context
	userId := ctx.Get(AuthUserId)
	if userId == nil {
		return 0, errors.ErrUnauthenticated
	}

	// Convert user id to int
	intUserId, ok := userId.(int64)
	if !ok {
		return 0, errors.ErrUnauthenticated
	}

	return int(intUserId), nil
}
