package handlers

import (
	"github.com/gofiber/fiber/v2"
	"notcoder_exe__blog/database"
	"notcoder_exe__blog/models"
)

func FieldCreate(c *fiber.Ctx) error {
	var field = models.Field{}

	if err := c.BodyParser(&field); err != nil {
		return err
	}

	database.Db.Create(&field)
	return c.JSON(fiber.Map{
		"success": true,
		"data":    &field,
	})
}

func FieldGet(c *fiber.Ctx) error {
	var field []models.Field

	database.Db.Find(&field)

	return c.JSON(field)
}

//func FieldGetByID(c *fiber.Ctx) error{
//
//}
//
//func FieldUpdate(c *fiber.Ctx) error{
//
//}
//
//func FieldDelete(c *fiber.Ctx) error{
//
//}
