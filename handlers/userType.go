package handlers

import (
	"notcoder_exe__blog/database"
	"notcoder_exe__blog/models"

	"github.com/gofiber/fiber/v2"
)

// UserCreate registers a user
func UserTypeCreate(c *fiber.Ctx) error {

	var usertype models.UserType

	if err := c.BodyParser(&usertype); err != nil {
		return err
	}

	database.InsertUserType(&usertype)

	return c.JSON(fiber.Map{
		"success": true,
		"data":    &usertype,
	})
}

func UserTypeGet(c *fiber.Ctx) error {

	return nil
}

// NotFound returns custom 404 page
func NotFound(c *fiber.Ctx) error {
	return c.Status(404).SendFile("./static/private/404.html")
}
