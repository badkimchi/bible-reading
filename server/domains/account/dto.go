package account

type UserInfoDto struct {
	Sub           string `json:"sub"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	EmailVerified bool   `json:"email_verified"`
	Picture       string `json:"picture"`
	Email         string `json:"email"`
	Locale        string `json:"locale"`
	Jwt           Jwt    `json:"jwt"`
}

type OAuthRequest struct {
	Token string `json:"token"`
}

type Jwt struct {
	Token        string       `json:"token"`
	Expiration   string       `json:"expiration"` // expiration is already in the token but useful for frontend
	RefreshToken RefreshToken `json:"refreshToken"`
}

type RefreshToken struct {
	Token      string `json:"token"`
	Expiration string `json:"expiration"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken"`
}
