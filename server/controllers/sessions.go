package controllers

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
)

func GenerateSessionId(length int) (string, error) {
	// Make an empty slice of bytes
	bytes := make([]byte, length)

	// Fill it with random bytes that represent chars
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	// Turn the slice into a string
	return base64.URLEncoding.EncodeToString(bytes)[:length], nil
}
func AddSessionID(ID string, UserID int) error {
	err := DeleteOldSession(ID)
	if err != nil {
		return err
	}

	// Insert a new session into the DB
	_, err = database.Exec("INSERT INTO sessions (ID, UserId, ExpiresAt) VALUES (?, ?, ?)", ID, UserID, time.Now().Add(7*24*time.Hour))
	if err != nil {
		return err
	}

	return nil
}
func DeleteOldSession(ID string) error {
	_, err := database.Exec("DELETE FROM sessions WHERE ID = ?", ID)
	if err != nil {
		return err
	}
	return nil
}
func VerifySessionID(ctx *gin.Context) (int, error) {
	// Get cookie
	cookie, err := ctx.Cookie("logged_in_cookie")

	// If cookie isn't found, return -1 as id
	if err != nil {
		return -1, err
	}

	var UserId int
	var ExpiresAtStr string

	// Select the UserID and Expiring date from db
	row := database.QueryRow("SELECT UserId, ExpiresAt FROM sessions WHERE ID = ?", cookie)
	err = row.Scan(&UserId, &ExpiresAtStr)

	if err != nil {
		return -1, err
	}

	dateFormat := "2006-01-02 15:04:05"

	// Parse the expiration date into a time.Time data type
	ExpiresAt, err := time.Parse(dateFormat, ExpiresAtStr)
	if err != nil {
		return -1, err
	}
	// Check if session ID is expired
	if ExpiresAt.Before(time.Now()) {
		return -1, errors.New("Expired session ID")
	}

	return UserId, nil
}
