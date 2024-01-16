package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Email    string  `json:"email" validate:"email"`
	Password string  `json:"password"`
	Username string  `json:"username"`
	Currency string  `json:"currency"`
	Balance  float64 `json:"balance"`
}

func Register(ctx *gin.Context) {
	var user User
	// Decode body into user var
	err := ctx.BindJSON(&user)
	if err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "error", err.Error())
		return
	}

	var emailExists int

	// Check if email already exists
	err = database.QueryRow("SELECT COUNT(*) FROM users WHERE Email = ?", user.Email).Scan(&emailExists)
	if err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "error", err.Error())
		return
	}
	if emailExists >= 1 {
		SendResponse(ctx, http.StatusNotAcceptable, "error", "Email already exists")
		return
	}

	// If the email is distinct, hash password and create new account
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "error", err.Error())
		return
	}

	// Insert account into database
	_, err = database.Exec("INSERT INTO users (Email, Password, Username, Currency, Balance) VALUES (?, ?, ?, ?, ?)", user.Email, hashedPassword, user.Username, user.Currency, user.Balance)
	if err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "error", err.Error())
		return
	}

	SendResponse(ctx, http.StatusCreated, "success", "Sucesfully registred account!")
}
