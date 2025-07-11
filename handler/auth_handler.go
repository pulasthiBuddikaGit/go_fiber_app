// handler/auth_handler.go
package handler

import (
	"time"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pulasthiBuddikaGit/go_fiber_app/repository"
	//"github.com/pulasthiBuddikaGit/go_fiber_app/model"
	"golang.org/x/crypto/bcrypt"
	"os"
)

//get the JWT secret from environment variable
// Alternatively, you can use a config package to load this from a config file
var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func LoginHandler(c *fiber.Ctx) error {
	var creds struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&creds); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// Get user from DB
	user, err := repository.GetUserByEmail(creds.Email)
	if err != nil || user == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	// Create JWT
	claims := jwt.MapClaims{
		"user_id": user.ID.Hex(),
		"exp":     time.Now().Add(time.Hour * 72).Unix(), // expires in 3 days
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//here token get signed by our secret key
	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create token"})
	}

	// Return token
	return c.JSON(fiber.Map{"token": signedToken})
}
