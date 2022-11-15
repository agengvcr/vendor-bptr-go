package response

import "vendor.bptr.co.id/api-auth/model"

type GetUserResponse struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Name     string `json:"name"`
}

func NewGetUserResponseFromUser(user model.User) GetUserResponse {
	return GetUserResponse{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
		Name:     user.Name,
	}
}
