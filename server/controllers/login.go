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

	//! First() method only works if the dbdata struct has the UserID and Password attribute with these exact names
	//TODO: maybe find out why the fuck that happens
	var dbdata = struct {
		UserID   int
		Password string
	}{}

	err = DB.Table("users").Where("email = ?", user.Email).Select("user_id", "password").First(&dbdata).Error

	// Check if email exists
	if err != nil && err != sql.ErrNoRows {
		SendResponse(ctx, http.StatusInternalServerError, "error", err.Error())
		return err
	}

	// Check if the passwords match
	if !ComparePassword(user.Password, dbdata.Password) {
		SendResponse(ctx, http.StatusBadRequest, "error", "Invalid email or password")
		return err
	}

	// Generate a session ID and put it into the database
	sessionId, err := GenerateSessionId(32)
	if err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "error", err.Error())
		return err
	}

	err = AddSessionID(sessionId, dbdata.UserID)
	if err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "error", err.Error())
		return err
	}

	err = DeleteSessionAndCookie(ctx)
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
