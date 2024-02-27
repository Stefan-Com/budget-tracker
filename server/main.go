package main

import (
	"log"
	"server/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

var PORT = "8000"

func main() {
	router := fiber.New()
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	controllers.DB, err = controllers.InitDB()
	if err != nil {
		log.Fatalf("There has been an error while initalizing database: \n%v\n", err)
	}

	// Auto migrate structs into tables
	controllers.DB.Table("categories").AutoMigrate(&controllers.Category{})
	controllers.DB.Table("expenses").AutoMigrate(&controllers.Transaction{})
	controllers.DB.Table("incomes").AutoMigrate(&controllers.Transaction{})
	controllers.DB.Table("sessions").AutoMigrate(&controllers.Session{})
	controllers.DB.Table("users").AutoMigrate(&controllers.User{})

	router.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:9000",
		AllowMethods:     "GET, POST, PATCH, DELETE",
		AllowHeaders:     "Origin, Content-Type",
		AllowCredentials: true,
	}))

	api := router.Group("/api")

	//Authenctication
	api.Post("/register", controllers.Register)
	api.Post("/login", controllers.Login)
	api.Delete("/logout", controllers.Logout)

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

	err = router.Listen(":" + PORT)
	if err != nil {
		log.Panicf("There has been an error while starting the server: \n%v\n", err)
	}
}
