package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var EnvVars struct {
	DB_HOST     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
	DB_PORT     string
	DB_SSLMODE  string

	EMAIL_HOST     string
	EMAIL_PORT     string
	EMAIL_USER     string
	EMAIL_PASSWORD string
}

func InitConfig() {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Error loading .env file: %v", err)
	}

	// Database variables
	db_host := os.Getenv("DB_HOST")
	db_user := os.Getenv("DB_USER")
	db_password := os.Getenv("DB_PASSWORD")
	db_name := os.Getenv("DB_NAME")
	db_port := os.Getenv("DB_PORT")
	db_sslmode := os.Getenv("DB_SSLMODE")

	EnvVars.DB_HOST = db_host
	EnvVars.DB_USER = db_user
	EnvVars.DB_PASSWORD = db_password
	EnvVars.DB_NAME = db_name
	EnvVars.DB_PORT = db_port
	EnvVars.DB_SSLMODE = db_sslmode

	// Mail variables
	email_host := os.Getenv("EMAIL_HOST")
	email_port := os.Getenv("EMAIL_PORT")
	email_user := os.Getenv("EMAIL_USER")
	email_password := os.Getenv("EMAIL_PASSWORD")

	EnvVars.EMAIL_HOST = email_host
	EnvVars.EMAIL_PORT = email_port
	EnvVars.EMAIL_USER = email_user
	EnvVars.EMAIL_PASSWORD = email_password

}
