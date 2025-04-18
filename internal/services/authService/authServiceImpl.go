package authService

import (
	"context"
	"encoding/json"
	"github.com/Square-POC/SquarePosBE/configurations"
	"github.com/Square-POC/SquarePosBE/internal/schemas/responseDtos"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var googleOauthConfig = &oauth2.Config{
	Scopes:   []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
	Endpoint: google.Endpoint,
}

type authServiceImpl struct {
	authConf *configurations.AuthConfigurations
}

func NewAuthService(
	authConf *configurations.AuthConfigurations,
) AuthService {
	googleOauthConfig.ClientID = authConf.OAuthClientId
	googleOauthConfig.ClientSecret = authConf.OAuthClientSecret
	googleOauthConfig.RedirectURL = authConf.OAuthRedirectUrl
	return &authServiceImpl{
		authConf: authConf,
	}
}

func (a *authServiceImpl) OAuthLogin() string {
	return googleOauthConfig.AuthCodeURL("randomstate")
}

func (a *authServiceImpl) OAuthCallBack(code string) (*responseDtos.OAuthLoginResponse, error) {

	// Exchange code for token
	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, err
	}

	// Get user info from Google
	client := googleOauthConfig.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	var user responseDtos.User

	b, _ := json.Marshal(result)
	_ = json.Unmarshal(b, &user)

	userInfo := responseDtos.OAuthLoginResponse{
		Message: "AuthSuccess",
		User:    user,
	}

	return &userInfo, nil
}
