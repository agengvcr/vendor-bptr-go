package response

import "vendor.bptr.co.id/api-auth/model"

type VerifyTokenResponse struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Name     string `json:"name"`
}

func NewVerifyTokenResponseFromUser(user model.User) VerifyTokenResponse {
	return VerifyTokenResponse{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
		Name:     user.Name,
	}
}
