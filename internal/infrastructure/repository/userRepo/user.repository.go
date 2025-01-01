package userRepo

import (
	"context"
	"github.com/tdatIT/who-sent-api/internal/domain/models"
	"github.com/tdatIT/who-sent-api/pkgs/utils/pagable"
)

type UserRepository interface {
	FindAndCountUsers(ctx context.Context, query *pagable.Query) ([]*models.User, int64, error)
	FindByIdWithRelations(ctx context.Context, id int, relations ...string) (*models.User, error)
	FindByEmail(ctx context.Context, email string) (*models.User, error)
	Create(ctx context.Context, user *models.User) (*models.User, error)
	UpdateByMap(ctx context.Context, id int, data map[string]interface{}) error
	Delete(ctx context.Context, id int) error
}
