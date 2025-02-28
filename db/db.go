package db

import (
	"os"
	"fmt"
	"gorm.io/gorm"
	"log"
	"gorm.io/driver/postgres"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "host=localhost user=postgres password=yourpassword dbname=dialecto port=5432 sslmode=disable"
	}

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal("Failed to get DB instance:", err)
	}
	if err = sqlDB.Ping(); err != nil {
		log.Fatal("Database ping failed:", err)
	}

	fmt.Println("Connected to PostgreSQL!")
}