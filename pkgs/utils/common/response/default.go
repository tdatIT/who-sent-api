package responses

import (
	"github.com/tdatIT/who-sent-api/pkgs/utils/common/enum"
)

var (
	DefaultSuccess = General{
		Status:  200,
		Code:    enum.Ok,
		Message: "success",
		Data:    nil,
	}

	DefaultError = General{
		Status:  500,
		Code:    enum.Internal,
		Message: "Internal server error",
		Data:    nil,
	}
)
