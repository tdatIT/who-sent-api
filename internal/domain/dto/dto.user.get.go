package dto

type GetUserByIdReq struct {
	ID int `param:"id" validate:"required"`
}

type GetUserByIdResp struct {
	UserDTO
}
