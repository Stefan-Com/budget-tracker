package controllers

import (
	"database/sql"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Category struct {
	Id          int     `json:"id" gorm:"primaryKey autoIncrement"`
	ParentId    int     `json:"parentid" gorm:"not null"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Currency    string  `json:"currency"`
	Budget      float64 `json:"budget" gorm:"default:0"`
	Spent       float64 `json:"spent" gorm:"default:0"`
	Gotten      float64 `json:"gotten" gorm:"default:0"`
	Type        string  `json:"type" gorm:"default:income"`
	Budgeted    bool    `json:"budgeted"`
}

func GetCategories(ctx *fiber.Ctx) error {
	var categories []Category

	// Get user ID from session ID
	UserId, err := VerifySessionID(ctx)
	if err == http.ErrNoCookie || err == sql.ErrNoRows {
		SendResponse(ctx, http.StatusNotFound, "error", "Could find cookie or session ID")
		return err
	} else if err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "error", err.Error())
		return err
	}

	// Select categories from db with same ParentId as the UserId
	err = DB.Table("categories").Find(&categories, "parent_id = ?", UserId).Error
	if err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "error", err.Error())
		return err
	}

	// Send data back to the client
	return SendResponse(ctx, http.StatusOK, "success", categories)
}

func AddCategory(ctx *fiber.Ctx) error {
	var category Category

	// Decode body into category var, using a ptr
	// err := ctx.BindJSON(&category)
	err := ctx.BodyParser(&category)
	if err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "error", err.Error())
		return err
	}

	category.ParentId, err = VerifySessionID(ctx)
	if err == sql.ErrNoRows || err == http.ErrNoCookie {
		// Error when the session ID doesn't exist in the db or cookie is not found
		SendResponse(ctx, http.StatusNotFound, "error", "Couldn't find cookie or session ID")
		return err
	} else if err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "error", err.Error())
		return err
	}

	// Insert category into database with placeholders
	err = DB.Table("categories").Create(&category).Error
	if err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "error", err.Error())
		return err
	}

	// Send response to the client
	return SendResponse(ctx, http.StatusCreated, "success", "Succesfully added category!")
}

func DeleteCategory(ctx *fiber.Ctx) error {
	var category Category
	err := ctx.BodyParser(&category)
	if err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "error", err.Error())
		return err
	}

	category.ParentId, err = VerifySessionID(ctx)
	if err == sql.ErrNoRows || err == http.ErrNoCookie {
		// Error when the session ID doesn't exist in the db or cookie is not found
		SendResponse(ctx, http.StatusNotFound, "error", "Couldn't find cookie or session ID")
		return err
	} else if err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "error", err.Error())
		return err
	}

	// Delete from table by id and parentid (respective user's id)
	err = DB.Table("categories").Delete(&category).Error
	if err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "error", err.Error())
		return err
	}

	// Send response to the client
	return SendResponse(ctx, http.StatusOK, "success", "Succesfully deleted category!")
}

func EditCategory(ctx *fiber.Ctx) error {
	var category Category
	err := ctx.BodyParser(&category)
	if err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "error", err.Error())
		return err
	}

	category.ParentId, err = VerifySessionID(ctx)
	if err == sql.ErrNoRows || err == http.ErrNoCookie {
		SendResponse(ctx, http.StatusNotFound, "error", "Couldn't find cookie or session ID")
		return err
	} else if err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "error", err.Error())
		return err
	}

	err = DB.Table("categories").Where("id = ? AND parent_id = ?", category.Id, category.ParentId).Updates(&category).Error
	if err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "error", err.Error())
		return err
	}

	return SendResponse(ctx, http.StatusOK, "success", "Succesfully edited category!")
}
