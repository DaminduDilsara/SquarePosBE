package loyaltyService

import (
	"github.com/Square-POC/SquarePosBE/internal/schemas/requestDtos"
	"github.com/Square-POC/SquarePosBE/internal/schemas/responseDtos"
)

type LoyaltyService interface {
	AccumulateLoyaltyService(request requestDtos.AccumulateLoyaltyRequestDto, authHeader string) (*responseDtos.AccumulateLoyaltyResponseDto, error)
	CreateLoyaltyRewardService(authHeader string) (*responseDtos.CreateLoyaltyRewardResponseDto, error)
	RedeemLoyaltyRewardService(authHeader string, rewardId string) (*responseDtos.RedeemLoyaltyResponseDto, error)
	RetrieveLoyaltyAccountService(authHeader string) (*responseDtos.RetrieveLoyaltyAccountResponseDto, error)
	SearchLoyaltyRewards(authHeader string, status string) (*responseDtos.SearchLoyaltyRewardsResponseDto, error)
}
