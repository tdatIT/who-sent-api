package authHandle

import (
	"github.com/labstack/echo/v4"
	"github.com/tdatIT/who-sent-api/internal/biz/authServ"
	"github.com/tdatIT/who-sent-api/internal/domain/dto"
	"github.com/tdatIT/who-sent-api/pkgs/logger"
	responses "github.com/tdatIT/who-sent-api/pkgs/utils/common/response"
	errors "github.com/tdatIT/who-sent-api/pkgs/utils/common/servErr"
)

type authHandleImpl struct {
	authServ authServ.AuthService
}

func NewAuthHandle(authServ authServ.AuthService) AuthHandle {
	return &authHandleImpl{authServ: authServ}
}

func (a authHandleImpl) LoginByUsernameAndPassword(ctx echo.Context) error {
	req := new(dto.LoginByUserPasswordReq)

	if err := ctx.Bind(req); err != nil {
		logger.Errorf("error while binding request: %v", err)
		return errors.ErrBadRequest
	}

	if err := ctx.Validate(req); err != nil {
		logger.Errorf("error while validating request: %v", err)
		return err
	}

	userData, err := a.authServ.LoginByUsernameAndPassword(ctx.Request().Context(), req)
	if err != nil {
		logger.Errorf("error while logging in: %v", err)
		return err
	}

	response := responses.DefaultSuccess
	response.Data = userData

	return response.JSON(ctx)
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

	userData, err := a.authServ.RegisterNewUserByEmail(ctx.Request().Context(), req)
	if err != nil {
		logger.Errorf("error while registering user: %v", err)
		return err
	}

	response := responses.DefaultSuccess
	response.Data = userData

	return response.JSON(ctx)
}
