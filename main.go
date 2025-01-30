package main

import (
	"github.com/MohamedOuhami/JobInternshipTrackerGO/initializers"
	"github.com/MohamedOuhami/JobInternshipTrackerGO/routes"
	"github.com/gin-gonic/gin"

)

// Before starting the program
func init() {

	// Load the Environment variables
	initializers.LoadEnv()

	// Connect to the database
	initializers.InitializeDB()

	// Migrate the database
	initializers.SyncDatabase()
}

// The main function
func main() {

	r := gin.Default()

  // Auth Routes
  routes.SetupAuthRouter(r)

  // Opportunties Routes
  routes.SetupOpportunityRoutes(r)

  // Job Routes
  routes.SetupJobRoutes(r)

	r.Run()

}
