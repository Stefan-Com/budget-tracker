package controllers

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Session struct {
	ID        string
	UserID    int
	CreatedAt time.Time
	ExpiresAt time.Time
}

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
	err = DB.Table("sessions").Create(&Session{ID: ID, UserID: UserID, ExpiresAt: time.Now().Add(7 * 24 * time.Hour)}).Error
	if err != nil {
		return err
	}

	return nil
}

func DeleteOldSession(ID string) error {
	err := DB.Table("sessions").Delete(&Session{}, "id = ?", ID).Error
	if err != nil {
		return err
	}
	return nil
}

func VerifySessionID(ctx *fiber.Ctx) (int, error) {
	// Get cookie
	cookie := ctx.Cookies("logged_in_cookie")

	// If cookie isn't found, return -1 as id
	if cookie == "" {
		return -1, http.ErrNoCookie
	}

	var session Session

	// Select the UserID and Expiring date from db
	err := DB.Table("sessions").First(&session, "id = ?", cookie).Error

	if err != nil {
		return -1, err
	}

	// Check if session ID is expired
	if session.ExpiresAt.Before(time.Now()) {
		return -1, errors.New("expired session ID")
	}

	return session.UserID, nil
}
