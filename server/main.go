package main

import (
	"net/http"
	"server/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

var PORT = "8000"

func main() {
	router := fiber.New()
	err := godotenv.Load("../.env")
	if err != nil {
		panic(err)
	}

	err = controllers.Main()
	if err != nil {
		panic(err)
	}

	router.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173",
		AllowMethods:     "GET, POST, PATCH, DELETE",
		AllowHeaders:     "Origin, Content-Type",
		AllowCredentials: true,
	}))

	router.Post("/register", controllers.Register)
	router.Post("/login", controllers.Login)
	// Categories
	router.Get("/categories", controllers.GetCategories)
	router.Post("/categories", controllers.AddCategory)
	router.Patch("/categories", controllers.EditCategory)
	router.Delete("/categories", controllers.DeleteCategory)
	// Transactions
	router.Get("/transactions/:type", controllers.GetTransactions)
	router.Post("/transactions/:type", controllers.AddTransaction)
	router.Patch("/transactions/:type", controllers.EditTransaction)
	router.Delete("/transactions/:type", controllers.DeleteTransaction)
	// Cookies
	router.Delete("/logout", func(ctx *fiber.Ctx) error {
		err := controllers.Logout(ctx)
		if err != nil {
			controllers.SendResponse(ctx, http.StatusInternalServerError, "error", err.Error())
			return err
		}
		return controllers.SendResponse(ctx, http.StatusOK, "success", "Succesfully logged out!")
	})

	err = router.Listen(":" + PORT)
	if err != nil {
		panic(err)
	}
}
