package dto

type LoginByUserPasswordReq struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginByUserPasswordResp struct {
	User UserDTO `json:"user"`
	GetAccessTokenData
}

type GetAccessTokenData struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
}
