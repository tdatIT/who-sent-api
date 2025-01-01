package userServ

import (
	"context"
	"github.com/tdatIT/who-sent-api/internal/domain/dto"
)

type UserService interface {
	//Queries
	GetUserByID(ctx context.Context, req *dto.GetUserByIdReq) (*dto.GetUserByIdResp, error)

	//Commands
	RegisterUser(ctx context.Context, req *dto.UserRegisterReq) (*dto.UserRegisterResp, error)
	LoginByUsernameAndPassword(ctx context.Context, req *dto.LoginByUserPasswordReq) (*dto.LoginByUserPasswordResp, error)
}
