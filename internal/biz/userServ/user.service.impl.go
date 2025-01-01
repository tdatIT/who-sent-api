package userServ

import (
	"context"
	"github.com/tdatIT/who-sent-api/config"
	"github.com/tdatIT/who-sent-api/internal/domain/dto"
	"github.com/tdatIT/who-sent-api/internal/domain/models"
	"github.com/tdatIT/who-sent-api/internal/infrastructure/repository/userRepo"
	"github.com/tdatIT/who-sent-api/pkgs/logger"
	"github.com/tdatIT/who-sent-api/pkgs/utils/hpwd"
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

func (u userServImpl) RegisterUser(ctx context.Context, req *dto.UserRegisterReq) (*dto.UserRegisterResp, error) {
	hashPass, err := hpwd.HashPassword(req.Password)
	if err != nil {
		logger.Errorf("[RegisterUser] Error while hashing password: %v", err)
		return nil, err
	}

	entity := &models.User{
		Firstname:   req.Firstname,
		IsActivated: models.UserActiveStatus,
		Email:       req.Email,
		Password:    hashPass,
	}

	savedUser, err := u.userRepos.Create(ctx, entity)
	if err != nil {
		logger.Errorf("[RegisterUser] Error while creating user: %v", err)
		return nil, err
	}

	resp := new(dto.UserRegisterResp)
	if err := mapper.BindingStruct(savedUser, resp); err != nil {
		logger.Errorf("[RegisterUser] Error while binding struct: %v", err)
		return nil, err
	}

	return resp, nil
}

func (u userServImpl) LoginByUsernameAndPassword(ctx context.Context, req *dto.LoginByUserPasswordReq) (*dto.LoginByUserPasswordResp, error) {
	//TODO implement me
	panic("implement me")
}
