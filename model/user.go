package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


// User represents a user document in the "users" collection
type User struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name  string             `bson:"name" json:"name"`
	Email string             `bson:"email" json:"email" unique:"true"` // Unique index on email
	Password  string              `bson:"password,omitempty" json:"password,omitempty"` // Omit from JSON responses by default
	CreatedAt time.Time           `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time           `bson:"updatedAt" json:"updatedAt"`
	DeletedAt *time.Time          `bson:"deletedAt,omitempty" json:"deletedAt,omitempty"` // Nil if not deleted (soft delete)
}
