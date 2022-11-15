package request

type VerifyTokenRequest struct {
	Token string `form:"token" binding:"required"`
	Type  string `form:"type" binding:"required"`
}
