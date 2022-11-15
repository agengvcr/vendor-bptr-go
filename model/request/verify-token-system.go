package request

type VerifyTokenSystemRequest struct {
	Token  string `form:"token" binding:"required"`
	Server string `form:"server" binding:"required"`
}
