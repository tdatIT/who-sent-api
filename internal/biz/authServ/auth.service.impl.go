package authServ

import (
	"context"
	"errors"
	"github.com/tdatIT/who-sent-api/config"
	"github.com/tdatIT/who-sent-api/internal/domain/dto"
	"github.com/tdatIT/who-sent-api/internal/domain/models"
	"github.com/tdatIT/who-sent-api/internal/infrastructure/adapter/auth"
	"github.com/tdatIT/who-sent-api/internal/infrastructure/repository/userRepo"
	"github.com/tdatIT/who-sent-api/pkgs/logger"
	"github.com/tdatIT/who-sent-api/pkgs/utils/common/servErr"
	"github.com/tdatIT/who-sent-api/pkgs/utils/hpwd"
	"github.com/tdatIT/who-sent-api/pkgs/utils/mapper"
	"gorm.io/gorm"
)

type authService struct {
	cfg          *config.AppConfig
	authProvider *auth.AuthJwtProvider
	userRepos    userRepo.UserRepository
}

func NewAuthService(cfg *config.AppConfig, authProvider *auth.AuthJwtProvider, userRepos userRepo.UserRepository) AuthService {
	return &authService{cfg: cfg, authProvider: authProvider, userRepos: userRepos}
}

func (a authService) RegisterNewUserByEmail(ctx context.Context, req *dto.UserRegisterReq) (*dto.UserRegisterResp, error) {
	user, err := a.userRepos.FindByEmail(ctx, req.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logger.Errorf("[RegisterNewUserByEmail] Error while finding user by email: %v", err)
		return nil, err
	}

	if user != nil {
		return nil, servErr.ErrEmailAlreadyExists
	}
	hashPwd, _ := hpwd.HashPassword(req.Password)
	newUser := &models.User{
		Firstname:   req.Firstname,
		IsActivated: true,
		Email:       req.Email,
		Password:    hashPwd,
		Roles:       []models.Role{{ID: a.cfg.OtherKM.DefaultRoleID}},
	}

	savedUser, err := a.userRepos.Create(ctx, newUser)
	if err != nil {
		logger.Errorf("[RegisterNewUserByEmail] Error while creating user: %v", err)
		return nil, err
	}

	userDTO := dto.UserDTO{}
	if err := mapper.BindingStruct(savedUser, &userDTO); err != nil {
		logger.Errorf("[RegisterNewUserByEmail] Error while binding struct: %v", err)
		return nil, err
	}

	return &dto.UserRegisterResp{UserDTO: userDTO}, nil
}

func (a authService) LoginByUsernameAndPassword(ctx context.Context, req *dto.LoginByUserPasswordReq) (*dto.LoginByUserPasswordResp, error) {
	isValid, user, err := a.validateUserLoginWorkflow(ctx, req.Username, req.Password)
	if err != nil {
		logger.Errorf("[LoginByUsernameAndPassword] Error while validating user login workflow: %v", err)
		return nil, err
	}

	if !isValid {
		return nil, err
	}

	accessToken, err := a.authProvider.GenerateAccessToken(user.GetUserIdStr(), user.Email, user.GetRolesStringSlice())
	if err != nil {
		logger.Errorf("[LoginByUsernameAndPassword] Error while generating access token: %v", err)
		return nil, err
	}

	refreshToken, err := a.authProvider.GenerateRefreshToken(user.GetUserIdStr())
	if err != nil {
		logger.Errorf("[LoginByUsernameAndPassword] Error while generating refresh token: %v", err)
		return nil, err
	}

	userDTO := dto.UserDTO{}
	if err := mapper.BindingStruct(user, &userDTO); err != nil {
		logger.Errorf("[LoginByUsernameAndPassword] Error while binding struct: %v", err)
		return nil, err
	}
	return &dto.LoginByUserPasswordResp{
		User: userDTO,
		GetAccessTokenData: dto.GetAccessTokenData{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
			ExpiresIn:    int(a.cfg.Adapter.Auth.AccessExp.Seconds()),
		},
	}, nil
}

func (a authService) validateUserLoginWorkflow(ctx context.Context, email, pwd string) (bool, *models.User, error) {
	user, err := a.userRepos.FindByEmail(ctx, email)
	if err != nil {
		return false, nil, err
	}

	if !user.IsActivated {
		return false, user, servErr.ErrUserIsNotActivated
	}

	if !hpwd.CheckPasswordHash(pwd, user.Password) {
		return false, user, servErr.ErrInvalidCredentials
	}

	return true, user, nil
}
