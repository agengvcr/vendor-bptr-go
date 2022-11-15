package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"vendor.bptr.co.id/api-auth/lib"
	"vendor.bptr.co.id/api-auth/model/request"
	response "vendor.bptr.co.id/api-auth/model/response"
	repository "vendor.bptr.co.id/api-auth/repository"
)

type AuthController struct {
	userRepo *repository.VendorUserRepository
}

func NewAuthController(userRepo *repository.VendorUserRepository) *AuthController {
	return &AuthController{
		userRepo: userRepo,
	}
}

func (ac *AuthController) GetToken(context *gin.Context) {
	var req request.GetTokenRequest
	errinput := context.ShouldBind(&req) // check empty

	// validasi form login
	if errinput != nil {
		context.JSON(http.StatusBadRequest, response.ErrorResponse(errinput, response.ErrData{StatusCode: http.StatusBadRequest}))
		return
	}

	// Check credential is exist
	user := ac.userRepo.GetUserByUsername(req.Username)
	// if user.Id == 0 {
	// 	context.JSON(http.StatusOK, response.ErrorResponse(errors.New("invalid credential"), response.ErrData{StatusCode: http.StatusUnauthorized}))
	// 	return
	// }

	// Unauthorized user
	checkPassword := lib.HashPassword(req.Password, user.Password)
	if !checkPassword {
		context.JSON(http.StatusOK, response.ErrorResponse(errors.New("invalid credential"), response.ErrData{StatusCode: http.StatusUnauthorized}))
		return
	}

	// generate token paseto
	accessToken, accessPayload, refreshToken, refreshTokenPayload, err := lib.Init().GetToken(
		user.Id,
		user.Username,
	)

	if err != nil {
		context.JSON(http.StatusOK, response.ErrorResponse(err, response.ErrData{StatusCode: http.StatusInternalServerError}))
		return
	}

	data := response.NewGetTokenResponseFromGetToken(accessToken, accessPayload, refreshToken, refreshTokenPayload, user)
	context.JSON(http.StatusOK, response.SuccessResponse(data))
}
