package loyaltyService

import (
	"github.com/Square-POC/SquarePosBE/configurations"
	"github.com/Square-POC/SquarePosBE/internal/clients"
	"github.com/Square-POC/SquarePosBE/internal/schemas/requestDtos"
	"github.com/Square-POC/SquarePosBE/internal/schemas/responseDtos"
	"github.com/google/uuid"
	"log"
)

const loyaltyServiceLogPrefix = "loyalty_service_impl"

type loyaltyServiceImpl struct {
	client     clients.LoyaltyClient
	squareConf *configurations.SquareConfigurations
}

func NewLoyaltyService(
	client clients.LoyaltyClient,
	squareConf *configurations.SquareConfigurations,
) LoyaltyService {
	return &loyaltyServiceImpl{
		client:     client,
		squareConf: squareConf,
	}
}

func (l *loyaltyServiceImpl) AccumulateLoyalty(request requestDtos.AccumulateLoyaltyIncomingRequestDto, authHeader string) (*responseDtos.AccumulateLoyaltyResponseDto, error) {

	squareReq := requestDtos.AccumulateLoyaltySquareRequestDto{
		AccumulatePoints: requestDtos.AccumulatePoints{
			Points: request.Points,
		},
		IdempotencyKey: uuid.NewString(),
		LocationId:     l.squareConf.LocationId,
	}

	squareResp, err := l.client.AccumulatePoints(squareReq, authHeader)
	if err != nil {
		log.Printf("%v - Error: %v", loyaltyServiceLogPrefix, err)
		return nil, err
	}

	outgoingRep := responseDtos.AccumulateLoyaltyResponseDto{
		Points: squareResp.Events[0].AccumulatePoints.Points,
	}

	return &outgoingRep, nil
}
