package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func RequireOwner(c *gin.Context){

  
  user,exists := c.Get("user")

  if !exists {
    fmt.Println("The user instance does not exist")
  }

  fmt.Println("============= This is the user ======== ")
  fmt.Println(user)

  c.Next()


}
 
