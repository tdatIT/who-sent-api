package userServ

import (
	"context"
	"github.com/tdatIT/who-sent-api/config"
	"github.com/tdatIT/who-sent-api/internal/domain/dto"
	"github.com/tdatIT/who-sent-api/internal/infrastructure/repository/userRepo"
	"github.com/tdatIT/who-sent-api/pkgs/logger"
	"github.com/tdatIT/who-sent-api/pkgs/utils/mapper"
)

type userServImpl struct {
	cfg       *config.AppConfig
	userRepos userRepo.UserRepository
}

func NewUserService(cfg *config.AppConfig, userRepos userRepo.UserRepository) UserService {
	return &userServImpl{cfg: cfg, userRepos: userRepos}
}

func (u userServImpl) GetUserByID(ctx context.Context, req *dto.GetUserByIdReq) (*dto.GetUserByIdResp, error) {
	entity, err := u.userRepos.FindByIdWithRelations(ctx, req.ID, "Roles")
	if err != nil {
		logger.Errorf("[GetUserByID] Error while finding user by id: %v", err)
		return nil, err
	}

	resp := new(dto.GetUserByIdResp)
	if err := mapper.BindingStruct(entity, resp); err != nil {
		logger.Errorf("[GetUserByID] Error while binding struct: %v", err)
		return nil, err
	}

	return resp, nil
}
