package userServ

import (
	"context"
	"github.com/tdatIT/who-sent-api/internal/domain/dto"
)

type UserService interface {
	//Queries
	GetUserByID(ctx context.Context, req *dto.GetUserByIdReq) (*dto.GetUserByIdResp, error)
}
