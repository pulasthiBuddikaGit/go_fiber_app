package middleware

import (
	"os"
	"fmt"
	//I am using golang-jwt/jwt/v5 for token creation and validation , gofiber/jwt/v3 using for middleware to protect routes.

	//I renamed the import to jwtware to avoid conflict with github.com/golang-jwt/jwt/v5 if you use both in the same file. This is optional, but good for clarity.
	jwt "github.com/gofiber/jwt/v3"

	//jwtware "github.com/gofiber/jwt/v3"
	"github.com/gofiber/fiber/v2"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func JWTProtected() fiber.Handler {
	return jwt.New(jwt.Config{
		SigningKey: jwtSecret,
		//ErrorHandler is a function that will be called when the JWT validation fails.
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			fmt.Println("JWT ERROR:", err) // log the exact issue
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Invalid or expired JWT",
			})
		},
	})
}


// func JWTProtected() fiber.Handler {
// 	jwtSecret := []byte(os.Getenv("JWT_SECRET"))

// 	return jwtware.New(jwtware.Config{
// 		SigningKey: jwtSecret,
// 	})
// }
