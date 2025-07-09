package handler

import (
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"

	"github.com/gofiber/fiber/v2"
	"github.com/pulasthiBuddikaGit/go_fiber_app/model"
	"github.com/pulasthiBuddikaGit/go_fiber_app/repository"
	"go.mongodb.org/mongo-driver/bson"
)


//bind the both user and phone numbers to one endpoint and call the CreateUser and CreateUserPhone functions to save both details in separate collections
func CreateUserHandler(c *fiber.Ctx) error {
	var req model.UserWithPhoneRequest

	//req body has both user details and phone numbers
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request format",
		})
	}

	// Validate required fields (optional but recommended)
	if req.Password == "" || req.Email == "" || req.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Name, Email, and Password are required",
		})
	}

	 	// Hash the password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to hash password",
		})
	}

	// Step 1: Create the user
	//this user struct is used to save the user details in the users collection
	user := model.User{
		Name:      req.Name,
		Email:     req.Email,
		Password:  string(hashedPassword), // Store the hashed password
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	userResult, err := repository.CreateUser(&user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}

	// Step 2: Save phone numbers

	//take the insertedID from the userResult and use it to save the phone numbers
	insertedID := userResult.InsertedID.(primitive.ObjectID)

	// Loop through the phone numbers and create UserPhone docs
	for _, phone := range req.PhoneNumbers {
		userPhone := model.UserPhone{
			UserID:      insertedID,
			PhoneNumber: phone,
		}
		_, err := repository.CreateUserPhone(&userPhone)
		if err != nil {
			// Optionally log or rollback here
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to create phone number",
			})
		}
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User and phone(s) created successfully",
	})
}


func GetUserByIDHandler(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	// Fetch the user
	user, err := repository.GetUserByID(id)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	// Fetch user's phone numbers
	//phones, err := repository.GetPhoneNumbersByUserID(id)
	phones, err := repository.GetPhoneNumbersByUserID(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve phone numbers",
		})
	}

	// Combine response
	//this response struct is used to send the user details and phone numbers in one response
	//It is mirror UserWithPhones model
	response := model.UserWithPhones{
		//ID:           user.ID.Hex(),
		Name:         user.Name,
		Email:        user.Email,
		PhoneNumbers: phones,
	}

	return ctx.JSON(response)
}


func GetAllUsersHandler(c *fiber.Ctx) error {
	log.Println("📄 GetAllUsersHandler called")

	users, err := repository.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch users with phone numbers",
		})
	}

	return c.Status(fiber.StatusOK).JSON(users)
}

// UpdateUserHandler handles PUT /users/:id
func UpdateUserHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	var updateUser model.User
	if err := c.BodyParser(&updateUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	updateData := bson.M{}
	if updateUser.Name != "" {
		updateData["name"] = updateUser.Name
	}
	if updateUser.Email != "" {
		updateData["email"] = updateUser.Email
	}

	//update the updatedAt field to the current time
	updateData["updatedAt"] = time.Now()

	result, err := repository.UpdateUser(id, updateData)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update user",
		})
	}

	if result.MatchedCount == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	return c.JSON(fiber.Map{
		"message": "User updated successfully",
	})
}

// DeleteUserHandler handles DELETE /users/:id
func DeleteUserHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	result, err := repository.DeleteUser(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete user",
		})
	}

	if result.DeletedCount == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	return c.JSON(fiber.Map{
		"message": "User deleted successfully",
	})
}



