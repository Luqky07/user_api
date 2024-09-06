package controllers

import (
	"fmt"
	"net/http"

	"github.com/Luqky07/user_api/models/request"
	"github.com/Luqky07/user_api/models/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Function to create a new client
func CreateNewClient(c *gin.Context) {
	var req request.CreateUserRequest

	//Validate CreateUserRequest model
	if err := c.ShouldBindJSON(&req); err != nil {
		//Validating data
		if validationErrs, ok := err.(validator.ValidationErrors); ok {
			var errorMessages []string
			for _, fieldError := range validationErrs {
				//Custom validating error
				switch fieldError.Tag() {
				case "required":
					errorMessages = append(errorMessages, fmt.Sprintf("El campo %s es obligatorio.", fieldError.Field()))
				case "validDocument":
					errorMessages = append(errorMessages, "El documento de identidad no es correcto")
				case "email":
					errorMessages = append(errorMessages, fmt.Sprintf("El campo %s debe ser un correo electrónico válido.", fieldError.Field()))
				case "validPhone":
					errorMessages = append(errorMessages, "El teléfono no tiene un formato correcto")
				default:
					errorMessages = append(errorMessages, fmt.Sprintf("El campo %s tiene un error de validación.", fieldError.Field()))
				}
			}
			//Return validating error message
			c.JSON(http.StatusBadRequest, gin.H{"errors": errorMessages})
			return
		}

		//Return other type error
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res := response.NewCreateUserResponse(req)

	c.JSON(http.StatusCreated, res)
}
