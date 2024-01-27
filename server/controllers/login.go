package controllers

import (
	"database/sql"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func Login(ctx *fiber.Ctx) error {
	var user User

	// Decode body into user var
	err := ctx.BodyParser(&user)
	if err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "error", err.Error())
		return err
	}

	// Check if email exists
	row := database.QueryRow("SELECT UserId, Password FROM users WHERE Email = ?", user.Email)
	var UserID int
	var HashedPass string
	err = row.Scan(&UserID, &HashedPass)

	// Check if there aren't any rows found
	if err != nil && err != sql.ErrNoRows {
		SendResponse(ctx, http.StatusInternalServerError, "error", err.Error())
		return err
	}

	// Check if the passwords match
	if !ComparePassword(user.Password, HashedPass) || err == sql.ErrNoRows {
		SendResponse(ctx, http.StatusBadRequest, "error", "Invalid email or password")
		return err
	}

	// Generate a session ID and put it into the database
	sessionId, err := GenerateSessionId(32)
	if err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "error", err.Error())
		return err
	}

	err = AddSessionID(sessionId, UserID)
	if err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "error", err.Error())
		return err
	}

	err = Logout(ctx)
	ctx.Cookie(&fiber.Cookie{
		Name:     "logged_in_cookie",
		Value:    sessionId,
		MaxAge:   60 * 60 * 24 * 7,
		Path:     "/",
		Domain:   "localhost",
		Secure:   true,
		HTTPOnly: true,
	})

	if err != http.ErrNoCookie && err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "error", err.Error())
		return err
	}

	// Send success message
	return SendResponse(ctx, http.StatusOK, "success", "Sucesfully logged in!")
}
