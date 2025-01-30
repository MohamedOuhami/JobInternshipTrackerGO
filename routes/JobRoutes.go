package routes

import (
	"github.com/MohamedOuhami/JobInternshipTrackerGO/controllers"
	"github.com/MohamedOuhami/JobInternshipTrackerGO/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupJobRoutes(r *gin.Engine){

  jobRoutes := r.Group("/api/v1/jobs")
  
  {
    jobRoutes.GET("/",middlewares.RequireAuth,middlewares.RequireOwner,controllers.GetAllJobs)
    jobRoutes.GET("/:id",middlewares.RequireAuth,middlewares.RequireOwner,controllers.GetJobById)
    jobRoutes.GET("/byOwner/:ownerId",middlewares.RequireAuth,middlewares.RequireOwner,controllers.GetJobsByOwner)
    jobRoutes.POST("/",middlewares.RequireAuth,middlewares.RequireOwner,controllers.CreateJob)
    jobRoutes.PUT("/:id",middlewares.RequireAuth,middlewares.RequireOwner,controllers.EditJob)
    jobRoutes.DELETE("/:id",middlewares.RequireAuth,middlewares.RequireOwner,controllers.DeleteJob)
    jobRoutes.DELETE("/massDelete",middlewares.RequireAuth,middlewares.RequireOwner,controllers.MassDeleteJobs)

  }
}
