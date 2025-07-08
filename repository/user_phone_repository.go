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

//this function returns string array of phone numbers by userID
func GetPhoneNumbersByUserID(userID string) ([]string, error) {
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

	var phoneNumbers []string
	//Iterate through the found docs and append phone.numbers to the phoneNumbers array
	for cursor.Next(ctx) {
		var phone model.UserPhone
		if err := cursor.Decode(&phone); err != nil {
			return nil, err
		}
		//append every phone number to the phoneNumbers array
		phoneNumbers = append(phoneNumbers, phone.PhoneNumber)
	}

	return phoneNumbers, nil
}


// UpdateUserPhoneByID updates the phone number of a user by their ID
func UpdateUserPhoneByID(id string, newPhoneNumber string) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objID}
	update := bson.M{"$set": bson.M{"phoneNumber": newPhoneNumber}}

	return userPhoneCollection.UpdateOne(ctx, filter, update)
}

// DeleteUserPhoneByID deletes a user phone document by its ID
func DeleteUserPhoneByID(id string) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	return userPhoneCollection.DeleteOne(ctx, bson.M{"_id": objID})
}
