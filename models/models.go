package models

import (
	"fmt"

	"github.com/tecolotedev/stori_back/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {

	// Create dns to connect with the database
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		config.EnvVars.DB_HOST,
		config.EnvVars.DB_USER,
		config.EnvVars.DB_PASSWORD,
		config.EnvVars.DB_NAME,
		config.EnvVars.DB_PORT,
		config.EnvVars.DB_SSLMODE,
	)

	fmt.Println("initiating db")

	// test connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// migrate models
	db.AutoMigrate(&Newsletter{})

	DB = db
}
