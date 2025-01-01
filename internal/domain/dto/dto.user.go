package dto

type UserDTO struct {
	ID          int      `json:"id"`
	Firstname   string   `json:"firstname"`
	Lastname    string   `json:"lastname"`
	IsActivated bool     `json:"is_activated"`
	IsVerified  bool     `json:"is_verified"`
	Roles       []string `json:"roles,omitempty"`
	Phone       string   `json:"phone"`
	Email       string   `json:"email"`
	AvatarUrl   string   `json:"avatar_url"`
	CreatedAt   string   `json:"created_at"`
	UpdatedAt   string   `json:"updated_at"`
}
