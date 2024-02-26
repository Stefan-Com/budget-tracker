package controllers

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	// Get env vars
	var user = os.Getenv("DB_USER")
	var password = os.Getenv("DB_PASSWORD")
	var host = os.Getenv("DB_HOST")
	var port = os.Getenv("DB_PORT")
	var name = os.Getenv("DB_NAME")
	var sslmode = os.Getenv("SSLMODE")

	var dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, user, password, name, port, sslmode)
	// Open database connection
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return nil, err
	}
	// Return the database and no error
	return db, nil
}
