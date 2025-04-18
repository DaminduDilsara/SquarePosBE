package responseDtos

type OAuthLoginResponse struct {
	Message string `json:"message"`
	User    User   `json:"user"`
}

type User struct {
	Email         string `json:"email"`
	FamilyName    string `json:"family_name"`
	GivenName     string `json:"given_name"`
	Id            string `json:"id"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	VerifiedEmail bool   `json:"verified_email"`
}
