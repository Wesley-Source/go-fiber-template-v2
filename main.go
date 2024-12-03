package main

import (
	"go-fiber-template-v2/app/database"
	"go-fiber-template-v2/app/middleware"
	"go-fiber-template-v2/app/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
)

func main() {
	database.ConnectDatabase()
	middleware.ConnectSessionsDB()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	engine := html.New("./app/views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Corrigindo o caminho dos arquivos estáticos
	app.Static("/", "./app/public")

	setupRoutes(app)

	log.Fatal(app.Listen(os.Getenv("PORT")))
}

func setupRoutes(app *fiber.App) {
	app.Get("/", middleware.Auth, routes.Index)

	app.Get("/login", middleware.Auth, routes.Login)
	app.Post("/login", middleware.Auth, routes.Login)

	app.Get("/register", middleware.Auth, routes.Register)
	app.Post("/register", middleware.Auth, routes.Register)

	app.Get("/logout", middleware.Auth, routes.Logout)

	app.Use(routes.UnknownRoute)
}
