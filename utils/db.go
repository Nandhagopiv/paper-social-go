package utils

import (
    "github.com/Nandhagopiv/paper-social-go/models" // Import the models package
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "log"
)

var DB *gorm.DB

// ConnectDatabase establishes the connection to the PostgreSQL database
func ConnectDatabase() {
    // DSN (Data Source Name) for PostgreSQL
    dsn := "host=localhost user=postgres password=pass-paper-social dbname=paper-social port=5432 sslmode=disable"
    // Connect to PostgreSQL using GORM
    database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    // Assign the DB instance to a global variable
    DB = database
    // Auto migrate the models (creates the tables)
    DB.AutoMigrate(&models.User{}, &models.Tweet{})
}
