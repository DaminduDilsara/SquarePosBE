package responseDtos

import "time"

type AccumulateLoyaltySquareResponseDto struct {
	Events []Events `json:"events"`
}

type Events struct {
	Id               string           `json:"id"`
	Type             string           `json:"type"`
	CreatedAt        time.Time        `json:"created_at"`
	AccumulatePoints AccumulatePoints `json:"accumulate_points"`
	LoyaltyAccountId string           `json:"loyalty_account_id"`
	LocationId       string           `json:"location_id"`
	Source           string           `json:"source"`
}

type AccumulatePoints struct {
	LoyaltyProgramId string `json:"loyalty_program_id"`
	Points           int    `json:"points"`
}

type AccumulateLoyaltyResponseDto struct {
	Points int `json:"points_earned"`
}
