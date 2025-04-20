package clients

import (
	"github.com/Square-POC/SquarePosBE/internal/schemas/requestDtos"
	"github.com/Square-POC/SquarePosBE/internal/schemas/responseDtos"
)

type LoyaltyClient interface {
	AccumulatePoints(request requestDtos.AccumulateLoyaltySquareRequestDto, authHeader string) (*responseDtos.AccumulateLoyaltySquareResponseDto, error)
	CreateLoyaltyReward(request requestDtos.CreateLoyaltyRewardSquareRequestDto, authHeader string) (*responseDtos.CreateLoyaltyRewardSquareResponseDto, error)
	RedeemLoyaltyReward(request requestDtos.RedeemLoyaltySquareRequestDto, authHeader string, rewardId string) (*responseDtos.RedeemLoyaltySquareResponseDto, error)
	RetrieveLoyaltyAccount(authHeader string, accountId string) (*responseDtos.RetrieveLoyaltyAccountSquareResponseDto, error)
	SearchLoyaltyRewards(dto *requestDtos.SearchLoyaltyRewardsSquareRequestDto, authHeader string) (*responseDtos.SearchLoyaltyRewardsSquareResponseDto, error)
}
