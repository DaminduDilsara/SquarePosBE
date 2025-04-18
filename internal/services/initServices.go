package services

import (
	"github.com/Square-POC/SquarePosBE/configurations"
	"github.com/Square-POC/SquarePosBE/internal/clients"
	"github.com/Square-POC/SquarePosBE/internal/services/authService"
	"github.com/Square-POC/SquarePosBE/internal/services/loyaltyService"
)

type ServiceCollection struct {
	AuthSvc    authService.AuthService
	LoyaltySvc loyaltyService.LoyaltyService
}

func InitServices(
	conf *configurations.Config,
	loyaltyClient clients.LoyaltyClient,
) *ServiceCollection {

	return &ServiceCollection{
		AuthSvc:    authService.NewAuthService(conf.AuthConfig),
		LoyaltySvc: loyaltyService.NewLoyaltyService(loyaltyClient, conf.SquareConfig),
	}
}
