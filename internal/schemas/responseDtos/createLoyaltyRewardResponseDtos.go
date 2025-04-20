package responseDtos

import "time"

type CreateLoyaltyRewardSquareResponseDto struct {
	Reward Reward `json:"reward"`
}

type CreateLoyaltyRewardResponseDto struct {
	RewardId string `json:"reward_id"`
	Points   int    `json:"points"`
	Status   string `json:"status"`
}

type Reward struct {
	Id               string    `json:"id"`
	Status           string    `json:"status"`
	LoyaltyAccountId string    `json:"loyalty_account_id"`
	RewardTierId     string    `json:"reward_tier_id"`
	Points           int       `json:"points"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
