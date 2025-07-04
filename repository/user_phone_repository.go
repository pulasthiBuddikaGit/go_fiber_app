package repository

import (
	"context"
	"time"

	"github.com/pulasthiBuddikaGit/go_fiber_app/model"
	"go.mongodb.org/mongo-driver/mongo"
)

var userPhoneCollection *mongo.Collection

// InitUserPhoneRepository initializes the user_phones collection
func InitUserPhoneRepository(db *mongo.Database) {
	userPhoneCollection = db.Collection("user_phones")
}

// CreateUserPhone inserts a new user phone document into the collection
func CreateUserPhone(phone *model.UserPhone) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return userPhoneCollection.InsertOne(ctx, phone)
}
