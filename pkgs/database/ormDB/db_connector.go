package ormDB

import (
	"fmt"
	"github.com/tdatIT/who-sent-api/config"
	"github.com/tdatIT/who-sent-api/pkgs/logger"
)

func NewDBConnection(config *config.AppConfig) Gorm {
	dataSourceName := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.DB.Postgres.Host,
		config.DB.Postgres.Port,
		config.DB.Postgres.UserName,
		config.DB.Postgres.Password,
		config.DB.Postgres.Database,
	)
	cfg := Config{
		AutoMigrateMode: config.DB.AutoMigrate,
		DSN:             dataSourceName,
		MaxOpenConns:    config.DB.Postgres.MaxOpenConn,
		MaxIdleConns:    config.DB.Postgres.MaxIdleConn,
		ConnMaxLifetime: config.DB.Postgres.ConnMaxLifetime,
		ConnMaxIdleTime: config.DB.Postgres.ConnMaxIdleTime,
		Debug:           true,
		CacheMode:       config.Cache.QueryCache,
		RedisConfig:     &config.Cache.Redis,
		LogFormat:       config.LogConfig.Encoding,
	}
	conn, err := New(cfg)
	if err != nil {
		logger.Errorf("Error while creating database connection: %v", err)
		panic(err)
	}

	logger.Info("Gorm has created database connection - POSTGRES")
	return conn
}
