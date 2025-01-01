package healthcheck

import (
	"fmt"
	"github.com/hellofresh/health-go/v5"
	healthPostgres "github.com/hellofresh/health-go/v5/checks/postgres"
	healthRedis "github.com/hellofresh/health-go/v5/checks/redis"
	"github.com/tdatIT/who-sent-api/config"

	"time"
)

func NewHealthCheckService(config *config.AppConfig) (*health.Health, error) {
	h, err := health.New(health.WithComponent(health.Component{
		Name:    "funong-service",
		Version: "v1.0.0",
	}))
	if err != nil {
		return nil, err
	}

	//db check
	err = h.Register(health.Config{
		Name:      "postgres",
		Timeout:   time.Second * 2,
		SkipOnErr: false,
		Check: healthPostgres.New(healthPostgres.Config{
			DSN: fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
				config.DB.Postgres.Host,
				config.DB.Postgres.Port,
				config.DB.Postgres.UserName,
				config.DB.Postgres.Password,
				config.DB.Postgres.Database,
			),
		}),
	})
	if err != nil {
		return nil, err
	}

	//redis check
	err = h.Register(health.Config{
		Name:      "redis",
		Timeout:   time.Second * 2,
		SkipOnErr: false,
		Check: healthRedis.New(healthRedis.Config{
			DSN: fmt.Sprintf("redis://%s", config.Cache.Redis.Address[0]),
		}),
	})
	if err != nil {
		return nil, err
	}

	return h, nil
}
