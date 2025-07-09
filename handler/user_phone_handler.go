package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pulasthiBuddikaGit/go_fiber_app/repository"
)


//Needed
//I am using this struct because user is only updating phone number
type UpdatePhoneRequest struct {
	PhoneNumber string `json:"phoneNumber"`
}

// UpdateUserPhoneHandler handles PUT /user-phones/:id
//Needed
func UpdateUserPhoneHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	var body UpdatePhoneRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	result, err := repository.UpdateUserPhoneByID(id, body.PhoneNumber)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update phone number",
		})
	}

	if result.MatchedCount == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Phone number record not found",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Phone number updated successfully",
	})
}


//Needed
// DeleteUserPhoneHandler handles DELETE /user-phones/:id
func DeleteUserPhoneHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	result, err := repository.DeleteUserPhoneByID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete phone number",
		})
	}

	if result.DeletedCount == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Phone number not found",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Phone number deleted successfully",
	})
}
