

package repository

import (
	"context"
	"time"
	"log"

	"github.com/pulasthiBuddikaGit/go_fiber_app/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var userCollection *mongo.Collection

// InitUserRepository initializes the user collection with the given MongoDB database
func InitUserRepository(db *mongo.Database) {
	userCollection = db.Collection("users")

		// Create unique index on email field
	indexModel := mongo.IndexModel{
		Keys: bson.D{{Key: "email", Value: 1}}, // Index key
		Options: options.Index().SetUnique(true),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//try to create this index when user inserting a new user
	_, err := userCollection.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		log.Fatal("❌ Failed to create unique index on email:", err)
	}

	log.Println("✅ Unique index on email created (or already exists)")
}

// CreateUser inserts a new user document into the MongoDB "users" collection
func CreateUser(user *model.User) (*mongo.InsertOneResult, error) {
	// Set a timeout for the DB operation
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Set the CreatedAt and UpdatedAt fields to the current time
	//mn meka methna update kale mkd user me deatils 
	

	return userCollection.InsertOne(ctx, user)
}

// GetUserByID fetches a user document by its MongoDB ObjectID
func GetUserByID(id string) (*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Convert the hex string ID to MongoDB ObjectID
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var user model.User
	err = userCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}


func GetAllUsers() ([]model.UserWithPhones, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := userCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []model.UserWithPhones

	for cursor.Next(ctx) {
		var user model.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}

		// Fetch phone numbers for each user
		phones, err := GetPhoneNumbersByUserID(user.ID.Hex())
		if err != nil {
			return nil, err
		}

		result := model.UserWithPhones{
			//ID:           user.ID.Hex(),
			Name:         user.Name,
			Email:        user.Email,
			PhoneNumbers: phones,
		}
		results = append(results, result)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

// repository/user_repository.go
//because of this function we had to make email field unique
func GetUserByEmail(email string) (*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user model.User
	err := userCollection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}



// UpdateUser updates a user's fields by ID
func UpdateUser(id string, updateData bson.M) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	update := bson.M{"$set": updateData}

	//UpdateOne finds the user and updates it with the provided data
	result, err := userCollection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		return nil, err
	}

	/*This function returns struct provided by the MongoDB driver
	That struct name is UpdateResult
	this struct contains attributes like:
	result.MatchedCount  // is 1 if the user was found, otherwise 0
	result.ModifiedCount // is 1 if the update actually changed any fields */

	return result, nil
}

// DeleteUser deletes a user by ID
func DeleteUser(id string) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	result, err := userCollection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return nil, err
	}

	return result, nil
}


