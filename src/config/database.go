package config

import (
	"fmt"
	"os"
	"time"

	"ticket-api/src/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupDatabase() *gorm.DB {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPass, dbHost, dbPort, dbName)

	var db *gorm.DB
	var err error
	maxRetries := 10

	for i := 1; i <= maxRetries; i++ {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			sqlDB, errPing := db.DB()
			if errPing == nil && sqlDB.Ping() == nil {
				fmt.Println("✅ Connected to database.")
				break
			}
		}

		fmt.Printf("🔁 Retry connecting to DB (%d/%d): %v\n", i, maxRetries, err)
		time.Sleep(3 * time.Second)
	}

	if err != nil {
		panic("❌ Failed to connect to database after retries: " + err.Error())
	}

	// migrate
	err = db.AutoMigrate(
		&entity.User{},
		&entity.Category{},
		&entity.Event{},
		&entity.Ticket{},
	)
	if err != nil {
		panic("Failed to migrate database: " + err.Error())
	}

	DB = db
	return db
}
