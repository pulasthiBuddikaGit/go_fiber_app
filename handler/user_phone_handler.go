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

// GetUserPhoneNumbersByUserIDHandler handles GET /user-phones/:userId
// func GetUserPhoneNumbersByUserIDHandler(c *fiber.Ctx) error {
// 	userID := c.Params("userId")

// 	phones, err := repository.GetUserPhoneNumbersByUserID(userID)
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"error": "Failed to retrieve phone numbers",
// 		})
// 	}

// 	return c.JSON(phones)
// }

//I am using this struct because user is only updating phone number
type UpdatePhoneRequest struct {
	PhoneNumber string `json:"phoneNumber"`
}

// UpdateUserPhoneHandler handles PUT /user-phones/:id
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
