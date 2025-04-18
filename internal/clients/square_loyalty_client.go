package clients

import (
	"github.com/Square-POC/SquarePosBE/internal/schemas/requestDtos"
	"github.com/Square-POC/SquarePosBE/internal/schemas/responseDtos"
)

type LoyaltyClient interface {
	AccumulatePoints(request requestDtos.AccumulateLoyaltySquareRequestDto, authHeader string) (*responseDtos.AccumulateLoyaltySquareResponseDto, error)
}
