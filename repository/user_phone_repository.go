package repository

import (
	"context"
	"time"

	"github.com/pulasthiBuddikaGit/go_fiber_app/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

// GetUserPhoneNumbersByUserID fetches all phone numbers associated with a specific user ID
func GetUserPhoneNumbersByUserID(userID string) ([]model.UserPhone, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"userId": objID}
	cursor, err := userPhoneCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var phones []model.UserPhone
	for cursor.Next(ctx) {
		var phone model.UserPhone
		if err := cursor.Decode(&phone); err != nil {
			return nil, err
		}
		phones = append(phones, phone)
	}

	return phones, nil
}