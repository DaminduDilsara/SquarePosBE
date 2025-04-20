package requestDtos

type RedeemLoyaltySquareRequestDto struct {
	IdempotencyKey string `json:"idempotency_key"`
	LocationId     string `json:"location_id"`
}

type RedeemLoyaltyRequestDto struct {
	LoyaltyRewardId string `json:"loyalty_reward_id"`
}
