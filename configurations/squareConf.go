package configurations

type SquareConfigurations struct {
	BaseUrl       string `yaml:"base_url"`
	SquareVersion string `yaml:"square_version"`
	LocationId    string `yaml:"location_id"`
	AccountId     string `yaml:"account_id"`
	LoyaltyTierId string `yaml:"loyalty_tier_id"`
}
