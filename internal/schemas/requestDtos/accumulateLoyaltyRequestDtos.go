package requestDtos

type AccumulateLoyaltySquareRequestDto struct {
	AccumulatePoints AccumulatePoints `json:"accumulate_points"`
	IdempotencyKey   string           `json:"idempotency_key"`
	LocationId       string           `json:"location_id"`
}

type AccumulatePoints struct {
	Points int `json:"points"`
}

type AccumulateLoyaltyRequestDto struct {
	Points int `json:"points"`
}
