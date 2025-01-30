package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/MohamedOuhami/JobInternshipTrackerGO/initializers"
	"github.com/MohamedOuhami/JobInternshipTrackerGO/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {

	var user models.User

	// Get the token from the request cookies
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// Return the secret key for validation
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	// Error handling for parsing the token
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Validate the claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		// Check token expiration
		if isTokenExpired(claims) {
			fmt.Println("Token expired")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		//
		fmt.Println("========== Claims =========")
		fmt.Println(claims["sub"])

		initializers.DB.First(&user, claims["sub"])

		// Check if the user exists
		if user.ID == 0 {
			fmt.Println("No subject")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Attach to req
		c.Set("user", user)

		// Continue the actual request
		c.Next()

	} else {
		// Invalid token claims
		fmt.Println("Invalid Token claims")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

}

// Helper function to check if the token is expired
func isTokenExpired(claims jwt.MapClaims) bool {
	expirationTime := float64(time.Now().Unix())
	if float64(claims["exp"].(float64)) < expirationTime {
		return true
	}
	return false
}
