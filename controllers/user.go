package controllers

import (
	"github.com/KaioMarxDEV/gamestop/database"
	"github.com/KaioMarxDEV/gamestop/models"
	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	db := database.DB
	usr := new(models.User)

	// get body data from request and copies to usr var
	if err := c.BodyParser(usr); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Data type incorrect", "data": err})
	}

	// queries the database to insert this new user and verifies if it succeed
	err := db.Create(&usr).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": err})
	}

	return c.Status(201).JSON(fiber.Map{
		"status": "success",
		"message": fiber.Map{
			"userid": usr.ID,
		},
	})
}
