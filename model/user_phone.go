package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// UserPhone represents a phone document in the "user_phones" collection
type UserPhone struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID      primitive.ObjectID `bson:"userId" json:"userId"` // reference to User
	PhoneNumber string             `bson:"phoneNumber" json:"phoneNumber"`
}
