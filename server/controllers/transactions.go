package controllers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Transaction struct {
	Table         string  `json:"table"`
	Id            int     `json:"id"`
	ParentId      int     `json:"parentid"`
	Title         string  `json:"title"`
	Description   string  `json:"description"`
	Currency      string  `json:"currency"`
	PaymentMethod string  `json:"paymentmethod"`
	Amount        float64 `json:"amount"`
	Participant   string  `json:"participant"`
	Recurring     bool    `json:"recurring"`
	Interval      string  `json:"interval"`
	Category      string  `json:"category"`
	Tax           float64 `json:"tax"`
	Taxxed        bool    `json:"taxxed"`
	Fulfilled     bool    `json:"fulfilled"`
	DateCreated   string  `json:"datecreated"`
	FileURL       string  `json:"fileurl"`
}

func GetTransactions(ctx *gin.Context) {
	var transactionType = ctx.Param("type")

	UserId, err := VerifySessionID(ctx)
	if err == sql.ErrNoRows || err == http.ErrNoCookie {
		SendResponse(ctx, http.StatusNotFound, "error", "Couldn't find cookie or session ID")
		return
	} else if err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "error", err.Error())
		return
	}
	var transactions []Transaction
	// Select every row that has the ParentId = to the respective UserId
	rows, err := database.Query("SELECT * FROM "+transactionType+" WHERE ParentId = ?", UserId)
	if err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "error", err.Error())
		return
	}

	// Loop over the selected rows
	for rows.Next() {
		var transaction Transaction
		// Get the values from each column
		err = rows.Scan(
			&transaction.Id, &transaction.ParentId, &transaction.Title,
			&transaction.Description, &transaction.Currency, &transaction.PaymentMethod,
			&transaction.Amount, &transaction.DateCreated, &transaction.Participant,
			&transaction.Recurring, &transaction.Interval, &transaction.Category,
			&transaction.FileURL, &transaction.Taxxed, &transaction.Tax, &transaction.Fulfilled,
		)
		transaction.Table = transactionType
		if err != nil {
			SendResponse(ctx, http.StatusInternalServerError, "error", err.Error())
			return
		}

		transactions = append(transactions, transaction)
	}

	// Send transactions back to the client
	SendResponse(ctx, http.StatusOK, "success", transactions)
}

func AddTransaction(ctx *gin.Context) {
	var transactionType = ctx.Param("type")
	var transaction Transaction
	var err error = nil
	err = ctx.BindJSON(&transaction)
	if err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "error", err.Error())
		return
	}

	transaction.ParentId, err = VerifySessionID(ctx)
	if err == sql.ErrNoRows || err == http.ErrNoCookie {
		SendResponse(ctx, http.StatusNotFound, "error", "Couldn't find cookie or session ID")
		return
	} else if err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "error", err.Error())
		return
	}

	var query string = "INSERT INTO " + transactionType + " (ParentId, Title, Description, Currency, PaymentMethod, Amount, Participant, Recurring, `Interval`, Category, FileURL, Taxxed, Tax, Fulfilled) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	// Insert the transaction into the database
	_, err = database.Exec(query,
		&transaction.ParentId, &transaction.Title,
		&transaction.Description, &transaction.Currency,
		&transaction.PaymentMethod, &transaction.Amount,
		&transaction.Participant, &transaction.Recurring,
		&transaction.Interval, &transaction.Category,
		&transaction.FileURL, &transaction.Taxxed,
		&transaction.Tax, &transaction.Fulfilled,
	)

	if err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "error", err.Error())
		return
	}

	// Set status to OK and send success msg to user
	SendResponse(ctx, http.StatusCreated, "success", "Succesfully added "+transactionType+"!")
}
func DeleteTransaction(ctx *gin.Context) {
	var transactionType = ctx.Param("type")
	var transaction Transaction
	// Decode body into transaction var
	err := ctx.BindJSON(&transaction)
	if err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "error", err.Error())
		return
	}

	transaction.ParentId, err = VerifySessionID(ctx)
	if err == sql.ErrNoRows || err == http.ErrNoCookie {
		// Error when the session ID doesn't exist in the db or cookie isn't found
		SendResponse(ctx, http.StatusNotFound, "error", "Couldn't find cookie or session ID")
		return
	} else if err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "error", err.Error())
		return
	}

	// Delete the transaction that has the coresponding ID and parentid (respective user id)
	_, err = database.Exec("DELETE FROM "+transaction.Table+" WHERE Id = ? AND ParentId = ?", transaction.Id, transaction.ParentId)
	if err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "error", err.Error())
		return
	}

	// Set status to OK and send success msg
	SendResponse(ctx, http.StatusOK, "success", "Succesfully deleted "+transactionType+"!")
}

func EditTransaction(ctx *gin.Context) {
	var transactionType = ctx.Param("type")
	var transaction Transaction
	err := ctx.BindJSON(&transaction)
	if err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "error", err.Error())
		return
	}

	transaction.ParentId, err = VerifySessionID(ctx)
	if err == sql.ErrNoRows || err == http.ErrNoCookie {
		SendResponse(ctx, http.StatusNotFound, "error", "Couldn't find cookie or session ID")
		return
	} else if err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "error", err.Error())
		return
	}
	fmt.Println(transaction.Table)
	var query string = "UPDATE " + transaction.Table + " SET `Title` = ?, `Description` = ?, `Currency` = ?, `PaymentMethod` = ?, `Amount` = ?, `Participant` = ?, `Recurring` = ?, `Interval` = ?, `Category` = ?, `FileURL` = ?, `Taxxed` = ?, `Tax` = ?, `Fulfilled` = ? WHERE `Id` = ? AND `ParentId` = ?"
	_, err = database.Exec(query,
		transaction.Title, transaction.Description,
		transaction.Currency, transaction.PaymentMethod,
		transaction.Amount, transaction.Participant,
		transaction.Recurring, transaction.Interval,
		transaction.Category, transaction.FileURL,
		transaction.Taxxed, transaction.Tax,
		transaction.Fulfilled, transaction.Id, transaction.ParentId,
	)
	if err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "error", err.Error())
		return
	}

	SendResponse(ctx, http.StatusOK, "success", "Succesfully edited "+transactionType+"!")
}
