package model

//Used to accept input from the client (e.g., POST /users)
type UserWithPhoneRequest struct {
	Name        string   `json:"name"`
	Email       string   `json:"email"`
	Password  string      `bson:"password,omitempty" json:"password,omitempty"` // Omit from JSON responses by default
	PhoneNumbers []string `json:"phoneNumbers"` // allow multiple numbers
	
}
