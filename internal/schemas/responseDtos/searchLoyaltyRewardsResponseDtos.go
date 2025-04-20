package responseDtos

import "time"

type SearchLoyaltyRewardsSquareResponseDto struct {
	Rewards []Rewards `json:"rewards"`
}

type Rewards struct {
	Id               string    `json:"id"`
	Status           string    `json:"status"`
	LoyaltyAccountId string    `json:"loyalty_account_id"`
	RewardTierId     string    `json:"reward_tier_id"`
	Points           int       `json:"points"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	RedeemedAt       time.Time `json:"redeemed_at,omitempty"`
}

type SearchLoyaltyRewardsResponseDto struct {
	Rewards []Rewards `json:"rewards"`
}
