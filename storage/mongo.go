package storage

import (
	"context"
	"log"
	"time"

	"github.com/pulasthiBuddikaGit/go_fiber_app/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client       // exported Mongo client
var UserDB *mongo.Database     // exported DB instance

func InitMongo(cfg *config.Config) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOpts := options.Client().ApplyURI(cfg.MongoURI)

	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Fatalf("❌ Failed to connect to MongoDB: %v", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("❌ MongoDB ping failed: %v", err)
	}

	log.Println("✅ MongoDB connection established")

	Client = client
	UserDB = client.Database(cfg.Database)
}
