package requestDtos

type CreateLoyaltyRewardSquareRequestDto struct {
	IdempotencyKey string `json:"idempotency_key"`
	Reward         Reward `json:"reward"`
}

type Reward struct {
	LoyaltyAccountId string `json:"loyalty_account_id"`
	RewardTierId     string `json:"reward_tier_id"`
}
