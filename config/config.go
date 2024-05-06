package config

import "os"

var EnvVars struct {
	DB_HOST     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
	DB_PORT     string
	DB_SSLMODE  string
}

func InitConfig() {

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

}
