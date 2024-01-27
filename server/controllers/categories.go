package controllers

import (
	"database/sql"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Category struct {
	Id          int     `json:"id"`
	ParentId    int     `json:"parentid"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Currency    string  `json:"currency"`
	Budget      float64 `json:"budget"`
	Spent       float64 `json:"spent"`
	Gotten      float64 `json:"gotten"`
	Type        string  `json:"type"`
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
	rows, err := database.Query("SELECT * FROM categories WHERE ParentId = ?", UserId)
	if err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "error", err.Error())
		return err
	}

	// Close rows
	defer rows.Close()

	// Loop over the rows and put each of them in the "categories" slice
	for rows.Next() {
		var category Category
		err = rows.Scan(&category.Id, &category.ParentId, &category.Title, &category.Description, &category.Currency, &category.Budget, &category.Spent, &category.Gotten, &category.Type, &category.Budgeted)
		if err != nil {
			SendResponse(ctx, http.StatusInternalServerError, "error", err.Error())
			return err
		}
		categories = append(categories, category)
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
	_, err = database.Exec("INSERT INTO categories (ParentId, Title, Description, Currency, Budget, Spent, Gotten, Type, Budgeted) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)",
		&category.ParentId, &category.Title,
		&category.Description, &category.Currency,
		&category.Budget, &category.Spent,
		&category.Gotten, &category.Type,
		&category.Budgeted)
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
	_, err = database.Exec("DELETE FROM categories WHERE Id = ? AND ParentId = ?", category.Id, category.ParentId)
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

	_, err = database.Exec("UPDATE categories SET Title = ?, Description = ?, Currency = ?, Budget = ?, Spent = ?, Gotten = ?, Type = ?, Budgeted = ? WHERE Id = ? AND ParentId = ?",
		category.Title, category.Description, category.Currency, category.Budget, category.Spent, category.Gotten, category.Type, category.Budgeted, category.Id, category.ParentId)
	if err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "error", err.Error())
		return err
	}

	return SendResponse(ctx, http.StatusOK, "success", "Succesfully edited category!")
}
