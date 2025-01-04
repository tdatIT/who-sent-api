package userRepo

import (
	"context"
	"github.com/tdatIT/who-sent-api/internal/domain/models"
	"github.com/tdatIT/who-sent-api/pkgs/database/ormDB"
	"github.com/tdatIT/who-sent-api/pkgs/logger"
	"github.com/tdatIT/who-sent-api/pkgs/utils/pagable"

	"gorm.io/gorm"
)

func NewUserRepositoryImpl(dbClient ormDB.Gorm) UserRepository {
	return &userRepositoryImpl{dbClient: dbClient}
}

type userRepositoryImpl struct {
	dbClient ormDB.Gorm
}

func (u userRepositoryImpl) FindAndCountUsers(ctx context.Context, query *pagable.Query) ([]*models.User, int64, error) {
	var users []*models.User
	var total int64

	err := u.dbClient.ExecWithContext(func(db *gorm.DB) error {
		return db.Model(&models.User{}).
			Order("created_at desc").
			Count(&total).
			Limit(query.GetLimit()).
			Offset(query.GetOffset()).
			Find(&users).Error
	}, ctx)
	if err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

func (u userRepositoryImpl) FindByIdWithRelations(ctx context.Context, id int, relations ...string) (*models.User, error) {
	var entity *models.User
	result := u.dbClient.ExecWithContext(func(tx *gorm.DB) error {
		query := tx.Model(&models.User{})
		// Loop over each relation and preload it
		for _, relation := range relations {
			query = query.Preload(relation)
		}
		return query.Where("id = ?", id).First(entity).Error
	}, ctx)

	if result != nil {
		logger.Errorf("Error while finding identity by id with relations: %v", result)
		return nil, result
	}

	return entity, nil
}

func (u userRepositoryImpl) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	var entity models.User
	result := u.dbClient.ExecWithContext(func(tx *gorm.DB) error {
		return tx.Model(&models.User{}).First(&entity, "email = ? and deleted_at is null", email).Error
	}, ctx)

	if result != nil {
		logger.Errorf("Error while finding identity by email: %v", result)
		return nil, result
	}

	return &entity, nil
}

func (u userRepositoryImpl) Create(ctx context.Context, user *models.User) (*models.User, error) {
	result := u.dbClient.Transaction(func(tx *gorm.DB) error {
		return tx.Create(user).Error
	})
	if result != nil {
		logger.Errorf("Error while creating identity: %v", result)
		return nil, result
	}

	return user, nil
}

func (u userRepositoryImpl) UpdateByMap(ctx context.Context, id int, data map[string]interface{}) error {
	result := u.dbClient.Transaction(func(tx *gorm.DB) error {
		return tx.Model(&models.User{}).Where("id = ?", id).Updates(data).Error
	})
	if result != nil {
		logger.Errorf("Error while updating identity by map: %v", result)
		return result
	}

	return nil
}

func (u userRepositoryImpl) Delete(ctx context.Context, id int) error {
	result := u.dbClient.Transaction(func(tx *gorm.DB) error {
		return tx.Delete(&models.User{}, id).Error
	})
	if result != nil {
		logger.Errorf("Error while deleting identity: %v", result)
		return result
	}

	return nil
}
