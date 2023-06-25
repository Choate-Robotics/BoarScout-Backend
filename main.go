package main

import (
	"antibotapi/routes"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()

	register(app)
	listener(app)

}

func register(app *fiber.App) {
	api := app.Group("/api/v2/choate")
	routes.ChoateRobotics(api)
}

func listener(app *fiber.App) {
	err := app.Listen(":3000")
	if err != nil {
		fmt.Println("Error initializing BoarScout Backend")
	}
}
