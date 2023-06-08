package config

import (
	"fmt"
	"os"

	"github.com/mstgnz/services/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// OpenDatabase is creating a new connection to our database
func OpenDatabase() *gorm.DB {

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbZone := os.Getenv("DB_ZONE")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=%s", dbHost, dbPort, dbUser, dbPass, dbName, dbZone)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to create a connection to database")
	}
	_ = db.AutoMigrate(&entity.Blog{}, &entity.Comment{})
	return db
}

// CloseDatabase method is closing a connection between your app and your db
func CloseDatabase(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}
	_ = dbSQL.Close()
}
