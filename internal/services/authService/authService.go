package authService

import "github.com/Square-POC/SquarePosBE/internal/schemas/responseDtos"

type AuthService interface {
	OAuthLogin() string
	OAuthCallBack(code string) (*responseDtos.OAuthLoginResponse, error)
}
