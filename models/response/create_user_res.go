package response

import (
	"time"

	"github.com/Luqky07/user_api/models/tables"
)

// Struct to manage the response of the users creation
type CreateUserResponse struct {
	UserId     string    `json:"userId"`
	Name       string    `json:"name"`
	LastName   string    `json:"lastname"`
	Birthdate  time.Time `json:"birthdate"`
	Email      string    `json:"email"`
	DocumentId string    `json:"documentid"`
	PhoneId    string    `json:"phoneId"`
}

func NewCreateUserResponse(client tables.Client, phone tables.ClientPhone) CreateUserResponse {
	return CreateUserResponse{
		UserId:     client.Id,
		Name:       client.Name,
		LastName:   client.LastName,
		Birthdate:  client.Birthdate,
		Email:      client.Email,
		DocumentId: client.DocumentId,
		PhoneId:    phone.PhoneId,
	}
}
