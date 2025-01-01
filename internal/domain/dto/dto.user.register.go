package dto

type UserRegisterReq struct {
	Firstname string `json:"firstname" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
}

type UserRegisterResp struct {
	UserDTO
}
