package routes

import (

	"github.com/gofiber/fiber/v2"
	"github.com/pulasthiBuddikaGit/go_fiber_app/handler"
)

// routes/auth_routes.go
func RegisterAuthRoutes(app fiber.Router) {
	app.Post("/login", handler.LoginHandler)
}
