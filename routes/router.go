package routes

import (
	"antibotapi/robotics"
	"github.com/gofiber/fiber/v2"
)

func ChoateRobotics(app fiber.Router) {
	app.Post("/auth", robotics.AuthenticateUser)
	app.Post("/session", robotics.CollectData)
	app.Get("/teams", robotics.GetTeams)
	app.Post("/sheets", robotics.AppendSheet)
}
