package services

import (
	"github.com/Luqky07/user_api/data"
	"github.com/Luqky07/user_api/models/datamodels"
	"github.com/Luqky07/user_api/models/request"
	"github.com/Luqky07/user_api/models/response"
	"github.com/Luqky07/user_api/models/tables"
)

// Function to create a new User
func CreateNewClient(req request.CreateUserRequest) (*response.CreateUserResponse, datamodels.ApiError) {

	//Search if exist an user using documentId an email
	var client tables.Client
	result := data.DB.First(&client, "DocumentId = ?", req.DocumentId)

	if result.RowsAffected > 0 {
		return nil, datamodels.ApiError{StatusCode: 409, Message: "Ya existe un usuario asociado a ese documento de identidad"}
	}

	result = data.DB.First(&client, "Email = ?", req.Email)

	if result.RowsAffected > 0 {
		return nil, datamodels.ApiError{StatusCode: 409, Message: "Ya existe un usuario asociado a email"}
	}

	//Create de Client element in database
	client = tables.Client{
		Name:             req.Name,
		LastName:         req.LastName,
		Birthdate:        req.Birthdate,
		Email:            req.Email,
		DocumentId:       req.DocumentId,
		CreationEmployee: data.API_ADMIN,
	}
	result = data.DB.Create(&client)

	if result.Error != nil {
		return nil, datamodels.ApiError{StatusCode: 500, Message: result.Error.Error()}
	}

	var phone tables.Phone
	var clientPhone tables.ClientPhone

	//Search if the phone exist in database
	result = data.DB.First(&phone, "PhoneNumber = ?", req.Phone)

	if result.RowsAffected == 0 {

		//If not exist
		phone = tables.Phone{
			PhoneNumber:      req.Phone,
			CreationEmployee: data.API_ADMIN,
		}

		result = data.DB.Create(&phone)

		if result.Error != nil {
			return nil, datamodels.ApiError{StatusCode: 500, Message: result.Error.Error()}
		}
	}

	clientPhone = tables.ClientPhone{
		ClientId:         client.Id,
		PhoneId:          phone.Id,
		PhoneTypeId:      1,
		CreationEmployee: data.API_ADMIN,
	}

	result = data.DB.Create(&clientPhone)

	if result.Error != nil {
		return nil, datamodels.ApiError{StatusCode: 500, Message: result.Error.Error()}
	}

	response := response.NewCreateUserResponse(client, clientPhone)

	err := SendMail(req.Email, "Creación usuaio", "Se ha creado un nuevo usuario a tu nombre en shashastudio")
	if err != nil {
		return nil, datamodels.ApiError{StatusCode: 500, Message: err.Error()}
	}

	return &response, datamodels.ApiError{StatusCode: 200, Message: ""}
}
