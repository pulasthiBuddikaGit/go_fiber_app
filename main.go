package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/pulasthiBuddikaGit/go_fiber_app/config"
	//"github.com/pulasthiBuddikaGit/go_fiber_app/repository"
	"github.com/pulasthiBuddikaGit/go_fiber_app/storage"
)

func main() {
	// Load config (e.g., from .env)
	cfg := config.LoadConfig()

	// Initialize MongoDB connection
	storage.InitMongo(cfg)

	// Initialize repositories
	//repository.InitUserRepository(storage.Client.Database(cfg.Database))

	// Create a new Fiber app
	app := fiber.New()

	// 🟡 You can register routes later here...
	// routes.RegisterUserRoutes(app)

	// Start the Fiber server
	log.Println("🚀 Server is running on http://localhost:8080")
	if err := app.Listen(":8080"); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
}
