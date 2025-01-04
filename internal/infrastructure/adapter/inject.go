package adapter

import (
	"github.com/google/wire"
	"github.com/tdatIT/who-sent-api/internal/infrastructure/adapter/auth"
)

var Set = wire.NewSet(
	auth.NewAuthJwtProvider,
)
