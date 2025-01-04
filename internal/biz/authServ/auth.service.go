package authServ

import (
	"context"
	"github.com/tdatIT/who-sent-api/internal/domain/dto"
)

type AuthService interface {
	RegisterNewUserByEmail(ctx context.Context, req *dto.UserRegisterReq) (*dto.UserRegisterResp, error)
	LoginByUsernameAndPassword(ctx context.Context, req *dto.LoginByUserPasswordReq) (*dto.LoginByUserPasswordResp, error)
}
