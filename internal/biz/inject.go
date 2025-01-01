package biz

import (
	"github.com/google/wire"
	"github.com/tdatIT/who-sent-api/internal/biz/userServ"
)

var Set = wire.NewSet(userServ.NewUserService)
