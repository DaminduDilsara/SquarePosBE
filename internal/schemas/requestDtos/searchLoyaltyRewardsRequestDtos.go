package requestDtos

type SearchLoyaltyRewardsSquareRequestDto struct {
	Query Query `json:"query"`
}

type Query struct {
	LoyaltyAccountId string `json:"loyalty_account_id"`
	Status           string `json:"status"`
}
