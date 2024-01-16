package controllers

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var database *sql.DB

func DB() (*sql.DB, error) {
	var user = os.Getenv("DATABASE_USER")
	var password = os.Getenv("DATABASE_PASSWORD")
	var url = os.Getenv("DATABASE_URL")
	var name = os.Getenv("DATABASE_NAME")
	var statement = fmt.Sprintf("%v:%v@tcp(%v)/%v", user, password, url, name)
	// Open database connection
	db, err := sql.Open("mysql", statement)
	if err != nil {
		return nil, err
	}
	// Test database connection
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	// Return the database and no error
	return db, nil
}

// Function so set up the database
func Main() error {
	var err error
	database, err = DB()
	if err != nil {
		defer database.Close()
		return err
	}
	return nil
}
