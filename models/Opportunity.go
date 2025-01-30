package models

import (
	"time"

	"gorm.io/gorm"
)

type Opportunity struct {

  gorm.Model

  CompanyName string
  JobType string
  PostName string
  City string
  Status string
  NextInterview *time.Time
  UserID uint `gorm:"column:owner_id"`


}
