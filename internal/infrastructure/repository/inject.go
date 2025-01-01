package repository

import (
	"github.com/google/wire"
	"github.com/tdatIT/who-sent-api/internal/infrastructure/repository/userRepo"
)

var Set = wire.NewSet(
	userRepo.NewUserRepositoryImpl,
)
