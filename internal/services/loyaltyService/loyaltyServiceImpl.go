package loyaltyService

import (
	"github.com/Square-POC/SquarePosBE/configurations"
	"github.com/Square-POC/SquarePosBE/internal/clients"
	"github.com/Square-POC/SquarePosBE/internal/schemas/requestDtos"
	"github.com/Square-POC/SquarePosBE/internal/schemas/responseDtos"
	"github.com/google/uuid"
	"log"
)

const loyaltyServiceLogPrefix = "loyalty_service_impl"

type loyaltyServiceImpl struct {
	client     clients.LoyaltyClient
	squareConf *configurations.SquareConfigurations
}

func NewLoyaltyService(
	client clients.LoyaltyClient,
	squareConf *configurations.SquareConfigurations,
) LoyaltyService {
	return &loyaltyServiceImpl{
		client:     client,
		squareConf: squareConf,
	}
}

func (l *loyaltyServiceImpl) AccumulateLoyaltyService(request requestDtos.AccumulateLoyaltyRequestDto, authHeader string) (*responseDtos.AccumulateLoyaltyResponseDto, error) {

	squareReq := requestDtos.AccumulateLoyaltySquareRequestDto{
		AccumulatePoints: requestDtos.AccumulatePoints{
			Points: request.Points,
		},
		IdempotencyKey: uuid.NewString(),
		LocationId:     l.squareConf.LocationId,
	}

	squareResp, err := l.client.AccumulatePoints(squareReq, authHeader)
	if err != nil {
		log.Printf("%v - Error: %v", loyaltyServiceLogPrefix, err)
		return nil, err
	}

	outgoingRep := responseDtos.AccumulateLoyaltyResponseDto{
		Points: squareResp.Events[0].AccumulatePoints.Points,
	}

	return &outgoingRep, nil
}

func (l *loyaltyServiceImpl) CreateLoyaltyRewardService(authHeader string) (*responseDtos.CreateLoyaltyRewardResponseDto, error) {

	squareReq := requestDtos.CreateLoyaltyRewardSquareRequestDto{
		IdempotencyKey: uuid.NewString(),
		Reward: requestDtos.Reward{
			LoyaltyAccountId: l.squareConf.LocationId,
			RewardTierId:     l.squareConf.LoyaltyTierId,
		},
	}

	squareResp, err := l.client.CreateLoyaltyReward(squareReq, authHeader)
	if err != nil {
		log.Printf("%v - Error: %v", loyaltyServiceLogPrefix, err)
		return nil, err
	}

	outgoingRep := responseDtos.CreateLoyaltyRewardResponseDto{
		RewardId: squareResp.Reward.Id,
		Points:   squareResp.Reward.Points,
		Status:   squareResp.Reward.Status,
	}
	return &outgoingRep, nil
}

func (l *loyaltyServiceImpl) RedeemLoyaltyRewardService(authHeader string, rewardId string) (*responseDtos.RedeemLoyaltyResponseDto, error) {

	squareReq := requestDtos.RedeemLoyaltySquareRequestDto{
		IdempotencyKey: uuid.NewString(),
		LocationId:     l.squareConf.LocationId,
	}

	_, err := l.client.RedeemLoyaltyReward(squareReq, authHeader, rewardId)
	if err != nil {
		log.Printf("%v - Error: %v", loyaltyServiceLogPrefix, err)
		return nil, err
	}

	outgoingRep := responseDtos.RedeemLoyaltyResponseDto{
		Status: "Success",
	}

	return &outgoingRep, nil
}

func (l *loyaltyServiceImpl) RetrieveLoyaltyAccountService(authHeader string) (*responseDtos.RetrieveLoyaltyAccountResponseDto, error) {

	squareResp, err := l.client.RetrieveLoyaltyAccount(authHeader, l.squareConf.AccountId)
	if err != nil {
		log.Printf("%v - Error: %v", loyaltyServiceLogPrefix, err)
		return nil, err
	}

	outgoingRep := responseDtos.RetrieveLoyaltyAccountResponseDto{
		Balance:        squareResp.LoyaltyAccount.Balance,
		LifetimePoints: squareResp.LoyaltyAccount.LifetimePoints,
		EnrolledAt:     squareResp.LoyaltyAccount.EnrolledAt,
	}

	return &outgoingRep, nil
}

func (l *loyaltyServiceImpl) SearchLoyaltyRewards(authHeader string, status string) (*responseDtos.SearchLoyaltyRewardsResponseDto, error) {
	squareReq := requestDtos.SearchLoyaltyRewardsSquareRequestDto{
		Query: requestDtos.Query{
			LoyaltyAccountId: l.squareConf.AccountId,
			Status:           status,
		},
	}

	sqareResp, err := l.client.SearchLoyaltyRewards(&squareReq, authHeader)
	if err != nil {
		log.Printf("%v - Error: %v", loyaltyServiceLogPrefix, err)
		return nil, err
	}

	outgoingRep := responseDtos.SearchLoyaltyRewardsResponseDto{
		Rewards: sqareResp.Rewards,
	}

	return &outgoingRep, nil
}
