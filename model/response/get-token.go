package response

import (
	"time"

	"vendor.bptr.co.id/api-auth/model"
)

type GetTokenUserResponse struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Name     string `json:"name"`
}

type GetTokenResponse struct {
	Token                 string               `json:"access_token"`
	AccessTokenExpiresAt  time.Time            `json:"access_token_expires_at"`
	RefreshToken          string               `json:"refresh_token"`
	RefreshTokenExpiresAt time.Time            `json:"refresh_token_expires_at"`
	User                  GetTokenUserResponse `json:"user"`
}

func NewGetTokenResponseFromGetToken(accessToken string, accessPayload *model.Payload, refreshToken string, refreshTokenPayload *model.Payload, user model.User) GetTokenResponse {
	getTokenUserResponse := GetTokenUserResponse{
		Id:       user.Id,
		Username: user.Username,
		Name:     user.Name,
		Email:    user.Email,
	}
	return GetTokenResponse{
		Token:                 accessToken,
		AccessTokenExpiresAt:  accessPayload.ExpiredAt,
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: refreshTokenPayload.ExpiredAt,
		User:                  getTokenUserResponse,
	}
}
