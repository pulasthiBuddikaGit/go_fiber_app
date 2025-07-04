

package repository

import (
	"context"
	"time"

	"github.com/pulasthiBuddikaGit/go_fiber_app/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var userCollection *mongo.Collection

// InitUserRepository initializes the user collection with the given MongoDB database
func InitUserRepository(db *mongo.Database) {
	userCollection = db.Collection("users")
}

// CreateUser inserts a new user document into the MongoDB "users" collection
func CreateUser(user *model.User) (*mongo.InsertOneResult, error) {
	// Set a timeout for the DB operation
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

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

// GetAllUsers returns all users from the collection
func GetAllUsers() ([]model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := userCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []model.User
	for cursor.Next(ctx) {
		var user model.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return users, nil
}






// package repository

// import (
// 	"context"
// 	"time"

// 	"github.com/pulasthiBuddikaGit/go_fiber_app/model"
// 	//"github.com/pulasthiBuddikaGit/go_fiber_app/storage"

// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"go.mongodb.org/mongo-driver/mongo"
// )

// //var userCollection *mongo.Collection = storage.UserDB.Collection("users")
// var userCollection *mongo.Collection

// func InitUserRepository(db *mongo.Database) {
// 	userCollection = db.Collection("users")
// }

// // CreateUser inserts a new user into the database
// func CreateUser(user *model.User) (*mongo.InsertOneResult, error) {

// 	//why ctx and cancel?
// 	// Context is used to control the timeout for the operation

// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	return userCollection.InsertOne(ctx, user)
// }

// // GetUserByID fetches a user by its ObjectID
// func GetUserByID(id string) (*model.User, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	objID, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var user model.User
// 	err = userCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &user, nil
// }

// // GetAllUsers returns all users in the collection
// func GetAllUsers() ([]model.User, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	cursor, err := userCollection.Find(ctx, bson.M{})
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer cursor.Close(ctx)

// 	var users []model.User
// 	for cursor.Next(ctx) {
// 		var user model.User
// 		if err := cursor.Decode(&user); err != nil {
// 			return nil, err
// 		}
// 		users = append(users, user)
// 	}

// 	return users, nil
// }

// // UpdateUser updates user fields by ID
// func UpdateUser(id string, updateData bson.M) (*mongo.UpdateResult, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	objID, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return userCollection.UpdateOne(
// 		ctx,
// 		bson.M{"_id": objID},
// 		bson.M{"$set": updateData},
// 	)
// }

// // DeleteUser deletes a user by ID
// func DeleteUser(id string) (*mongo.DeleteResult, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	objID, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return userCollection.DeleteOne(ctx, bson.M{"_id": objID})
// }
