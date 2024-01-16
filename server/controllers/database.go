package controllers

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var database *sql.DB

func DB() (*sql.DB, error) {
	var user, password, url, name = os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASSWORD"), os.Getenv("URL"), os.Getenv("DATABASE_NAME")
	var statement = user + ":" + password + "@tcp(" + url + ")/" + name
	// Open database connection
	db, err := sql.Open("mysql", statement)
	if err != nil {
		return nil, fmt.Errorf("Failed to open database, \n%v\n ", err)
	}
	// Test database connection
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("Failed to ping database, \n%v\n", err)
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
		fmt.Printf("There has been an error: %v\n", err)
		return err
	}
	return nil
}
