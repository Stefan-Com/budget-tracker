package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Logout(ctx *gin.Context) error {
	// Get the value of the cookie
	cookie, err := ctx.Cookie("logged_in_cookie")
	if err == http.ErrNoCookie || err != nil {
		return err
	}
	err = DeleteOldSession(cookie)
	if err != nil {
		return err
	}

	// Remove cookie
	ctx.SetCookie("logged_in_cookie", "", -1, "/", "localhost", true, true)
	return nil
}
