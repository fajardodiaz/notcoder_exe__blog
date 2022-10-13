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

	// GORM function to create a record
	database.Db.Create(&usertype)

	// Return a response to user
	return c.JSON(fiber.Map{
		"success": true,
		"data":    &usertype,
	})
}

func UserTypeGet(c *fiber.Ctx) error {
	var usertype []models.UserType

	database.Db.Find(&usertype)
	return c.JSON(usertype)
}

func UserTypeGetByID(c *fiber.Ctx) error {
	var usertype models.UserType
	id, err := c.ParamsInt("id")
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"error": err,
		})
	}

	database.Db.First(&usertype, id)

	if usertype.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "User not found",
		})
	}

	return c.JSON(usertype)
}

func UserTypeDelete(c *fiber.Ctx) error {
	var usertype models.UserType

	id, err := c.ParamsInt("id")

	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"error": err,
		})
	}

	database.Db.First(&usertype, id)

	if usertype.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "User not found",
		})
	}

	database.Db.Delete(&usertype, id)
	return c.JSON(usertype)
}

func UserTypeUpdate(c *fiber.Ctx) error {

	return nil
}

// NotFound returns custom 404 page
func NotFound(c *fiber.Ctx) error {
	return c.Status(404).SendFile("./static/private/404.html")
}
