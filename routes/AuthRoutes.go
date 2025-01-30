package routes

import (
	"github.com/MohamedOuhami/JobInternshipTrackerGO/controllers"
	"github.com/gin-gonic/gin"
)

func SetupAuthRouter(r *gin.Engine) {

	authRoutes := r.Group("/api/v1/auth")
	{
		// Register Users
		authRoutes.POST("/signup", controllers.Register)

		// Login the user endpoint
		authRoutes.POST("/login", controllers.Login)
	}

}
