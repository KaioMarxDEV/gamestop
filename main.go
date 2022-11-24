package main

import (
	"log"

	"github.com/KaioMarxDEV/gamestop/database"
	"github.com/KaioMarxDEV/gamestop/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// conn to docker database - workaround of container "ready" load
	// if you want to understand the why behind this: https://docs.docker.com/compose/startup-order/
	var err error
	for err != nil {
		err = database.ConnectDB()
	}

	app := fiber.New()  // http server instance
	app.Use(cors.New()) // enable cors if you want to call this API from js runtime Front-End

	routes.SetupRoutes(app) // router abstraction for http methods and routes

	log.Fatal(app.Listen(":3333")) // listener on localhost:3333 with exit log around it
}
