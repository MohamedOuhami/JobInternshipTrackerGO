package controllers

import (
	"fmt"
	"math"
	"net/http"
	"os"
	"time"

	"github.com/MohamedOuhami/JobInternshipTrackerGO/dto"
	"github.com/MohamedOuhami/JobInternshipTrackerGO/initializers"
	"github.com/MohamedOuhami/JobInternshipTrackerGO/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// Register the new User
func Register(c *gin.Context) {

	// Get the info from the body
	var body dto.UserRegisterDto

	// Bind the body to the Context
	err := c.Bind(&body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "There was an error binding the body",
		})

		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "There was an error hashing the password",
		})

		return
	}

	age := math.Floor(time.Now().Sub(*body.Dob).Hours() / 24 / 365)
	// Store the new User
	newUser := models.User{FirstName: body.FirstName, Dob: body.Dob, Age: int(age), LastName: body.LastName, Username: body.Username, Email: body.Email, Password: string(hashedPassword)}

	result := initializers.DB.Create(&newUser)

	InsertErr := result.Error

	if InsertErr != nil {

		println(InsertErr.Error())
		if InsertErr.Error() == `ERROR: duplicate key value violates unique constraint "uni_users_username" (SQLSTATE 23505)` {

			c.JSON(http.StatusBadRequest, gin.H{
				"message": "User already exists",
			})

			return
		}
	}

	c.JSON(http.StatusOK, gin.H{})

	// Return the response
}

// Login the new User
func Login(c *gin.Context) {

	// Get the user's info
	var body dto.UserLoginDto

	var user models.User

	c.Bind(&body)

	// Check if the user exist
	initializers.DB.Where("email = ?", body.Email).First(&user)
	fmt.Println("This is the user's email :", body.Email)

	if user.ID == 0 {

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "User Does not exist",
		})

		return
	}

	// Check if the password is correct
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Wrong credentials",
		})

		return
	}

	expirationTime := time.Now().Add(24 * time.Hour).Unix() // 24 hours from now in Unix timestamp

	// Generate the jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": float64(expirationTime),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error generating the JWT",
		})

		return
	}

	// Return the Response

	// Return the token as a Cookie
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{})
}
