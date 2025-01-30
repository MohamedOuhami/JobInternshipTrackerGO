package initializers

import (
	"fmt"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
    panic("Error loading the .env file")
	}

  fmt.Println("Loaded the env succesfully")
}
