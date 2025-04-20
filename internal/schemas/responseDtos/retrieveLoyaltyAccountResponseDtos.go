package responseDtos

import "time"

type RetrieveLoyaltyAccountSquareResponseDto struct {
	LoyaltyAccount LoyaltyAccount `json:"loyalty_account"`
}

type Mapping struct {
	Id          string    `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	PhoneNumber string    `json:"phone_number"`
}

type LoyaltyAccount struct {
	Id             string    `json:"id"`
	ProgramId      string    `json:"program_id"`
	Balance        int       `json:"balance"`
	LifetimePoints int       `json:"lifetime_points"`
	CustomerId     string    `json:"customer_id"`
	EnrolledAt     time.Time `json:"enrolled_at"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	Mapping        Mapping   `json:"mapping"`
}

type RetrieveLoyaltyAccountResponseDto struct {
	Balance        int       `json:"balance"`
	LifetimePoints int       `json:"lifetime_points"`
	EnrolledAt     time.Time `json:"enrolled_at"`
}
