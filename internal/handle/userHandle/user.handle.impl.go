package userHandle

import (
	"github.com/labstack/echo/v4"
	"github.com/tdatIT/who-sent-api/internal/biz/userServ"
	"github.com/tdatIT/who-sent-api/internal/domain/dto"
	"github.com/tdatIT/who-sent-api/pkgs/logger"
	responses "github.com/tdatIT/who-sent-api/pkgs/utils/common/response"
	errors "github.com/tdatIT/who-sent-api/pkgs/utils/common/servErr"
)

type userHandleImpl struct {
	userServ userServ.UserService
}

func NewUserHandle(userServ userServ.UserService) UserHandle {
	return &userHandleImpl{userServ: userServ}
}

func (a userHandleImpl) GetUserById(ctx echo.Context) error {
	req := new(dto.GetUserByIdReq)

	if err := ctx.Bind(req); err != nil {
		logger.Errorf("error while binding request: %v", err)
		return errors.ErrBadRequest
	}

	if err := ctx.Validate(req); err != nil {
		logger.Errorf("error while validating request: %v", err)
		return err
	}

	userData, err := a.userServ.GetUserByID(ctx.Request().Context(), req)
	if err != nil {
		logger.Errorf("error while registering user: %v", err)
		return err
	}

	response := responses.DefaultSuccess
	response.Data = userData

	return response.JSON(ctx)
}
