package routes

import (
	"github.com/MohamedOuhami/JobInternshipTrackerGO/controllers"
	"github.com/MohamedOuhami/JobInternshipTrackerGO/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupOpportunityRoutes(r *gin.Engine) {

	jobRoutes := r.Group("/api/v1/opportunities")

	{
		jobRoutes.GET("/", middlewares.RequireAuth, middlewares.RequireOwner, controllers.GetAllOpportunities)
		jobRoutes.GET("/:id", middlewares.RequireAuth, middlewares.RequireOwner, controllers.GetopportunityById)
		jobRoutes.GET("/byOwner/:ownerId", middlewares.RequireAuth, middlewares.RequireOwner, controllers.GetOpportunitiesByOwner)
		jobRoutes.POST("/", middlewares.RequireAuth, middlewares.RequireOwner, controllers.Createopportunity)
		jobRoutes.PUT("/:id", middlewares.RequireAuth, middlewares.RequireOwner, controllers.Editopportunity)
		jobRoutes.DELETE("/:id", middlewares.RequireAuth, middlewares.RequireOwner, controllers.Deleteopportunity)
		jobRoutes.DELETE("/massDelete", middlewares.RequireAuth, middlewares.RequireOwner, controllers.Deleteopportunity)
		jobRoutes.PUT("/turnToJob/:id", middlewares.RequireAuth, middlewares.RequireOwner, controllers.TurnOpportunityToJob)

	}
}
