package request

type VerifyTokenUserRequest struct {
	Token  string `form:"token" binding:"required"`
	Access string `form:"access" binding:"required"`
}
