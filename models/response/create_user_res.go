package response

import (
	"time"

	"github.com/Luqky07/user_api/models/request"
)

type CreateUserResponse struct {
	Guid       string    `json:"guid"`
	Name       string    `json:"name"`
	LastName   string    `json:"lastname"`
	Birthdate  time.Time `json:"birthdate"`
	Email      string    `json:"email"`
	DocumentId string    `json:"documentid"`
	Phone      string    `json:"phone"`
}

func NewCreateUserResponse(req request.CreateUserRequest) CreateUserResponse {
	return CreateUserResponse{
		"1234567890",
		req.Name,
		req.LastName,
		req.Birthdate,
		req.Email,
		req.DocumentId,
		req.Phone,
	}
}
