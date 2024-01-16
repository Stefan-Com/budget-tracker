package main

import (
	"net/http"
	"server/controllers"

	"github.com/gin-gonic/gin"
)

var PORT = "8000"

func main() {
	router := gin.Default()
	err := controllers.Main()
	if err != nil {
		panic(err)
	}

	router.Use(CORSMiddleware())

	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)
	// Categories
	router.GET("/categories", controllers.GetCategories)
	router.POST("/categories", controllers.AddCategory)
	router.PATCH("/categories", controllers.EditCategory)
	router.DELETE("/categories", controllers.DeleteCategory)
	// Transactions
	router.PUT("/transactions", controllers.GetTransactions)
	router.POST("/transactions", controllers.AddTransaction)
	router.PATCH("/transactions", controllers.EditTransaction)
	router.DELETE("/transactions", controllers.DeleteTransaction)
	// Cookies
	router.DELETE("/logout", func(ctx *gin.Context) {
		err := controllers.Logout(ctx)
		if err != nil {
			controllers.SendResponse(ctx, http.StatusInternalServerError, "error", err.Error())
			return
		}
		controllers.SendResponse(ctx, http.StatusOK, "success", "Succesfully logged out!")
	})

	//TODO: add CORS
	// CORSHandler := cors.New(cors.Options{
	// 	AllowedOrigins:   []string{"http://localhost:5173"},
	// 	AllowedMethods:   []string{"GET", "POST", "PATCH", "DELETE"},
	// 	AllowCredentials: true,
	// 	Debug:            true,
	// })

	// router.Use(CORSHandler())

	err = router.Run(":" + PORT)
	if err != nil {
		panic(err)
	}
}
