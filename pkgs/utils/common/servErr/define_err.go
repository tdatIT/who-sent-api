package servErr

import "github.com/tdatIT/who-sent-api/pkgs/utils/common/enum"

var (
	ErrInvalidCredentials = &ServError{
		Status:  401,
		Code:    enum.InvalidCredentials,
		Message: "invalid username or password",
	}

	ErrEmailAlreadyExists = &ServError{
		Status:  409,
		Code:    enum.ErrEmailAlreadyExists,
		Message: "email already exists",
	}

	ErrUserIsNotActivated = &ServError{
		Status:  401,
		Code:    enum.UserNotActivated,
		Message: "user is not activated",
	}
)
