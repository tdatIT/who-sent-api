package ormDB

import (
	"context"
	"database/sql"
	"github.com/tdatIT/who-sent-api/pkgs/logger"

	"gorm.io/gorm"
)

func (g *_gorm) SqlDB() *sql.DB {
	return g.sqlDB
}

func (g *_gorm) ExecWithContext(fc func(tx *gorm.DB) error, ctx context.Context) (err error) {
	tx := g.db.WithContext(ctx)
	err = fc(tx)
	return err
}

func (g *_gorm) DB() *gorm.DB {
	return g.db
}

func (g *_gorm) Close() error {
	logger.Info("Close database connection")
	return g.sqlDB.Close()
}

func (g *_gorm) Transaction(fc func(tx *gorm.DB) error) (err error) {
	panicked := true
	tx := g.db.Begin()
	defer func() {
		if panicked || err != nil {
			tx.Rollback()
		}
	}()

	err = fc(tx)

	if err == nil {
		err = tx.Commit().Error
	}

	panicked = false
	return
}
