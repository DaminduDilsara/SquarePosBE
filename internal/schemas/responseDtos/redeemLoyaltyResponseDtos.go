package responseDtos

import "time"

type RedeemLoyaltyResponseDto struct {
	Status string `json:"status"`
}

type RedeemLoyaltySquareResponseDto struct {
	Event Event `json:"event"`
}

type Event struct {
	Id               string       `json:"id"`
	Type             string       `json:"type"`
	CreatedAt        time.Time    `json:"created_at"`
	RedeemReward     RedeemReward `json:"redeem_reward"`
	LoyaltyAccountId string       `json:"loyalty_account_id"`
	LocationId       string       `json:"location_id"`
	Source           string       `json:"source"`
}

type RedeemReward struct {
	LoyaltyProgramId string `json:"loyalty_program_id"`
	RewardId         string `json:"reward_id"`
}
