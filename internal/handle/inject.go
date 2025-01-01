package handle

import (
	"github.com/google/wire"
	"github.com/tdatIT/who-sent-api/internal/handle/authHandle"
	"github.com/tdatIT/who-sent-api/internal/handle/userHandle"
)

var Set = wire.NewSet(
	authHandle.NewAuthHandle,
	userHandle.NewUserHandle,
)
