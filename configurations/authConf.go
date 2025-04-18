package configurations

type AuthConfigurations struct {
	TokenExpireTime   int    `yaml:"token_expire_time"`
	JwtSecret         string `yaml:"jwt_secret"`
	OAuthClientId     string `yaml:"o_auth_client_id"`
	OAuthClientSecret string `yaml:"o_auth_client_secret"`
	OAuthRedirectUrl  string `yaml:"o_auth_redirect_url"`
}
