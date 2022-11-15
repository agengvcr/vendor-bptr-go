package request

type GetRefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}
