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
		AllowOrigins:     "http://localhost:9000",
		AllowMethods:     "GET, POST, PATCH, DELETE",
		AllowHeaders:     "Origin, Content-Type",
		AllowCredentials: true,
	}))

	api := router.Group("/api")
	api.Post("/register", controllers.Register)
	api.Post("/login", controllers.Login)
	// Categories
	api.Get("/categories", controllers.GetCategories)
	api.Post("/categories", controllers.AddCategory)
	api.Patch("/categories", controllers.EditCategory)
	api.Delete("/categories", controllers.DeleteCategory)
	// Transactions
	api.Get("/transactions/:type", controllers.GetTransactions)
	api.Post("/transactions/:type", controllers.AddTransaction)
	api.Patch("/transactions/:type", controllers.EditTransaction)
	api.Delete("/transactions/:type", controllers.DeleteTransaction)
	// Cookies
	api.Delete("/logout", func(ctx *fiber.Ctx) error {
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
