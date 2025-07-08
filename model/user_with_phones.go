package model

//Used to send safe response to the client (e.g., GET /users/:id)
type UserWithPhones struct {
	//ID          string   `json:"id"`
	Name        string   `json:"name"`
	Email       string   `json:"email"`
	PhoneNumbers []string `json:"phoneNumbers"`
}
