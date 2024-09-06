package request

import "time"

// Struct to manage the users creation
type CreateUserRequest struct {
	Name       string    `json:"name" binding:"required"`
	LastName   string    `json:"lastname" binding:"required"`
	Birthdate  time.Time `json:"birthdate" binding:"required"`
	Email      string    `json:"email" binding:"required,email"`
	DocumentId string    `json:"documentid" binding:"required,validDocument"`
	Phone      string    `json:"phone" binding:"required,validPhone"`
}
