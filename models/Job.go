package models

import "gorm.io/gorm"


type Job struct {

  gorm.Model

	CompanyName   string
	JobType       string
	PostName      string
	City          string
  UserID uint `gorm:"column:owner_id"`


}
