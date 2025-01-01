package ormDB

import (
	"context"
	"database/sql"
	"github.com/tdatIT/who-sent-api/config"
	"github.com/tdatIT/who-sent-api/pkgs/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

// Gorm defines interface for access the database.
type Gorm interface {
	DB() *gorm.DB
	SqlDB() *sql.DB
	ExecWithContext(fc func(tx *gorm.DB) error, ctx context.Context) (err error)
	Transaction(fc func(tx *gorm.DB) error) (err error)
	Close() error
}

// Config GORM Config
type Config struct {
	Debug           bool
	DBType          string
	DSN             string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
	ConnMaxIdleTime time.Duration
	TablePrefix     string
	CacheMode       bool
	LogFormat       string
	AutoMigrateMode bool
	RedisConfig     *config.Redis
}

// _gorm orm struct
type _gorm struct {
	db    *gorm.DB
	sqlDB *sql.DB
}

// New Create orm.DB and  instance
func New(c Config) (Gorm, error) {
	dial := postgres.Open(c.DSN)
	gConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   c.TablePrefix,
			SingularTable: true,
		},
	}

	db, err := gorm.Open(dial, gConfig)
	if err != nil {
		return nil, err
	}

	if c.Debug {
		db = db.Debug()
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	if c.MaxOpenConns != 0 {
		sqlDB.SetMaxOpenConns(c.MaxOpenConns)
	}
	if c.ConnMaxLifetime != 0 {
		sqlDB.SetConnMaxLifetime(c.ConnMaxLifetime)
	}
	if c.MaxIdleConns != 0 {
		sqlDB.SetMaxIdleConns(c.MaxIdleConns)
	}
	if c.ConnMaxIdleTime != 0 {
		sqlDB.SetConnMaxIdleTime(c.ConnMaxIdleTime)
	}

	err = db.AutoMigrate(
	// Add your models here
	)
	if err != nil {
		logger.Errorf("AutoMigrate error: %v", err)
		return nil, err
	}

	return &_gorm{
		db:    db,
		sqlDB: sqlDB,
	}, nil
}
