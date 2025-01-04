package auth

import "github.com/golang-jwt/jwt/v5"

type AuthClaims struct {
	UserId   string   `json:"user_id"`
	Username string   `json:"username"`
	Roles    []string `json:"roles"`
	jwt.RegisteredClaims
}
