package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// User represents a user document in the "users" collection
type User struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name  string             `bson:"name" json:"name"`
	Email string             `bson:"email" json:"email"`
}
