package responseDtos

type SquareErrorResponseDto struct {
	Errors []struct {
		Code     string `json:"code"`
		Detail   string `json:"detail"`
		Category string `json:"category"`
	} `json:"errors"`
}
