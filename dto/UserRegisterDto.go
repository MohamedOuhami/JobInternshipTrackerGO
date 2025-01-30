package dto

import "time"

type UserRegisterDto struct {

  FirstName string
  LastName string
  Username string 
  Email string
  Password string
  Dob *time.Time `json:"dob"`

}
