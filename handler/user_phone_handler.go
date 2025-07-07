package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pulasthiBuddikaGit/go_fiber_app/model"
	"github.com/pulasthiBuddikaGit/go_fiber_app/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateUserPhoneHandler handles the creation of a user phone
func CreateUserPhoneHandler(c *fiber.Ctx) error {
	var input struct {
		UserID      string `json:"userId"`
		PhoneNumber string `json:"phoneNumber"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Convert string ID to ObjectID
	userObjectID, err := primitive.ObjectIDFromHex(input.UserID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}

	// Create the UserPhone model
	phone := model.UserPhone{
		UserID:      userObjectID,
		PhoneNumber: input.PhoneNumber,
	}

	result, err := repository.CreateUserPhone(&phone)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to save user phone",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User phone saved successfully",
		"id":      result.InsertedID,
	})
}
