package services

import (
	"github.com/Square-POC/SquarePosBE/configurations"
	"github.com/Square-POC/SquarePosBE/internal/services/authService"
)

type ServiceCollection struct {
	AuthSvc authService.AuthService
}

func InitServices(
	conf *configurations.Config,
) *ServiceCollection {

	return &ServiceCollection{
		AuthSvc: authService.NewAuthService(conf.AuthConfig),
	}
}
