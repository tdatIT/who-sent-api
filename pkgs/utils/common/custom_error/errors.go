package servErr

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/tdatIT/who-sent-api/pkgs/utils/common/enum"
	responses "github.com/tdatIT/who-sent-api/pkgs/utils/common/response"
	"net/http"
	"os"
)

type ServError struct {
	Status               int
	InternalErrorMessage string
	Code                 string `json:"code"`
	Message              string `json:"message"`
}

func (e *ServError) Error() string {
	if e.InternalErrorMessage != "" {
		return e.InternalErrorMessage
	}
	return e.Message
}

func Is(err, target error) bool {
	return errors.Is(err, target)
}

func New(msg string) error {
	return errors.New(msg)
}

func Unwrap(err error) error {
	return errors.Unwrap(err)
}

func As(err error, target any) bool {
	return errors.As(err, &target)
}

func (e *ServError) WithInternalError(inErr error) *ServError {
	e.InternalErrorMessage = inErr.Error()
	return e
}

func CustomErrorHandler(err error, ctx echo.Context) {
	msg := responses.DefaultError

	var (
		customErr *ServError
		echoErr   *echo.HTTPError
	)

	switch {
	case errors.As(err, &echoErr):
		msg.Status = echoErr.Code
		msg.Code = fmt.Sprintf("%d", echoErr.Code)
		msg.Message = fmt.Sprintf("%v", echoErr.Message)
	case errors.As(err, &customErr):
		msg.Status = customErr.Status
		msg.Code = customErr.Code
		msg.Message = customErr.Message
	case errors.As(err, &validator.ValidationErrors{}):
		msg.Status = http.StatusBadRequest
		msg.Code = enum.InvalidArgument
		if os.Getenv("SERVER_MODE") == "prod" {
			msg.Message = "Invalid parameter"
		} else {
			msg.Message = err.Error()
		}

	default:
		msg.Status = http.StatusInternalServerError
		msg.Code = enum.Internal
		msg.Message = "Internal server error"
	}

	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	err = ctx.JSON(msg.Status, msg)
	if err != nil {
		ctx.Logger().Error(err)
	}
}
