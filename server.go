package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"vendor.bptr.co.id/api-auth/config"
	controller "vendor.bptr.co.id/api-auth/controller"
	globalvar "vendor.bptr.co.id/api-auth/globalvar"
	response "vendor.bptr.co.id/api-auth/model/response"
	repository "vendor.bptr.co.id/api-auth/repository"
)

var (
	// ini database
	db       *gorm.DB                            = config.SetupDatabaseConnection()
	userRepo *repository.VendorUserRepository    = repository.NewVendorUserRepository(db)
	authCont controller.VendorUserAuthController = *controller.NewVendorUserAuthController(userRepo)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	v1Routes := r.Group("/v1")
	{
		v1Routes.GET("/version", func(c *gin.Context) {
			c.JSON(http.StatusOK, response.SuccessResponse(response.Version{
				Version:  globalvar.VERSION,
				Database: os.Getenv("DB_DATABASE"),
			}))
		})

		// Auth
		authRoutes := v1Routes.Group("/auth")
		{
			tokenRoutes := authRoutes.Group("/token")
			{
				tokenRoutes.POST("/get", authCont.GetToken)
				// tokenRoutes.GET("/verify", authCont.VerifyToken)
				// tokenRoutes.POST("/refresh", authCont.GetRefreshToken)
			}
		}
	}

	servicePort := os.Getenv("SERVICE_PORT")
	err := r.Run(":" + servicePort)
	if err != nil {
		log.Println("Error Running on Port - " + servicePort)
		return
	}
}
