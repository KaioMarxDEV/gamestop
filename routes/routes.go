package routes

import (
	"github.com/KaioMarxDEV/gamestop/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/gamestop/api")
	api.Get("/", controllers.Hello)

	users := api.Group("/user")
	users.Post("/", controllers.CreateUser)
}
