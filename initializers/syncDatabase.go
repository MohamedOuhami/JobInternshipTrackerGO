package initializers

import (
	"fmt"

	"github.com/MohamedOuhami/JobInternshipTrackerGO/models"
)

func SyncDatabase() {

	err := DB.AutoMigrate(&models.User{},&models.Job{},&models.Opportunity{})

  if err != nil {
    panic("There was an error in creating the table")
  }

  fmt.Println("Succesfully created the tables")

}
