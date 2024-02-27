package controllers

import (
	"database/sql"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Transaction struct {
	Id            int     `json:"id" gorm:"primaryKey;autoIncrement"`
	ParentId      int     `json:"parentid" gorm:"not null"`
	Title         string  `json:"title"`
	Description   string  `json:"description"`
	Currency      string  `json:"currency"`
	PaymentMethod string  `json:"paymentmethod"`
	Amount        float64 `json:"amount" gorm:"default:0"`
	Participant   string  `json:"participant"`
	Recurring     bool    `json:"recurring"`
	Interval      string  `json:"interval"`
	Category      string  `json:"category"`
	Tax           float64 `json:"tax" gorm:"default:0"`
	Taxxed        bool    `json:"taxxed"`
	Fulfilled     bool    `json:"fulfilled"`
	DateCreated   string  `json:"datecreated"`
	FileURL       string  `json:"fileurl"`
	Table         string  `json:"table"`
}

func GetTransactions(ctx *fiber.Ctx) error {
	var transactionType = ctx.Params("type")

	UserId, err := VerifySessionID(ctx)
	if err == sql.ErrNoRows || err == http.ErrNoCookie {
		SendResponse(ctx, http.StatusNotFound, "error", "Couldn't find cookie or session ID")
		return err
	} else if err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "error", err.Error())
		return err
	}
	var transactions []Transaction

	// Select every row that has the ParentId = to the respective UserId
	err = DB.Table(transactionType).Where("parent_id = ?", UserId).Find(&transactions).Error
	if err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "error", err.Error())
		return err
	}

	// Send transactions back to the client
	return SendResponse(ctx, http.StatusOK, "success", transactions)
}

func AddTransaction(ctx *fiber.Ctx) error {
	var transactionType = ctx.Params("type")
	var transaction Transaction
	var err error = nil
	err = ctx.BodyParser(&transaction)
	if err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "error", err.Error())
		return err
	}

	transaction.ParentId, err = VerifySessionID(ctx)
	if err == sql.ErrNoRows || err == http.ErrNoCookie {
		SendResponse(ctx, http.StatusNotFound, "error", "Couldn't find cookie or session ID")
		return err
	} else if err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "error", err.Error())
		return err
	}

	// Insert the transaction into the database
	err = DB.Table(transactionType).Create(&transaction).Error

	if err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "error", err.Error())
		return err
	}

	// Set status to OK and send success msg to user
	return SendResponse(ctx, http.StatusCreated, "success", "Succesfully added "+transactionType[:len(transactionType)-1]+"!")
}

func DeleteTransaction(ctx *fiber.Ctx) error {
	var transactionType = ctx.Params("type")
	var transaction Transaction
	// Decode body into transaction var
	err := ctx.BodyParser(&transaction)
	if err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "error", err.Error())
		return err
	}

	transaction.ParentId, err = VerifySessionID(ctx)
	if err == sql.ErrNoRows || err == http.ErrNoCookie {
		// Error when the session ID doesn't exist in the db or cookie isn't found
		SendResponse(ctx, http.StatusNotFound, "error", "Couldn't find cookie or session ID")
		return err
	} else if err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "error", err.Error())
		return err
	}

	// Delete the transaction that has the coresponding ID and parentid (respective user id)
	err = DB.Table(transactionType).Where("id = ? AND parent_id = ?", transaction.Id, transaction.ParentId).Delete(&Transaction{}).Error
	if err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "error", err.Error())
		return err
	}

	// Set status to OK and send success msg
	return SendResponse(ctx, http.StatusOK, "success", "Succesfully deleted "+transactionType[:len(transactionType)-1]+"!")
}

func EditTransaction(ctx *fiber.Ctx) error {
	var transactionType = ctx.Params("type")
	var transaction Transaction
	err := ctx.BodyParser(&transaction)
	if err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "error", err.Error())
		return err
	}

	transaction.ParentId, err = VerifySessionID(ctx)
	if err == sql.ErrNoRows || err == http.ErrNoCookie {
		SendResponse(ctx, http.StatusNotFound, "error", "Couldn't find cookie or session ID")
		return err
	} else if err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "error", err.Error())
		return err
	}
	err = DB.Table(transactionType).Where("id = ? AND parent_id = ?", transaction.Id, transaction.ParentId).Updates(&transaction).Error
	if err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "error", err.Error())
		return err
	}

	return SendResponse(ctx, http.StatusOK, "success", "Succesfully edited "+transactionType[:len(transactionType)-1]+"!")
}
