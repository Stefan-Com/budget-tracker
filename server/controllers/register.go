package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	UserId   int     `json:"userid" gorm:"primaryKey;autoIncrement"`
	Email    string  `json:"email" validate:"email" gorm:"unique;not null;"`
	Password string  `json:"password" gorm:"not null"`
	Username string  `json:"username" gorm:"not null"`
	Currency string  `json:"currency" gorm:"not null"`
	Balance  float64 `json:"balance" gorm:"default:0"`
}

func Register(ctx *fiber.Ctx) error {
	var user User
	// Decode body into user var
	err := ctx.BodyParser(&user)
	if err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "error", err.Error())
		return err
	}

	var emailExists int64

	// Check if email already exists
	err = DB.Table("users").Where("email = ?", user.Email).Count(&emailExists).Error
	if err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "error", err.Error())
		return fiber.ErrBadRequest
	}
	if emailExists >= 1 {
		SendResponse(ctx, http.StatusNotAcceptable, "error", "Email already exists")
		return err
	}

	// If the email is distinct, hash password and create new account
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "error", err.Error())
		return err
	}

	// Insert account into database
	err = DB.Table("users").Create(&User{Email: user.Email, Password: hashedPassword, Username: user.Username, Currency: user.Currency, Balance: user.Balance}).Error
	if err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "error", err.Error())
		return err
	}

	return SendResponse(ctx, http.StatusCreated, "success", "Sucesfully registred account!")
}
