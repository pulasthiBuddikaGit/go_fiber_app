package storage

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/pulasthiBuddikaGit/go_fiber_app/config" // replace with your actual module name

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client              // global MongoDB client
var UserDB *mongo.Database          // global MongoDB database instance

func InitMongo(cfg *config.Config) {
	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB
	clientOptions := options.Client().ApplyURI(cfg.MongoURI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("MongoDB connection error:", err)
	}

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("MongoDB ping failed:", err)
	}

	fmt.Println("✅ Connected to MongoDB")

	Client = client
	UserDB = client.Database(cfg.Database)
}
