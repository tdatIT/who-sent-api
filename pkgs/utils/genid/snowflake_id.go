package genid

import (
	"github.com/bwmarrin/snowflake"
	"github.com/tdatIT/who-sent-api/pkgs/logger"

	"os"
	"strconv"
)

var _node *snowflake.Node

// GetSnowLakeIns singleton
func GetSnowLakeIns() *snowflake.Node {
	if _node == nil {
		nodeEnv := os.Getenv("NODE_ID")
		nodeNumber := 1
		if nodeEnv != "" {
			nodeNumber, _ = strconv.Atoi(nodeEnv)
		}

		_n, err := snowflake.NewNode(int64(nodeNumber))
		if err != nil {
			logger.Fatalf("GetSnowLakeIns create snowlake instance: %v", err)
		}
		_node = _n
	}

	return _node
}
