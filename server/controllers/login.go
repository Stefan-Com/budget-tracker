package controllers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	var user User

	// Decode body into user var
	err := ctx.BindJSON(&user)
	if err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "error", err.Error())
		return
	}

	// Check if email exists
	row := database.QueryRow("SELECT UserId, Password FROM users WHERE Email = ?", user.Email)
	var UserID int
	var HashedPass string
	err = row.Scan(&UserID, &HashedPass)

	// Check if there aren't any rows found
	if err != nil && err != sql.ErrNoRows {
		SendResponse(ctx, http.StatusInternalServerError, "error", err.Error())
		return
	}

	// Check if the passwords match
	if !ComparePassword(user.Password, HashedPass) || err == sql.ErrNoRows {
		SendResponse(ctx, http.StatusBadRequest, "error", "Invalid email or password")
		return
	}

	// Generate a session ID and put it into the database
	sessionId, err := GenerateSessionId(32)
	if err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "error", err.Error())
		return
	}

	err = AddSessionID(sessionId, UserID)
	if err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "error", err.Error())
		return
	}

	err = Logout(ctx)
	ctx.SetCookie("logged_in_cookie", sessionId, 60*60*24*7, "/", "localhost", true, true)

	if err != http.ErrNoCookie && err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "error", err.Error())
		return
	}

	// Send success message
	SendResponse(ctx, http.StatusOK, "success", "Sucesfully logged in!")
}