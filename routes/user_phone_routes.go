package routes

import (

	"github.com/gofiber/fiber/v2"
	"github.com/pulasthiBuddikaGit/go_fiber_app/handler"
)

// RegisterUserRoutes defines routes related to users
func RegisterUserPhoneRoutes(app fiber.Router) {
	userPhoneGroup := app.Group("/user-phones")

	userPhoneGroup.Post("/", handler.CreateUserPhoneHandler)
	userPhoneGroup.Get("/:userId", handler.GetUserPhoneNumbersByUserIDHandler)
	userPhoneGroup.Put("/:id", handler.UpdateUserPhoneHandler)
	// userGroup.Get("/", handler.GetAllUsersHandler)
	// userGroup.Get("/:id", handler.GetUserByIDHandler)
	// userGroup.Put("/:id", handler.UpdateUserHandler)
	// userGroup.Delete("/:id", handler.DeleteUserHandler)
}