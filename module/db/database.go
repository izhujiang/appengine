package db

import (
	"fmt"
	"log"
	"os"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Instance *gorm.DB

func Connect() (*gorm.DB, error) {
	if Instance != nil {
		return Instance, nil
	}

	var err error

	driver := os.Getenv("DB_DRIVER")
	host := os.Getenv("DB_HOST")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	databaseName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := "test.db" // default, using sqlite
	switch strings.ToLower(driver) {
	case "postgres":
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, username, password, databaseName, port)
	}
	Instance, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	} else {
		log.Println("Successfully connected to the database: ", dsn)
	}

	return Instance, err
}
