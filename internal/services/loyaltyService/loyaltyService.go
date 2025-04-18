package loyaltyService

import (
	"github.com/Square-POC/SquarePosBE/internal/schemas/requestDtos"
	"github.com/Square-POC/SquarePosBE/internal/schemas/responseDtos"
)

type LoyaltyService interface {
	AccumulateLoyalty(request requestDtos.AccumulateLoyaltyIncomingRequestDto, authHeader string) (*responseDtos.AccumulateLoyaltyResponseDto, error)
}
