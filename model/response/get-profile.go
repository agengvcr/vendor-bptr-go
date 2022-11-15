package response

import "vendor.bptr.co.id/api-auth/model"

type GetProfileResponse struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Name     string `json:"name"`
}

func NewGetProfileResponseFromUser(user model.User) GetProfileResponse {
	return GetProfileResponse{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
		Name:     user.Name,
	}
}
