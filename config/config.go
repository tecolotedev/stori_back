package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type EnvVarsType struct {
	PORT string
}

var EnvVars EnvVarsType

// Initiate env vars
func SetUpConfig() {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Error loading .env file: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	EnvVars.PORT = port
}
