package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Square-POC/SquarePosBE/configurations"
	"github.com/Square-POC/SquarePosBE/internal/schemas/requestDtos"
	"github.com/Square-POC/SquarePosBE/internal/schemas/responseDtos"
	"io/ioutil"
	"log"
	"net/http"
)

const loyaltyClientLogPrefix = "square_loyalty_client_impl"

type loyaltyClientImpl struct {
	squareConf *configurations.SquareConfigurations
}

func NewLoyaltyClient(squareConf *configurations.SquareConfigurations) LoyaltyClient {
	return &loyaltyClientImpl{
		squareConf: squareConf,
	}
}

func (l *loyaltyClientImpl) AccumulatePoints(request requestDtos.AccumulateLoyaltySquareRequestDto, authHeader string) (*responseDtos.AccumulateLoyaltySquareResponseDto, error) {
	url := fmt.Sprintf("%v/loyalty/accounts/%v/accumulate", l.squareConf.BaseUrl, l.squareConf.AccountId)
	method := "POST"

	payload, err := json.Marshal(request)
	if err != nil {
		log.Printf("%v - Error marshalling JSON: %v", loyaltyClientLogPrefix, err)
		return nil, err
	}

	log.Printf("%v - Making request to %s with payload %s", loyaltyClientLogPrefix, url, string(payload))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
	if err != nil {
		log.Printf("%v - Error: %v", loyaltyClientLogPrefix, err)
		return nil, err
	}
	req.Header.Add("Square-Version", l.squareConf.SquareVersion)
	req.Header.Add("Authorization", fmt.Sprintf("%v", authHeader))
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Printf("%v - Error: %v", loyaltyClientLogPrefix, err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("%v - Error: %v", loyaltyClientLogPrefix, err)
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		var errorResponse responseDtos.SquareErrorResponseDto
		if err := json.Unmarshal(body, &errorResponse); err != nil {
			log.Printf("%v - Error: %v", loyaltyClientLogPrefix, err)
			return nil, err
		}
		return nil, fmt.Errorf("square API error: %v", errorResponse.Errors)
	}

	var internalResp responseDtos.AccumulateLoyaltySquareResponseDto
	if err := json.Unmarshal(body, &internalResp); err != nil {
		log.Printf("%v - Error: %v", loyaltyClientLogPrefix, err)
		return nil, err
	}

	return &internalResp, nil

}

func (l *loyaltyClientImpl) CreateLoyaltyReward(request requestDtos.CreateLoyaltyRewardSquareRequestDto, authHeader string) (*responseDtos.CreateLoyaltyRewardSquareResponseDto, error) {
	url := fmt.Sprintf("%v/loyalty/rewards/%v/create", l.squareConf.BaseUrl)
	method := "POST"

	payload, err := json.Marshal(request)
	if err != nil {
		log.Printf("%v - Error marshalling JSON: %v", loyaltyClientLogPrefix, err)
		return nil, err
	}

	log.Printf("%v - Making request to %s with payload %s", loyaltyClientLogPrefix, url, string(payload))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
	if err != nil {
		log.Printf("%v - Error: %v", loyaltyClientLogPrefix, err)
		return nil, err
	}

	req.Header.Add("Square-Version", l.squareConf.SquareVersion)
	req.Header.Add("Authorization", fmt.Sprintf("%v", authHeader))
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Printf("%v - Error: %v", loyaltyClientLogPrefix, err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("%v - Error: %v", loyaltyClientLogPrefix, err)
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		var errorResponse responseDtos.SquareErrorResponseDto
		if err := json.Unmarshal(body, &errorResponse); err != nil {
			log.Printf("%v - Error: %v", loyaltyClientLogPrefix, err)
		}
	}

	var internalResp responseDtos.CreateLoyaltyRewardSquareResponseDto
	if err := json.Unmarshal(body, &internalResp); err != nil {
		log.Printf("%v - Error: %v", loyaltyClientLogPrefix, err)
		return nil, err
	}

	return &internalResp, nil
}

func (l *loyaltyClientImpl) RedeemLoyaltyReward(request requestDtos.RedeemLoyaltySquareRequestDto, authHeader string, rewardId string) (*responseDtos.RedeemLoyaltySquareResponseDto, error) {
	url := fmt.Sprintf("%v/loyalty/rewards/%v/redeem", l.squareConf.BaseUrl, rewardId)
	method := "POST"

	payload, err := json.Marshal(request)
	if err != nil {
		log.Printf("%v - Error marshalling JSON: %v", loyaltyClientLogPrefix, err)
		return nil, err
	}

	log.Printf("%v - Making request to %s with payload %s", loyaltyClientLogPrefix, url, string(payload))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
	if err != nil {
		log.Printf("%v - Error: %v", loyaltyClientLogPrefix, err)
		return nil, err
	}

	req.Header.Add("Square-Version", l.squareConf.SquareVersion)
	req.Header.Add("Authorization", fmt.Sprintf("%v", authHeader))
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Printf("%v - Error: %v", loyaltyClientLogPrefix, err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("%v - Error: %v", loyaltyClientLogPrefix, err)
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		var errorResponse responseDtos.SquareErrorResponseDto
		if err := json.Unmarshal(body, &errorResponse); err != nil {
			log.Printf("%v - Error: %v", loyaltyClientLogPrefix, err)
		}
	}

	var internalResp responseDtos.RedeemLoyaltySquareResponseDto
	if err := json.Unmarshal(body, &internalResp); err != nil {
		log.Printf("%v - Error: %v", loyaltyClientLogPrefix, err)
		return nil, err
	}

	return &internalResp, nil
}

func (l *loyaltyClientImpl) RetrieveLoyaltyAccount(authHeader string, accountId string) (*responseDtos.RetrieveLoyaltyAccountSquareResponseDto, error) {
	url := fmt.Sprintf("%v/loyalty/accounts/%v", l.squareConf.BaseUrl, accountId)
	method := "GET"

	log.Printf("%v - Making request to %s", loyaltyClientLogPrefix, url)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Printf("%v - Error: %v", loyaltyClientLogPrefix, err)
		return nil, err
	}

	req.Header.Add("Square-Version", l.squareConf.SquareVersion)
	req.Header.Add("Authorization", fmt.Sprintf("%v", authHeader))

	res, err := client.Do(req)
	if err != nil {
		log.Printf("%v - Error: %v", loyaltyClientLogPrefix, err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("%v - Error: %v", loyaltyClientLogPrefix, err)
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		var errorResponse responseDtos.SquareErrorResponseDto
		if err := json.Unmarshal(body, &errorResponse); err != nil {
			log.Printf("%v - Error: %v", loyaltyClientLogPrefix, err)
		}
	}

	var internalResp responseDtos.RetrieveLoyaltyAccountSquareResponseDto
	if err := json.Unmarshal(body, &internalResp); err != nil {
		log.Printf("%v - Error: %v", loyaltyClientLogPrefix, err)
		return nil, err
	}

	return &internalResp, nil
}

func (l *loyaltyClientImpl) SearchLoyaltyRewards(dto *requestDtos.SearchLoyaltyRewardsSquareRequestDto, authHeader string) (*responseDtos.SearchLoyaltyRewardsSquareResponseDto, error) {
	url := fmt.Sprintf("%v/loyalty/rewards/search", l.squareConf.BaseUrl)
	method := "POST"

	payload, err := json.Marshal(dto)
	if err != nil {
		log.Printf("%v - Error marshalling JSON: %v", loyaltyClientLogPrefix, err)
		return nil, err
	}

	log.Printf("%v - Making request to %s with payload %s", loyaltyClientLogPrefix, url, string(payload))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
	if err != nil {
		log.Printf("%v - Error: %v", loyaltyClientLogPrefix, err)
		return nil, err
	}

	req.Header.Add("Square-Version", l.squareConf.SquareVersion)
	req.Header.Add("Authorization", fmt.Sprintf("%v", authHeader))
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Printf("%v - Error: %v", loyaltyClientLogPrefix, err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("%v - Error: %v", loyaltyClientLogPrefix, err)
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		var errorResponse responseDtos.SquareErrorResponseDto
		if err := json.Unmarshal(body, &errorResponse); err != nil {
			log.Printf("%v - Error: %v", loyaltyClientLogPrefix, err)
		}
		return nil, fmt.Errorf("square API error: %v", errorResponse.Errors)
	}

	var internalResp responseDtos.SearchLoyaltyRewardsSquareResponseDto
	if err := json.Unmarshal(body, &internalResp); err != nil {
		log.Printf("%v - Error: %v", loyaltyClientLogPrefix, err)
		return nil, err
	}

	return &internalResp, nil

}
