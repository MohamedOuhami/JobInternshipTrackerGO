package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {

  gorm.Model

  FirstName string
  LastName string
  Username string `gorm:"unique"`
  Email string `gorm:"unique"`
  Password string
  Dob *time.Time
  Age int

  // Associations
  Jobs []Job
  Opportunities []Opportunity

}
