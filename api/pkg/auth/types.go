package auth

type AuthenticateResult struct {
	// User Tokens
	Data *AuthTokens `json:"data"`
}

// Auth tokens have access and refresh token for user
type AuthTokens struct {
	// Access Token
	Access *Token `json:"access"`
	// Refresh Token
	Refresh *Token `json:"refresh"`
}

// Token includes the JWT, Expire Duration & Time
type Token struct {
	// JWT
	Token string `json:"token"`
	// Duration the token will Expire In
	RefreshInterval string `json:"refreshInterval"`
	// Time the token will expires at
	ExpiresAt int64 `json:"expiresAt"`
}
