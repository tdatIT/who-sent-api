package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/tdatIT/who-sent-api/config"
	"time"
)

type AuthJwtProvider struct {
	cfg *config.AppConfig
}

func NewAuthJwtProvider(cfg *config.AppConfig) *AuthJwtProvider {
	return &AuthJwtProvider{cfg: cfg}
}

// GenerateAccessToken creates a new access token
func (ap *AuthJwtProvider) GenerateAccessToken(userId string, username string, roles []string) (string, error) {
	claims := AuthClaims{
		UserId:   userId,
		Username: username,
		Roles:    roles,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ap.cfg.Adapter.Auth.AccessExp)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   userId,
			ID:        uuid.NewString(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(ap.cfg.Adapter.Auth.JwtSecret))
}

// GenerateRefreshToken creates a new refresh token
func (ap *AuthJwtProvider) GenerateRefreshToken(userId string) (string, error) {
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(ap.cfg.Adapter.Auth.RefreshExp)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Subject:   userId,
		ID:        uuid.NewString(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(ap.cfg.Adapter.Auth.JwtSecret))
}

// ValidateAccessToken validates an access token and returns claims
func (ap *AuthJwtProvider) ValidateAccessToken(tokenString string) (*AuthClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(ap.cfg.Adapter.Auth.JwtSecret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*AuthClaims)
	if !ok || !token.Valid {
		return nil, jwt.ErrTokenMalformed
	}

	return claims, nil
}

// ValidateRefreshToken validates a refresh token and returns claims
func (ap *AuthJwtProvider) ValidateRefreshToken(tokenString string) (*jwt.RegisteredClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(ap.cfg.Adapter.Auth.JwtSecret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok || !token.Valid {
		return nil, jwt.ErrTokenMalformed
	}

	return claims, nil
}

// GenerateAccessTokenFromRefreshToken generates a new access token using a valid refresh token
func (ap *AuthJwtProvider) GenerateAccessTokenFromRefreshToken(refreshToken string, username string, roles []string) (string, error) {
	claims, err := ap.ValidateRefreshToken(refreshToken)
	if err != nil {
		return "", err
	}

	userId := claims.Subject
	if userId == "" {
		return "", jwt.ErrTokenMalformed
	}

	return ap.GenerateAccessToken(userId, username, roles)
}
