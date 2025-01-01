package database

import (
	"github.com/google/wire"
	"github.com/tdatIT/who-sent-api/pkgs/database/cacheDB"
	"github.com/tdatIT/who-sent-api/pkgs/database/ormDB"
)

var Set = wire.NewSet(
	ormDB.NewDBConnection,
	cacheDB.NewCacheEngine,
)
