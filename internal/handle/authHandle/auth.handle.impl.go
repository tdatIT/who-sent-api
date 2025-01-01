package authHandle

import (
	"github.com/labstack/echo/v4"
	"github.com/tdatIT/who-sent-api/internal/biz/userServ"
	"github.com/tdatIT/who-sent-api/internal/domain/dto"
	"github.com/tdatIT/who-sent-api/pkgs/logger"
	errors "github.com/tdatIT/who-sent-api/pkgs/utils/common/custom_error"
	responses "github.com/tdatIT/who-sent-api/pkgs/utils/common/response"
)

type authHandleImpl struct {
	userServ userServ.UserService
}

func NewAuthHandle(userServ userServ.UserService) AuthHandle {
	return &authHandleImpl{userServ: userServ}
}

func (a authHandleImpl) RegisterByEmail(ctx echo.Context) error {
	req := new(dto.UserRegisterReq)

	if err := ctx.Bind(req); err != nil {
		logger.Errorf("error while binding request: %v", err)
		return errors.ErrBadRequest
	}

	if err := ctx.Validate(req); err != nil {
		logger.Errorf("error while validating request: %v", err)
		return err
	}

	userData, err := a.userServ.RegisterUser(ctx.Request().Context(), req)
	if err != nil {
		logger.Errorf("error while registering user: %v", err)
		return err
	}

	response := responses.DefaultSuccess
	response.Data = userData

	return response.JSON(ctx)
}
