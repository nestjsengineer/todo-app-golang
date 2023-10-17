package db

import (
	"log"
	"os"
	"sync"
	"todo-app/pkg/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

func InitDB() (*gorm.DB, error) {
	once.Do(func() {
		err := godotenv.Load() // Load the .env file
		if err != nil {
			log.Fatal("Error loading .env file")
		}

		host := os.Getenv("DB_HOST")
		port := os.Getenv("DB_PORT")
		user := os.Getenv("DB_USER")
		password := os.Getenv("DB_PASSWORD")
		dbname := os.Getenv("DB_NAME")
		sslMode := os.Getenv("DB_SSL_MODE")

		dsn := "host=" + host + " port=" + port + " user=" + user + " password=" + password +
			" dbname=" + dbname + " sslmode=" + sslMode

		// var err error
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err) // or handle the error as needed
		}
		db.AutoMigrate(&models.Todo{})
	})

	return db, nil
}

func GetDB() *gorm.DB {
	return db
}

func AutoMigrate() {
	db.AutoMigrate(&models.Todo{})
}
