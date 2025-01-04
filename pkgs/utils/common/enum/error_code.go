package enum

// common error
const (
	Ok                 = "0"
	Cancelled          = "001"
	Unknown            = "002"
	InvalidArgument    = "003"
	DeadlineExceeded   = "004"
	NotFound           = "005"
	AlreadyExists      = "006"
	PermissionDenied   = "007"
	Unauthenticated    = "016"
	ResourceExhausted  = "008"
	FailedPrecondition = "009"
	OutOfRange         = "011"
	Unimplemented      = "012"
	Internal           = "013"
	Unavailable        = "014"
	DataLoss           = "015"
	NotFoundEntity     = "017"
)

// custom error

// Invalid error
const (
	InvalidCredentials    = "100"
	InvalidToken          = "102"
	InvalidRefreshToken   = "103"
	InvalidOldPassword    = "104"
	InvalidNewPassword    = "105"
	InvalidUser           = "106"
	InvalidRole           = "107"
	InvalidPermission     = "108"
	InvalidUserStatus     = "109"
	InvalidUserType       = "110"
	InvalidUserEmail      = "111"
	InvalidUserPhone      = "112"
	ErrEmailAlreadyExists = "113"
	UserNotActivated      = "114"
)
