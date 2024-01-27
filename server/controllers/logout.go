package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func Logout(ctx *fiber.Ctx) error {
	// Get the value of the cookie
	cookie := ctx.Cookies("logged_in_cookie")
	if cookie == "" {
		return http.ErrNoCookie
	}
	err := DeleteOldSession(cookie)
	if err != nil {
		return err
	}

	// Remove cookie
	ctx.Cookie(&fiber.Cookie{
		Name:     "logged_in_cookie",
		Value:    "",
		MaxAge:   -1,
		Path:     "/",
		Domain:   "localhost",
		Secure:   true,
		HTTPOnly: true,
	})
	return nil
}
