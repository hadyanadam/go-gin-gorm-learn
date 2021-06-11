package config

import (
	"fmt"
	"os"

	"github.com/hadyanadam/go-gorm-gin/entity"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// SetupDatabaseConnection is creating new connection to our database
func SetupDatabaseConnection() *gorm.DB {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", dbHost, dbUser, dbPass, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to create a connection to database")
	}

	db.AutoMigrate(&entity.Book{}, &entity.User{})
	return db
}


func CloseDatabaseConnection(db *gorm.DB) {
	conn, err := db.DB()
	if err != nil {
		panic("Failed to close connection with database")
	}

	conn.Close()
}